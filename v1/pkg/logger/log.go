// Package logger 是对标准库 log/slog 的轻量级、生产级封装。
// 额外提供了：
//   - 日志级别 Trace / Fatal；
//   - 动态调整日志级别与输出；
//   - 可选调用栈 / 来源打印；
//   - 包级默认 Logger。
package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"sync/atomic"
	"time"

)

// handlerBox 把 slog.Handler 做一次装箱，方便原子级整体替换。
type handlerBox struct{ h slog.Handler }

// Level 只是 slog.Level 的别名，便于外部统一使用。
type Level = slog.Level

// 预定义日志级别，向下兼容 slog。
const (
	LevelTrace Level = slog.LevelDebug - 4 // 比 Debug 更细
	LevelDebug Level = slog.LevelDebug
	LevelInfo  Level = slog.LevelInfo
	LevelWarn  Level = slog.LevelWarn
	LevelError Level = slog.LevelError
	LevelFatal Level = slog.LevelError + 4 // 最高级别，打印后进程退出
)

// Logger 是增强版日志器，内部通过原子操作保证并发安全。
type Logger struct {
	app    string                     // 应用名，会作为固定属性写入每条日志
	addSrc int32                      // 原子标志：是否打印调用位置（0/1）
	skip  int                        //
	level  atomic.Int32               // 当前全局日志级别，可运行时修改
	hptr   atomic.Pointer[handlerBox] // 指向当前真正工作的 slog.Handler
}

// New 创建一个 Logger，默认级别 Info、JSON 格式输出到 os.Stdout。
// 可通过 Option 覆盖默认行为。
func New(app string, opts ...Option) *Logger {
	l := &Logger{app: app}
	l.level.Store(int32(LevelInfo))
	l.skip = 2
	// 初始 handler：JSON + Info
	box := &handlerBox{
		h: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	}
	l.hptr.Store(box)

	// 应用用户选项
	for _, o := range opts {
		o(l)
	}
	return l
}

// Option 定义函数式选项模式。
type Option func(*Logger)

// WithLevel 设置初始日志级别。
func WithLevel(lv Level) Option {
	return func(l *Logger) { l.SetLevel(lv) }
}

// WithOutput 设置输出目的地。
func WithOutput(w io.Writer) Option {
	return func(l *Logger) { l.SetOutput(w) }
}

// swapHandler 是一个内部工具：
// 根据给定工厂函数和 io.Writer 构建新的 handlerBox 并原子替换。
func (l *Logger) swapHandler(fn func(io.Writer) slog.Handler, w io.Writer) {
	box := &handlerBox{h: fn(w)}
	l.hptr.Store(box)
}

// WithText 将日志格式切换为 Text（人类可读）。
func WithText() Option {
	return func(l *Logger) {
		l.swapHandler(func(w io.Writer) slog.Handler {
			return slog.NewTextHandler(w, &slog.HandlerOptions{
				Level: Level(l.level.Load()),
			})
		}, l.getOutput())
	}
}

// WithJSON 将日志格式切换为 JSON（机器可读）。
func WithJSON() Option {
	return func(l *Logger) {
		l.swapHandler(func(w io.Writer) slog.Handler {
			return slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: Level(l.level.Load()),
			})
		}, l.getOutput())
	}
}

// SetOutput 动态修改输出目的地，同时保持当前格式（Text/JSON）。
func (l *Logger) SetOutput(w io.Writer) {
	cur := l.hptr.Load().h
	switch cur.(type) {
	case *slog.TextHandler:
		l.swapHandler(func(w io.Writer) slog.Handler {
			return slog.NewTextHandler(w, &slog.HandlerOptions{
				Level: Level(l.level.Load()),
			})
		}, w)
	default:
		l.swapHandler(func(w io.Writer) slog.Handler {
			return slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: Level(l.level.Load()),
			})
		}, w)
	}
}

// WithAddSource 设置是否打印调用位置（文件:行号）。
func WithAddSource(v bool) Option {
	return func(l *Logger) { l.SetAddSource(v) }
}

func WithSkip(v int) Option {
	return func(l *Logger) { l.skip = v }
}

// getOutput 获取当前 handler 关联的 io.Writer，拿不到则退回 os.Stdout。
func (l *Logger) getOutput() io.Writer {
	if h, ok := l.hptr.Load().h.(interface{ Writer() io.Writer }); ok {
		return h.Writer()
	}
	return os.Stdout
}

// SetLevel 运行时安全地修改全局日志级别。
func (l *Logger) SetLevel(lv Level) { l.level.Store(int32(lv)) }

// SetAddSource 运行时开关调用位置打印。
func (l *Logger) SetAddSource(v bool) {
	if v {
		atomic.StoreInt32(&l.addSrc, 1)
	} else {
		atomic.StoreInt32(&l.addSrc, 0)
	}
}

// Enabled 实现 slog.Handler 接口，用于判断某条日志是否需要打印。
func (l *Logger) Enabled(ctx context.Context, lv Level) bool {
	return lv >= Level(l.level.Load())
}

// log 是内部统一写日志的入口。
func (l *Logger) log(ctx context.Context, lv Level, msg string, args ...any) {
	if !l.Enabled(ctx, lv) {
		return
	}
	h := l.hptr.Load().h
	record := slog.NewRecord(now(), lv, msg, l.pc())
	record.AddAttrs(l.buildAttrs(ctx, args...)...)
	_ = h.Handle(ctx, record)
}

// buildAttrs 把用户传入的 k-v 对转换成 slog.Attr 数组，同时附加固定字段。
func (l *Logger) buildAttrs(ctx context.Context, kv ...any) []slog.Attr {
	var a []slog.Attr

	// 调用位置
	if atomic.LoadInt32(&l.addSrc) == 1 {
		a = append(a, slog.String("source", l.caller()))
	}

	// 应用名
	a = append(a, slog.String("app", l.app))

	// 从 ctx 中尝试拿 request_id
	if ctx != nil {
		if v := ctx.Value("request_id"); v != nil {
			a = append(a, slog.String("request_id", fmt.Sprint(v)))
		}
	}

	// 解析 kv 列表
	for i := 0; i < len(kv); {
		switch v := kv[i].(type) {
		case slog.Attr:
			a = append(a, v)
			i++
		default:
			if i+1 < len(kv) {
				key, ok := kv[i].(string)
				if !ok {
					key = "key"
				}
				a = append(a, slog.Any(key, kv[i+1]))
				i += 2
			} else {
				a = append(a, slog.Any("key", kv[i]))
				i++
			}
		}
	}
	return a
}

// caller 返回调用位置字符串：file:line。
// func (l *Logger) caller() string {
// 	var pcs [1]uintptr
// 	runtime.Callers(3, pcs[:])
// 	frame, _ := runtime.CallersFrames(pcs[:]).Next()
// 	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
// }

func (l *Logger) caller() string {
	skip := 2 + l.skip // 2 = runtime.Callers + caller 自身
	const maxDepth = 10
	var pcs [maxDepth]uintptr
	n := runtime.Callers(skip, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	for i := 0; i <= l.skip; i++ {
		frame, more := frames.Next()
		if !more {
			return "<unknown>"
		}
		if i == l.skip {
			return fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
	}
	return "<unknown>"
}

// pc 返回调用位置程序计数器，用于 slog.Record。
func (l *Logger) pc() uintptr {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	return pcs[0]
}

// now 返回当前时间，便于单测打桩。
func now() time.Time { return time.Now() }

// -------------- 分级快捷方法 --------------

func (l *Logger) Trace(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelTrace, msg, kv...)
}
func (l *Logger) Debug(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelDebug, msg, kv...)
}
func (l *Logger) Info(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelInfo, msg, kv...)
}
func (l *Logger) Warn(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelWarn, msg, kv...)
}
func (l *Logger) Error(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelError, msg, kv...)
}
func (l *Logger) Fatal(ctx context.Context, msg string, kv ...any) {
	l.log(ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}

// Must 在 err != nil 时 Fatal 并附带额外 kv，否则静默。
func (l *Logger) Must(ctx context.Context, err error, kv ...any) {
	if err != nil {
		l.Fatal(ctx, err.Error(), kv...)
	}
}

// -------------- 包级默认 Logger --------------

// Default 是开箱即用的全局 Logger，应用名 "default"。
var Default = New("default", WithAddSource(true),WithSkip(2))

// 下面暴露包级快捷函数，内部都转发到 Default。
func Trace(ctx context.Context, msg string, kv ...any) { Default.Trace(ctx, msg, kv...) }
func Debug(ctx context.Context, msg string, kv ...any) { Default.Debug(ctx, msg, kv...) }
func Info(ctx context.Context, msg string, kv ...any)  { Default.Info(ctx, msg, kv...) }
func Warn(ctx context.Context, msg string, kv ...any)  { Default.Warn(ctx, msg, kv...) }
func Error(ctx context.Context, msg string, kv ...any) { Default.Error(ctx, msg, kv...) }
func Fatal(ctx context.Context, msg string, kv ...any) { Default.Fatal(ctx, msg, kv...) }
func Must(ctx context.Context, err error, kv ...any)   { Default.Must(ctx, err, kv...) }

// -------------- Sprintf 风格扩展 --------------

// Logf 提供 fmt.Sprintf 风格的格式化消息，并可携带结构化字段。
// 例：logx.Logf(ctx, logx.LevelInfo, "user %d login from %s", uid, ip, "role", "admin")
func (l *Logger) Logf(ctx context.Context, lv Level, format string, a ...any) {
	if !l.Enabled(ctx, lv) {
		return
	}

	// 前 len(a)-len(kv) 个元素用于 Sprintf，其余是 kv 对
	// 为了简单，我们约定：当 a 长度为奇数时，最后一个元素会被当成“孤立值”处理
	n := len(a)
	msg := fmt.Sprintf(format, a[:n]...) // 整条消息
	l.log(ctx, lv, msg)                  // kv 为空
}

// 以下六个分级便利方法把 Logf 的 lv 参数固化，调用更短。

func (l *Logger) Tracef(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelTrace, format, a...)
}
func (l *Logger) Debugf(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelDebug, format, a...)
}
func (l *Logger) Infof(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelInfo, format, a...)
}
func (l *Logger) Warnf(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelWarn, format, a...)
}
func (l *Logger) Errorf(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelError, format, a...)
}
func (l *Logger) Fatalf(ctx context.Context, format string, a ...any) {
	l.Logf(ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// -------------- 包级默认 Logger 的 Sprintf 快捷方法 --------------

func Logf(ctx context.Context, lv Level, format string, a ...any) {
	Default.Logf(ctx, lv, format, a...)
}
func Tracef(ctx context.Context, format string, a ...any) { Default.Tracef(ctx, format, a...) }
func Debugf(ctx context.Context, format string, a ...any) { Default.Debugf(ctx, format, a...) }
func Infof(ctx context.Context, format string, a ...any)  { Default.Infof(ctx, format, a...) }
func Warnf(ctx context.Context, format string, a ...any)  { Default.Warnf(ctx, format, a...) }
func Errorf(ctx context.Context, format string, a ...any) { Default.Errorf(ctx, format, a...) }
func Fatalf(ctx context.Context, format string, a ...any) { Default.Fatalf(ctx, format, a...) }

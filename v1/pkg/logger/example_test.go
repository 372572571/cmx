package logger

import (
	"context"
	"os"
	"testing"
)

func Example() {
	ctx := context.Background()
	// 1) 创建带选项的 logger
	log := New("billing",
		WithLevel(LevelDebug),
		WithOutput(os.Stderr),
		WithText(),
		WithAddSource(true),
	)

	log.Info(ctx, "hello", "user", "alice", "count", 3)
	log.Must(ctx, nil) // 不会 panic
	defer func(l *Logger) {
		if r := recover(); r != nil {
			l.Error(ctx, "panic", "error", r)
		}
	}(log)
	log.Fatalf(ctx, "hello", "user", "alice", "count", 3)
	// Output:
	// time=... level=INFO source=... app=billing msg=hello user=alice count=3
}

func Test_Example(t *testing.T) {
	Example()
}

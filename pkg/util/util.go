package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// LoadDirAllFile
func LoadDirAllFile(root string, suffix []string) []string {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("LoadDirAllFile: root: %s suffix: %v err: %v \n", root, suffix, err)
		}
	}()
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, s := range suffix {
				if filepath.Ext(path) == s {
					files = append(files, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
func MustSucc[T any](s T, err error) T {
	if err != nil {
		NoError(err)
	}
	return s
}

func ToCamelCasing(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	s += "."
	n := strings.Builder{}
	n.Grow(len(s))
	temp := strings.Builder{}
	temp.Grow(len(s))
	wordFirst := true
	for _, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if wordFirst && vIsLow {
			v -= 'a' - 'A'
		}

		if vIsCap || vIsLow {
			temp.WriteByte(v)
			wordFirst = false
		} else {
			isNum := v >= '0' && v <= '9'
			wordFirst = isNum || v == '_' || v == '/' || v == ' ' || v == '-' || v == '.' || v == ':'
			if temp.Len() > 0 && wordFirst {
				word := temp.String()
				// upper := strings.ToUpper(word)
				// if _, ok := acronym[upper]; ok {
				// 	n.WriteString(upper)
				// } else {
				n.WriteString(word)
				// }
				temp.Reset()
			}
			if isNum {
				n.WriteByte(v)
			}
		}
	}
	return n.String()
}

func FirstLowerCamelCasing(s string) string {
	result := ToCamelCasing(s)
	if len(result) > 0 {
		r := []rune(result)
		r[0] = unicode.ToLower(r[0])
		return string(r)
	}
	return result
}

func IsHaveFile(p string) bool {
	return isHave(p)
}

func IsHaveDir(p string) bool {
	return isHave(p)
}

// is exist
func isHave(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// 注释字符串
func Comment(s string) string {
	return "// " + strings.ReplaceAll(s, "\n", "\n// ")
}

// 转义换行符号
func Escape(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}

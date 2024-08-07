package util

import (
	"fmt"
	"testing"
)

func TestLoadDirAllFile(t *testing.T) {
	path := "test/"

	fmt.Println(LoadDirAllFile(path, []string{".txt"}))
}

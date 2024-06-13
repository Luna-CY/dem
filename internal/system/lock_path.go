package system

import (
	"os"
)

// Executable 判断文件是否可执行
func Executable(file string) bool {
	fi, err := os.Stat(file)
	if nil != err {
		return false
	}

	if fi.IsDir() {
		return false
	}

	if 0 == fi.Mode()&0111 {
		return false
	}

	return true
}

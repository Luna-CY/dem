package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// GetStringFromEnv 获取环境变量
// 如果value不为空直接返回
// 否则判断环境变量是否存在
// 如果不存在返回默认值
func GetStringFromEnv(value string, key string, d string) string {
	if "" != value {
		return value
	}

	value = os.Getenv(key)
	if "" != value {
		return value
	}

	return d
}

// RemoveAll 删除路径
func RemoveAll(path string) error {
	// 提权
	if err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}

		if err := os.Chmod(path, 0777); nil != err {
			return err
		}

		return nil
	}); nil != err && !os.IsNotExist(err) {
		return err
	}

	return os.RemoveAll(path)
}

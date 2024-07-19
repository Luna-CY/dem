package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
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

// Checksum 计算文件的校验和
func Checksum(path string) (string, error) {
	var file, err = os.Open(path)
	if nil != err {
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	var hash = sha256.New()
	if _, err := io.Copy(hash, file); nil != err {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

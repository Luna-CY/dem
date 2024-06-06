package utils

import "os"

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

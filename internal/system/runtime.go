package system

import "runtime"

// GetSystemArch 获取系统架构类型
func GetSystemArch() string {
	return runtime.GOOS + "-" + runtime.GOARCH
}

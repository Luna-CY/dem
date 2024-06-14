package system

import (
	"github.com/Luna-CY/dem/internal/utils"
	"path/filepath"
	"strings"
)

const DemRootPath = "/opt/godem"

// GetRootPath 获取DEM系统根目录
func GetRootPath() string {
	return utils.GetStringFromEnv("", DemEnvPrefix+"ROOT_PATH", DemRootPath)
}

// GetIndexPath 获取DEM索引路径
func GetIndexPath() string {
	return filepath.Join(GetRootPath(), "index")
}

// GetPkgPath 获取包路径
func GetPkgPath() string {
	return filepath.Join(GetRootPath(), "packages")
}

// ReplaceVariables 替换变量
func ReplaceVariables(s string, name string, value string) string {
	return strings.ReplaceAll(s, name, value)
}

// GetSystemEnvironmentPath 获取DEM系统环境文件路径
func GetSystemEnvironmentPath() string {
	return filepath.Join(GetRootPath(), "config", "environment.json")
}

// GetPackageRootPath 获取包的根目录
func GetPackageRootPath(pkg string) string {
	return filepath.Join(GetPkgPath(), pkg)
}

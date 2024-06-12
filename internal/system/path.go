package system

import (
	"github.com/Luna-CY/dem/internal/utils"
	"path/filepath"
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

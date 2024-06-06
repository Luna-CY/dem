package system

import "github.com/Luna-CY/dem/internal/utils"

const DemRootPath = "/opt/godem"

// GetRootPath 获取DEM系统根目录
func GetRootPath() string {
	return utils.GetStringFromEnv("", DemEnvPrefix+"ROOT_PATH", DemRootPath)
}

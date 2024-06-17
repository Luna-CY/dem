package pkg

import (
	"context"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
)

// Uninstall 移除安装的工具包
func Uninstall(ctx context.Context, name string) error {
	ind, err := index.Lookup(name)
	if nil != err {
		return echo.Error("查找工具包失败: %s", err)
	}

	var path = system.GetPackageRootPath(ind.PackageName)

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return echo.Error("工具包[%s]不支持当前平台: %s", name, system.GetSystemArch())
	}

	for _, cmd := range platform.Uninstall {
		cmd = system.ReplaceVariables(cmd, "{ROOT}", path)

		if err := utils.ExecuteShellCommand(ctx, cmd, nil); nil != err {
			return echo.Error("移除安装的工具包[%s]失败: %s", name, err)
		}
	}

	return nil
}

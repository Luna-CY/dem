package pkg

import (
	"context"
	"fmt"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
)

// Uninstall 移除安装的工具包
func Uninstall(ctx context.Context, name string) error {
	ind, err := index.Lookup(name)
	if nil != err {
		return fmt.Errorf("find package failed: %s", err)
	}

	var path = system.GetPackageRootPath(ind.PackageName)

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return fmt.Errorf("package[%s] does not support current platform: %s", name, system.GetSystemArch())
	}

	for _, cmd := range platform.Uninstall {
		cmd = system.ReplaceVariables(cmd, "{ROOT}", path)

		if err := utils.ExecuteShellCommand(ctx, cmd, nil); nil != err {
			return fmt.Errorf("uninstall package[%s] failed: %s", name, err)
		}
	}

	return nil
}

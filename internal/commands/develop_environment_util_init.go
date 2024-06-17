package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentUtilInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初始化当前项目环境",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			pe, err := environment.GetProjectEnvironment()
			if nil != err {
				return echo.Error("查询当前项目环境信息失败: %s", err)
			}

			for name, version := range pe.Packages {
				var pkgName = name + "@" + version

				installed, err := pkg.Installed(pkgName)
				if nil != err {
					return echo.Error("检查工具包[%s]安装状态失败: %s", pkgName, err)
				}

				if installed {
					continue
				}

				_ = echo.Info("安装工具包[%s]...", pkgName)
				if err := pkg.Install(cmd.Context(), pkgName); nil != err {
					return echo.Error("安装工具包[%s]失败: %s", pkgName, err)
				}

				_ = echo.Info("工具包[%s]安装成功", name)
			}

			return nil
		},
	}
}

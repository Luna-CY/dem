package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentManagementUtilUninstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall pkg [pkg [...]]",
		Short: "移除安装的工具包",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, name := range args {
				installed, err := pkg.Installed(name)
				if nil != err {
					_ = echo.Error("检查工具包[%s]安装状态失败: %s", name, err)

					continue
				}

				if !installed {
					_ = echo.Info("工具包[%s]未安装，跳过", name)

					continue
				}

				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					_ = echo.Error("查询项目环境配置失败: %s", err)

					continue
				}

				if _, ok := pe.Packages[name]; ok {
					delete(pe.Packages, name)

					if err := pe.Save(); nil != err {
						_ = echo.Error("清理工具包[%s]的环境配置失败: %s", name, err)

						continue
					}
				}

				se, err := environment.GetSystemEnvironment()
				if nil != err {
					_ = echo.Error("查询系统环境配置失败: %s", err)

					continue
				}

				if _, ok := se.Packages[name]; ok {
					delete(se.Packages, name)

					if err := se.Save(); nil != err {
						_ = echo.Error("清理工具包[%s]的环境配置失败: %s", name, err)

						continue
					}
				}

				_ = echo.Info("移除安装的工具包[%s]...", name)
				if err := pkg.Uninstall(cmd.Context(), name); nil != err {
					_ = echo.Error(err.Error())

					continue
				}

				_ = echo.Info("工具包[%s]移除完成", name)
			}

			return nil
		},
	}
}

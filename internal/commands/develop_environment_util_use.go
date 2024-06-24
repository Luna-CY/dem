package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	system2 "github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewDevelopEnvironmentUtilUseCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "use [options] package",
		Short: "设定环境工具包",
		RunE: func(cmd *cobra.Command, args []string) error {
			var tokens = strings.Split(args[0], "@")
			if 2 != len(tokens) {
				_ = echo.Error("工具包名称无效")

				os.Exit(1)
			}

			fi, err := os.Stat(filepath.Join(system2.GetPackageRootPath(args[0]), ".installed"))
			if nil != err {
				if os.IsNotExist(err) {
					_ = echo.Error("工具包[%s]未安装", args[0])

					os.Exit(1)
				}

				_ = echo.Error("检查工具包[%s]安装状态失败: %s", args[0], err)

				os.Exit(1)
			}

			if fi.IsDir() {
				_ = echo.Error("检查工具包[%s]状态无效，请重新安装", args[0])

				os.Exit(1)
			}

			var env *environment.Environment

			if system {
				se, err := environment.GetSystemEnvironment()
				if nil != err {
					_ = echo.Error("查询全局环境配置失败: %s", err)

					os.Exit(1)
				}

				env = se
			} else {
				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					_ = echo.Error("查询项目环境配置失败: %s", err)

					os.Exit(1)
				}

				env = pe
			}

			if err := env.UsePackage(tokens[0], tokens[1]); nil != err {
				_ = echo.Error("设定环境工具包失败: %s", err)

				os.Exit(1)
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "设置为全局环境")

	return command
}

package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDevelopEnvironmentUtilEnvironmentSetCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "set [options] K=V [K=V [...]]",
		Short: "设定环境变量",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			for _, variable := range args {
				var tokens = strings.Split(variable, "=")
				if 2 != len(tokens) {
					_ = echo.Error("环境变量[%s]无效", variable)

					os.Exit(1)
				}

				env.Environments[tokens[0]] = tokens[1]
			}

			if err := env.Save(); nil != err {
				_ = echo.Error("设定环境变量失败: %s", err)

				os.Exit(1)
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "设置为全局环境")

	return command
}

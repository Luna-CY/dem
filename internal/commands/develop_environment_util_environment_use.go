package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"strings"
)

func NewDevelopEnvironmentUtilEnvironmentUseCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "use [options] package",
		Short: "设定环境工具包",
		RunE: func(cmd *cobra.Command, args []string) error {
			var tokens = strings.Split(args[0], "@")
			if 2 != len(tokens) {
				fmt.Println("工具包名称无效")

				return nil
			}

			var env *environment.Environment

			if system {
				se, err := environment.GetSystemEnvironment()
				if nil != err {
					fmt.Printf("查询全局环境配置失败: %s\n", err)

					return nil
				}

				env = se
			} else {
				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					fmt.Printf("查询项目环境配置失败: %s\n", err)

					return nil
				}

				env = pe
			}

			if err := env.UsePackage(tokens[0], tokens[1]); nil != err {
				fmt.Printf("设定环境工具包失败: %s\n", err)

				return nil
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "设置为全局环境")

	return command
}

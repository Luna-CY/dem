package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentUtilEnvironmentUnsetCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "unset [options] K [K [...]]",
		Short: "移除环境变量",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			for _, name := range args {
				env.Environments[name] = environment.ValueNotSet
			}

			if err := env.Save(); nil != err {
				fmt.Printf("移除环境变量失败: %s\n", err)

				return nil
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "移除全局环境")

	return command
}

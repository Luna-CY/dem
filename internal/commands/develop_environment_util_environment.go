package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentUtilEnvironmentCommand() *cobra.Command {
	var showIndexEnvironments bool

	var command = &cobra.Command{
		Use:   "env [subcommand] [options]",
		Short: "查看环境信息",
		RunE: func(cmd *cobra.Command, args []string) error {
			me, err := environment.GetMixedEnvironment()
			if nil != err {
				fmt.Printf("查询环境配置失败: %s\n", err)

				return nil
			}

			// 覆盖变量: 项目、全局、索引
			var environments = make(map[string]string)
			if showIndexEnvironments {
				for pkg, version := range me.Packages {
					ind, err := index.Lookup(pkg + "@" + version)
					if nil != err {
						fmt.Printf("查询工具包[%s@%s]索引失败: %s\n", pkg, version, err)

						return nil
					}

					for k, v := range ind.Platforms[system.GetSystemArch()].Environments {
						environments[k] = system.ReplaceVariables(v, system.GetPackageRootPath(ind.PackageName))
					}
				}
			}

			for k, v := range me.Environments {
				environments[k] = v
			}

			fmt.Printf("启用的工具包\n")
			for pkg, version := range me.Packages {
				fmt.Printf("%s@%s\t", pkg, version)
			}

			fmt.Print("\n\n")

			fmt.Printf("设置的环境变量表\n")
			for k, v := range environments {
				if environment.ValueNotSet == v {
					continue
				}

				fmt.Printf("%s=%s\n", k, v)
			}

			fmt.Print("\n")

			return nil
		},
	}

	command.AddCommand(NewDevelopEnvironmentUtilEnvironmentUseCommand(), NewDevelopEnvironmentUtilEnvironmentSetCommand(), NewDevelopEnvironmentUtilEnvironmentUnsetCommand())
	command.Flags().BoolVar(&showIndexEnvironments, "show-index-envs", false, "显示索引中配置的环境变量")

	return command
}

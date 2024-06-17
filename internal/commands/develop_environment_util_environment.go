package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
						environments[k] = system.ReplaceVariables(v, "{ROOT}", system.GetPackageRootPath(ind.PackageName))
					}
				}
			}

			for k, v := range me.Environments {
				environments[k] = v
			}

			pkgs, err := os.ReadDir(system.GetPkgPath())
			if nil != err {
				fmt.Printf("查询系统工具包信息失败: %s\n", err)

				return nil
			}

			fmt.Println("已安装的工具包")

			for _, pkg := range pkgs {
				fi, err := os.Stat(filepath.Join(system.GetPkgPath(), pkg.Name(), ".installed"))
				if nil != err {
					continue
				}

				if fi.IsDir() {
					continue
				}

				fmt.Printf("%s\t", pkg.Name())
			}

			fmt.Print("\n\n")

			fmt.Println("当前项目或全局启用的工具包")
			for pkg, version := range me.Packages {
				fmt.Printf("%s@%s\t", pkg, version)
			}

			fmt.Print("\n\n")

			fmt.Println("当前项目或全局设置的环境变量表")
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

	command.AddCommand(NewDevelopEnvironmentUtilEnvironmentSetCommand(), NewDevelopEnvironmentUtilEnvironmentUnsetCommand())
	command.Flags().BoolVar(&showIndexEnvironments, "show-index-envs", false, "显示索引中配置的环境变量")

	return command
}

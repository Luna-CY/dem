package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"strings"
)

func NewDevelopEnvironmentUtilInfoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "info pkg",
		Short: "查看工具包信息",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ind, err := index.Lookup(args[0])
			if nil != err {
				return echo.Error("查询工具包[%s]索引失败: %s", args[0], err)
			}

			installed, err := pkg.Installed(args[0])
			if nil != err {
				return echo.Error("检查工具包[%s]安装状态失败: %s", args[0], err)
			}

			var platform = ind.Platforms[system.GetSystemArch()]

			fmt.Println(ind.PackageName)
			fmt.Println(ind.HomePage)
			fmt.Print("\n")
			fmt.Println(ind.Description)
			if "" != ind.Description {
				fmt.Print("\n")
			}

			if 0 != len(platform.Depends) {
				fmt.Printf("依赖的工具包: %s\n", strings.Join(platform.Depends, ", "))
			}

			fmt.Printf("是否已安装: %t\n", installed)
			if installed {
				fmt.Printf("安装路径: %s\n", system.GetPackageRootPath(ind.PackageName))

				var paths []string
				for _, p := range platform.Paths {
					paths = append(paths, system.ReplaceVariables(p, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
				}

				if 0 != len(paths) {
					fmt.Printf("命令搜索路径: %s\n", strings.Join(paths, ":"))
				}

				if 0 != len(platform.Environments) {
					fmt.Print("\n")
					fmt.Printf("预配置的环境变量: \n")
					for k, v := range platform.Environments {
						fmt.Printf("%s=%s\n", k, system.ReplaceVariables(v, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
					}
				}
			}

			return nil
		},
	}
}

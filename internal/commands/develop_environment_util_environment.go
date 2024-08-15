package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
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
		Short: "show current environment info",
		Run: func(cmd *cobra.Command, args []string) {
			me, err := environment.GetMixedEnvironment()
			if nil != err {
				echo.Errorln("find environment configuration failed: %s", true, err)

				os.Exit(1)
			}

			// 覆盖变量: 项目、全局、索引
			var environments = make(map[string]string)
			if showIndexEnvironments {
				for pkg, version := range me.Packages {
					ind, err := index.Lookup(pkg + "@" + version)
					if nil != err {
						echo.Errorln("find index of package[%s@%s] failed: %s", true, pkg, version, err)

						os.Exit(1)
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
				echo.Errorln("find system packages failed: %s", true, err)

				os.Exit(1)
			}

			echo.Infoln("installed packages")

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

			echo.Infoln("current project or system enabled packages")
			for pkg, version := range me.Packages {
				fmt.Printf("%s@%s\t", pkg, version)
			}

			fmt.Print("\n\n")

			echo.Infoln("current project or system enabled environments")
			for k, v := range environments {
				if environment.ValueNotSet == v {
					continue
				}

				fmt.Printf("%s=%s\n", k, v)
			}

			fmt.Print("\n")
		},
	}

	command.AddCommand(NewDevelopEnvironmentUtilEnvironmentSetCommand(), NewDevelopEnvironmentUtilEnvironmentUnsetCommand())
	command.Flags().BoolVar(&showIndexEnvironments, "show-index-envs", false, "show pre-defined environments in index")

	return command
}

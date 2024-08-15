package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDevelopEnvironmentUtilInfoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "info pkg",
		Short: "show package info",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ind, err := index.Lookup(args[0])
			if nil != err {
				echo.Errorln("find package index failed: %s", true, err)

				os.Exit(1)
			}

			installed, err := pkg.Installed(args[0])
			if nil != err {
				echo.Errorln("check package[%s] installed status failed: %s", true, args[0], err)

				os.Exit(1)
			}

			var platform = ind.Platforms[system.GetSystemArch()]

			fmt.Println(ind.PackageName)
			fmt.Println(ind.HomePage)
			fmt.Print("\n")
			fmt.Println(ind.Description)
			if "" != ind.Description {
				fmt.Print("\n")
			}

			echo.Infoln("installed: %t", installed)
			if 0 != len(platform.Depends) {
				echo.Infoln("depends: %s", strings.Join(platform.Depends, ", "))
			}

			if installed {
				echo.Infoln("install path: %s", system.GetPackageRootPath(ind.PackageName))

				var paths []string
				for _, p := range platform.Paths {
					paths = append(paths, system.ReplaceVariables(p, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
				}

				if 0 != len(paths) {
					fmt.Printf("command find path: %s\n", strings.Join(paths, ":"))
				}

				if 0 != len(platform.Environments) {
					fmt.Print("\n")
					fmt.Printf("pre-defined environment variables: \n")
					for k, v := range platform.Environments {
						fmt.Printf("%s=%s\n", k, system.ReplaceVariables(v, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
					}
				}
			}
		},
	}
}

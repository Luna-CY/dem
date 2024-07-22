package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/spf13/cobra"
	"os"
)

func NewDevelopEnvironmentUtilInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "initialize current project environment",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			pe, err := environment.GetProjectEnvironment()
			if nil != err {
				echo.Errorln("find current project environment failed: %s", true, err)

				os.Exit(1)
			}

			for name, version := range pe.Packages {
				var pkgName = name + "@" + version

				installed, err := pkg.Installed(pkgName)
				if nil != err {
					echo.Errorln("check package[%s] installed status failed: %s", true, pkgName, err)

					os.Exit(1)
				}

				if installed {
					continue
				}

				echo.Infoln("install package[%s]...", pkgName)
				if err := pkg.Install(cmd.Context(), pkgName); nil != err {
					echo.Errorln("package[%s] installed failed: %s", true, pkgName, err)

					os.Exit(1)
				}

				echo.Infoln("package[%s] installed successfully", pkgName)
			}
		},
	}
}

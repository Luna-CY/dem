package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentManagementUtilUninstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall pkg [pkg [...]]",
		Short: "remove installed packages",
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range args {
				installed, err := pkg.Installed(name)
				if nil != err {
					echo.Errorln("check package[%s] installed status failed: %s", true, name, err)

					continue
				}

				if !installed {
					echo.Infoln("package[%s] is not installed, skip", name)

					continue
				}

				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					echo.Errorln("get project environment failed: %s", true, err)

					continue
				}

				if _, ok := pe.Packages[name]; ok {
					delete(pe.Packages, name)

					if err := pe.Save(); nil != err {
						echo.Errorln("clean package[%s] environment configuration failed: %s", true, name, err)

						continue
					}
				}

				se, err := environment.GetSystemEnvironment()
				if nil != err {
					echo.Errorln("get system environment failed: %s", true, err)

					continue
				}

				if _, ok := se.Packages[name]; ok {
					delete(se.Packages, name)

					if err := se.Save(); nil != err {
						echo.Errorln("clean package[%s] environment configuration failed: %s", true, name, err)

						continue
					}
				}

				echo.Infoln("package[%s] removing...", name)
				if err := pkg.Uninstall(cmd.Context(), name); nil != err {
					echo.Errorln(err.Error(), true)

					continue
				}

				echo.Infoln("package[%s] removed", name)
			}
		},
	}
}

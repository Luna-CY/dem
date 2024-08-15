package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	system2 "github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewDevelopEnvironmentUtilUseCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "use [options] package",
		Short: "use a package",
		Run: func(cmd *cobra.Command, args []string) {
			var tokens = strings.Split(args[0], "@")
			if 2 != len(tokens) {
				echo.Errorln("invalid package name", false)

				os.Exit(1)
			}

			fi, err := os.Stat(filepath.Join(system2.GetPackageRootPath(args[0]), ".installed"))
			if nil != err {
				if os.IsNotExist(err) {
					echo.Errorln("package[%s] is not installed", false, args[0])

					os.Exit(1)
				}

				echo.Errorln("check package[%s] installed status failed: %s", false, args[0], err)

				os.Exit(1)
			}

			if fi.IsDir() {
				echo.Errorln("invalid package[%s], reinstall it", false, args[0])

				os.Exit(1)
			}

			var env *environment.Environment

			if system {
				se, err := environment.GetSystemEnvironment()
				if nil != err {
					echo.Errorln("get system environment failed: %s", true, err)

					os.Exit(1)
				}

				env = se
			} else {
				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					echo.Errorln("get project environment failed: %s", true, err)

					os.Exit(1)
				}

				env = pe
			}

			if err := env.UsePackage(tokens[0], tokens[1]); nil != err {
				echo.Errorln("use package failed: %s", true, err)

				os.Exit(1)
			}
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "set to system environment")

	return command
}

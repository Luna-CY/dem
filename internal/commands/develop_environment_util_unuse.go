package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDevelopEnvironmentUtilUnuseCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "unuse package",
		Short: "un-use a package",
		Run: func(cmd *cobra.Command, args []string) {
			var tokens = strings.Split(args[0], "@")
			if 2 != len(tokens) {
				echo.Errorln("invalid package name", false)

				os.Exit(1)
			}

			pe, err := environment.GetProjectEnvironment()
			if nil != err {
				echo.Errorln("get project environment failed: %s", false, err)

				os.Exit(1)
			}

			version, ok := pe.Packages[tokens[0]]
			if !ok || version != tokens[1] {
				return
			}

			delete(pe.Packages, tokens[0])

			if err := pe.Save(); nil != err {
				echo.Errorln("save project environment failed: %s", false, err)

				os.Exit(1)
			}
		},
	}

	return command
}

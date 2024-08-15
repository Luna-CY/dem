package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDevelopEnvironmentUtilEnvironmentSetCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "set [options] K=V [K=V [...]]",
		Short: "set environment variables",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var env *environment.Environment

			if system {
				se, err := environment.GetSystemEnvironment()
				if nil != err {
					echo.Errorln("find global environment configuration failed: %s", true, err)

					os.Exit(1)
				}

				env = se
			} else {
				pe, err := environment.GetProjectEnvironment()
				if nil != err {
					echo.Errorln("find project environment configuration failed: %s", true, err)

					os.Exit(1)
				}

				env = pe
			}

			for _, variable := range args {
				var tokens = strings.Split(variable, "=")
				if 2 != len(tokens) {
					echo.Errorln("invalid environment variable: %s", false, variable)

					os.Exit(1)
				}

				env.Environments[tokens[0]] = tokens[1]
			}

			if err := env.Save(); nil != err {
				echo.Errorln("set environment variables failed: %s", true, err)

				os.Exit(1)
			}
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "set system environment variables")

	return command
}

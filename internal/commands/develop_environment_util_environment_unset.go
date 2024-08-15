package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"os"
)

func NewDevelopEnvironmentUtilEnvironmentUnsetCommand() *cobra.Command {
	var system bool

	var command = &cobra.Command{
		Use:   "unset [options] K [K [...]]",
		Short: "remove environment variables",
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

			for _, name := range args {
				env.Environments[name] = environment.ValueNotSet
			}

			if err := env.Save(); nil != err {
				echo.Errorln("remove environment variables failed: %s", true, err)

				os.Exit(1)
			}
		},
	}

	command.Flags().BoolVarP(&system, "system", "s", false, "remove environment variables from system")

	return command
}

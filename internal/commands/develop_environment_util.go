package commands

import "github.com/spf13/cobra"

func NewDevelopEnvironmentUtilCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "deu subcommand [options] [args]",
		Short: "develop environment management utils",
	}

	command.AddCommand(
		NewDevelopEnvironmentUtilUpdateCommand(),
		NewDevelopEnvironmentUtilSearchCommand(),
		NewDevelopEnvironmentUtilInstallCommand(),
		NewDevelopEnvironmentUtilEnvironmentCommand(),
		NewDevelopEnvironmentUtilUseCommand(),
		NewDevelopEnvironmentUtilUnuseCommand(),
		NewDevelopEnvironmentManagementUtilUninstallCommand(),
		NewDevelopEnvironmentUtilInitCommand(),
		NewDevelopEnvironmentUtilInfoCommand(),
	)

	return command
}

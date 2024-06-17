package commands

import "github.com/spf13/cobra"

func NewDevelopEnvironmentUtilCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "deu subcommand [options] [args]",
		Short: "DEM环境管理工具",
	}

	command.AddCommand(NewDevelopEnvironmentUtilUpdateCommand(), NewDevelopEnvironmentUtilSearchCommand(), NewDevelopEnvironmentUtilInstallCommand(), NewDevelopEnvironmentUtilEnvironmentCommand(), NewDevelopEnvironmentUtilUseCommand(), NewDevelopEnvironmentUtilUnuseCommand(), NewDevelopEnvironmentManagementUtilUninstallCommand())

	return command
}

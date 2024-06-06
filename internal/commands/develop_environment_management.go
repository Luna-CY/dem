package commands

import "github.com/spf13/cobra"

func NewDevelopEnvironmentManagementCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "dem command [options] [args]",
		Short: "通过DEM运行命令，[options]和[args]将被传递给command",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

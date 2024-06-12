package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentUtilInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "install [options] lib/name",
		Short: "安装工具包",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ind, err := index.Lookup(args[0])
			if nil != err {
				return echo.Error("查找工具包失败: %s", err)
			}

			fmt.Printf("%+v\n", ind)

			return nil
		},
	}
}

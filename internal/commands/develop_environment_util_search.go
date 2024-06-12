package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/spf13/cobra"
)

func NewDevelopEnvironmentUtilSearchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "search keyword",
		Short: "搜索索引库",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			indexes, err := index.Search(args[0])
			if nil != err {
				return echo.Error("搜索索引库失败: %s", err)
			}

			for _, index := range indexes {
				fmt.Printf("%s\t", index)
			}

			fmt.Print("\n")

			return nil
		},
	}
}

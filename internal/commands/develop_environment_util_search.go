package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/spf13/cobra"
	"os"
)

func NewDevelopEnvironmentUtilSearchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "search keyword",
		Short: "搜索索引库",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			indexes, err := index.Search(args[0])
			if nil != err {
				_ = echo.Error("搜索索引库失败: %s", err)

				os.Exit(1)
			}

			for _, ind := range indexes {
				fmt.Printf("%s\t", ind)
			}

			fmt.Print("\n")

			return nil
		},
	}
}

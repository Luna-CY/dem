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
		Short: "search package by keyword",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			indexes, err := index.Search(args[0])
			if nil != err {
				echo.Errorln("search index failed: %s", true, err)

				os.Exit(1)
			}

			for _, ind := range indexes {
				fmt.Printf("%s\t", ind)
			}

			fmt.Print("\n")
		},
	}
}

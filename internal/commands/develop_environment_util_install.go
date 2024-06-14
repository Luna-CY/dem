package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/spf13/cobra"
	"strings"
)

func NewDevelopEnvironmentUtilInstallCommand() *cobra.Command {
	var overwrite bool

	var command = &cobra.Command{
		Use:   "install [options] package [package [...]]",
		Short: "安装工具包",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var pkgs []string

			for _, name := range args {
				installed, err := pkg.Installed(name)
				if nil != err {
					cmd.PrintErrln(err)

					return nil
				}

				if installed && !overwrite {
					_ = echo.Info("工具包[%s]已安装，跳过", name)

					continue
				}

				pkgs = append(pkgs, name)

			}

			fmt.Printf("需要安装的工具包: [%s]\n", strings.Join(pkgs, ","))
			for _, name := range pkgs {
				if err := pkg.Install(cmd.Context(), name); nil != err {
					cmd.PrintErrf("安装工具包[%s]失败: %s\n", name, err)

					return nil
				}
			}

			return nil
		},
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装")

	return command
}

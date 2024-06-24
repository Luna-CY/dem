package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDevelopEnvironmentUtilUnuseCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "unuse package",
		Short: "取消设定当前环境工具包",
		RunE: func(cmd *cobra.Command, args []string) error {
			var tokens = strings.Split(args[0], "@")
			if 2 != len(tokens) {
				_ = echo.Error("工具包名称无效")

				os.Exit(1)
			}

			pe, err := environment.GetProjectEnvironment()
			if nil != err {
				_ = echo.Error("查询项目环境配置失败: %s", err)

				os.Exit(1)
			}

			version, ok := pe.Packages[tokens[0]]
			if !ok || version != tokens[1] {
				return nil
			}

			delete(pe.Packages, tokens[0])

			if err := pe.Save(); nil != err {
				_ = echo.Error("取消设定当前环境工具包失败: %s", err)

				os.Exit(1)
			}

			return nil
		},
	}

	return command
}

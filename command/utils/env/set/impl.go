// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package set

import (
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var project bool

func NewSetCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "set NAME KEY=VALUE [KEY=VALUE [...]]",
		Short: "设置工具的环境变量",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var kvs = map[string]string{}
			for _, kv := range args[1:] {
				var tokens = strings.SplitN(kv, "=", 2)
				if 2 != len(tokens) {
					continue
				}

				kvs[tokens[0]] = tokens[1]
			}

			if err := environment.SetEnvironments(args[0], kvs, project); nil != err {
				echo.ErrorLN(err)

				os.Exit(1)
			}
		},
	}

	command.Flags().BoolVarP(&project, "project", "p", false, "仅当前项目")

	return command
}

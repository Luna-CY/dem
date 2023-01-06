// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package command

import (
	"context"
	"fmt"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/system"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func MainCommandExecute(ctx context.Context) error {
	main.SetUsageTemplate(MainCommandUsage)
	main.DisableFlagParsing = true

	return main.ExecuteContext(ctx)
}

var main = &cobra.Command{
	Use: "dem",
	Run: func(cmd *cobra.Command, args []string) {
		if 0 == len(args) {
			fmt.Println("未指定执行的命令")

			os.Exit(1)
		}

		var used = environment.GetUsed()

		var params []string
		if 2 <= len(args) {
			params = args[1:]
		}

		var environments = map[string]string{}
		for _, env := range os.Environ() {
			var tokens = strings.SplitN(env, "=", 2)
			if 2 != len(tokens) {
				continue
			}

			environments[tokens[0]] = tokens[1]
		}

		var paths []string
		for _, item := range used {
			for _, env := range item.Environments {
				var tokens = strings.SplitN(env, "=", 2)
				if 2 != len(tokens) {
					continue
				}

				environments[tokens[0]] = tokens[1]
			}

			for _, path := range item.Paths {
				paths = append(paths, path)
			}
		}

		environments["PATH"] = fmt.Sprintf("PATH=%s:%s", strings.Join(paths, ":"), os.Getenv("PATH"))

		var name, err = system.LockPath(args[0], paths)
		if nil != err {
			fmt.Println(err)

			os.Exit(1)
		}

		var command = exec.CommandContext(cmd.Context(), name, params...)
		command.Env = []string{}
		for k, v := range environments {
			command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
		}

		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); nil != err {
			echo.ErrorLN(err)
		}
	},
}

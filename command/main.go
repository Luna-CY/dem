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
	"github.com/Luna-CY/cobra"
	"github.com/Luna-CY/dem/core"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/mapping"
	"github.com/Luna-CY/dem/util/system"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func MainCommandExecute(ctx context.Context) error {
	main.DisableFlagParsing = true

	return main.ExecuteContext(ctx)
}

var main = &cobra.Command{
	Use: "dem",
	Run: func(cmd *cobra.Command, args []string) {
		if 0 == len(args) {
			var used = environment.GetUsed()
			if 0 == len(used) {
				fmt.Println("当前环境未配置工具")
				fmt.Println()
				fmt.Println("若要获取所有可用的工具，可使用命令 dem-utils index list 获取可用的所有工具信息")
				fmt.Println("若未安装工具，可使用命令 dem-utils install TOOL VERSION 安装需要的工具及版本")
				fmt.Println("如果已安装了工具，可使用命令 dem-utils switch-to TOOL VERSION ENV_TAG 可以切换当前环境生效的工具、版本及环境标签")
				fmt.Println("其他可用命令可通过 dem-utils --help 获取")

				return
			}

			var names = mapping.Keys(used)
			sort.Strings(names)

			fmt.Printf("Version %s\n", core.Version)
			fmt.Println()
			fmt.Println("Usage:\n  dem CMD [options] [args]")
			fmt.Println()
			fmt.Println()

			fmt.Println("当前环境所有工具及其可用的命令表:")
			for _, name := range names {
				var tool = used[name]

				var version, ok = index.GetVersion(name, tool.Version)
				if !ok {
					continue
				}

				var target = filepath.Join(core.Root, name, version.Version)
				var keywords = []string{"{ROOT}", target, "{VERSION}", version.Version}

				var commands []string
				for _, path := range version.Paths {
					path = strings.NewReplacer(keywords...).Replace(path)
					files, err := filepath.Glob(filepath.Join(path, "*"))
					if nil != err {
						echo.ErrorLN(err)

						os.Exit(1)
					}

					for _, name := range files {
						if nil == system.Executable(name) {
							commands = append(commands, filepath.Base(name))
						}
					}
				}

				fmt.Printf("\t%-20s%v\n", name, commands)
			}

			return
		}

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
		for name, used := range environment.GetUsed() {
			var version, ok = index.GetVersion(name, used.Version)
			if !ok {
				continue
			}

			var target = filepath.Join(core.Root, name, version.Version)
			var keywords = []string{"{ROOT}", target, "{VERSION}", version.Version}

			for _, path := range version.Paths {
				paths = append(paths, strings.NewReplacer(keywords...).Replace(path))
			}

			for _, environment := range version.Environments {
				var tokens = strings.SplitN(environment, "=", 2)
				if 2 != len(tokens) {
					continue
				}

				environments[tokens[0]] = strings.NewReplacer(keywords...).Replace(tokens[1])
			}

			for k, v := range environment.GetEnvironments(name, used.Version, used.Tag) {
				environments[k] = strings.NewReplacer(keywords...).Replace(v)
			}
		}

		paths = append(paths, filepath.SplitList(os.Getenv("PATH"))...)
		environments["PATH"] = strings.Join(paths, string(filepath.ListSeparator))

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

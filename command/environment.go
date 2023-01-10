// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package command

import (
	"fmt"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/mapping"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
)

var environmentCommand = &cobra.Command{
	Use:   "env",
	Short: "运行环境管理器",
	Args:  cobra.NoArgs,
}

var environmentSetCommand = &cobra.Command{
	Use:   "set",
	Short: "设置环境变量",
	Long:  "设置环境变量，前三个参数分别指定 {工具名称} {工具版本} {环境标签}，第四个及之后的所有参数为环境变量的KV对",
	Run: func(cmd *cobra.Command, args []string) {
		if 4 > len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		if version, ok := index.GetVersion(args[0], args[1]); ok {
			args[1] = version.Version
		}

		var kvs = map[string]string{}
		for _, kv := range args[3:] {
			var tokens = strings.SplitN(kv, "=", 2)
			if 2 != len(tokens) {
				continue
			}

			kvs[tokens[0]] = tokens[1]
		}

		if err := environment.SetEnvironments(args[0], args[1], args[2], kvs); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

var environmentGetCommand = &cobra.Command{
	Use:   "get",
	Short: "获取工具指定标签配置的环境变量列表",
	Run: func(cmd *cobra.Command, args []string) {
		if 3 != len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		var environments = environment.GetEnvironments(args[0], args[1], args[2])
		var keys = mapping.Keys(environments)

		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("%s=%q\n", key, environments[key])
		}
	},
}

var environmentCopyCommand = &cobra.Command{
	Use:   "copy",
	Short: "拷贝环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		if 5 != len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		if version, ok := index.GetVersion(args[0], args[1]); ok {
			args[1] = version.Version
		}

		var environments = environment.GetEnvironments(args[0], args[1], args[2])
		if 0 == len(environments) {
			echo.InfoLN("源环境标签内没有配置任何环境变量，拷贝取消")

			return
		}

		if version, ok := index.GetVersion(args[0], args[3]); ok {
			args[3] = version.Version
		}

		if err := environment.SetEnvironments(args[0], args[3], args[4], environments); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

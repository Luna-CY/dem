// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package env

import (
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var set = &cobra.Command{
	Use:   "set NAME VERSION TAG KEY=VALUE [KEY=VALUE [...]]",
	Short: "设置环境变量",
	Long:  "设置环境变量，前三个参数分别指定 {工具名称} {工具版本} {环境标签}，第四个及之后的所有参数为环境变量的KV对",
	Args:  cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		if 4 > len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
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

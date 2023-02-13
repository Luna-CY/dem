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
)

var rem = &cobra.Command{
	Use:     "rem NAME VERSION TAG",
	Aliases: []string{"remove"},
	Short:   "移除环境变量标签",
	Long:    "移除环境变量标签，该操作将删除该标签下所有的环境变量",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if 3 > len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		if project {
			if err := environment.DelProjectEnvironments(args[0], args[1], args[2]); nil != err {
				echo.ErrorLN(err)

				os.Exit(1)
			}

			return
		}

		if err := environment.DelEnvironments(args[0], args[1], args[2]); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

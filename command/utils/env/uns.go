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

var uns = &cobra.Command{
	Use:   "uns NAME VERSION TAG KEY [KEY [...]]",
	Short: "移除环境变量",
	Long:  "移除环境变量，前三个参数分别指定 {工具名称} {工具版本} {环境标签}，第四个及之后的所有参数为环境变量的KEY",
	Args:  cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		if 4 > len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		if err := environment.UnSetEnvironments(args[0], args[1], args[2], args[3:]); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

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
	"github.com/spf13/cobra"
	"os"
)

var switchToCommand = &cobra.Command{
	Use:   "switch-to",
	Short: "切换工具的版本及环境",
	Run: func(cmd *cobra.Command, args []string) {
		if 3 != len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
		}

		var _, ok = index.GetVersion(args[0], args[1])
		if !ok {
			echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

			return
		}

		if err := environment.SwitchTo(args[0], args[1], args[2]); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

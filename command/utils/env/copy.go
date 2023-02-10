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
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/spf13/cobra"
	"os"
)

var cop = &cobra.Command{
	Use:     "cop NAME SOURCE_VERSION SOURCE_TAG TARGET_VERSION TARGET_TAG",
	Aliases: []string{"copy"},
	Short:   "拷贝环境变量",
	Args:    cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		if 5 != len(args) {
			echo.ErrorLN("参数数量不足，可通过--help获取使用方法")

			return
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

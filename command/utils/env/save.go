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

var sav = &cobra.Command{
	Use:     "sav NAME VERSION [TAG:-]",
	Aliases: []string{"save"},
	Short:   "将全局环境变量保存到当前项目",
	Args:    cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		if 2 == len(args) {
			args = append(args, "-")
		}

		var environments = environment.GetEnvironments(args[0], args[1], args[2])
		if err := environment.SetProjectEnvironments(args[0], args[1], args[2], environments); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

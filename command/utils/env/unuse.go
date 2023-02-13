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

var unu = &cobra.Command{
	Use:     "unu NAME",
	Aliases: []string{"unuse"},
	Short:   "移除工具的版本选择，使其为未设置状态",
	Args:    cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		if 2 == len(args) {
			args = append(args, "-")
		}

		if project {
			if err := environment.UnsetToProject(args[0]); nil != err {
				echo.ErrorLN(err)

				os.Exit(1)
			}

			return
		}

		if err := environment.UnsetTo(args[0]); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	},
}

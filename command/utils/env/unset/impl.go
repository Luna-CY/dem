// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package unset

import (
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/spf13/cobra"
	"os"
)

func New() *cobra.Command {
	var command = &cobra.Command{
		Use:   "unset NAME KEY [KEY [...]]",
		Short: "移除环境变量",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := environment.UnsetEnvironments(args[0], args[1:], false); nil != err {
				echo.ErrorLN(err)

				os.Exit(1)
			}
		},
	}

	return command
}

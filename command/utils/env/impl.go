// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package env

import "github.com/spf13/cobra"

var project bool

func NewEnvCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "env",
		Short: "运行环境管理器",
		Args:  cobra.NoArgs,
	}

	use.Flags().BoolVarP(&project, "project", "p", false, "仅当前项目（当前目录）")
	unu.Flags().BoolVarP(&project, "project", "p", false, "仅当前项目（当前目录）")
	command.AddCommand(get, set, cop, inf, use, unu, uns, rem)

	return command
}

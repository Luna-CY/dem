// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package use

import (
	"github.com/spf13/cobra"
)

var project bool

func NewUseCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "use NAME VERSION",
		Short: "切换工具的版本及环境",
		Args:  cobra.ExactArgs(2),
		Run:   run,
	}

	command.Flags().BoolVarP(&project, "project", "p", false, "仅当前项目")

	return command
}

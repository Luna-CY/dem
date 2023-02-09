// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package index

import (
	"github.com/spf13/cobra"
)

func NewIndexCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "ind",
		Aliases: []string{"index"},
		Short:   "索引管理器",
		Args:    cobra.NoArgs,
	}

	command.AddCommand(list, update)

	return command
}

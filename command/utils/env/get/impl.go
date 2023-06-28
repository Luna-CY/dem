// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package get

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/util/mapping"
	"github.com/spf13/cobra"
	"sort"
)

func NewGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get NAME",
		Short: "获取工具的环境变量列表",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var environments = environment.GetEnvironments(args[0])
			var keys = mapping.Keys(environments)
			sort.Strings(keys)

			for _, key := range keys {
				fmt.Printf("%s=%q\n", key, environments[key])
			}
		},
	}
}

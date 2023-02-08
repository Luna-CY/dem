// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package env

import (
	"fmt"
	"github.com/Luna-CY/cobra"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/util/mapping"
	"sort"
)

var copyCommand = &cobra.Command{
	Use:       "get",
	Short:     "获取工具指定标签配置的环境变量列表",
	Args:      cobra.RangeArgs(2, 3),
	ValidArgs: []string{"NAME", "VERSION", "TAG:-"},
	Run: func(cmd *cobra.Command, args []string) {
		if 2 == len(args) {
			args = append(args, "-")
		}

		var environments = environment.GetEnvironments(args[0], args[1], args[2])
		var keys = mapping.Keys(environments)

		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("%s=%q\n", key, environments[key])
		}
	},
}

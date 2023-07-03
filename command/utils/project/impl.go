// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package project

import (
	"fmt"
	"github.com/Luna-CY/dem/command/utils/env/get"
	"github.com/Luna-CY/dem/command/utils/project/initproject"
	"github.com/Luna-CY/dem/command/utils/project/set"
	"github.com/Luna-CY/dem/command/utils/project/unset"
	"github.com/Luna-CY/dem/command/utils/project/unuse"
	"github.com/Luna-CY/dem/command/utils/project/use"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/util/mapping"
	"github.com/spf13/cobra"
	"sort"
)

func New() *cobra.Command {
	var command = &cobra.Command{
		Use:   "project",
		Short: "项目管理命令",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("当前环境使用的工具及版本信息:")

			var software = environment.GetSoftware(false)
			var names = mapping.Keys(software)
			sort.Strings(names)

			for _, name := range names {
				var version = software[name]

				var v, _ = index.GetSoftwareVersion(name, version)
				fmt.Printf("\t名称: %-30s 版本: %-30s\n", name, v.Version)
			}
		},
	}

	command.AddCommand(initproject.New(), get.New(), set.New(), unset.New(), unuse.New(), use.New())

	return command
}

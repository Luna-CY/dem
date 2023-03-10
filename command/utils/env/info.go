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
	"github.com/Luna-CY/dem/core"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/mapping"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sort"
)

var inf = &cobra.Command{
	Use:     "inf",
	Aliases: []string{"info"},
	Short:   "查看当前环境配置的所有工具",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("全局环境使用的工具及版本信息:")

		var used = environment.GetGlobalUsed()
		var names = mapping.Keys(used)
		sort.Strings(names)

		for _, name := range names {
			var info = used[name]

			var v, _ = index.GetVersion(name, info.Version)
			var showVersion = v.Version + fmt.Sprintf("[%v]", info.Version)

			fmt.Printf("\t名称: %-30s 版本[别名]: %-30s 环境标签: %s\n", name, showVersion, info.Tag)
		}

		if 0 != len(environment.GetProjectUsed()) {
			fmt.Println("当前项目环境使用的工具及版本信息:")

			var used = environment.GetProjectUsed()
			var names = mapping.Keys(used)
			sort.Strings(names)

			for _, name := range names {
				var info = used[name]

				var v, _ = index.GetVersion(name, info.Version)
				var showVersion = v.Version + fmt.Sprintf("[%v]", info.Version)

				fmt.Printf("\t名称: %-30s 版本[别名]: %-30s 环境标签: %s\n", name, showVersion, info.Tag)
			}
		}

		fmt.Println("已安装的工具及版本信息:")

		var tools = index.GetVersions()
		names = mapping.Keys(tools)
		sort.Strings(names)

		for _, name := range names {
			var versions = tools[name]

			for _, version := range versions {
				var fs, err = os.Stat(filepath.Join(core.Root, name, version))
				if nil != err && !os.IsNotExist(err) {
					echo.ErrorLN(err)

					continue
				}

				if nil == fs {
					continue
				}

				if fs.IsDir() {
					var v, _ = index.GetVersion(name, version)
					var showVersion = version + fmt.Sprintf("%v", v.Alias)

					fmt.Printf("\t名称: %-30s 版本[别名]: %-30s 安装路径: %s\n", name, showVersion, filepath.Join(core.Root, name, version))
				}
			}
		}
	},
}

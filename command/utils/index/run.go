// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package index

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/Luna-CY/dem/internal/util/mapping"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sort"
)

func run(_ *cobra.Command, _ []string) {
	var tools = index.GetSoftwareVersions()

	var names = mapping.Keys(tools)
	sort.Strings(names)

	for _, name := range names {
		var versions = tools[name]
		var installed = make([]string, 0)
		var available = make([]string, 0)

		for _, version := range versions {
			var v, _ = index.GetSoftwareVersion(name, version)

			var fs, err = os.Stat(filepath.Join(core.Software, name, v.Version))
			if nil != err && !os.IsNotExist(err) {
				echo.ErrorLN(err)

				continue
			}

			if nil == fs {
				available = append(available, version)

				continue
			}

			if fs.IsDir() {
				installed = append(installed, version)
			}
		}

		var showInstalled = fmt.Sprintf("%v", installed)
		var showAvailable = fmt.Sprintf("%v", available)

		fmt.Printf("名称:%-30s 已安装:%-60s 可用:%-60v\n", name, showInstalled, showAvailable)
	}
}

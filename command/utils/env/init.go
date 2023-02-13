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
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/installer"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/system"
	"github.com/spf13/cobra"
	"os"
)

var ini = &cobra.Command{
	Use:     "ini",
	Aliases: []string{"init"},
	Short:   "初始化当前项目（当前目录）的运行环境",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var st, err = os.Stat(".dem")
		if nil != err {
			if os.IsNotExist(err) {
				echo.InfoLN("当前目录下未找到dem项目配置信息，无法初始化项目环境")

				return
			}
		}

		if !st.IsDir() {
			echo.ErrorLN("当前目录下的dem项目配置信息无效，请检查项目目录是否正确")

			return
		}

		var used = environment.GetProjectUsed()
		if 0 == len(used) {
			echo.InfoLN("当前项目没有自定义环境配置，初始化结束")

			return
		}

		for name, info := range used {
			var version, ok = index.GetVersion(name, info.Version)
			if !ok {
				echo.ErrorLN(fmt.Sprintf("无效的自定义环境配置，未找到工具[%s]的[%s]版本信息，请检查项目配置是否正确", name, info.Version))

				continue
			}

			if system.Installed(name, version.Version) {
				echo.InfoLN(fmt.Sprintf("工具[%s]的[%s]版本已安装", name, info.Version))

				continue
			}

			if err := system.Install(cmd.Context(), name, version); nil != err {
				if installer.RemotePackageNotExists != err {
					echo.ErrorLN(err)
				}

				continue
			}
		}

		echo.InfoLN("项目环境初始化完成")
	},
}

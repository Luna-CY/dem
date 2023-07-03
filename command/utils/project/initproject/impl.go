// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package initproject

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/installer"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/Luna-CY/dem/internal/util/system"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初始化当前项目的运行环境",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var software = environment.GetSoftware(true)
			if 0 == len(software) {
				return
			}

			for name, v := range software {
				var version, ok = index.GetSoftwareVersion(name, v)
				if !ok {
					echo.ErrorLN(fmt.Sprintf("无效的自定义环境配置，未找到工具[%s]的[%s]版本信息，请检查项目配置是否正确", name, v))

					continue
				}

				if environment.Installed(name, version.Version) {
					if err := environment.SwitchTo(name, version.Version, true); nil != err {
						echo.ErrorLN(fmt.Sprintf("切换工具[%s]的[%s]版本失败，请重新尝试初始化", name, v))
					}

					continue
				}

				if err := system.Install(cmd.Context(), name, version); nil != err {
					if installer.RemotePackageNotExists != err {
						echo.ErrorLN(err)
					}

					continue
				}

				if err := environment.SwitchTo(name, version.Version, true); nil != err {
					echo.ErrorLN(fmt.Sprintf("切换工具[%s]的[%s]版本失败，请重新尝试初始化", name, v))
				}
			}

			echo.InfoLN("项目环境初始化完成")
		},
	}
}

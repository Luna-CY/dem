// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package command

import (
	"fmt"
	"github.com/Luna-CY/dem/dem"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/installer"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var overwrite bool

var installCommand = &cobra.Command{
	Use:   "install",
	Short: "工具安装器",
	Run: func(cmd *cobra.Command, args []string) {
		if 2 != len(args) {
			echo.ErrorLN("未指定工具名称或工具版本，可通过--help获取使用方法")

			return
		}

		var version, ok = index.GetVersion(args[0], args[1])
		if !ok {
			echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

			return
		}

		var target = filepath.Join(dem.Root, args[0], version.Version)

		// 检测是否已安装
		st, err := os.Stat(target)
		if nil != err && !os.IsNotExist(err) {
			echo.ErrorLN(err)

			return
		}

		if nil != st {
			if st.IsDir() && !overwrite {
				echo.InfoLN(fmt.Sprintf("工具[%s]的[%s]版本已存在，若要重新安装可设置--overwrite参数", args[0], args[1]))

				return
			}

			// 删除之前需要先提权，避免某些文件在只读权限下由于权限不足而失败
			if err := system.Chmod(target, 0777); nil != err {
				echo.ErrorLN(err)

				return
			}

			if err := os.RemoveAll(target); nil != err {
				echo.ErrorLN(err)

				return
			}
		}

		var isFail bool
		defer func() {
			if isFail {
				_ = os.RemoveAll(target)
			}
		}()

		if err := os.MkdirAll(target, os.ModeDir|0755); nil != err {
			isFail = true

			return
		}

		if version.Archive.Enable {
			// 通过打包的方式安装
			if err := installer.Archive(cmd.Context(), target, version); nil != err {
				if installer.RemotePackageNotExists != err || !version.Source.Enable {
					isFail = true
					if installer.RemotePackageNotExists != err {
						echo.ErrorLN(err)
					}

					return
				}
			}
		}

		if version.Source.Enable {
			// 通过源码的方式安装
			if err := installer.Source(cmd.Context(), target, version); nil != err {
				isFail = true
				echo.ErrorLN(err)

				return
			}
		}

		if !version.Archive.Enable && !version.Source.Enable {
			isFail = true
			echo.ErrorLN(fmt.Sprintf("工具[%s]的[%s]版本未配置有效的安装方式，请更新本地索引或进行反馈", args[0], args[1]))

			return
		}

		echo.InfoLN("安装完成")
	},
}

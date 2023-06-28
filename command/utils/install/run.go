// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package install

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/installer"
	echo2 "github.com/Luna-CY/dem/internal/util/echo"
	system2 "github.com/Luna-CY/dem/internal/util/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func run(cmd *cobra.Command, args []string) {
	if 2 != len(args) {
		echo2.ErrorLN("未指定工具名称或工具版本，可通过--help获取使用方法")

		return
	}

	var version, ok = index.GetSoftwareVersion(args[0], args[1])
	if !ok {
		echo2.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

		return
	}

	var target = filepath.Join(core.Software, args[0], version.Version)

	// 检测是否已安装
	st, err := os.Stat(target)
	if nil != err && !os.IsNotExist(err) {
		echo2.ErrorLN(err)

		return
	}

	if nil != st {
		if st.IsDir() && !overwrite {
			echo2.InfoLN(fmt.Sprintf("工具[%s]的[%s]版本已存在，若要重新安装可设置--overwrite参数", args[0], args[1]))

			return
		}

		// 删除之前需要先提权，避免某些文件在只读权限下由于权限不足而失败
		if err := system2.Chmod(target, 0777); nil != err {
			echo2.ErrorLN(err)

			return
		}

		if err := os.RemoveAll(target); nil != err {
			echo2.ErrorLN(err)

			return
		}
	}

	if err := system2.Install(cmd.Context(), args[0], version); nil != err {
		if installer.RemotePackageNotExists != err {
			echo2.ErrorLN(err)
		}

		return
	}

	echo2.InfoLN("安装完成")
	if !environment.IsSet(args[0]) {
		echo2.InfoLN("检测到该工具未配置运行时环境，将自动设置当前版本为运行时环境")

		if err := environment.SwitchTo(args[0], args[1], false); nil != err {
			echo2.ErrorLN(err)

			os.Exit(1)
		}
	}

	if switchTo || switchToProject {
		if err := environment.SwitchTo(args[0], args[1], switchToProject); nil != err {
			echo2.ErrorLN(err)

			os.Exit(1)
		}

		echo2.InfoLN(fmt.Sprintf("已将运行环境切换为工具[%s]的[%s]版本", args[0], args[1]))
	}
}

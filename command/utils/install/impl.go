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
	"github.com/Luna-CY/dem/core"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/installer"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var overwrite bool
var switchTo bool

func NewInstallCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "ins NAME VERSION",
		Aliases: []string{"install"},
		Short:   "安装指定的工具到本地环境",
		Args:    cobra.ExactArgs(2),
		Run:     run,
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装，设置该参数时将完全移除已安装的内容并重新安装，请谨慎使用")
	command.Flags().BoolVar(&switchTo, "switch-to", false, "安装完成后设置到运行时环境")

	return command
}

func run(cmd *cobra.Command, args []string) {
	if 2 != len(args) {
		echo.ErrorLN("未指定工具名称或工具版本，可通过--help获取使用方法")

		return
	}

	var version, ok = index.GetVersion(args[0], args[1])
	if !ok {
		echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

		return
	}

	var target = filepath.Join(core.Root, args[0], version.Version)

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

	if err := system.Install(cmd.Context(), args[0], version); nil != err {
		if installer.RemotePackageNotExists != err {
			echo.ErrorLN(err)
		}

		return
	}

	echo.InfoLN("安装完成")
	if !environment.IsSetUsedEnvironment(args[0]) {
		echo.InfoLN("检测到该工具未配置运行时环境，将自动设置当前版本为运行时环境")

		if err := environment.SwitchTo(args[0], args[1], "-"); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	}

	if switchTo {
		if err := environment.SwitchTo(args[0], args[1], "-"); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		echo.InfoLN(fmt.Sprintf("已将运行环境切换为工具[%s]的[%s]版本", args[0], args[1]))
	}
}

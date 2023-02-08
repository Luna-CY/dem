// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package switch_to

import (
	"fmt"
	"github.com/Luna-CY/cobra"
	"github.com/Luna-CY/dem/core"
	"github.com/Luna-CY/dem/environment"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"os"
	"path/filepath"
)

func NewSwitchToCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:       "switch-to",
		Aliases:   []string{"s", "st"},
		Short:     "切换工具的版本及环境",
		Args:      cobra.RangeArgs(2, 3),
		ValidArgs: []string{"NAME", "VERSION", "TAG:-"},
		Run:       run,
	}

	return command
}

func run(_ *cobra.Command, args []string) {
	if 2 == len(args) {
		args = append(args, "-")
	}

	var version, ok = index.GetVersion(args[0], args[1])
	if !ok {
		echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

		return
	}

	var _, err = os.Stat(filepath.Join(core.Root, args[0], version.Version))
	if nil != err {
		if !os.IsNotExist(err) {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		echo.InfoLN(fmt.Sprintf("当前环境未安装工具[%s]的[%s]版本", args[0], args[1]))
		echo.InfoLN(fmt.Sprintf("若要安装请使用 dem-utils install --switch-to %s %s", args[0], args[1]))

		return
	}

	if err := environment.SwitchTo(args[0], version.Version, args[2]); nil != err {
		echo.ErrorLN(err)

		os.Exit(1)
	}
}

// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package remove

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/Luna-CY/dem/internal/util/execute"
	"github.com/Luna-CY/dem/internal/util/system"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewRemoveCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "remove NAME VERSION",
		Short: "从本地移除已安装的工具",
		Args:  cobra.ExactArgs(2),
		Run:   run,
	}

	return command
}

func run(cmd *cobra.Command, args []string) {
	var version, ok = index.GetSoftwareVersion(args[0], args[1])
	if !ok {
		echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", args[0], args[1]))

		return
	}

	if !environment.Installed(args[0], args[1]) {
		return
	}

	var target = filepath.Join(core.Software, args[0], version.Version)
	var keywords = []string{"{VERSION}", version.Version, "{ROOT}", target}

	// 执行删除前的脚本
	if 0 != len(version.Archive.Script.Remove.Before) {
		echo.InfoLN("执行删除前脚本...")
		for _, command := range version.Archive.Script.Remove.Before {
			if err := execute.RunCommand(cmd.Context(), target, strings.NewReplacer(keywords...).Replace(command)); nil != err {
				echo.ErrorLN(fmt.Sprintf("执行删除前脚本失败: %s", err))
			}
		}
	}

	// 删除之前需要先提权，避免某些文件在只读权限下由于权限不足而失败
	_ = system.Chmod(target, 0777)
	if err := os.RemoveAll(target); nil != err {
		echo.ErrorLN(err)

		return
	}

	// 执行删除前的脚本
	if 0 != len(version.Archive.Script.Remove.After) {
		echo.InfoLN("执行删除后脚本...")
		for _, command := range version.Archive.Script.Remove.After {
			if err := execute.RunCommand(cmd.Context(), target, strings.NewReplacer(keywords...).Replace(command)); nil != err {
				echo.ErrorLN(fmt.Sprintf("执行删除后脚本失败: %s", err))
			}
		}
	}

	echo.InfoLN(fmt.Sprintf("工具[%s]的版本[%s]已移除", args[0], args[1]))
}

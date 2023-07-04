// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package install

import (
	"github.com/spf13/cobra"
)

var overwrite bool
var source bool
var switchTo bool
var switchToProject bool

func New() *cobra.Command {
	var command = &cobra.Command{
		Use:   "install NAME [VERSION]",
		Short: "安装指定的工具到本地环境，版本号可选，未指定版本时自动安装最新版本",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装，设置该参数时将完全移除已安装的内容并重新安装，请谨慎使用")
	command.Flags().BoolVar(&source, "source", false, "从源码安装，这将跳过预构建的安装包，这仅对一些允许从源码编译的工具有效")
	command.Flags().BoolVar(&switchTo, "switch-to", false, "安装完成后设置到运行时环境")
	command.Flags().BoolVar(&switchToProject, "switch-to-project", false, "安装完成后设置到当前项目的运行时环境")

	return command
}

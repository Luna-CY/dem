// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package command

import (
	"context"
	"github.com/spf13/cobra"
)

func ToolsCommandExecute(ctx context.Context) error {
	environmentSetCommand.SetUsageTemplate(EnvironmentSetCommandUsage)
	environmentListCommand.SetUsageTemplate(EnvironmentListCommandUsage)
	environmentCopyCommand.SetUsageTemplate(EnvironmentCopyCommandUsage)
	environmentCommand.AddCommand(environmentSetCommand, environmentListCommand, environmentCopyCommand)

	indexCommand.AddCommand(indexListCommand, indexUpdateCommand)
	installCommand.SetUsageTemplate(InstallCommandUsage)
	installCommand.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装，设置该参数时将完全移除已安装的内容并重新安装，请谨慎使用")
	installCommand.Flags().BoolVar(&switchTo, "switch-to", false, "安装完成后设置到运行时环境")
	removeCommand.SetUsageTemplate(InstallCommandUsage)
	switchToCommand.SetUsageTemplate(SwitchToCommandUsage)
	tools.AddCommand(environmentCommand, indexCommand, installCommand, removeCommand, switchToCommand)

	return tools.ExecuteContext(ctx)
}

var tools = &cobra.Command{
	Use:   "dem-utils",
	Short: "环境管理工具集",
	Args:  cobra.NoArgs,
}

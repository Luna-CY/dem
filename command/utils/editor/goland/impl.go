// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package goland

import "github.com/spf13/cobra"

// NewGolandCommand 创建goland命令
func NewGolandCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "goland [VERSION]",
		Short: "设置GoLand的GO相关配置",
		Args:  cobra.RangeArgs(0, 1),
		Run:   run,
	}
}

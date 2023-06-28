// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package update

import (
	"github.com/spf13/cobra"
)

var proxy bool

func NewUpdateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "update",
		Short: "更新本地索引",
		Args:  cobra.NoArgs,
		Run:   run,
	}

	command.Flags().BoolVar(&proxy, "proxy", false, "通过[https://ghproxy.com]进行代理，常规代理请设置SHELL的HTTP_PROXY与HTTPS_PROXY环境变量")

	return command
}

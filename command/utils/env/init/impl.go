// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package init

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	echo2 "github.com/Luna-CY/dem/internal/util/echo"
	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初始化当前项目（当前目录）的运行环境",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := environment.InitProject(); nil != err {
				echo2.ErrorLN(fmt.Sprintf("保存项目配置信息失败: %s", err))

				return
			}

			echo2.InfoLN("项目环境初始化完成")
		},
	}
}

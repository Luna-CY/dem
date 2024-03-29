// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package utils

import (
	"github.com/Luna-CY/dem/command/utils/editor"
	"github.com/Luna-CY/dem/command/utils/env"
	"github.com/Luna-CY/dem/command/utils/index"
	"github.com/Luna-CY/dem/command/utils/install"
	"github.com/Luna-CY/dem/command/utils/remove"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/spf13/cobra"
)

func NewUtilsCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "dem-utils",
		Short:   "环境管理工具集",
		Args:    cobra.NoArgs,
		Version: core.Version,
	}

	command.AddCommand(index.NewIndexCommand(), env.NewEnvCommand(), install.NewInstallCommand(), remove.NewRemoveCommand(), editor.NewEditorCommand())

	return command
}

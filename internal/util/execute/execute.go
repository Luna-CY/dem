// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package execute

import (
	"context"
	"github.com/Luna-CY/dem/internal/util/echo"
	"os"
	"os/exec"
)

// RunCommand 执行命令
func RunCommand(ctx context.Context, working string, command string) error {
	var commands = []string{"/bin/bash", "-c", command}

	echo.InfoLN(commands)
	var c = exec.CommandContext(ctx, commands[0], commands[1:]...)

	c.Dir = working
	for _, env := range os.Environ() {
		c.Env = append(c.Env, env)
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

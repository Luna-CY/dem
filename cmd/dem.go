// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"context"
	"github.com/Luna-CY/dem/command"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/util/echo"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	var initializer = []func() error{environment.Load, index.Load}
	for _, f := range initializer {
		if err := f(); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}
	}
}

func main() {
	var ctx, cancel = context.WithCancel(context.Background())

	go func() {
		var ch = make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	if err := command.MainCommandExecute(ctx); nil != err {
		echo.ErrorLN(err)
	}
}

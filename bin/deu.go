package main

import (
	"context"
	"fmt"
	"github.com/Luna-CY/dem/internal/commands"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ctx, cancel = context.WithCancel(context.Background())

	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	go func() {
		<-ch
		cancel()
	}()

	if err := commands.NewDevelopEnvironmentUtilCommand().ExecuteContext(ctx); nil != err {
		fmt.Println(err)
	}
}

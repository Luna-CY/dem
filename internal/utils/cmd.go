package utils

import (
	"context"
	"github.com/Luna-CY/dem/internal/echo"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// ExecuteShellCommand 执行Shell命令
func ExecuteShellCommand(ctx context.Context, command string, output io.Writer) error {
	var shell = filepath.Base(os.Getenv("SHELL"))
	if !Executable(shell) {
		shell = "/bin/sh"
	}

	echo.Infoln("run shell command: %s -c %s", shell, command)
	var cmd = exec.CommandContext(ctx, shell, "-c", command)

	if nil != output {
		cmd.Stdout = output
		cmd.Stderr = output
	}

	return cmd.Run()
}

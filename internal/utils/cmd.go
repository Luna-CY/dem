package utils

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// ExecuteShellCommand 执行Shell命令
func ExecuteShellCommand(ctx context.Context, command string, output io.Writer) error {
	var cmd = exec.CommandContext(ctx, filepath.Base(os.Getenv("SHELL")), "-c", command)

	if nil != output {
		cmd.Stdout = output
		cmd.Stderr = output
	}

	return cmd.Run()
}

package utils

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

// ExecuteShellCommand 执行Shell命令
func ExecuteShellCommand(ctx context.Context, command string) error {
	var cmd = exec.CommandContext(ctx, filepath.Base(os.Getenv("SHELL")), "-c", command)

	return cmd.Run()
}

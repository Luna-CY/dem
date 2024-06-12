package utils

import (
	"context"
	"os/exec"
)

// ExecuteShellCommand 执行Shell命令
func ExecuteShellCommand(ctx context.Context, command string) error {
	var cmd = exec.CommandContext(ctx, "sh", "-c", command)

	return cmd.Run()
}

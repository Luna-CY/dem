package commands

import (
	"context"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func NewDevelopEnvironmentUtilInstallCommand() *cobra.Command {
	var overwrite bool

	var command = &cobra.Command{
		Use:   "install [options] package [package [...]]",
		Short: "安装工具包",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, name := range args {
				if err := install(cmd.Context(), name, overwrite); nil != err {
					return err
				}
			}

			return nil
		},
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装")

	return command
}

func install(ctx context.Context, name string, overwrite bool) error {
	ind, err := index.Lookup(name)
	if nil != err {
		return echo.Error("查找工具包失败: %s", err)
	}

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return echo.Error("工具包[%s]不支持当前平台: %s", name, system.GetSystemArch())
	}

	var path = system.GetPackageRootPath(ind.PackageName)
	var installed = filepath.Join(path, ".installed")

	fi, err := os.Stat(installed)
	if nil != err && !os.IsNotExist(err) {
		return echo.Error("安装工具包[%s]失败: %s", name, err)
	}

	if nil != fi && !fi.IsDir() && !overwrite {
		return echo.Info("工具包[%s]已安装，跳过", name)
	}

	// 移除路径
	if err := utils.RemoveAll(path); nil != err {
		return echo.Error("安装工具包[%s]失败: %s", name, err)
	}

	// 重建路径
	if err := os.MkdirAll(path, 0755); nil != err {
		return echo.Error("安装工具包[%s]失败: %s", name, err)
	}

	_ = echo.Info("安装工具包[%s]...", name)

	_ = echo.Info("下载[%s]所需的资源...", name)
	for _, download := range platform.Downloads {
		if err := utils.DownloadRemoteWithProgress(ctx, download.Name, system.ReplaceVariables(download.Target, path), download.Url); nil != err {
			return err
		}
	}

	_ = echo.Info("工具包[%s]安装中...", name)
	for _, cmd := range platform.Install {
		if err := utils.ExecuteShellCommand(ctx, system.ReplaceVariables(cmd, path)); nil != err {
			return err
		}
	}

	installedFile, err := os.Create(installed)
	if nil != err {
		return echo.Error("安装工具包[%s]失败: %s", name, err)
	}

	defer func() {
		_ = installedFile.Close()
	}()

	return echo.Info("工具包[%s]安装成功", name)
}

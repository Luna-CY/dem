package pkg

import (
	"context"
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"os"
	"path/filepath"
)

// Installed 检查工具包是否已安装
func Installed(name string) (bool, error) {
	ind, err := index.Lookup(name)
	if nil != err {
		return false, fmt.Errorf("查找工具包失败: %s", err)
	}

	var path = system.GetPackageRootPath(ind.PackageName)
	var installed = filepath.Join(path, ".installed")

	fi, err := os.Stat(installed)
	if nil != err && !os.IsNotExist(err) {
		return false, fmt.Errorf("检查工具包[%s]状态失败: %s", name, err)
	}

	return nil != fi && !fi.IsDir(), nil
}

// Install 安装工具包
func Install(ctx context.Context, name string) error {
	ind, err := index.Lookup(name)
	if nil != err {
		return fmt.Errorf("查找工具包失败: %s", err)
	}

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return fmt.Errorf("工具包[%s]不支持当前平台: %s", name, system.GetSystemArch())
	}

	var path = system.GetPackageRootPath(ind.PackageName)
	var installed = filepath.Join(path, ".installed")

	// 移除路径
	if err := utils.RemoveAll(path); nil != err {
		return fmt.Errorf("安装工具包[%s]失败: %s", name, err)
	}

	// 重建路径
	if err := os.MkdirAll(path, 0755); nil != err {
		return fmt.Errorf("安装工具包[%s]失败: %s", name, err)
	}

	_ = echo.Info("下载[%s]所需的资源...", name)
	for _, download := range platform.Downloads {
		if err := utils.DownloadRemoteWithProgress(ctx, download.Name, system.ReplaceVariables(download.Target, "{ROOT}", path), download.Url); nil != err {
			return err
		}
	}

	_ = echo.Info("工具包[%s]安装中...", name)
	lf, err := os.OpenFile(filepath.Join(path, "install.log"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return fmt.Errorf("安装工具包[%s]失败: %s", name, err)
	}

	defer func() {
		_ = lf.Close()
	}()

	for _, cmd := range platform.Install {
		cmd = system.ReplaceVariables(cmd, "{ROOT}", path)

		for _, dep := range platform.Depends {
			cmd = system.ReplaceVariables(cmd, fmt.Sprintf("{PKG:%s}", dep), system.GetPackageRootPath(dep))
		}

		if _, err := lf.WriteString(cmd + "\n"); nil != err {
			_ = echo.Error("安装工具包[%s]失败: %s", name, err)

			return fmt.Errorf("如果需要更多的详细信息，请查看安装日志[%s]", filepath.Join(path, "install.log"))
		}

		if err := utils.ExecuteShellCommand(ctx, cmd, lf); nil != err {
			_ = echo.Error("安装工具包[%s]失败: %s", name, err)

			return fmt.Errorf("如果需要更多的详细信息，请查看安装日志[%s]", filepath.Join(path, "install.log"))
		}
	}

	installedFile, err := os.Create(installed)
	if nil != err {
		return fmt.Errorf("安装工具包[%s]失败: %s", name, err)
	}

	defer func() {
		_ = installedFile.Close()
	}()

	return nil
}

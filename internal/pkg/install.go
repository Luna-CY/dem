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
		return false, fmt.Errorf("find package failed: %s", err)
	}

	var path = system.GetPackageRootPath(ind.PackageName)
	var installed = filepath.Join(path, ".installed")

	fi, err := os.Stat(installed)
	if nil != err && !os.IsNotExist(err) {
		return false, fmt.Errorf("check package[%s] status failed: %s", name, err)
	}

	return nil != fi && !fi.IsDir(), nil
}

// Install 安装工具包
func Install(ctx context.Context, name string) error {
	ind, err := index.Lookup(name)
	if nil != err {
		return fmt.Errorf("find package failed: %s", err)
	}

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return fmt.Errorf("package[%s] not support current platform: %s", name, system.GetSystemArch())
	}

	var path = system.GetPackageRootPath(ind.PackageName)
	var installed = filepath.Join(path, ".installed")

	// 移除路径
	if err := utils.RemoveAll(path); nil != err {
		return fmt.Errorf("install package[%s] failed: %s", name, err)
	}

	// 重建路径
	if err := os.MkdirAll(path, 0755); nil != err {
		return fmt.Errorf("install package[%s] failed: %s", name, err)
	}

	echo.Infoln("download package[%s] resources...", name)
	for _, download := range platform.Downloads {
		if err := utils.DownloadRemoteWithProgress(ctx, download.Name, system.ReplaceVariables(download.Target, "{ROOT}", path), download.Url, download.Checksum); nil != err {
			return err
		}
	}

	echo.Infoln("package[%s] installing...", name)
	lf, err := os.OpenFile(filepath.Join(path, "install.log"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return fmt.Errorf("install package[%s] failed: %s", name, err)
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
			echo.Errorln("install package[%s] failed: %s", false, name, err)
			echo.Errorln("if you need more details, please check the install log[%s]", false, filepath.Join(path, "install.log"))

			return fmt.Errorf("install package[%s] failed: %s", name, err)
		}

		if err := utils.ExecuteShellCommand(ctx, cmd, lf); nil != err {
			echo.Errorln("install package[%s] failed: %s", false, name, err)
			echo.Errorln("if you need more details, please check the install log[%s]", false, filepath.Join(path, "install.log"))

			return fmt.Errorf("install package[%s] failed: %s", name, err)
		}
	}

	installedFile, err := os.Create(installed)
	if nil != err {
		return fmt.Errorf("install package[%s] failed: %s", name, err)
	}

	defer func() {
		_ = installedFile.Close()
	}()

	return nil
}

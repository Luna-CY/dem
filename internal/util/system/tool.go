// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package system

import (
	"context"
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/index"
	installer2 "github.com/Luna-CY/dem/internal/installer"
	"os"
	"path/filepath"
)

// Installed 检查是否已安装
func Installed(name string, version string) bool {
	var target = filepath.Join(core.Software, name, version)

	st, err := os.Stat(target)
	if nil != err && !os.IsNotExist(err) {
		return false
	}

	if nil == st {
		return false
	}

	return st.IsDir()
}

// Install 安装工具
func Install(ctx context.Context, name string, version index.Version) error {
	var target = filepath.Join(core.Software, name, version.Version)

	var isFail bool
	defer func() {
		if isFail {
			_ = os.RemoveAll(target)
		}
	}()

	if err := os.MkdirAll(target, os.ModeDir|0755); nil != err {
		isFail = true

		return err
	}

	var installed = false
	if version.Archive.Enable {
		installed = true
		// 通过打包的方式安装
		if err := installer2.Archive(ctx, target, version); nil != err {
			installed = false
			if installer2.RemotePackageNotExists != err || !version.Source.Enable {
				isFail = true

				return err
			}
		}
	}

	if version.Source.Enable && !installed {
		// 通过源码的方式安装
		if err := installer2.Source(ctx, target, version); nil != err {
			isFail = true

			return err
		}
	}

	if !version.Archive.Enable && !version.Source.Enable {
		isFail = true

		return fmt.Errorf("工具[%s]的[%s]版本未配置有效的安装方式，请更新本地索引或进行反馈", name, version.Version)
	}

	return nil
}

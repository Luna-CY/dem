// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package installer

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/downloader"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/execute"
	"strings"
)

var RemotePackageNotExists = errors.New("remote package not exists")

// Archive 打包安装
func Archive(ctx context.Context, target string, version index.Version) error {
	if !version.Archive.Enable {
		return errors.New("未启用打包安装")
	}

	var keywords = []string{"{VERSION}", version.Version, "{ROOT}", target}
	var url = strings.NewReplacer(keywords...).Replace(version.Archive.Package)

	ok, err := downloader.RemotePackageExists(url)
	if nil != err {
		return err
	}

	if !ok {
		echo.ErrorLN(fmt.Sprintf("远程包文件[%s]不存在", url))

		return RemotePackageNotExists
	}

	// 执行安装后的脚本
	if 0 != len(version.Archive.Script.Install.Before) {
		for _, command := range version.Archive.Script.Install.Before {
			if err := execute.RunCommand(ctx, target, strings.NewReplacer(keywords...).Replace(command)); nil != err {
				return err
			}
		}
	}

	if err := downloader.DownloadAndDecompress(ctx, url, target); nil != err {
		return err
	}

	// 执行安装后的脚本
	if 0 != len(version.Archive.Script.Install.After) {
		for _, command := range version.Archive.Script.Install.After {
			if err := execute.RunCommand(ctx, target, strings.NewReplacer(keywords...).Replace(command)); nil != err {
				return err
			}
		}
	}

	return nil
}

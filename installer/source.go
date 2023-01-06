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
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/downloader"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/execute"
	"os"
	"path/filepath"
	"strings"
)

// Source 源码安装
func Source(ctx context.Context, target string, version index.Version) error {
	var keywords = []string{"{VERSION}", version.Version, "{ROOT}", target, "{DEPEND}", filepath.Join(target, "depends")}
	var url = strings.NewReplacer(keywords...).Replace(version.Source.Package)

	// 源码下载到临时目录
	temp, err := os.MkdirTemp(target, "")
	if nil != err {
		return err
	}

	// 清理资源
	defer os.RemoveAll(temp)
	keywords = append(keywords, "{TEMP}", temp)

	// 安装依赖
	for _, depend := range version.Source.Depends {
		if depend.Archive.Enable {
			if err := Archive(ctx, target, depend); nil != err {
				return err
			}
		} else if depend.Source.Enable {
			if err := Source(ctx, target, depend); nil != err {
				return err
			}
		}
	}

	// 下载源码
	if err := downloader.DownloadAndDecompress(ctx, url, temp); nil != err {
		return err
	}

	// build
	echo.InfoLN("执行编译脚本，此过程可能等待时间较长，请耐心等待...")
	for _, command := range version.Source.Build.Chains {
		if err := execute.RunCommand(ctx, temp, strings.NewReplacer(keywords...).Replace(command)); nil != err {
			return err
		}
	}

	return nil
}

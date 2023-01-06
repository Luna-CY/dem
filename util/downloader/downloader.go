// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package downloader

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/pack"
	"github.com/cheggaaa/pb/v3"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// RemotePackageExists 检查远程包文件是否存在
func RemotePackageExists(url string) (bool, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		return false, err
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		return false, err
	}
	defer response.Body.Close()

	return http.StatusNotFound != response.StatusCode, nil
}

// DownloadAndDecompress 下载gzip压缩包并解压到目标位置下
func DownloadAndDecompress(ctx context.Context, url string, target string) error {
	_, filename := path.Split(url)
	var targetFilepath = filepath.Join(target, filename)

	file, err := os.Create(targetFilepath)
	if nil != err {
		return err
	}

	defer file.Close()

	echo.InfoLN("下载目标文件[", url, "] -> [", targetFilepath, "]")
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if nil != err {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	if http.StatusOK != response.StatusCode {
		return errors.New(response.Status)
	}

	var bar = pb.Full.Start64(response.ContentLength)
	if n, err := io.Copy(bar.NewProxyWriter(file), response.Body); nil != err {
		return fmt.Errorf("拷贝文件失败. 已拷贝字节: %d 错误: %v", n, err)
	}

	bar.Finish()
	if _, err := file.Seek(0, io.SeekStart); nil != err {
		return err
	}

	echo.InfoLN(fmt.Sprintf("下载完成，解压[%s]...", file.Name()))
	switch true {
	case strings.HasSuffix(url, "tar.gz"), strings.HasSuffix(url, "tgz"):
		if err := pack.UnTarGzip(file, response.ContentLength, target); nil != err {
			return err
		}
	case strings.HasSuffix(url, "tar.xz"), strings.HasSuffix(url, "txz"):
		if err := pack.UnTarXZip(file, response.ContentLength, target); nil != err {
			return err
		}
	case strings.HasSuffix(url, "zip"):
		if err := pack.UnZip(file, response.ContentLength, target); nil != err {
			return err
		}
	}

	return nil
}

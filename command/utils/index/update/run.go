// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package update

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func run(cmd *cobra.Command, _ []string) {
	echo.InfoLN("读取元数据信息")

	var manifest = fmt.Sprintf("https://raw.githubusercontent.com/Luna-CY/dem-repo/%s/index/manifest.yaml", core.Version)
	if proxy {
		manifest = core.GithubProxy + manifest
	}

	request, err := http.NewRequestWithContext(cmd.Context(), http.MethodGet, manifest, nil)
	if nil != err {
		echo.ErrorLN(err)

		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		echo.ErrorLN(err)

		os.Exit(1)
	}

	if http.StatusOK != response.StatusCode {
		echo.ErrorLN(response.Status)

		os.Exit(1)
	}

	defer response.Body.Close()

	var indexes struct {
		Index map[string]map[string][]string `yaml:"index"`
	}
	if err := yaml.NewDecoder(response.Body).Decode(&indexes); nil != err {
		echo.ErrorLN(err)

		os.Exit(1)
	}

	var arch, ok = indexes.Index[runtime.GOOS]
	if !ok {
		echo.InfoLN("未支持该系统")

		return
	}

	names, ok := arch[runtime.GOARCH]
	if !ok {
		echo.InfoLN("未支持该系统架构")

		return
	}

	for _, name := range names {
		var tokens = strings.Split(name, ".")
		var localName = strings.Join([]string{tokens[0], tokens[2]}, ".")

		var remotePath = fmt.Sprintf("https://raw.githubusercontent.com/Luna-CY/dem-repo/%s/index/%s/%s/%s", core.Version, runtime.GOOS, runtime.GOARCH, name)
		if proxy {
			remotePath = core.GithubProxy + remotePath
		}

		var localFilepath = filepath.Join(core.Index, localName)

		echo.InfoLN(fmt.Sprintf("更新索引文件[%s] -> [%s]", remotePath, localFilepath))
		request, err := http.NewRequestWithContext(cmd.Context(), http.MethodGet, remotePath, nil)
		if nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		response, err := http.DefaultClient.Do(request)
		if nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		if http.StatusNotFound == response.StatusCode {
			continue
		}

		if http.StatusOK != response.StatusCode {
			echo.ErrorLN(response.Status)

			os.Exit(1)
		}

		f, err := os.OpenFile(localFilepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		_, err = io.Copy(f, response.Body)
		if nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		_ = f.Close()
		_ = response.Body.Close()
	}

	echo.InfoLN("更新索引完成")
}

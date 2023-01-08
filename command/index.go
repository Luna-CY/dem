// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package command

import (
	"fmt"
	"github.com/Luna-CY/dem/core"
	"github.com/Luna-CY/dem/index"
	"github.com/Luna-CY/dem/util/echo"
	"github.com/Luna-CY/dem/util/mapping"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

var indexCommand = &cobra.Command{
	Use:   "index",
	Short: "索引管理器",
	Args:  cobra.NoArgs,
}

var indexListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取索引列表",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var versions = index.GetVersions()

		var names = mapping.Keys(versions)
		sort.Strings(names)
		for _, name := range names {
			fmt.Println(name, versions[name])
		}
	},
}

var indexUpdateCommand = &cobra.Command{
	Use:   "update",
	Short: "更新本地索引",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		echo.InfoLN("读取元数据信息")
		var source = fmt.Sprintf("https://raw.githubusercontent.com/Luna-CY/dem-repo/%s/index/.metadata.yaml", core.Version)

		request, err := http.NewRequestWithContext(cmd.Context(), http.MethodGet, source, nil)
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
			Version string   `yaml:"version"`
			Files   []string `yaml:"files"`
		}
		if err := yaml.NewDecoder(response.Body).Decode(&indexes); nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		u, err := url.ParseRequestURI(source)
		if nil != err {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		for _, name := range indexes.Files {
			var remoteFilename = fmt.Sprintf("%s.%s.%s.%s.yaml", name, runtime.GOOS, runtime.GOARCH, indexes.Version)
			var localFilename = fmt.Sprintf("%s.%s.%s.yaml", name, runtime.GOOS, runtime.GOARCH)
			var localFilepath = filepath.Join(core.Home, "index", localFilename)

			u.Path = filepath.Join(filepath.Dir(u.Path), remoteFilename)

			echo.InfoLN(fmt.Sprintf("更新索引文件[%s] -> [%s]", u.String(), localFilepath))
			request, err := http.NewRequestWithContext(cmd.Context(), http.MethodGet, u.String(), nil)
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
	},
}

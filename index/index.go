// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package index

import (
	"github.com/Luna-CY/dem/core"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
)

func Init() error {
	index.versions = map[string]map[string]Version{}
	index.versionStringList = map[string][]string{}

	return filepath.Walk(filepath.Join(core.Home, "index"), func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}

		if info.IsDir() || '.' == info.Name()[0] || ".yaml" != filepath.Ext(info.Name()) {
			return nil
		}

		file, err := os.Open(path)
		if nil != err {
			return err
		}
		defer file.Close()

		var i Tool
		if err := yaml.NewDecoder(file).Decode(&i); nil != err {
			return err
		}

		if _, ok := index.versionStringList[i.Name]; !ok {
			index.versionStringList[i.Name] = []string{}
		}

		if _, ok := index.versions[i.Name]; !ok {
			index.versions[i.Name] = map[string]Version{}
		}

		for _, version := range i.Versions {
			for _, alias := range version.Alias {
				index.versionStringList[i.Name] = append(index.versionStringList[i.Name], alias)
				index.versions[i.Name][alias] = version
			}

			index.versionStringList[i.Name] = append(index.versionStringList[i.Name], version.Version)
			index.versions[i.Name][version.Version] = version
		}

		return nil
	})
}

var index struct {
	versionStringList map[string][]string
	versions          map[string]map[string]Version // NAME -> VERSION -> VERSION-STRUCT
}

func GetVersions() map[string][]string {
	return index.versionStringList
}

func GetVersion(name string, version string) (Version, bool) {
	var versions, ok = index.versions[name]
	if !ok {
		return Version{}, false
	}

	if version, ok := versions[version]; ok {
		return version, true
	}

	return Version{}, false
}

type Tool struct {
	Name     string    `yaml:"name"`
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Version      string   `yaml:"version"`      // 版本号
	Alias        []string `yaml:"alias"`        // 别名
	Paths        []string `yaml:"paths"`        // 搜索路径
	Environments []string `yaml:"environments"` // 环境变量
	Archive      struct {
		Enable  bool   `yaml:"enable"`  // 是否启用打包安装
		Package string `yaml:"package"` // 下载路径
		Script  struct {
			Install struct {
				Before []string `yaml:"before"` // 安装前执行的Shell命令列表
				After  []string `yaml:"after"`  // 安装后执行的Shell命令列表
			} `yaml:"install"` // 安装前后的钩子，source模式安装时此配置无效
			Remove struct {
				Before []string `yaml:"before"` // 移除前执行的Shell命令列表
				After  []string `yaml:"after"`  // 移除后执行的Shell命令列表
			} `yaml:"remove"` // 移除前后的钩子
		} `yaml:"script"` // 脚本配置
	} `yaml:"archive"` // 打包安装需要的配置
	Source struct {
		Enable  bool   `yaml:"enable"`  // 是否启用源码安装
		Package string `yaml:"package"` // 下载路径
		Build   struct {
			Chains []string `yaml:"chains"` // 操作链，任何一个失败都会结束
		} `yaml:"build"` // 编译配置
		Depends []Version `yaml:"depends"` // 依赖列表
	} `yaml:"source"` // 源代码编译安装时的配置，如果未配置则不允许从源代码安装
}

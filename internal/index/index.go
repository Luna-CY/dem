// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package index

import (
	"github.com/Luna-CY/dem/internal/core"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
)

// Load 加载索引数据
func Load() error {
	index.software = map[string]map[string]Version{}
	index.versions = map[string][]string{}

	if err := os.MkdirAll(core.Index, 0755); nil != err {
		return err
	}

	return filepath.Walk(core.Index, func(path string, info fs.FileInfo, err error) error {
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

		var i Software
		if err := yaml.NewDecoder(file).Decode(&i); nil != err {
			return err
		}

		if _, ok := index.versions[i.Name]; !ok {
			index.versions[i.Name] = []string{}
		}

		if _, ok := index.software[i.Name]; !ok {
			index.software[i.Name] = map[string]Version{}
		}

		for _, version := range i.Versions {
			index.versions[i.Name] = append(index.versions[i.Name], version.Version)
			index.software[i.Name][version.Version] = version
		}

		return nil
	})
}

// index 索引数据
var index struct {
	versions map[string][]string
	software map[string]map[string]Version // NAME -> VERSION -> VERSION-STRUCT
}

// GetSoftwareVersions 获取软件版本列表
func GetSoftwareVersions() map[string][]string {
	return index.versions
}

// GetSoftwareVersion 获取软件版本对象
func GetSoftwareVersion(name string, version string) (Version, bool) {
	var versions, ok = index.software[name]
	if !ok {
		return Version{}, false
	}

	if version, ok := versions[version]; ok {
		return version, true
	}

	return Version{}, false
}

// Software 软件结构定义
type Software struct {
	Name     string    `yaml:"name"`
	Versions []Version `yaml:"versions"`
}

// Version 版本结构定义
type Version struct {
	Version      string   `yaml:"version"`      // 版本号
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

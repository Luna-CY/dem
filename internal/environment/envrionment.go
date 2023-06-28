// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package environment

import (
	"encoding/json"
	"github.com/Luna-CY/dem/internal/core"
	"os"
	"path/filepath"
)

// Load 加载环境配置信息
func Load() error {
	{
		global.Software = map[string]string{}
		global.Environments = map[string]map[string]string{}

		var file, err = os.Open(filepath.Join(core.Root, "environment.json"))
		if nil != err && !os.IsNotExist(err) {
			return err
		}

		if nil != file {
			defer file.Close()

			if err := json.NewDecoder(file).Decode(&global); nil != err {
				return err
			}
		}
	}

	{
		project.Software = map[string]string{}
		project.Environments = map[string]map[string]string{}

		var file, err = os.Open(".dem.json")
		if nil != err && !os.IsNotExist(err) {
			return err
		}

		if nil != file {
			defer file.Close()

			if err := json.NewDecoder(file).Decode(&project); nil != err {
				return err
			}
		}
	}

	return nil
}

type Used struct {
	Version string `json:"version"`
	Tag     string `json:"tag"`
}

var global struct {
	// NAME -> VERSION
	Software map[string]string `json:"software"`
	// NAME -> ENVIRONMENT
	Environments map[string]map[string]string `json:"environments"`
}

var project struct {
	// NAME -> STRUCT
	Software map[string]string `json:"software"`
	// NAME -> ENVIRONMENT
	Environments map[string]map[string]string `json:"environments"`
}

// GetSoftware 获取当前环境使用的软件信息
func GetSoftware(p bool) map[string]string {
	var software = map[string]string{}
	if !p {
		software = global.Software
	}

	for name, version := range project.Software {
		software[name] = version
	}

	return software
}

// Installed 检查工具是否已安装
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

// IsSet 当前环境是否已启用工具的某个版本
func IsSet(name string) bool {
	var _, ok = global.Software[name]

	return ok
}

// GetEnvironments 获取工具的环境变量配置
func GetEnvironments(name string) map[string]string {
	var environments = map[string]string{}

	// 全局
	if envs, ok := global.Environments[name]; ok {
		for key, value := range envs {
			environments[key] = value
		}
	}

	// 项目
	if envs, ok := project.Environments[name]; ok {
		for key, value := range envs {
			environments[key] = value
		}
	}

	return environments
}

// SetEnvironments 设置工具的环境变量
func SetEnvironments(name string, kvs map[string]string, p bool) error {
	if !p {
		environments, ok := global.Environments[name]
		if !ok {
			environments = map[string]string{}
		}

		for k, v := range kvs {
			environments[k] = v
		}

		global.Environments[name] = environments

		return sync()
	}

	environments, ok := project.Environments[name]
	if !ok {
		environments = map[string]string{}
	}

	for k, v := range kvs {
		environments[k] = v
	}

	project.Environments[name] = environments

	return sync()
}

// UnsetEnvironments 移除环境变量
func UnsetEnvironments(name string, ks []string, p bool) error {
	if !p {
		environments, ok := global.Environments[name]
		if !ok {
			return nil
		}

		for _, k := range ks {
			delete(environments, k)
		}

		global.Environments[name] = environments

		return sync()
	}

	environments, ok := project.Environments[name]
	if !ok {
		return nil
	}

	for _, k := range ks {
		delete(environments, k)
	}

	project.Environments[name] = environments

	return sync()
}

// SwitchTo 切换工具版本
func SwitchTo(name string, version string, p bool) error {
	if !p {
		global.Software[name] = version

		return sync()
	}

	project.Software[name] = version

	return sync()
}

// Remove 移除工具，仅支持当前项目内移除
func Remove(name string) error {
	delete(project.Software, name)

	return sync()
}

// InitProject 初始化项目级配置
func InitProject() error {
	project.Software = global.Software
	project.Environments = global.Environments

	return sync()
}

// sync 同步更改到磁盘
func sync() error {
	{
		var file, err = os.OpenFile(filepath.Join(core.Root, "environment.json"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if nil != err {
			return err
		}
		defer file.Close()

		if err := json.NewEncoder(file).Encode(global); nil != err {
			return err
		}
	}

	{
		if 0 != len(project.Software) {
			var file, err = os.OpenFile(".dem.json", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
			if nil != err {
				return err
			}
			defer file.Close()

			if err := json.NewEncoder(file).Encode(project); nil != err {
				return err
			}
		} else {
			_ = os.RemoveAll(".dem.json")
		}
	}

	return nil
}

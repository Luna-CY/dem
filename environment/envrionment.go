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
	"github.com/Luna-CY/dem/core"
	"os"
	"path/filepath"
)

func Init() error {
	var path = filepath.Join(core.Home, "environment.json")

	var file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if nil != err {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if nil != err {
		return err
	}

	if 0 < stat.Size() {
		if err := json.NewDecoder(file).Decode(&global); nil != err {
			return err
		}
	}

	if st, err := os.Stat(filepath.Join(".dem", "environment.json")); nil == err && !st.IsDir() && 0 < st.Size() {
		var content, err = os.ReadFile(filepath.Join(".dem", "environment.json"))
		if nil != err {
			return err
		}

		if err := json.Unmarshal(content, &project); nil != err {
			return err
		}
	}

	return nil
}

type Used struct {
	Version string `json:"version"`
	Tag     string `json:"tag"`
}

var global struct {
	// NAME -> STRUCT
	Used map[string]Used `json:"used"`
	// NAME -> VERSION -> TAG -> KEY -> VALUE
	Environments map[string]map[string]map[string]map[string]string `json:"environments"`
}

var project struct {
	enabled bool
	// NAME -> STRUCT
	Used map[string]Used `json:"used"`
}

func GetGlobalUsed() map[string]Used {
	return global.Used
}

func GetProjectUsed() map[string]Used {
	return project.Used
}

func IsSetUsedEnvironment(name string) bool {
	var _, ok = global.Used[name]

	return ok
}

func GetEnvironments(name string, version string, tag string) map[string]string {
	if nil == global.Environments {
		return map[string]string{}
	}

	versions, ok := global.Environments[name]
	if !ok {
		return map[string]string{}
	}

	tags, ok := versions[version]
	if !ok {
		return map[string]string{}
	}

	environments, ok := tags[tag]
	if !ok {
		return map[string]string{}
	}

	return environments
}

func SetEnvironments(name string, version string, tag string, kvs map[string]string) error {
	if nil == global.Environments {
		global.Environments = map[string]map[string]map[string]map[string]string{}
	}

	if _, ok := global.Environments[name]; !ok {
		global.Environments[name] = map[string]map[string]map[string]string{}
	}

	if _, ok := global.Environments[name][version]; !ok {
		global.Environments[name][version] = map[string]map[string]string{}
	}

	if _, ok := global.Environments[name][version][tag]; !ok {
		global.Environments[name][version][tag] = map[string]string{}
	}

	for k, v := range kvs {
		global.Environments[name][version][tag][k] = v
	}

	return sync()
}

func UnSetEnvironments(name string, version string, tag string, keys []string) error {
	if nil == global.Environments {
		return nil
	}

	if _, ok := global.Environments[name]; !ok {
		return nil
	}

	if _, ok := global.Environments[name][version]; !ok {
		return nil
	}

	if _, ok := global.Environments[name][version][tag]; !ok {
		return nil
	}

	for _, key := range keys {
		delete(global.Environments[name][version][tag], key)
	}

	return sync()
}

func DelEnvironments(name string, version string, tag string) error {
	if nil == global.Environments {
		return nil
	}

	if _, ok := global.Environments[name]; !ok {
		return nil
	}

	if _, ok := global.Environments[name][version]; !ok {
		return nil
	}

	delete(global.Environments[name][version], tag)

	return sync()
}

func SwitchTo(name string, version string, tag string) error {
	if nil == global.Used {
		global.Used = map[string]Used{}
	}

	global.Used[name] = Used{Version: version, Tag: tag}

	return sync()
}

func UnsetTo(name string) error {
	if nil == global.Used {
		return nil
	}

	delete(global.Used, name)

	return sync()
}

func SwitchToProject(name string, version string, tag string) error {
	if nil == project.Used {
		project.Used = map[string]Used{}
	}

	project.enabled = true
	project.Used[name] = Used{Version: version, Tag: tag}

	return sync()
}

func UnsetToProject(name string) error {
	if nil == project.Used {
		return nil
	}

	project.enabled = true
	delete(project.Used, name)

	return sync()
}

func sync() error {
	{
		var path = filepath.Join(core.Home, "environment.json")

		var file, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if nil != err {
			return err
		}
		defer file.Close()

		if err := json.NewEncoder(file).Encode(global); nil != err {
			return err
		}
	}

	{
		if project.enabled {
			var path = filepath.Join(".dem", "environment.json")
			if err := os.MkdirAll(".dem", 0755); nil != err {
				return err
			}

			var file, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
			if nil != err {
				return err
			}
			defer file.Close()

			if err := json.NewEncoder(file).Encode(project); nil != err {
				return err
			}
		}
	}

	return nil
}

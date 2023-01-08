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
		if err := json.NewDecoder(file).Decode(&environment); nil != err {
			return err
		}
	}

	return nil
}

type Used struct {
	Version string `json:"version"`
	Tag     string `json:"tag"`
}

var environment struct {
	// NAME -> STRUCT
	Used map[string]Used `json:"used"`
	// NAME -> VERSION -> TAG -> KEY -> VALUE
	Environments map[string]map[string]map[string]map[string]string `json:"environments"`
}

func GetUsed() map[string]Used {
	return environment.Used
}

func IsSetUsedEnvironment(name string) bool {
	var _, ok = environment.Used[name]

	return ok
}

func GetEnvironments(name string, version string, tag string) map[string]string {
	if nil == environment.Environments {
		return map[string]string{}
	}

	versions, ok := environment.Environments[name]
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
	if nil == environment.Environments {
		environment.Environments = map[string]map[string]map[string]map[string]string{}
	}

	if _, ok := environment.Environments[name]; !ok {
		environment.Environments[name] = map[string]map[string]map[string]string{}
	}

	if _, ok := environment.Environments[name][version]; !ok {
		environment.Environments[name][version] = map[string]map[string]string{}
	}

	if _, ok := environment.Environments[name][version][tag]; !ok {
		environment.Environments[name][version][tag] = map[string]string{}
	}

	for k, v := range kvs {
		environment.Environments[name][version][tag][k] = v
	}

	return sync()
}

func SwitchTo(name string, version string, tag string) error {
	if nil == environment.Used {
		environment.Used = map[string]Used{}
	}

	environment.Used[name] = Used{Version: version, Tag: tag}

	return sync()
}

func sync() error {
	var path = filepath.Join(core.Home, "environment.json")

	var file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if nil != err {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(environment); nil != err {
		return err
	}

	return nil
}

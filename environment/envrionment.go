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
	"fmt"
	"github.com/Luna-CY/dem/dem"
	"os"
	"path/filepath"
)

func Init() error {
	var path = filepath.Join(dem.Home, "environment.json")

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
	Paths        []string `json:"paths"`
	Environments []string `json:"environments"`
}

var environment struct {
	// NAME -> VERSION -> TAG
	Used map[string]map[string]Used `json:"used"`
	// NAME -> VERSION -> TAG -> KEY -> VALUE
	Environments map[string]map[string]map[string]map[string]string `json:"environments"`
}

func GetUsed() []Used {
	var res []Used

	for _, version := range environment.Used {
		for _, used := range version {
			res = append(res, used)
		}
	}

	return res
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

func SwitchTo(name string, version string, tag string, paths []string, environments []string) error {
	if nil == environment.Used {
		environment.Used = map[string]map[string]Used{}
	}

	if _, ok := environment.Used[name]; !ok {
		environment.Used[name] = map[string]Used{}
	}

	var toolEnvironments = GetEnvironments(name, version, tag)
	for k, v := range toolEnvironments {
		environments = append(environments, fmt.Sprintf("%s=%s", k, v))
	}

	var used = Used{Paths: paths, Environments: environments}
	environment.Used[name][version] = used

	return sync()
}

func sync() error {
	var path = filepath.Join(dem.Home, "environment.json")

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

// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package goland

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/core"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/util/echo"
	"github.com/beevik/etree"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func run(_ *cobra.Command, args []string) {
	if 0 == len(args) {
		if v, ok := environment.GetSoftware()["golang"]; ok {
			args = append(args, v)
		}

		if 0 == len(args) {
			echo.ErrorLN("当前环境未配置有效的Golang版本")

			os.Exit(1)
		}
	}

	var version, ok = index.GetSoftwareVersion("golang", args[0])
	if !ok {
		echo.ErrorLN(fmt.Sprintf("未找到[%s]的[%s]版本，请检查安装的工具名称与版本是否正确，或更新本地索引", "golang", args[0]))

		return
	}

	var root = filepath.Join(core.Software, "golang", version.Version)
	if _, err := os.Stat(root); nil != err {
		if !os.IsNotExist(err) {
			echo.ErrorLN(err)

			os.Exit(1)
		}

		echo.InfoLN(fmt.Sprintf("当前环境未安装工具[%s]的[%s]版本", "golang", args[0]))
		echo.InfoLN(fmt.Sprintf("若要安装请使用 dem-utils install %s %s", "golang", args[0]))

		return
	}

	var workspacePath = filepath.Join(".idea", "workspace.xml")
	var workspace, err = os.ReadFile(workspacePath)
	if nil != err {
		if os.IsNotExist(err) {
			echo.InfoLN("当前目录不是有效的GoLand项目，请在GoLand项目目录内执行此命令")

			return
		}

		echo.ErrorLN(err)
	}

	var document = etree.NewDocument()
	if err := document.ReadFromBytes(workspace); nil != err {
		echo.ErrorLN(err)

		return
	}

	var project = document.SelectElement("project")
	if "4" != project.SelectAttr("version").Value {
		echo.ErrorLN("未支持的GoLand版本")

		return
	}

	// 检查并设置Go Root
	setGoRoot(project, root)

	// 检查并设置Go Libraries
	setGoLibraries(project, root)

	// 设置环境变量
	var keywords = []string{"{ROOT}", root, "{VERSION}", version.Version}

	var environments = map[string]string{}
	for _, item := range version.Environments {
		var tokens = strings.SplitN(item, "=", 2)
		if 2 != len(tokens) {
			continue
		}

		environments[tokens[0]] = strings.NewReplacer(keywords...).Replace(tokens[1])
	}

	for k, v := range environment.GetEnvironments("golang") {
		environments[k] = strings.NewReplacer(keywords...).Replace(v)
	}

	setGoEnvironments(project, environments)

	document.Indent(2)
	if err := document.WriteToFile(workspacePath); nil != err {
		echo.ErrorLN(err)

		return
	}

	echo.InfoLN("配置完成")
}

func setGoRoot(root *etree.Element, path string) {
	var component = root.FindElement("//component[@name='GOROOT']")
	if nil == component {
		component = etree.NewElement("component")
		component.CreateAttr("name", "GOROOT")
		root.AddChild(component)
	}

	component.CreateAttr("url", "file://"+filepath.Join(path, "go"))
}

func setGoLibraries(root *etree.Element, path string) {
	var component = root.FindElement("//component[@name='GoLibraries']")
	if nil == component {
		component = etree.NewElement("component")
		component.CreateAttr("name", "GoLibraries")
		root.AddChild(component)
	}

	var pathOption = component.FindElement("//component[@name='GoLibraries']/option[@name='useGoPathFromSystemEnvironment']")
	if nil == pathOption {
		pathOption = etree.NewElement("option")
		pathOption.CreateAttr("name", "useGoPathFromSystemEnvironment")
		component.AddChild(pathOption)
	}

	pathOption.CreateAttr("value", "false")

	var urlOption = component.FindElement("//component[@name='GoLibraries']/option[@name='urls']")
	if nil == urlOption {
		urlOption = etree.NewElement("option")
		urlOption.CreateAttr("name", "urls")
		component.AddChild(urlOption)
	}

	var list = urlOption.FindElement("//component[@name='GoLibraries']/option[@name='urls']/list")
	if nil == list {
		list = etree.NewElement("list")
		urlOption.AddChild(list)
	}

	for _, e := range list.ChildElements() {
		list.RemoveChild(e)
	}

	var url = etree.NewElement("option")
	url.CreateAttr("value", "file://"+filepath.Join(path, "data"))

	list.AddChild(url)
}

func setGoEnvironments(root *etree.Element, environments map[string]string) {
	var component = root.FindElement("//component[@name='VgoProject']")
	if nil == component {
		component = etree.NewElement("component")
		component.CreateAttr("name", "VgoProject")
		root.AddChild(component)
	}

	var env = component.FindElement("//component[@name='VgoProject']/environment")
	if nil == env {
		env = etree.NewElement("environment")
		component.AddChild(env)
	}

	var em = env.FindElement("//component[@name='VgoProject']/environment/map")
	if nil == em {
		em = etree.NewElement("map")
		env.AddChild(em)
	}

	for _, e := range em.ChildElements() {
		em.RemoveChild(e)
	}

	for key, value := range environments {
		var entry = etree.NewElement("entry")
		entry.CreateAttr("key", key)
		entry.CreateAttr("value", value)

		em.AddChild(entry)
	}
}

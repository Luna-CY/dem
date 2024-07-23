# Develop Environment Management

## 简介

开发环境管理器，期望解决本地多工具版本的问题

比如Go语言，默认情况下共用依赖安装目录，同时安装多版本时可能会出现依赖包无法兼容的问题

此工具将所有版本的工具隔离在各自的根目录下，通过统一的命令入口预配置环境信息来解决依赖、缓存等共用的问题

## Introduction

Development environment manager, hoping to solve the problem of multiple local tool versions

For example, Go language shares the dependency installation directory by default, and when multiple versions are installed at the same time, there may be problems with incompatible dependency packages

This tool isolates all versions of tools in their respective root directories, and pre-configures environment information through a unified command entry to solve shared problems such as dependencies and caches

## 安装脚本(Install Script)

`curl -L https://raw.githubusercontent.com/Luna-CY/dem/main/install.sh | sh`

## 手动安装

- 下载源码 `git clone https://github.com/Luna-CY/dem.git`
- 编译 `make build`
- 安装 `make install`
- 创建根目录 `sudo mkdir -p /opt/godem/`

需要依赖 `go 1.22+` ，`make install` 会将命令安装到 `/usr/local/bin` 目录下，一般情况下需要 `root` 权限

## Manual Install

- download source code `git clone https://github.com/Luna-CY/dem.git`
- build `make build`
- install `make install`
- create root directory `sudo mkdir -p /opt/godem/`

Need to rely on `go 1.22+`, `make install` will install the command to the `/usr/local/bin` directory, generally need `root` permission

## 环境配置

- 更新本地索引信息 `deu update`
- 查找工具包 `deu search 工具包名称` ，示例 `deu search golang`
- 安装工具包 `deu install 工具包名称` ，示例 `deu install golang@1.22.4`
- 设置全局工具包版本 `deu use -s 工具包名称` ，示例 `deu use -s golang@1.22.4`
- 设置项目工具包版本 `deu use 工具包名称` ，示例 `deu use golang@1.22.4`，此命令将会在当前目录下创建 `.dem.json` 文件，用于管理此项目的特定环境配置
- 设置全局环境变量 `deu set -s 变量名称=变量值` ，示例 `deu set -s GOPROXY=https://goproxy.cn,direct`
- 设置项目环境变量 `deu set 变量名称=变量值` ，示例 `deu set GOPROXY=https://goproxy.cn,direct`

## Environment Configuration

- Update local index information `deu update`
- Find tool package `deu search [package name]`, example `deu search golang`
- Install tool package `deu install [package name]@[version]`, example `deu install golang@1.22.4`
- Set system tool package and version `deu use -s [package name]@[version]`, example `deu use -s golang@1.22.4`
- Set project tool package and version `deu use [package name]@[version]`, example `deu use golang@1.22.4`, this command will create a `.dem.json` file in the current directory to manage the specific environment configuration of this project
- Set system environment variable `deu set -s [variable name]=[variable value]`, example `deu set -s GOPROXY=https://goproxy.cn,direct`
- Set project environment variable `deu set [variable name]=[variable value]`, example `deu set GOPROXY=https://goproxy.cn,direct`

## 统一入口

提供统一入口命令 `dem` ，通过此命令来代理调用其他命令，用法如下( `dem` 命令没有任何自己的命令选项)

- `dem` ，直接调用 `dem` 命令可以打印当前的查找目录列表和使用帮助
- `dem go --version` ，实际调用的命令为 `go --version` ，使用的go语言为当前项目或全局配置的特定go版本

## Unified Entry

Provide a unified entry command `dem`, use this command to proxy call other commands, usage is as follows (`dem` command has no command options of its own)

- `dem`, directly calling the `dem` command can print the current search directory list and usage help
- `dem go --version`, the actual command called is `go --version`, using the specific go version configured for the current project or globally

## 其他说明

工具的索引位置在 `/opt/godem/index/` 目录中，通过工具安装的工具包将会放置到 `/opt/godem/packages/` 目录下，移除工具包可使用命令 `deu uninstall 工具包名称` 进行操作，不推荐直接删除 `/opt/godem/packages/` 目录下的文件夹

## Other Instructions

The index location of the tool is in the `/opt/godem/index/` directory, and the tool package installed through the tool will be placed in the `/opt/godem/packages/` directory. To remove the tool package, you can use the command `deu uninstall [package name]` to operate. It is not recommended to directly delete the folders under the `/opt/godem/packages/` directory


## 卸载DEM

删除可执行命令 `/usr/local/bin/deu` 和 `/usr/local/bin/dem` 即可，删除DEM可能会影响通过DEM安装的命令，若要移除全部本地数据可使用命令 `rm -rf /opt/godem/` 进行操作

## Uninstall DEM

Just delete the executable commands `/usr/local/bin/deu` and `/usr/local/bin/dem`, deleting DEM may affect the commands installed through DEM. If you want to remove all local data, you can use the command `rm -rf /opt/godem/` to operate

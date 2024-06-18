# Develop Environment Management

## 简介

开发环境管理器，期望解决本地多工具版本的问题

比如Go语言，默认情况下共用依赖安装目录，同时安装多版本时可能会出现依赖包无法兼容的问题

此工具将所有版本的工具隔离在各自的根目录下，通过统一的命令入口预配置环境信息来解决依赖、缓存等共用的问题

## 手动安装

- 下载源码 `git clone https://github.com/Luna-CY/dem.git`
- 编译 `make build`
- 安装 `make install`
- 创建根目录 `sudo mkdir -p /opt/godem/`

需要依赖 `go 1.22+` ，`make install` 会将命令安装到 `/usr/local/bin` 目录下，一般情况下需要 `root` 权限

## 环境配置

- 查找工具包 `deu search 工具包名称` ，示例 `deu search golang`
- 安装工具包 `deu install 工具包名称` ，示例 `deu install golang@1.22.4`
- 设置全局工具包版本 `deu use -s 工具包名称` ，示例 `deu use -s golang@1.22.4`
- 设置项目工具包版本 `deu use 工具包名称` ，示例 `deu use golang@1.22.4`，此命令将会在当前目录下创建 `.dem.json` 文件，用于管理此项目的特定环境配置
- 设置全局环境变量 `deu set -s 变量名称=变量值` ，示例 `deu set -s GOPROXY=https://goproxy.cn,direct`
- 设置项目环境变量 `deu set 变量名称=变量值` ，示例 `deu set GOPROXY=https://goproxy.cn,direct`

## 统一入口

提供统一入口命令 `dem` ，通过此命令来代理调用其他命令，用法如下( `dem` 命令没有任何自己的命令选项)

- `dem` ，直接调用 `dem` 命令可以打印当前的查找目录列表和使用帮助
- `dem go --version` ，实际调用的命令为 `go --version` ，使用的go语言为当前项目或全局配置的特定go版本

## 其他说明

工具的索引位置在 `/opt/godem/index/` 目录中，通过工具安装的工具包将会放置到 `/opt/godem/packages/` 目录下，移除工具包可使用命令 `deu uninstall pkg` 进行操作，不推荐直接删除 `/opt/godem/packages/` 目录下的文件夹


## 卸载DEM

删除可执行命令 `/usr/local/bin/deu` 和 `/usr/local/bin/dem` 即可，删除DEM可能会影响通过DEM安装的命令，若要移除全部本地数据可使用命令 `rm -rf /opt/godem/` 进行操作

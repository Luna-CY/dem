# Developer Environment Manager

开发者环境管理器，该项目旨在简化开发者本地多环境管理的工作

不同时期创建的项目可能有不同版本的开发工具依赖，多版本工作环境的管理视工具的不同有不同的成本。通过此项目，可以简化这种操作成本，快速在多个版本之间进行切换

更重要的是不同版本的工具在下载所需依赖时可能会使用不同版本的依赖，若不对依赖库进行版本隔离，很容易产生版本间的不兼容现象，消除这种现象是本项目的主要任务之一

同时，该项目通过统一工具的安装目录、数据存储目录等，尽可能地使各个工具减少对本地宿主环境的侵入及副作用

## 已支持的工具

更多详细信息可通过`dem-tools index list`获取

- Golang
- Python
- NodeJS
- Protoc
- OpenJDK
- Mongodb Shell
- Mongodb Command Tools
- Chrome Driver

**通过本工具安装的Python，使用pip安装工具包时不需要设置--user参数，默认安装在`/opt/dem/tools/python/{VERSION}/lib/python{SORT-VERSION}/site-packages`目录下**

## 部分工具使用说明

### Chrome Driver

Chrome Driver通常需要获取到安装路径，可以通过命令`dem which chromedriver`来获取命令路径

## 部分工具编译备注

### python

ubuntu 22.04环境下编译python源码需要确保本地安装了如下依赖，否则安装后部分模块会不可用

- libffi7 libbz2-dev libffi-dev libncurses5-dev libgdbm-compat-dev libsqlite3-dev uuid-dev lzma-dev liblzma-dev libgdbm6 libgdbm-dev zlib1g zlib1g-dev xz-utils

## 安装

`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Luna-CY/dem/main/install.sh)"`

## 使用

DEM支持全局级别的环境配置与项目级别的环境配置

### 基础命令

- `dem-tools ind lis`: 查看所有支持的工具及版本列表
- `dem-tools ind upd`: 从远程更新本地索引
- `dem-tools env inf`: 查看当前环境信息
- `dem-tools ins {工具名称} {版本号}`: 安装工具的版本
- `dem-tools rem {工具名称} {版本号}`: 移除已安装工具的版本

### 全局环境配置

- `dem-tools env set {工具名称} {版本号} {环境标签名称} {KEY}={VALUE}`: 为工具的版本设置`dem`命令执行时的环境变量，环境标签名称用于标记一组环境变量
- `dem-tools env uns {工具名称} {版本号} {环境标签名称} {KEY}`: 在环境标签内删除指定的环境变量KEY
- `dem-tools env copy {工具名称} {源版本号} {源环境标签名称} {目标版本号} {目标环境标签名称}`: 为工具的版本设置`dem`命令执行时的环境变量，环境标签名称用于标记一组环境变量
- `dem-tools env use {工具名称} {版本号} {环境标签名称}`: 将全局DEM环境中的某个工具切换到指定的版本及环境标签

### 项目级别环境配置

- `dem-tools env sav {工具名称} {版本号} {环境标签名称}`: 将全局环境指定标签的环境变量拷贝到当前项目
- `dem-tools env set -p {工具名称} {版本号} {环境标签名称} {KEY}={VALUE}`: 为当前项目设置工具的环境变量
- `dem-tools env uns -p {工具名称} {版本号} {环境标签名称} {KEY}`: 在环境标签内删除指定的环境变量KEY
- `dem-tools env use -p {工具名称} {版本号} {环境标签名称}`: 仅切换当前项目的工具版本

### 使用DEM环境执行命令

通过DEM环境执行命令时将自动设置执行时的环境变量和命令查找的PATH列表

- `dem CMD [FLAGS] [ARGS]`: 使用DEM环境执行命令

## 关于DEM的PATH查找路径优先级说明

通过DEM环境执行命令时，查找命令的优先级为：当前项目环境配置路径>全局环境配置路径>系统环境路径

## 关于DEM的环境变量优先级说明

通过DEM环境执行命令时，环境变量的覆盖顺序为：系统环境变量<全局环境变量<当前项目环境变量

## 贡献

如果您对该项目有兴趣并想为该项目贡献您的代码，请将该项目fork到您自己的仓库，提交代码后创建一个MR请求，在此对您表示感谢

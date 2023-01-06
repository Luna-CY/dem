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

**通过本工具安装的Python，使用pip安装工具包时不需要设置--user参数，默认安装在`/opt/dem/tools/python/{VERSION}/lib/python{SORT-VERSION}/site-packages`目录下**

## 各工具编译备注

### python

ubuntu 22.04环境下编译python源码需要确保本地安装了如下依赖，否则安装后部分模块会不可用

- libffi7 libbz2-dev libffi-dev libncurses5-dev libgdbm-compat-dev libsqlite3-dev uuid-dev lzma-dev liblzma-dev libgdbm6 libgdbm-dev

## 安装

`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Luna-CY/dem/main/install.sh)"`

## 使用

- `dem-tools index list`: 查看所有支持的工具及版本列表
- `dem-tools index update`: 从远程更新本地索引
- `dem-tools install {工具名称} {版本号}`: 安装工具的版本
- `dem-tools env set {工具名称} {版本号} {环境标签名称} {KEY}={VALUE}`: 为工具的版本设置`w`命令执行时的环境变量，环境标签名称用于标记一组环境变量
- `dem-tools env switch-to {工具名称} {版本号} {环境标签名称}`: 将DEM环境中的某个工具切换到指定的版本及环境标签
- `dem CMD [FLAGS] [ARGS]`: 使用DEM环境执行命令

## 贡献

如果您对该项目有兴趣并想为该项目贡献您的代码，请将该项目fork到您自己的仓库，提交代码后创建一个MR请求，在此提前对您表示感谢

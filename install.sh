# Copyright (c) 2023 by Luna <luna@cyl-mail.com>
# dem is licensed under Mulan PSL v2.
# You can use this software according to the terms and conditions of the Mulan PSL v2.
# You may obtain a copy of Mulan PSL v2 at:
#          http://license.coscl.org.cn/MulanPSL2
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
# See the Mulan PSL v2 for more details.

OS=`uname | tr '[:upper:]' '[:lower:]'`
ARCH=`uname -m | tr '[:upper:]' '[:lower:]'`
VERSION="v0.1.0"
USER=`id -u`
GROUP=`id -g`
ROOT_PATH=/opt/dem
PACKAGE="https://github.com/Luna-CY/dem/releases/download/${VERSION}/dem_${OS}_${ARCH}_${VERSION}.tar.gz"

echo "开始下载安装包[$PACKAGE}..."
curl -sSL -o /tmp/dem.tar.gz $PACKAGE
echo "下载完成，准备系统环境，可能会需要输入密码..."
if [[ 0 == $USER ]]; then
  mkdir -p $ROOT_PATH $ROOT_PATH/bin
  else
  sudo mkdir -p $ROOT_PATH $ROOT_PATH/bin
  sudo chown -R $USER:$GROUP $ROOT_PATH
fi

tar -zxf /tmp/dem.tar.gz -C $ROOT_PATH/bin

echo "安装完成，请将下面的export命令加入到当前的SHELL环境配置文件中"
echo "    export PATH=${ROOT_PATH}/bin:\$PATH"
echo "如果您使用BASH环境，可能的环境配置文件位置为${HOME}/.bashrc"
echo "如果您使用ZSH环境，可能的环境配置文件位置为${HOME}/.zshrc"
echo "祝您工作愉快"

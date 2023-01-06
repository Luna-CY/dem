// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package core

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const Version = "v0.1.0"

// Home DEM用户主目录
var Home = ""

// Root 根目录位置，根据不同系统取不同的值
var Root = "/opt/dem"

func Init() error {
	if "darwin" != runtime.GOOS && "linux" != runtime.GOOS {
		return errors.New("未受支持的操作系统")
	}

	var home, err = os.UserHomeDir()
	if nil != err {
		return err
	}

	Home = filepath.Join(home, ".dem")
	if err := os.MkdirAll(Home, 0755); nil != err {
		return err
	}

	return nil
}

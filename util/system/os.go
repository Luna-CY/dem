// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package system

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Chmod 支持递归修改权限
func Chmod(path string, mode os.FileMode) error {
	// 递归所有路径
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}

		// 修改权限，不存在的路径忽略
		if err := os.Chmod(path, mode); nil != err && !os.IsNotExist(err) {
			return err
		}

		return nil
	})
}

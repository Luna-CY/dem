// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package core

import (
	"path/filepath"
)

const Version = "v0.2.0"
const GithubProxy = "https://ghproxy.com/"

// Root 根目录位置
var Root = filepath.Join("/opt", "dem")

// Index 索引路径
var Index = filepath.Join(Root, "index")

// Software 软件工具路径
var Software = filepath.Join(Root, "software")

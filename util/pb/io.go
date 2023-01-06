// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package pb

import (
	"github.com/cheggaaa/pb/v3"
	"io"
)

func NewProxyReaderAt(bar *pb.ProgressBar, reader io.ReaderAt) io.ReaderAt {
	bar.Set(pb.Bytes, true)

	return &ReaderAt{reader, bar}
}

type ReaderAt struct {
	io.ReaderAt
	bar *pb.ProgressBar
}

func (r *ReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = r.ReaderAt.ReadAt(p, off)
	r.bar.Add(n)

	return
}

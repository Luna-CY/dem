// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package pack

import (
	"archive/tar"
	"github.com/cheggaaa/pb/v3"
	"github.com/ulikunitz/xz"
	"io"
	"os"
	"path/filepath"
)

// UnTarXZip 解压xzip格式的tar包
func UnTarXZip(file *os.File, size int64, target string) error {
	var xr, err = xz.NewReader(file)
	if nil != err {
		return err
	}

	var bar = pb.Full.Start64(size)
	var bxz = bar.NewProxyReader(xr)

	var reader = tar.NewReader(bxz)
	for hdr, err := reader.Next(); io.EOF != err; hdr, err = reader.Next() {
		if nil != err {
			return err
		}

		if err := os.MkdirAll(filepath.Join(target, filepath.Dir(hdr.Name)), 0755); nil != err {
			return err
		}

		if hdr.FileInfo().IsDir() {
			continue
		}

		f, err := os.OpenFile(filepath.Join(target, hdr.Name), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, hdr.FileInfo().Mode())
		if nil != err {
			return err
		}

		if _, err := io.Copy(f, reader); nil != err {
			return err
		}

		_ = f.Close()
	}

	bar.Finish()

	return nil
}

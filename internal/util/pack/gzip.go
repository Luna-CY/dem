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
	"compress/gzip"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"path/filepath"
)

// UnTarGzip 解压gzip格式的tar包
func UnTarGzip(file *os.File, size int64, target string) error {
	var bar = pb.Full.Start64(size)
	var proxy = bar.NewProxyReader(file)

	var gr, err = gzip.NewReader(proxy)
	if nil != err {
		return err
	}

	defer gr.Close()

	var reader = tar.NewReader(gr)
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

		if tar.TypeLink == hdr.Typeflag || tar.TypeSymlink == hdr.Typeflag {
			if err := os.Symlink(hdr.Linkname, filepath.Join(target, hdr.Name)); nil != err {
				return err
			}

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

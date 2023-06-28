// Copyright (c) 2023 by Luna <luna@cyl-mail.com>
// dem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package pack

import (
	"archive/zip"
	pb2 "github.com/Luna-CY/dem/internal/util/pb"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"path/filepath"
)

// UnZip 解压zip文件
func UnZip(file *os.File, size int64, target string) error {
	var bar = pb.Full.Start64(size)
	var reader, err = zip.NewReader(pb2.NewProxyReaderAt(bar, file), size)
	if nil != err {
		return err
	}

	for _, hdr := range reader.File {
		if err := os.MkdirAll(filepath.Join(target, filepath.Dir(hdr.Name)), 0755); nil != err {
			return err
		}

		if hdr.FileInfo().IsDir() {
			continue
		}

		ft, err := os.OpenFile(filepath.Join(target, hdr.Name), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, hdr.FileInfo().Mode())
		if nil != err {
			return err
		}

		fs, err := hdr.Open()
		if nil != err {
			return err
		}

		if _, err := io.Copy(ft, fs); nil != err {
			return err
		}

		_ = ft.Close()
		_ = fs.Close()
	}

	bar.Finish()

	return nil
}

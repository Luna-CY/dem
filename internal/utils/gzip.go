package utils

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"path/filepath"
)

// GzipDecompressWithProgress 解压文件
func GzipDecompressWithProgress(_ context.Context, target string, filename string, file *os.File, size int64) error {
	var bar = NewProgressWithBytes(size, fmt.Sprintf("%-50s", "Decompressing "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	var pr = progressbar.NewReader(file, bar)
	defer func() {
		_ = pr.Close()
	}()

	gz, err := gzip.NewReader(&pr)
	if nil != err {
		return err
	}

	defer func() {
		_ = gz.Close()
	}()

	var pkg = tar.NewReader(gz)
	for head, err := pkg.Next(); nil == err; head, err = pkg.Next() {
		var fi = head.FileInfo()

		if fi.IsDir() {
			if err := os.MkdirAll(filepath.Join(target, head.Name), fi.Mode()); nil != err {
				return err
			}

			continue
		}

		tf, err := os.OpenFile(filepath.Join(target, head.Name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fi.Mode())
		if nil != err {
			return err
		}

		_, err = io.Copy(tf, pkg)
		if nil != err {
			_ = tf.Close()

			return err
		}

		_ = tf.Close()
	}

	return nil
}

package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadRemoteWithProgress 下载远程文件
func DownloadRemoteWithProgress(ctx context.Context, filename string, target string, url string, checksum string) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if nil != err {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if 200 != response.StatusCode {
		return fmt.Errorf("下载[%s]失败: %s", filename, response.Status)
	}

	if err := os.MkdirAll(filepath.Dir(target), 0755); nil != err {
		return err
	}

	tf, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}

	defer func() {
		_ = tf.Close()
	}()

	var bar = NewProgressWithBytes(response.ContentLength, fmt.Sprintf("%-50s", "Downloading "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	_, err = io.Copy(io.MultiWriter(tf, bar), response.Body)
	if nil != err {
		return err
	}

	// 检查文件的checksum值
	if "" != checksum {
		cs, err := Checksum(target)
		if nil != err {
			return err
		}

		if cs != checksum {
			return fmt.Errorf("远程文件[%s]的checksum值不匹配", url)
		}
	}

	return nil
}

// DownloadLocalWithProgress 下载本地并显示进度条
func DownloadLocalWithProgress(_ context.Context, filename string, target string, path string) error {
	file, err := os.Open(path)
	if nil != err {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	fileInfo, err := file.Stat()
	if nil != err {
		return err
	}

	tf, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}

	var bar = NewProgressWithBytes(fileInfo.Size(), fmt.Sprintf("%-50s", "Coping "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	if _, err = io.Copy(io.MultiWriter(tf, bar), file); nil != err {
		return err
	}

	defer func() {
		_ = tf.Close()
	}()

	return nil
}

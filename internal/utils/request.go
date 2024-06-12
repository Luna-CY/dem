package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadRemoteWithTmpFileAndProgress 下载并显示进度条
// 返回临时文件，调用方需要负责删除临时文件
func DownloadRemoteWithTmpFileAndProgress(ctx context.Context, filename string, url string) (*os.File, int64, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if nil != err {
		return nil, 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		return nil, 0, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	tf, err := os.CreateTemp("", "")
	if nil != err {
		return nil, 0, err
	}

	var bar = NewProgressWithBytes(response.ContentLength, fmt.Sprintf("%-50s", "Downloading "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	_, err = io.Copy(io.MultiWriter(tf, bar), response.Body)
	if nil != err {
		return nil, 0, err
	}

	if _, err := tf.Seek(0, io.SeekStart); nil != err {
		return nil, 0, err
	}

	return tf, response.ContentLength, nil
}

// DownloadRemoteWithProgress 下载远程文件
func DownloadRemoteWithProgress(ctx context.Context, filename string, target string, url string) (*os.File, int64, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if nil != err {
		return nil, 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if nil != err {
		return nil, 0, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	tf, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return nil, 0, err
	}

	var bar = NewProgressWithBytes(response.ContentLength, fmt.Sprintf("%-50s", "Downloading "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	_, err = io.Copy(io.MultiWriter(tf, bar), response.Body)
	if nil != err {
		return nil, 0, err
	}

	if _, err := tf.Seek(0, io.SeekStart); nil != err {
		return nil, 0, err
	}

	return tf, response.ContentLength, nil
}

// DownloadLocalWithProgress 下载并显示进度条
// 返回临时文件，调用方需要负责删除临时文件
func DownloadLocalWithProgress(_ context.Context, filename string, path string) (*os.File, int64, error) {
	file, err := os.Open(path)
	if nil != err {
		return nil, 0, err
	}

	fileInfo, err := file.Stat()
	if nil != err {
		return nil, 0, err
	}

	tf, err := os.CreateTemp("", "")
	if nil != err {
		return nil, 0, err
	}

	var bar = NewProgressWithBytes(fileInfo.Size(), fmt.Sprintf("%-50s", "Downloading "+filename))
	defer func() {
		_ = bar.Finish()
	}()

	if _, err = io.Copy(io.MultiWriter(tf, bar), file); nil != err {
		return nil, 0, err
	}

	if _, err := tf.Seek(0, io.SeekStart); nil != err {
		return nil, 0, err
	}

	return tf, fileInfo.Size(), nil
}

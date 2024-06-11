package utils

import (
	"context"
	"io"
	"net/http"
	"os"
)

// DownloadRemoteWithProgress 下载并显示进度条
// 返回临时文件，调用方需要负责删除临时文件
func DownloadRemoteWithProgress(ctx context.Context, filename string, url string) (*os.File, int64, error) {
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

	var bar = NewProgressWithBytes(response.ContentLength, "Downloading "+filename)

	_, err = io.Copy(io.MultiWriter(tf, bar), response.Body)
	if nil != err {
		return nil, 0, err
	}

	_ = bar.Finish()

	return tf, response.ContentLength, nil
}

// DownloadLocalWithProgress 下载并显示进度条
// 返回临时文件，调用方需要负责删除临时文件
func DownloadLocalWithProgress(_ context.Context, path string) (*os.File, int64, error) {
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

	_, err = io.Copy(tf, file)
	if nil != err {
		return nil, 0, err
	}

	return tf, fileInfo.Size(), nil
}

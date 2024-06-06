package utils

import (
	"context"
	"github.com/Luna-CY/dem/internal/echo"
	"io"
	"net/http"
	"os"
)

// DownloadRemoteWithProgress 下载并显示进度条
// 返回临时文件，调用方需要负责删除临时文件
func DownloadRemoteWithProgress(ctx context.Context, message string, filename string, url string) (*os.File, int64, error) {
	echo.Info(message)

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
func DownloadLocalWithProgress(_ context.Context, message string, filename string, path string) (*os.File, int64, error) {
	echo.Info(message)

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

	var bar = NewProgressWithBytes(fileInfo.Size(), "Downloading "+filename)

	_, err = io.Copy(io.MultiWriter(tf, bar), file)
	if nil != err {
		return nil, 0, err
	}

	_ = bar.Finish()

	return tf, fileInfo.Size(), nil
}

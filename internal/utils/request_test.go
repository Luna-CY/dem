package utils

import (
	"context"
	"os"
	"testing"
)

func TestDownloadWithProgress(t *testing.T) {
	tf, _, err := DownloadRemoteWithProgress(context.Background(), "开始下载测试文件", "go1.14.2.src.tar.gz", "https://dl.google.com/go/go1.14.2.src.tar.gz")
	if nil != err {
		t.Fatal(err)
	}

	defer func() {
		_ = tf.Close()
		_ = os.Remove(tf.Name())
	}()
}

package utils

import (
	"context"
	"os"
	"testing"
)

func TestGzipDecompressWithProgress(t *testing.T) {
	f, err := os.Open("../../testdata/base.tar.gz")
	if nil != err {
		t.Fatal(err)
	}

	defer func() {
		_ = f.Close()
	}()

	err = GzipDecompressWithProgress(context.Background(), "../../testdata/output/", "base.tar.gz", "../../testdata/base.tar.gz")
	if nil != err {
		t.Fatal(err)
	}
}

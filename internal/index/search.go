package index

import (
	"github.com/Luna-CY/dem/internal/system"
	"io/fs"
	"path/filepath"
	"strings"
)

// Search 搜索索引库
func Search(keyword string) ([]string, error) {
	var indexes []string

	if err := filepath.Walk(system.GetIndexPath(), func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(info.Name(), keyword) {
			indexes = append(indexes, filepath.Join(filepath.Base(filepath.Dir(filepath.Dir(path))), info.Name()))
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return indexes, nil
}

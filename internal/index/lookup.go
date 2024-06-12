package index

import (
	"errors"
	"fmt"
	"github.com/Luna-CY/dem/internal/system"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Lookup 定位索引
func Lookup(name string) (*Index, error) {
	pkg, filename := filepath.Split(name)
	if "" == pkg || "" == filename {
		return nil, errors.New("不完整的包名，完整的包名格式应为: 索引库/包")
	}

	var path = filepath.Join(system.GetIndexPath(), pkg, string(filename[0]), filename+".yaml")
	f, err := os.Open(path)
	if nil != err {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("未找到工具包: %s", name)
		}
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	var i = new(Index)
	if err := yaml.NewDecoder(f).Decode(i); nil != err {
		return nil, err
	}

	return i, nil
}

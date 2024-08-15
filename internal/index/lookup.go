package index

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/system"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Lookup 定位索引
func Lookup(name string) (*Index, error) {
	libraries, err := os.ReadDir(system.GetIndexPath())
	if nil != err {
		return nil, err
	}

	for _, library := range libraries {
		if !library.IsDir() {
			continue
		}

		var path = filepath.Join(system.GetIndexPath(), filepath.Base(library.Name()), string(name[0]), name+".yaml")
		f, err := os.Open(path)
		if nil != err {
			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		var i = new(Index)
		if err := yaml.NewDecoder(f).Decode(i); nil != err {
			return nil, err
		}

		_ = f.Close()

		return i, nil
	}

	return nil, fmt.Errorf("package [%s] not found", name)
}

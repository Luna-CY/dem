package environment

import (
	"encoding/json"
	"github.com/Luna-CY/dem/internal/system"
	"os"
)

// GetSystemEnvironment 获取系统环境变量
func GetSystemEnvironment() (*Environment, error) {
	f, err := os.Open(system.GetSystemEnvironmentPath())
	if nil != err && !os.IsNotExist(err) {
		return nil, err
	}

	if os.IsNotExist(err) {
		return NewEnvironment(system.GetSystemEnvironmentPath()), nil
	}

	defer func() {
		_ = f.Close()
	}()

	var environment = NewEnvironment(system.GetSystemEnvironmentPath())
	if err := json.NewDecoder(f).Decode(&environment); nil != err {
		return nil, err
	}

	return environment, nil
}

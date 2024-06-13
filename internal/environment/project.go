package environment

import (
	"encoding/json"
	"github.com/Luna-CY/dem/internal/system"
	"os"
)

func GetProjectEnvironment() (*Environment, error) {
	f, err := os.Open(".dem.json")
	if nil != err && !os.IsNotExist(err) {
		return nil, err
	}

	if os.IsNotExist(err) {
		return NewEnvironment(".dem.json"), nil
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

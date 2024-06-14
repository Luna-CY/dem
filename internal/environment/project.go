package environment

import (
	"encoding/json"
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

	var environment = NewEnvironment(".dem.json")
	if err := json.NewDecoder(f).Decode(&environment); nil != err {
		return nil, err
	}

	return environment, nil
}

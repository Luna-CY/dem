package environment

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const ValueNotSet = "DEMNULL" // 环境变量未设置

func NewEnvironment(filepath string) *Environment {
	return &Environment{
		Filepath:     filepath,
		Packages:     make(map[string]string),
		Environments: make(map[string]string),
	}
}

type Environment struct {
	Filepath     string            `json:"-"`            // 文件路径
	Packages     map[string]string `json:"packages"`     // 启用的包
	Environments map[string]string `json:"environments"` // 环境变量表
}

func (cls *Environment) UsePackage(name string, version string) error {
	cls.Packages[name] = version

	return cls.Save()
}

func (cls *Environment) SetEnvironment(name string, value string) error {
	cls.Environments[name] = value

	return cls.Save()
}

func (cls *Environment) Save() error {
	if err := os.MkdirAll(filepath.Dir(cls.Filepath), 0755); nil != err {
		return err
	}

	f, err := os.OpenFile(cls.Filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	if err := json.NewEncoder(f).Encode(cls); nil != err {
		return err
	}

	return nil
}

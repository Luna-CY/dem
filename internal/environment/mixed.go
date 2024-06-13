package environment

// GetMixedEnvironment 获取混合环境信息
func GetMixedEnvironment() (*Environment, error) {
	se, err := GetSystemEnvironment()
	if nil != err {
		return nil, err
	}

	pe, err := GetProjectEnvironment()
	if nil != err {
		return nil, err
	}

	for pkg, version := range pe.Packages {
		se.Packages[pkg] = version
	}

	for k, v := range pe.Environments {
		se.Environments[k] = v
	}

	return se, nil
}

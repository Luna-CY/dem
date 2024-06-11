package index

type Index struct {
	HomePage    string `yaml:"home-page"`   // 主页
	Description string `yaml:"description"` // 描述
	Version     string `yaml:"version"`     // 版本号
	Platforms   map[string]struct {
		Downloads []struct {
			Name     string `yaml:"name"`     // 包名称
			Url      string `yaml:"url"`      // 下载地址
			Checksum string `yaml:"checksum"` // 校验和
		} `yaml:"downloads"` // 需要下载的包
		Install   []string `yaml:"install"`   // 安装步骤，shell命令
		Uninstall []string `yaml:"uninstall"` // 卸载步骤，shell命令
	} `yaml:"platforms"` // 支持的平台
}

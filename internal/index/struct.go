package index

type Index struct {
	PackageName string `yaml:"package-name"` // 包名
	HomePage    string `yaml:"home-page"`    // 主页
	Description string `yaml:"description"`  // 描述
	Platforms   map[string]struct {
		Paths        []string          `yaml:"paths"`        // 查找路径
		Environments map[string]string `yaml:"environments"` // 环境变量表
		Depends      []string          `yaml:"depends"`      // 依赖包列表
		Downloads    []struct {
			Name     string `yaml:"name"`     // 包名称
			Url      string `yaml:"url"`      // 下载地址
			Target   string `yaml:"target"`   // 保存位置
			Checksum string `yaml:"checksum"` // 校验和
		} `yaml:"downloads"` // 需要下载的包
		Install   []string `yaml:"install"`   // 安装步骤，shell命令
		Uninstall []string `yaml:"uninstall"` // 卸载步骤，shell命令
	} `yaml:"platforms"` // 支持的平台
}

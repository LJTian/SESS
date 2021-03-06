package config

type RegistrationCenter struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type DBInfo struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

// ServerConfig
type ServerConfig struct {
	Name string             `yaml:"name"`
	Tags []string           `yaml:"tags"`
	DB   DBInfo             `yaml:"db"`
	RC   RegistrationCenter `yaml:"rc"`
}

// 配置文件
type Config struct {
	IP          string `yaml:"ip"`
	Port        int    `yaml:"port"`
	NamespaceId string `yaml:"namespaceId"`
	User        string `yaml:"user"`
	PassWord    string `yaml:"passWord"`
	DataId      string `yaml:"dataId"`
	Group       string `yaml:"group"`
	LocalIp     string `yaml:"localIp"`
}

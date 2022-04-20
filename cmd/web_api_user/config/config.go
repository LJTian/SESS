package config

type UserSrvConfig struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Name string `yaml:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `yaml:"key"`
}

type RegistrationCenter struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type ServerConfig struct {
	Name        string             `yaml:"name"`
	Tags        []string           `yaml:"tags"`
	RC          RegistrationCenter `yaml:"rc"`
	JWTInfo     JWTConfig          `yaml:"jwt"`
	UserSrvInfo string             `yaml:"userServer"`
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

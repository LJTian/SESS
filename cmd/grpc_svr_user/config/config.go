package config

type NacOS struct {
	IP          string `yaml:"ip"`
	Port        int    `yaml:"port"`
	NamespaceId string `yaml:"namespaceId"`
}

type DBInfo struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

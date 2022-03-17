package rabbitmq

import (
	MyYaml "SESS/pkg/yaml"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func GetYamlInfo() {
	Cfg, err := GetConfig("/Users/ljtian/data/git/github.com/LJTian/SESS/configs/MqYaml.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", Cfg)
}

// 获取配置文件内容
func GetConfig(pathStr string) (setting MyYaml.MqYaml, err error) {
	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(config, &setting)
	if err != nil {
		return MyYaml.MqYaml{}, err
	}
	return
}

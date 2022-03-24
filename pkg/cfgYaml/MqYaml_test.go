package cfgYaml

import (
	"SESS/pkg/tools"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

// 测试配置文件是否读取成功
func TestDemo(t *testing.T) {
	pathStr := "/Users/ljtian/data/git/github.com/LJTian/SESS/configs/MqYaml.yaml"
	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return
	}
	var this MqYaml
	err = yaml.Unmarshal(config, &this)
	tools.FailOnError(err, "Unmarshal err")
	this.PrintYamlInfo()
}

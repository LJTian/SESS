package cfgYaml

import "fmt"

// MqConnect 链接配置
type MqConnect struct {
	User   string `yaml:"user"`
	PassWd string `yaml:"passwd"`
	Ip     string `yaml:"ip"`
	Port   string `yaml:"port"`
}

// Producer 生产者配置
type Producer struct {
	Name    string `yaml:"name"`
	RoutKey string `yaml:"routKey"`
	Num     int    `yaml:"num"`
}

// Consumer 消费者配置
type Consumer struct {
	Name    string `yaml:"name"`
	RoutKey string `yaml:"routKey"`
	AutoAck bool   `yaml:"autoAck"`
	Num     int    `yaml:"num"`
}

// Topic 主题配置
type Topic struct {
	Name     string     `yaml:"name"`
	Producer []Producer `yaml:"producer"`
	Consumer []Consumer `yaml:"consumer"`
}

// MqYaml 整体配置
type MqYaml struct {
	Name    string      `yaml:"name"`    // 通用域-名字
	Domain  string      `yaml:"domain"`  // 通用域-域路径
	Version string      `yaml:"version"` // 通用域-版本信息
	Connect []MqConnect `yaml:"connect"`
	Topic   []Topic     `yaml:"topic"`
}

func (this *MqYaml) PrintYamlInfo() {
	fmt.Printf("Name:[%s]\t Domain:[%s]\t Version:[%s]\n", this.Name, this.Domain, this.Version)
	for _, v := range this.Connect {
		fmt.Printf("User:[%s]\t PassWd:[%s]\t Ip:[%s]\t Port:[%s]\n",
			v.User, v.PassWd, v.Ip, v.Port)
	}

	for _, v := range this.Topic {
		fmt.Printf("Topic Name:[%s]\n",
			v.Name)
		fmt.Println("生产者：-----------------------------")
		for k, v := range v.Producer {
			fmt.Printf("生产者%d: Name:[%s]\t RoutKey:[%s]\t Num:[%d]\n",
				k, v.Name, v.RoutKey, v.Num)
		}

		fmt.Println("消费者：-----------------------------")
		for k, v := range v.Consumer {
			fmt.Printf("消费者%d: Name:[%s]\t RoutKey:[%s]\t AutoAck:[%v]\t Num:[%d]\n",
				k, v.Name, v.RoutKey, v.AutoAck, v.Num)
		}
	}
}

/*
name: "mq"
domain: "server.mq.msg"
version: "0.0.1"
connect:
  - user: "guest"
    passwd: "guest"
    ip: "10.211.55.3"
    port: "5672"
topic:
  - name: "msgTran"
    producer:
    - name: "ProJob1"
      routKey: "log.cron.job"
      num: 5
    - name: "ProApi1"
      routKey: "log.trans.api"
      num: 5
    consumer:
    - name: "ConCron"
      routKey: "*.cron.*"
      autoAck: yes
      num: 5

    - name: "ConLog"
      routKey: "log#"
      autoAck: no
      num: 5

    - name: "ConTrans"
      routKey: "*.trans.*"
      autoAck: yes
      num: 5
*/

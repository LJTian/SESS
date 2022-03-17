package yaml

// MqConnect 链接配置
type MqConnect struct {
	User   string `yaml:"user"`
	PassWd string `yaml:"passWd"`
	Ip     string `yaml:"ip"`
	Port   string `yaml:"port"`
}

// Producer 生产者配置
type Producer struct {
	Name    string `yaml:"name"`
	RoutKey string `yaml:"routKey"`
}

// Consumer 消费者配置
type Consumer struct {
	Name    string `yaml:"name"`
	RoutKey string `yaml:"routKey"`
	AutoAck bool   `yaml:"autoAck"`
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

/*
name: "mq"
domain: "server.mq.msg"
version: "0.0.1"
data:
  - connect:
      user: "guest"
      passwd: "guest"
      ip: "10.211.55.3"
      port: "5672"

  - topic:
    - name: "MsgTrans"
      - producer:
          name: "ProJob1"
          routKey: "log.cron.job"

      - producer:
          name: "ProApi1"
          routKey: "log.trans.api"

      - consumer:
        - name: "ConCron"
        - routKey: "*.cron.*"
        - autoAck: ture

      - consumer:
          - name: "ConLog"
          - routKey: "log#"
          - autoAck: false

      - consumer:
          - name: "ConTrans"
          - routKey: "*.trans.*"
          - autoAck: ture
*/

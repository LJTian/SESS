package rabbitMq

import (
	"SESS/cmd/mq"
	"SESS/pkg/cfgYaml"
	"SESS/pkg/rabbitmq"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

type Mq struct {
	cfgPath string
	cfg     *cfgYaml.MqYaml
	connect *rabbitmq.RabbitMQ

	keyValueMap map[string]interface{} // key: topic_name
	KV          sync.RWMutex           // 读写锁

	receiveMap map[string]interface{} // 接收Map key: consumerName
	RM         sync.RWMutex           // 读写锁
}

func NewMq(cfgPath string) *Mq {
	return &Mq{
		cfgPath: cfgPath,
	}
}

// setCfg2Map 将配置信息写到map中
func (this *Mq) setCfg2Map() (err error) {

	topics := this.cfg.Topic

	// 创建内部map
	this.keyValueMap = make(map[string]interface{}, 0)
	var key string

	// 上锁
	this.KV.Lock()
	defer this.KV.Unlock()

	for _, topicV := range topics {
		// 设置主题
		key = fmt.Sprintf("%s_0", topicV.Name)
		this.keyValueMap[key] = topicV

		// 设置生产者 key:topic_producer_name
		for _, producerV := range topicV.Producer {
			if producerV.Name == "0" {
				err = errors.New("0 is reserved word, Please modify the 'name'")
				return
			}
			key = fmt.Sprintf("%s_producer_%s", topicV.Name, producerV.Name)
			this.keyValueMap[key] = producerV
		}

		// 设置消费者 key:topic_consumer_name
		for _, consumerV := range topicV.Consumer {
			if consumerV.Name == "0" {
				err = errors.New("0 is reserved word, Please modify the 'name'")
				return
			}
			key = fmt.Sprintf("%s_consumer_%s", topicV.Name, consumerV.Name)
			this.keyValueMap[key] = consumerV
		}
	}
	return
}

// Start Mq
func (this *Mq) Start() (err error) {
	config, err := ioutil.ReadFile(this.cfgPath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(config, this.cfg)
	if err != nil {
		return
	}
	this.cfg.PrintYamlInfo()

	// 整理资源至map
	err = this.setCfg2Map()
	if err != nil {
		return
	}
	// 建立链接
	// 先不考虑多Mq多机情况
	v := this.cfg.Connect[0]
	this.connect = rabbitmq.NewRabbitMQ(v.User, v.PassWd, v.PassWd, v.Port)
	err = this.connect.Connect()
	if err != nil {
		return err
	}
	return
}

func (this *Mq) Stop() {
	this.connect.CloseConnect()
}

//Send 发送
func (this *Mq) Send(data mq.DataInfo) (err error) {

	// 将锁的范围缩小
	object := func() interface{} {
		this.KV.RLock()
		defer this.KV.RUnlock()
		return this.keyValueMap[data.Operator]
	}()

	producer, ok := object.(cfgYaml.Producer)
	if !ok {
		err = errors.New(fmt.Sprintf("get Producer [%s] err ", data.Operator))
		return
	}

	// 发送信息
	err = this.connect.Send(data.Topic, producer.RoutKey, data.Msg)
	if err != nil && !errors.As(err, rabbitmq.ErrDoesNotExist) {
		return
	} else if errors.As(err, rabbitmq.ErrDoesNotExist) {
		// 如果协转器不存在创建并重新发送
		err = this.creatExchange(data)
		if err != nil {
			return
		}
		err = this.connect.Send(data.Topic, producer.RoutKey, data.Msg)
		if err != nil {
			return
		}
	}
	return
}

// 接收
func (this *Mq) Receive(data mq.DataInfo) (body chan []byte, err error) {

	// 将锁的范围缩小
	object := func() interface{} {
		this.RM.RLock()
		defer this.RM.RUnlock()
		return this.receiveMap[data.Operator]
	}()

	dev := object.(<-chan amqp.Delivery)
	for v := range dev {
		body <- v.Body
	}
	return
}

// 接收
func (this *Mq) ReceiveAck(data mq.DataInfo) (err error) {

	return
}

// 创建消费者
func (this *Mq) CreatConsumer(data mq.DataInfo) (err error) {

	// 将锁的范围缩小
	object := func() interface{} {
		this.KV.RLock()
		defer this.KV.RUnlock()
		return this.keyValueMap[data.Operator]
	}()

	consumer, ok := object.(cfgYaml.Consumer)
	if !ok {
		err = errors.New(fmt.Sprintf("get Consumer [%s] err ", data.Operator))
		return
	}

	queue, err := this.connect.CreatQueue(data.Topic, "")
	if err != nil {
		return
	}

	err = this.connect.BindQueue(data.Topic, queue.Name, consumer.RoutKey)
	if err != nil {
		return
	}

	msgs, err := this.connect.CreatConsume(data.Topic, queue.Name, consumer.Name, consumer.AutoAck)
	if err != nil {
		return
	}

	this.RM.Lock()
	defer this.RM.Unlock()
	this.receiveMap[data.Operator] = msgs

	return

}

func (this *Mq) creatExchange(data mq.DataInfo) (err error) {

	// 将锁的范围缩小
	object := func() interface{} {
		this.KV.RLock()
		defer this.KV.RUnlock()
		return this.keyValueMap[data.Operator]
	}()

	// 断言
	topic, ok := object.(cfgYaml.Topic)
	if !ok {
		err = errors.New(fmt.Sprintf("get Topic [%s] err ", data.Topic))
		return
	}

	// 开启协转器(目前只支持主题模式)
	err = this.connect.CreatExchange("topic", topic.Name)
	if err != nil {
		return err
	}
	return
}

func (this *Mq) GetConnect() *rabbitmq.RabbitMQ {
	return this.connect
}

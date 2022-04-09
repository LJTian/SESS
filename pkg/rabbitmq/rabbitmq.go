package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn       *amqp.Connection
	user       string
	passwd     string
	ip         string
	port       string
	channelMap map[string]*amqp.Channel
}

func NewRabbitMQ(user, passwd, ip, port string) *RabbitMQ {
	return &RabbitMQ{
		user:   user,
		passwd: passwd,
		ip:     ip,
		port:   port,
	}
}

// Connect 建立链接
func (this *RabbitMQ) Connect() (err error) {

	url := "amqp://" +
		this.user + ":" +
		this.passwd + "@" +
		this.ip + ":" +
		this.port + "/"

	this.conn, err = amqp.Dial(url)
	return
}

// CloseConnect 关闭链接
func (this *RabbitMQ) CloseConnect() (err error) {
	this.CloseChannels()
	return this.conn.Close()
}

func (this *RabbitMQ) CloseChannelByTopic(topic string) (err error) {
	// CloseChannelByTopic 基于事件主题 关闭管道
	if v, ok := this.channelMap[topic]; ok {
		v.ExchangeDelete(topic, false, true)
		v.Close()
		return
	}
	return errors.New(ErrDoesNotExist)
}

// CreatExchange 创建转换器
func (this *RabbitMQ) CreatExchange(exchangeType string, topic string) (err error) {

	if this.channelMap == nil {
		this.channelMap = make(map[string]*amqp.Channel)
	} else {
		if _, ok := this.channelMap[topic]; ok {
			return errors.New(ErrExisted)
		}
	}

	ch, err := this.conn.Channel()
	if err != nil {
		return
	}
	err = ch.ExchangeDeclare(
		topic,        // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	this.channelMap[topic] = ch
	return
}

// CloseChannels 关闭管道
func (this *RabbitMQ) CloseChannels() {
	for k, _ := range this.channelMap {
		this.CloseChannelByTopic(k)
	}
}

// CreatQueue 创建队列
func (this *RabbitMQ) CreatQueue(topic string, QueueName string) (queue *amqp.Queue, err error) {

	q, err := this.channelMap[topic].QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	queue = &q
	return
}

// BindQueue 绑定队列
func (this *RabbitMQ) BindQueue(topic string, queueName string, routingKey string) (err error) {
	err = this.channelMap[topic].QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		topic,      // exchange
		false,
		nil,
	)
	return
}

// Send 发送消息
func (this *RabbitMQ) Send(topic string, routingKey string, msg []byte) (err error) {

	if _, ok := this.channelMap[topic]; !ok {
		return errors.New(ErrDoesNotExist)
	}

	ch := this.channelMap[topic]
	err = ch.Publish(
		topic,      // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	return
}

// CreatConsume 创建消费者
func (this *RabbitMQ) CreatConsume(topic string, queueName string,
	ConsumeName string, aotoAck bool) (msgs <-chan amqp.Delivery, err error) {

	msgs, err = this.channelMap[topic].Consume(
		queueName,   // queue
		ConsumeName, // consumer
		aotoAck,     // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	return
}

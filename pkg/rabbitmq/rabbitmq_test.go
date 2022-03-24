package rabbitmq

import (
	"SESS/pkg/tools"
	"fmt"
	"log"
	"testing"
	"time"
)

func worker(ramq *RabbitMQ, topic string, routingKey string, name string) {

	q, err := ramq.CreatQueue(topic, "")
	if err != nil {
		tools.FailOnError(err, "Failed to declare a queue")
		return
	}

	err = ramq.BindQueue(topic, q.Name, routingKey)
	if err != nil {
		tools.FailOnError(err, "Failed to bind a queue")
		return
	}

	msgs, err := ramq.CreatConsume(topic, q.Name, "", true)
	if err != nil {
		tools.FailOnError(err, "Failed to register a consumer")
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("name:[%s] ---- [%s-%s]worker %s", name, topic, routingKey, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func producer(ramq *RabbitMQ, topic string, name string) {

	var err error

	forever1 := make(chan bool)

	for i := 0; i < 100; i++ {
		body := fmt.Sprintf("hello [%d]", i)
		if i%2 == 0 {
			err = ramq.Send(topic, "MngTran.Log", []byte(body))
		} else {
			err = ramq.Send(topic, "AcctTran.Log", []byte(body))
		}
		log.Printf("name[%s]  -- Send[%s]", name, body)

		if err != nil {
			tools.FailOnError(err, "Failed  Send")
		}
		time.Sleep(1 * time.Second)
	}

	<-forever1
}

// 测试Demo
func TestDemo(t *testing.T) {

	// 建立链接
	mq := NewRabbitMQ("guest", "guest", "10.211.55.3", "5672")
	err := mq.Connect()
	tools.FailOnError(err, "Connect err")

	topic := "tranLog"

	err = mq.CreatExchange("topic", topic)
	if err != nil {
		tools.FailOnError(err, "Failed exchange")
	}

	// 开启消费者
	go worker(mq, topic, "MngTran.*", "worker1")
	go worker(mq, topic, "*.Log", "worker2")
	go worker(mq, topic, "*.Log", "worker3")
	go worker(mq, topic, "AcctTran.*", "worker4")

	// 开启生产者
	go producer(mq, topic, "producer1")
	go producer(mq, topic, "producer2")

	stopCh := make(chan int)

	<-stopCh
	return
}

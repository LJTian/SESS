package mq

import (
	"SESS/cmd/internal/rabbitMq"
)

type DataInfo struct {
	Topic    string // 主题
	Operator string // 操作者名称
	Msg      []byte // 信息
}

// mq 接口
type mq interface {
	Start() error                     // 启动
	Stop()                            // 关闭
	Send(DataInfo) error              // 发送
	Receive(DataInfo) ([]byte, error) // 接收
	ReceiveAck(DataInfo) error        // 接受应答确认
	//CreatConsumer(DataInfo) (<-chan amqp.Delivery, error) // 创建消费者
}

// TODO 添加实现断言
var _ mq = (*rabbitMq.Mq)(nil)

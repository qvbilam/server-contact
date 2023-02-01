package initialize

import (
	"contact/business"
	"contact/global"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitQueue() {
	initChatQueue()
}

func initChatQueue() {
	user := global.ServerConfig.RabbitMQServerConfig.User
	password := global.ServerConfig.RabbitMQServerConfig.Password
	host := global.ServerConfig.RabbitMQServerConfig.Host
	port := global.ServerConfig.RabbitMQServerConfig.Port

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
		panic(any(err))
	}

	global.MessageQueueClient = conn

	// 创建私聊队列
	privateExchangeName := global.ServerConfig.RabbitMQServerConfig.ExchangeChatPrivate
	privateQueueName := fmt.Sprintf("%s-%s", privateExchangeName, global.ServerConfig.Name)
	business.CreateExchange(privateQueueName)
	business.CreateQueue(privateQueueName, privateExchangeName)

	// 创建群聊队列
	groupExchangeName := global.ServerConfig.RabbitMQServerConfig.ExchangeChatGroup
	groupQueueName := fmt.Sprintf("%s-%s", groupExchangeName, global.ServerConfig.Name)
	business.CreateExchange(groupQueueName)
	business.CreateQueue(groupQueueName, groupExchangeName)

	// 创建房间队列
	roomExchangeName := global.ServerConfig.RabbitMQServerConfig.ExchangeChatRoom
	roomQueueName := fmt.Sprintf("%s-%s", roomExchangeName, global.ServerConfig.Name)
	business.CreateExchange(roomQueueName)
	business.CreateQueue(roomQueueName, roomExchangeName)

	// 接受消息
	go business.ConsumeQueue(privateQueueName)
	go business.ConsumeQueue(groupQueueName)
	go business.ConsumeQueue(roomQueueName)
}

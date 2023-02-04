package business

import (
	"contact/enum"
	"contact/global"
	"contact/utils"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func CreateExchange(exchangeName string) {
	// 建立 amqp 通道
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "建立通道失败", err)
	}

	// 创建交换机(不存在创建)
	if err := ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil); err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列交换机", err)
	}
}

func CreateQueue(queueName, exchangeName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	q, _ := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)

	// 绑定交换机
	if exchangeName != "" {
		if err := ch.QueueBind(q.Name, "", exchangeName, false, nil); err != nil {
			zap.S().Fatalf("%s dial error: %s", "队列绑定交换机失败", err)
		}
	}
}

func DeleteQueue(queueName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	_, err = ch.QueueDelete(queueName, false, true, true)
	if err != nil {
		// 删除队列
		return
	}
	// 删除队列失败
}

// ConsumeQueue 消费消息
func ConsumeQueue(queueName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	deliveries, err := ch.Consume(queueName, "go-consumer", true, false, false, false, nil)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "消费消息失败", err)
	}

	for msg := range deliveries {
		fmt.Printf("read message: %s\n", msg.Body)
		// 传递到消息
		Dispatch(msg.Body)
	}
}

func Dispatch(data []byte) {
	fmt.Println("联系人服务分发原始数据", string(data))
	type C struct {
		Code    int64  `json:"code"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}
	type M struct {
		SenderId int64  `json:"send_user_id"`
		TargetId int64  `json:"target_id"`
		Type     string `json:"type"`
		C        C      `json:"content"`
	}
	m := M{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		zap.S().Errorf("错误的消息类型")
		return
	}
	// {SenderId:2 TargetId:1 Type:private C:{Code:0 Type:TextMsg Content:干什么呢}}
	fmt.Printf("%+v\n", m)
	allowType := []interface{}{enum.MsgTypeTxt, enum.MsgTypeImg, enum.MsgTypeGif, enum.MsgTypeLBS, enum.MsgTypeVoice, enum.MsgTypeVideo}
	if utils.InArray(m.C.Type, allowType) == false {
		return
	}

	var message string
	var tips *string
	switch m.C.Type {
	case enum.MsgTypeTxt:
		message = m.C.Content
	case enum.MsgTypeImg:
		message = "[表情]"
	case enum.MsgTypeGif:
		message = "[表情]"
	case enum.MsgTypeLBS:
		message = "[位置]"
	case enum.MsgTypeVoice:
		message = "[语音]"
	case enum.MsgTypeVideo:
		message = "[视频]"
	}

	b := ConversationBusiness{
		SenderID:        m.SenderId,
		UserID:          m.SenderId,
		ObjectType:      m.Type,
		ObjectID:        m.TargetId,
		LastMessage:     message,
		LastMessageType: m.C.Type,
		LastMessageTime: time.Now().UnixMilli(),
	}
	if tips != nil {
		b.Tips = tips
	}

	if err := b.Create(); err != nil {
		zap.S().Errorf("发送消息自动创建联系人失败")
	}
}

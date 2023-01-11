package business

import (
	userProto "contact/api/qvbilam/user/v1"
	"contact/global"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func initClient() {
	initDBClient()
	initUserClient()
}

func initUserClient() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", "127.0.0.1", 9801),
		grpc.WithInsecure())
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "grpc-user-server", err)
	}

	userClient := userProto.NewUserClient(conn)
	global.UserServerClient = userClient
}

func initDBClient() {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_contact"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	global.DB = db
}

func TestConversationBusiness_Create(t *testing.T) {
	initClient()
	b := ConversationBusiness{
		UserID:          1,
		ObjectType:      "group",
		ObjectID:        1,
		LastMessage:     "你好啊",
		LastMessageTime: time.Now().UnixMicro(),
	}

	err := b.Create()
	if err != nil {
		fmt.Println(err)
		return
	}

	bf := ConversationBusiness{
		UserID:          1,
		ObjectType:      "friend",
		ObjectID:        1,
		LastMessage:     "[表情]",
		LastMessageTime: time.Now().UnixMicro(),
	}

	err = bf.Create()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func TestConversationBusiness_List(t *testing.T) {
	initClient()
	b := ConversationBusiness{UserID: 1}
	total, list := b.List()
	fmt.Println(total)
	s, _ := json.Marshal(list)
	fmt.Printf("list: %s\n", s)
}

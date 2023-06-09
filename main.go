package main

import (
	"contact/api"
	proto "contact/api/qvbilam/contact/v1"
	"contact/global"
	"contact/initialize"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDatabase()
	initialize.InitRedis()
	initialize.InitServer()
	initialize.InitQueue()

	// 注册服务
	server := grpc.NewServer()
	proto.RegisterFriendServer(server, &api.FriendServer{})
	proto.RegisterGroupServer(server, &api.GroupServer{})
	proto.RegisterConversationServer(server, &api.ConversationServer{})

	Host := "0.0.0.0"
	Port := global.ServerConfig.Port

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Host, Port))
	if err != nil {
		zap.S().Panicf("listen port error: %s", err)
	}

	zap.S().Infof("start %s server, host: %s:%d", global.ServerConfig.Name, Host, Port)
	go func() {
		if err := server.Serve(lis); err != nil {
			zap.S().Panicf("start server error: %s", err)
		}
	}()

	// 监听结束
	quit := make(chan os.Signal)
	<-quit
}

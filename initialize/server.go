package initialize

import (
	messageProto "contact/api/qvbilam/message/v1"
	userProto "contact/api/qvbilam/user/v1"
	"contact/global"
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type dialConfig struct {
	host string
	port int64
	name string
}

type serverClientConfig struct {
	messageDialConfig *dialConfig
	userDialConfig    *dialConfig
	fileDialConfig    *dialConfig
}

func InitServer() {
	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
			name: global.ServerConfig.UserServerConfig.Name,
		},
		messageDialConfig: &dialConfig{
			host: global.ServerConfig.MessageServerConfig.Host,
			port: global.ServerConfig.MessageServerConfig.Port,
			name: global.ServerConfig.MessageServerConfig.Name,
		},
	}

	s.initUserServer()
	s.initMessageServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", s.userDialConfig.name, err)
	}

	userClient := userProto.NewUserClient(conn)

	global.UserServerClient = userClient
}

func (s *serverClientConfig) initMessageServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.messageDialConfig.host, s.messageDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", s.messageDialConfig.name, err)
	}

	messageClient := messageProto.NewMessageClient(conn)
	global.MessageServerClient = messageClient
}

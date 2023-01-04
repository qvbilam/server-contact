package global

import (
	messageProto "contact/api/qvbilam/message/v1"
	userProto "contact/api/qvbilam/user/v1"
	"contact/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	Redis               redis.Client
	ServerConfig        *config.ServerConfig
	UserServerClient    userProto.UserClient
	MessageServerClient messageProto.MessageClient
)

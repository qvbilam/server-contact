package initialize

import (
	"contact/config"
	"contact/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func InitConfig() {
	initEnvConfig()
	initViperConfig()
}

func initEnvConfig() {
	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	rabbitMQPort, _ := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	messageServerPort, _ := strconv.Atoi(os.Getenv("MESSAGE_SERVER_PORT"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER_SERVER_PORT"))

	if global.ServerConfig == nil {
		global.ServerConfig = &config.ServerConfig{}
	}
	// server
	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Port = serverPort
	// sql
	global.ServerConfig.DBConfig.Host = os.Getenv("DB_HOST")
	global.ServerConfig.DBConfig.Port = dbPort
	global.ServerConfig.DBConfig.User = os.Getenv("DB_USER")
	global.ServerConfig.DBConfig.Password = os.Getenv("DB_PASSWORD")
	global.ServerConfig.DBConfig.Database = os.Getenv("DB_DATABASE")
	// redis
	global.ServerConfig.RedisConfig.Host = os.Getenv("REDIS_HOST")
	global.ServerConfig.RedisConfig.Port = redisPort
	global.ServerConfig.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	global.ServerConfig.RedisConfig.Database = redisDb
	// rabbit
	global.ServerConfig.RabbitMQServerConfig.Host = os.Getenv("RABBITMQ_HOST")
	global.ServerConfig.RabbitMQServerConfig.Port = int64(rabbitMQPort)
	global.ServerConfig.RabbitMQServerConfig.Name = os.Getenv("RABBITMQ_NAME")
	global.ServerConfig.RabbitMQServerConfig.User = os.Getenv("RABBITMQ_USER")
	global.ServerConfig.RabbitMQServerConfig.Password = os.Getenv("RABBITMQ_PASSWORD")
	global.ServerConfig.RabbitMQServerConfig.Exchange = os.Getenv("RABBITMQ_EXCHANGE")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatPrivate = os.Getenv("RABBITMQ_EXCHANGE_CHAT_PRIVATE")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatGroup = os.Getenv("RABBITMQ_EXCHANGE_CHAT_GROUP")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatRoom = os.Getenv("RABBITMQ_EXCHANGE_CHAT_ROOM")
	global.ServerConfig.RabbitMQServerConfig.QueuePrefix = os.Getenv("RABBITMQ_QUEUE_SUFFIX")
	// message-server
	global.ServerConfig.MessageServerConfig.Name = os.Getenv("MESSAGE_SERVER_NAME")
	global.ServerConfig.MessageServerConfig.Host = os.Getenv("MESSAGE_SERVER_HOST")
	global.ServerConfig.MessageServerConfig.Port = int64(messageServerPort)
	// user-server
	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER_SERVER_NAME")
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER_SERVER_HOST")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
}

func initViperConfig() {
	file := "config.yaml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigFile(file)
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}

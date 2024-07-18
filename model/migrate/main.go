package main

import (
	"contact/global"
	"contact/initialize"
	"contact/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	initialize.InitConfig()
	user := global.ServerConfig.DBConfig.User
	password := global.ServerConfig.DBConfig.Password
	host := global.ServerConfig.DBConfig.Host
	port := global.ServerConfig.DBConfig.Port
	database := global.ServerConfig.DBConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键关联
	})

	if err != nil {
		panic(any(err))
	}
	_ = db.AutoMigrate(
		&model.Friend{},
		&model.FriendApply{},
		&model.Group{},
		&model.GroupMember{},
		&model.Conversation{}, // 创建这个表需要现将 object 注释,防止创建object表
	)

}

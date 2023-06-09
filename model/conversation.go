package model

import "time"

type Object struct {
	IDModel
	Name   string
	Avatar string
	Remark string
	IsDND  bool
}

type Conversation struct {
	IDModel
	UserID          int64 `gorm:"type:int not null default 0;comment:用户id;index:idx_user_object"`
	SenderID        int64 `gorm:"type:int not null default 0;comment:发送用户id;index:idx_sender"`
	Object          *Object
	ObjectID        int64  `gorm:"type:int not null default 0;comment:业务id;index:idx_user_object"`
	ObjectType      string `gorm:"type:varchar(128) not null default '';index:idx_user_object"`
	NewsCount       int64  `gorm:"type:int not null default 0;comment:新消息数量;"`
	Tips            string `gorm:"type:varchar(128) not null default '';"`
	LastMessage     string `gorm:"type:varchar(2048) not null default '';"`
	LastMessageType string `gorm:"type:varchar(128) not null default '';"`
	LastMessageAt   *time.Time
	DateModel
	DeletedModel
}

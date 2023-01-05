package model

import "time"

type Conversation struct {
	IDModel
	UserModel
	BusType       string `gorm:"type:varchar(128) not null default '';"`
	BusID         int64  `gorm:"type:int not null default 0;comment:业务id;index:idx_bus_id"`
	NewsCount     int64  `gorm:"type:int not null default 0;comment:新消息数量;"`
	LastMessage   string `gorm:"type:varchar(128) not null default '';"`
	LastMessageAt *time.Time
	DateModel
	DeletedModel
}

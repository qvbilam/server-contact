package model

import "time"

type Group struct {
	IDModel
	Code int64 `gorm:"type:int not null default 0;comment:群号;index:idx_code"`
	UserModel
	Name             string `gorm:"type:varchar(128) not null default '';comment:群名称"`
	Avatar           string `gorm:"type:varchar(255) not null default '';comment:群头像"`
	Cover            string `gorm:"type:varchar(255) not null default '';comment:群封面"`
	Introduce        string `gorm:"type:varchar(255) not null default '';comment:群介绍"`
	MemberCount      int64  `gorm:"type:int not null default 0;comment:群成员人数"`
	AllowMemberCount int64  `gorm:"type:int not null default 0;comment:允许群成员人数"`
	IsGlobalBanned   bool   `gorm:"type:tinyint(1) not null default 0;comment:开启群禁言"`
	BannedEndAt      *time.Time
	DateModel
	DeletedModel
}

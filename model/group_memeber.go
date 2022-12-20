package model

import "time"

type GroupMember struct {
	IDModel
	GroupID int64 `gorm:"type:int not null default 0;comment:群ID;index:idx_group_level"`
	UserModel
	NickName    string `gorm:"type:varchar(128) not null default '';comment:群名称"`
	Role        int64  `gorm:"type:int not null default 0;comment:角色: 0成员,1管理,2群主"`
	Level       int64  `gorm:"type:int not null default 0;comment:群等级;index:idx_group_level"`
	Exp         int64  `gorm:"type:int not null default 0;comment:群经验"`
	Remark      string `gorm:"type:varchar(128) not null default '';comment:群备注"`
	IsDnd       bool   `gorm:"type:tinyint(1) not null default 0;comment:是否开启免打扰"`
	IsBanned    bool   `gorm:"type:tinyint(1) not null default 0;comment:开启禁言"`
	BannedEndAt *time.Time
	DateModel
	DeletedModel
}

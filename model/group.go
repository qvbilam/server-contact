package model

import (
	"contact/enum"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

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
	Member           GroupMember   // 用户自身
	Members          []GroupMember // 所有成员
	//Conversation     Conversation  `gorm:"polymorphic:Object;"`
	DateModel
	DeletedModel
}

func (g Group) AfterCreate(tx *gorm.DB) error {
	// 创建群主
	if res := tx.Create(&GroupMember{
		GroupID:   g.ID,
		UserModel: UserModel{UserID: g.UserID},
		Role:      enum.GroupRoleOwner,
	}); res.RowsAffected == 0 || res.Error != nil {
		return status.Errorf(codes.Internal, "create group owner error: %s", res.Error)
	}

	return nil
}

func (g Group) AfterDelete(tx *gorm.DB) error {
	// 删除所有群
	res := tx.Delete(&GroupMember{GroupID: g.ID})
	if res.RowsAffected == 0 || res.Error != nil {
		return status.Errorf(codes.Internal, "delete group members error: %s", res.Error)
	}
	return nil
}

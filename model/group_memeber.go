package model

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type GroupMember struct {
	IDModel
	GroupID int64 `gorm:"type:int not null default 0;comment:群ID;index:idx_group_level"`
	UserModel
	Nickname    string `gorm:"type:varchar(128) not null default '';comment:群名称"`
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

func (g GroupMember) AfterCreate(tx *gorm.DB) error {
	// 添加人数
	groupModel := &Group{IDModel: IDModel{ID: g.GroupID}}
	if res := tx.Model(&groupModel).
		Update("member_count", gorm.Expr("member_count + ?", 1)); res.RowsAffected == 0 || res.Error != nil {
		return status.Errorf(codes.Internal, "create group owner error: %s", res.Error)
	}

	return nil
}

func (g GroupMember) AfterDelete(tx *gorm.DB) error {
	// 减少人数
	groupModel := &Group{IDModel: IDModel{ID: g.GroupID}}
	if res := tx.Model(&groupModel).
		Update("member_count", gorm.Expr("member_count - ?", 1)); res.RowsAffected == 0 || res.Error != nil {
		return status.Errorf(codes.Internal, "create group owner error: %s", res.Error)
	}

	return nil
}

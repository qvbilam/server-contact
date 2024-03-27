package model

type Friend struct {
	IDModel
	UserID       int64        `gorm:"type:int not null default 0;comment:用户id;index:idx_user_apply"`
	FriendUserID int64        `gorm:"type:int not null default 0;comment:好用用户id;index:idx_user_friend"`
	Remark       string       `gorm:"type:varchar(128) not null default '';comment:备注;"`
	Conversation Conversation `gorm:"polymorphic:Object;"`
	DateModel
	DeletedModel
}

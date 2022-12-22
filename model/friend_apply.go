package model

type FriendApply struct {
	IDModel
	UserID      int64  `gorm:"type:int not null default 0;comment:用户id;index:idx_user_apply"`
	ApplyUserID int64  `gorm:"type:int not null default 0;comment:申请用户id;index:idx_user_apply"`
	Content     string `gorm:"type:varchar(255) not null default 0;comment:内容"`
	Status      int64  `gorm:"type:tinyint(3) not null default 0;comment:-1拒绝,0等待,1同意"`
	DateModel
	DeletedModel
}

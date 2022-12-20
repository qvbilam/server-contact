package model

type Friend struct {
	IDModel
	UserModel
	FriendUserID int64 `gorm:"type:int not null default 0;comment:朋友id;"`
	DateModel
	DeletedModel
}

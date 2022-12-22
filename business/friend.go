package business

import (
	"contact/global"
	"contact/model"
	"gorm.io/gorm"
)

type FriendBusiness struct {
	ID           int64
	UserID       int64
	FriendUserID int64
	Remark       string
}

func (b *FriendBusiness) IsFriend(tx *gorm.DB) bool {
	var count int64
	if tx == nil {
		tx = global.DB
	}
	res := tx.Model(&model.Friend{}).Where(model.Friend{
		UserID:       b.UserID,
		FriendUserID: b.FriendUserID,
	}).Count(&count)
	if res.RowsAffected == 0 || res.Error != nil {
		return false
	}

	return count > 0
}

func (b *FriendBusiness) Users() []model.Friend {
	var users []model.Friend
	tx := global.DB
	res := tx.Where(&model.Friend{UserID: b.UserID}).Find(&users)
	if res.RowsAffected == 0 {
		return nil
	}
	return users
}

func (b *FriendBusiness) Delete() error {
	tx := global.DB
	tx.Begin()
	m := model.Friend{UserID: b.UserID, FriendUserID: b.FriendUserID}
	condition, entity := m, m
	if res := tx.Where(&condition).Delete(&entity); res.Error != nil || res.RowsAffected == 0 {
		tx.Commit()
		return nil
	}
	tx.Commit()
	return nil
}

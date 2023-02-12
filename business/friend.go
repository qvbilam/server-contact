package business

import (
	"contact/global"
	"contact/model"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FriendBusiness struct {
	ID            int64
	UserID        int64
	FriendUserID  int64
	FriendUserIds []int64
	Remark        *string
	Keyword       string
}

func (b *FriendBusiness) IsFriend() bool {
	var count int64
	fmt.Printf("%+v\n", b)
	res := global.DB.Debug().Model(&model.Friend{}).Where(model.Friend{
		UserID:       b.UserID,
		FriendUserID: b.FriendUserID,
	}).Count(&count)
	if res.RowsAffected == 0 || res.Error != nil {
		return false
	}

	return count > 0
}

func (b *FriendBusiness) Users() (int64, []model.Friend) {
	var users []model.Friend
	tx := global.DB
	tx = tx.Where(&model.Friend{UserID: b.UserID})
	if b.FriendUserIds != nil {
		tx = tx.Where("friend_user_id in ?", b.FriendUserIds)
	}
	if b.Keyword != "" {
		// todo 查用户 code, 昵称, 备注
	}

	res := tx.Find(&users)
	if res.RowsAffected == 0 {
		return 0, nil
	}

	return int64(len(users)), users
}

func (b *FriendBusiness) Update() error {
	entity := model.Friend{}
	tx := global.DB
	tx.Begin()
	res := tx.Where(model.Friend{IDModel: model.IDModel{ID: b.ID}}).First(&entity)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "未找到好友")
	}

	if b.UserID != entity.UserID {
		tx.Rollback()
		return status.Errorf(codes.Unauthenticated, "非法操作")
	}

	if b.Remark != nil {
		entity.Remark = *b.Remark
	}

	res = tx.Save(entity)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.Internal, "修改好友信息失败")
	}
	tx.Commit()
	return nil
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

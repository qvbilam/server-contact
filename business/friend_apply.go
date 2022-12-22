package business

import (
	"contact/enum"
	"contact/global"
	"contact/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FriendApplyBusiness struct {
	ID          int64
	UserID      int64
	ApplyUserID int64
	Content     string
}

func (b *FriendApplyBusiness) IsApply(tx *gorm.DB) bool {
	var count int64
	if tx == nil {
		tx = global.DB
	}
	res := tx.Model(&model.FriendApply{}).Where(model.FriendApply{
		UserID:      b.UserID,
		ApplyUserID: b.ApplyUserID,
		Status:      enum.FriendApplyStatusWait,
	}).Count(&count)

	if res.RowsAffected == 0 || res.Error != nil {
		return false
	}

	return count > 0
}

func (b *FriendApplyBusiness) Apply() error {
	tx := global.DB
	tx.Begin()
	fb := FriendBusiness{
		UserID:       b.ApplyUserID,
		FriendUserID: b.UserID,
	}
	if fb.IsFriend(tx) {
		tx.Rollback()
		return status.Errorf(codes.AlreadyExists, "已经是好友")
	}
	if b.IsApply(tx) {
		tx.Rollback()
		return status.Errorf(codes.AlreadyExists, "等待对方确认")
	}
	entity := model.FriendApply{
		UserID:      b.UserID,
		ApplyUserID: b.ApplyUserID,
		Content:     b.Content,
		Status:      enum.FriendApplyStatusReject,
	}

	tx.Create(&entity)
	tx.Commit()
	return nil
}

func (b *FriendApplyBusiness) Agree() error {
	tx := global.DB
	tx.Begin()
	entity := model.FriendApply{}
	res := tx.Where(&model.FriendApply{Status: enum.FriendApplyStatusWait}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&entity, b.ID)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "处理成功")
	}

	entity.Status = enum.FriendApplyStatusAgree
	tx.Save(&entity)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "添加失败")
	}

	var f = []model.Friend{{
		UserID:       b.UserID,
		FriendUserID: b.ApplyUserID,
	}, {
		UserID:       b.ApplyUserID,
		FriendUserID: b.UserID,
	}}

	res = tx.CreateInBatches(&f, 100)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.Internal, "添加失败")
	}

	tx.Commit()
	return nil
}

func (b *FriendApplyBusiness) Reject() error {
	tx := global.DB
	tx.Begin()
	entity := model.FriendApply{}
	res := tx.Where(&model.FriendApply{Status: enum.FriendApplyStatusWait}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&entity, b.ID)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "处理成功")
	}

	entity.Status = enum.FriendApplyStatusReject
	res = tx.Save(&entity)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "拒绝失败")
	}

	tx.Commit()
	return nil
}

func (b *FriendApplyBusiness) Users() []model.FriendApply {
	var users []model.FriendApply
	tx := global.DB
	res := tx.Where(&model.FriendApply{UserID: b.UserID}).Find(&users)
	if res.RowsAffected == 0 {
		return nil
	}
	return users
}

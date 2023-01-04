package business

import (
	"contact/enum"
	"contact/global"
	"contact/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm/clause"
	"time"
)

type GroupMemberBusiness struct {
	ID          *int64
	GroupID     *int64
	UserID      *int64
	Nickname    *string
	Role        *int64
	Level       *int64
	Exp         *int64
	Remark      *string
	IsDnd       *bool
	IsBanned    *bool
	BannedEndAt string
}

func (b *GroupMemberBusiness) Create() (int64, error) {
	m := b.ToModel()
	tx := global.DB.Begin()

	// 验证群人数
	groupEntity := model.Group{}
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&groupEntity, b.GroupID)
	if groupEntity.MemberCount >= groupEntity.AllowMemberCount {
		tx.Rollback()
		return 0, status.Errorf(codes.ResourceExhausted, "group is full")
	}

	// 验证用户是否存在
	if res := tx.Where(model.GroupMember{GroupID: *b.GroupID, UserModel: model.UserModel{UserID: *b.UserID}}).First(&model.GroupMember{}); res.RowsAffected != 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.AlreadyExists, "user is joined group")
	}

	// 添加成员
	if res := tx.Create(&m); res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "create group member error: %s", res.Error)
	}
	tx.Commit()
	return m.ID, nil
}

func (b *GroupMemberBusiness) Delete() (int64, error) {
	tx := global.DB.Begin()
	// 验证是否为群主
	if res := tx.Where(model.GroupMember{
		GroupID:   *b.GroupID,
		UserModel: model.UserModel{UserID: *b.UserID},
		Role:      enum.GroupRoleOwner,
	}).First(&model.GroupMember{}); res.RowsAffected != 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.AlreadyExists, "user is owner can not quit group")
	}

	// 注意: 删除实体进入afterDelete是获取不到ID的. 需要在模型中传入请求的参数
	m := b.ToModel()
	res := tx.Where(m).Delete(&m)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "delete group member error: %s", res.Error)
	}

	tx.Commit()
	return res.RowsAffected, nil
}

func (b *GroupMemberBusiness) Member() (*model.GroupMember, error) {
	var member *model.GroupMember
	res := global.DB.Where(&model.GroupMember{
		GroupID: *b.GroupID,
		UserModel: model.UserModel{
			UserID: *b.UserID,
		},
	}).First(&member)

	if res.RowsAffected == 0 || res.Error != nil || member == nil {
		return nil, status.Errorf(codes.NotFound, "用户未加入群")
	}

	return member, nil
}

func (b *GroupMemberBusiness) Members() []model.GroupMember {
	var members []model.GroupMember
	if res := global.DB.Where(model.GroupMember{GroupID: *b.GroupID}).Find(&members); res.RowsAffected == 0 {
		return nil
	}
	return members
}

func (b *GroupMemberBusiness) ToModel() model.GroupMember {
	m := model.GroupMember{}
	if b.ID != nil {
		m.ID = *b.ID
	}
	if b.GroupID != nil {
		m.GroupID = *b.GroupID
	}
	if b.UserID != nil {
		m.UserID = *b.UserID
	}
	if b.Nickname != nil {
		m.Nickname = *b.Nickname
	}
	if b.Role != nil {
		m.Role = *b.Role
	}
	if b.Level != nil {
		m.Level = *b.Level
	}
	if b.Exp != nil {
		m.Exp = *b.Exp
	}
	if b.Remark != nil {
		m.Remark = *b.Remark
	}
	if b.IsDnd != nil {
		m.IsDnd = *b.IsDnd
	}
	if b.IsBanned != nil {
		m.IsBanned = *b.IsBanned
	}
	if b.BannedEndAt != "" {
		date, _ := time.ParseInLocation("2006-01-02 15:04:05", b.BannedEndAt, time.Local)
		m.BannedEndAt = &date
	}

	return m
}

package business

import (
	"contact/global"
	"contact/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupBusiness struct {
	ID               *int64
	Code             *int64
	UserId           *int64
	Name             string
	Avatar           string
	Cover            string
	Introduce        string
	MemberCount      int64
	AllowMemberCount int64
	IsGlobalBanned   *bool
	BannedEndAt      int64
}

func (b *GroupBusiness) Create() (int64, error) {
	m := b.ToModel()
	tx := global.DB.Begin()
	// 创建群
	if res := tx.Create(&m); res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "create group error: %s", res.Error)
	}

	// 创建群成员

	tx.Commit()
	return m.ID, nil
}

func (b *GroupBusiness) Detail() *model.Group {
	m := model.Group{}
	if res := global.DB.First(&m, *b.ID); res.RowsAffected == 0 || res.Error != nil {
		return nil
	}
	return &m
}

func (b *GroupBusiness) ToModel() model.Group {
	m := model.Group{}
	if b.Code != nil {
		m.Code = *b.Code
	}
	if b.UserId != nil {
		m.UserID = *b.UserId
	}
	if b.Name != "" {
		m.Name = b.Name
	}
	if b.Avatar != "" {
		m.Avatar = b.Avatar
	}
	if b.Cover != "" {
		m.Cover = b.Cover
	}
	if b.Introduce != "" {
		m.Introduce = b.Introduce
	}
	if b.MemberCount != 0 {
		m.MemberCount = b.MemberCount
	}
	if b.AllowMemberCount != 0 {
		m.AllowMemberCount = b.AllowMemberCount
	}
	if b.IsGlobalBanned != nil {
		m.IsGlobalBanned = *b.IsGlobalBanned
		if *b.IsGlobalBanned == true && b.BannedEndAt != 0 {
			//date, _ := time.ParseInLocation("2006-01-02 15:04:05", strconv.FormatInt(b.BannedEndAt, 10), time.Local)
			//m.BannedEndAt = date
		}
	}

	return m
}

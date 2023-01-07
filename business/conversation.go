package business

import "C"
import (
	userProto "contact/api/qvbilam/user/v1"
	"contact/enum"
	"contact/global"
	"contact/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ConversationBusiness struct {
	ID              int64
	UserID          int64
	ObjectType      string
	ObjectID        int64
	NewsCount       int64
	Tips            *string
	LastMessage     string
	LastMessageTime int64
}

func (b *ConversationBusiness) NewCount() int64 {
	// 获取用户未开启免打扰的群
	var groupMembers []model.GroupMember
	var groupIds []int64
	global.DB.Where(&model.GroupMember{UserModel: model.UserModel{UserID: b.UserID}, IsDnd: false}).Select("group_id").Find(&groupMembers)
	for _, g := range groupMembers {
		groupIds = append(groupIds, g.GroupID)
	}

	// 查询未读消息数量
	var res struct {
		Count int64
	}
	if res := global.DB.Where(&model.Conversation{ObjectType: enum.ObjectTypeUser}).
		Or("object_type = ? and object_id in ?", enum.ObjectTypeGroup, groupIds).
		Select("sum(`news_count`) as count").First(&res); res.RowsAffected == 0 {
		return res.RowsAffected
	}

	return res.Count
}

func (b *ConversationBusiness) List() (int64, []*model.Conversation) {
	var conversations []*model.Conversation
	res := global.DB.
		Where(&model.Conversation{UserID: b.UserID}).
		Find(&conversations)
	if res.RowsAffected == 0 || res.Error != nil {
		return 0, nil
	}

	var userIds []int64
	for _, c := range conversations {
		if c.ObjectType == enum.ObjectTypeGroup {
			userIds = append(userIds, c.ObjectID)
		}
	}

	groupMap := make(map[int64]model.Group)
	friendMap := make(map[int64]model.Friend)
	userMap := make(map[int64]*userProto.UserResponse)
	users, _ := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{Id: userIds})

	gb := GroupBusiness{UserId: &b.UserID}
	fb := FriendBusiness{UserID: b.UserID}
	_, groups := gb.Mine()
	_, friends := fb.Users()
	for _, g := range groups {
		groupMap[g.ID] = g
	}
	for _, f := range friends {
		friendMap[f.FriendUserID] = f
	}
	if users != nil {
		for _, u := range users.Users {
			userMap[u.Id] = u
		}
	}

	for _, c := range conversations {
		if c.ObjectType == enum.ObjectTypeGroup {
			c.Object.ID = groupMap[c.ObjectID].ID
			c.Object.Name = groupMap[c.ObjectID].Name
			c.Object.Avatar = groupMap[c.ObjectID].Avatar
			c.Object.Remark = groupMap[c.ObjectID].Member.Remark
			c.Object.IsDND = groupMap[c.ObjectID].Member.IsDnd
		}
		if c.ObjectType == enum.ObjectTypeUser {
			remark := ""
			if _, ok := friendMap[c.ObjectID]; ok {
				remark = friendMap[c.ObjectID].Remark
			}
			if _, ok := userMap[c.ObjectID]; ok {
				c.Object.ID = userMap[c.ObjectID].Id
				c.Object.Name = userMap[c.ObjectID].Nickname
				c.Object.Avatar = userMap[c.ObjectID].Avatar
				c.Object.Remark = remark
			}
		}

	}
	return res.RowsAffected, conversations
}

func (b *ConversationBusiness) Create() error {
	var entity model.Conversation
	tx := global.DB.Begin()

	res := tx.Where(&model.Conversation{
		UserID:     b.UserID,
		ObjectType: b.ObjectType,
		ObjectID:   b.ObjectID,
	}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&entity)

	// 时间戳转 time.Time
	t := time.UnixMicro(b.LastMessageTime)

	if res.RowsAffected == 0 {
		entity.UserID = b.UserID
		entity.ObjectType = b.ObjectType
		entity.ObjectID = b.ObjectID
		entity.NewsCount = 1
		if b.Tips != nil {
			entity.Tips = *b.Tips
		}
		entity.LastMessage = b.LastMessage
		entity.LastMessageAt = &t
		res = tx.Create(&entity)
	} else {
		updates := map[string]interface{}{
			"news_count":      gorm.Expr("news_count + ?", 1),
			"last_message":    b.LastMessage,
			"last_message_at": &t,
		}
		if b.Tips != nil {
			updates["tips"] = *b.Tips
		}

		res = tx.Model(&model.Conversation{}).Where(&model.Conversation{IDModel: model.IDModel{ID: entity.ID}}).Updates(updates)
	}
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, res.Error.Error())
	}

	tx.Commit()
	return nil
}

func (b *ConversationBusiness) Delete() error {
	tx := global.DB.Begin()
	res := tx.Delete(&model.Conversation{IDModel: model.IDModel{ID: b.ID}}, b.ID)
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	return nil
}

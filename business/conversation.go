package business

import "C"
import (
	userProto "contact/api/qvbilam/user/v1"
	"contact/enum"
	"contact/global"
	"contact/model"
	"contact/utils"
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
	SenderID        int64
	ObjectType      string
	ObjectID        int64
	NewsCount       int64
	Tips            *string
	LastMessage     string
	LastMessageType string
	LastMessageTime int64
}

// 已读
func (b *ConversationBusiness) Read() error {
	global.DB.Model(&model.Conversation{}).Where(&model.Conversation{
		IDModel: model.IDModel{ID: b.ID},
		UserID:  b.UserID,
	}).Updates(map[string]interface{}{
		"news_count": 0,
		"tips":       "",
	})

	return nil
}

// NewCount 新消息数量
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
	if res := global.DB.Model(&model.Conversation{}).
		Where(&model.Conversation{ObjectType: enum.ObjectTypeUser}).
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
		Order("last_message_at desc").
		Find(&conversations)
	if res.RowsAffected == 0 || res.Error != nil {
		return 0, nil
	}

	var userIds []int64
	for _, c := range conversations {
		if c.ObjectType == enum.ObjectTypeUser {
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

			if _, ok := groupMap[c.ObjectID]; ok {
				c.Object = &model.Object{}
				c.Object.ID = groupMap[c.ObjectID].ID
				c.Object.Name = groupMap[c.ObjectID].Name
				c.Object.Avatar = groupMap[c.ObjectID].Avatar
				c.Object.Remark = groupMap[c.ObjectID].Member.Remark
				c.Object.IsDND = groupMap[c.ObjectID].Member.IsDnd
			}
		}
		if c.ObjectType == enum.ObjectTypeUser {
			remark := ""
			if _, ok := friendMap[c.ObjectID]; ok {
				remark = friendMap[c.ObjectID].Remark
			}
			if _, ok := userMap[c.ObjectID]; ok {
				c.Object = &model.Object{}
				c.Object.ID = userMap[c.ObjectID].Id
				c.Object.Name = userMap[c.ObjectID].Nickname
				c.Object.Avatar = userMap[c.ObjectID].Avatar
				c.Object.Remark = remark
			}
		}

	}

	return res.RowsAffected, conversations
}

// Create 创建/更新联系人
func (b *ConversationBusiness) Create() error {
	tx := global.DB.Begin()

	// 时间戳转 time.Time
	if b.LastMessageTime == 0 {
		b.LastMessageTime = time.Now().UnixMilli()
	}
	//t := time.UnixMicro(b.LastMessageTime)

	var createRes bool
	if b.ObjectType == enum.ObjectTypeGroup {
		createRes = b.createGroup(tx)
	} else if b.ObjectType == enum.ObjectTypeUser || b.ObjectType == enum.ObjectTypePrivate {
		createRes = b.createUser(tx)
	} else {
		tx.Rollback()
		return status.Errorf(codes.Internal, "不支持的会话对象类型: %s", b.ObjectType)
	}
	if createRes == false {
		tx.Rollback()
		return status.Errorf(codes.Internal, "创建会话失败")
	}

	tx.Commit()
	return nil
}

// Delete 删除会话
func (b *ConversationBusiness) Delete() error {
	tx := global.DB.Begin()
	res := tx.Where(&model.Conversation{UserID: b.UserID}).Delete(&model.Conversation{IDModel: model.IDModel{ID: b.ID}}, b.ID)
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	return nil
}

// 创建群聊会话
func (b *ConversationBusiness) createGroup(tx *gorm.DB) bool {
	var insertUserIds []int64
	var updateUserIds []int64
	var insertData []model.Conversation
	var conversationUsers []model.Conversation
	var conversationUserIds []int

	unix := time.UnixMilli(b.LastMessageTime)

	// 获取有当前群会话的所有用户
	tx.Model(&model.Conversation{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Conversation{
		ObjectID:   b.ObjectID,
		ObjectType: enum.ObjectTypeGroup,
	}).Order("user_id asc").Select("user_id").Find(&conversationUsers)

	for _, c := range conversationUsers {
		conversationUserIds = append(conversationUserIds, int(c.UserID))
	}

	// 获取群组所有用户
	gmb := GroupMemberBusiness{GroupID: &b.ObjectID}
	members, total := gmb.Members()
	if total == 0 {
		return false
	}

	for _, member := range members {
		if utils.BinarySearch(conversationUserIds, int(member.UserID)) == -1 {
			insertUserIds = append(insertUserIds, member.UserID)
		} else {
			updateUserIds = append(updateUserIds, member.UserID)
		}
	}

	ic := model.Conversation{
		SenderID:        b.SenderID,
		ObjectID:        b.ObjectID,
		ObjectType:      b.ObjectType,
		NewsCount:       1,
		LastMessage:     b.LastMessage,
		LastMessageType: b.LastMessageType,
		LastMessageAt:   &unix,
	}

	updateData := map[string]interface{}{
		"sender_id":         b.SenderID,
		"news_count":        gorm.Expr("news_count + ?", 1),
		"last_message":      b.LastMessage,
		"last_message_type": b.LastMessageType,
		"last_message_at":   &unix,
	}

	if b.Tips != nil {
		ic.Tips = *b.Tips
		updateData["tips"] = *b.Tips
	}

	for _, uId := range insertUserIds {
		i := ic
		i.UserID = uId
		insertData = append(insertData, i)
	}

	// 更新
	if len(updateUserIds) > 0 {
		if res := tx.Model(&model.Conversation{}).Where(&model.Conversation{
			ObjectID:   b.ObjectID,
			ObjectType: enum.ObjectTypeGroup,
		}).Where("user_id in ?", updateUserIds).Updates(updateData); res.RowsAffected == 0 {
			return false
		}
	}

	// 创建
	if len(insertData) > 0 {
		if res := tx.CreateInBatches(insertData, 1000); res.RowsAffected == 0 {
			return false
		}
	}

	return true
}

// 创建私聊会话
func (b *ConversationBusiness) createUser(tx *gorm.DB) bool {
	// A 对话 B， A 打开对话框自动创建会话， 所以只需要 更新AB, 增加 BA 的会话框
	var entity model.Conversation
	var senderEntity model.Conversation

	res := tx.Where(&model.Conversation{
		UserID:     b.ObjectID,
		ObjectType: enum.ObjectTypeUser,
		ObjectID:   b.SenderID,
	}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&entity)
	unix := time.UnixMilli(b.LastMessageTime)

	// 创建/修改 接受人会话
	if res.RowsAffected == 0 {
		entity.UserID = b.ObjectID
		entity.ObjectType = enum.ObjectTypeUser
		entity.ObjectID = b.SenderID
		entity.NewsCount = 1
		entity.SenderID = b.SenderID
		if b.Tips != nil {
			entity.Tips = *b.Tips
		}
		entity.LastMessage = b.LastMessage
		entity.LastMessageType = b.LastMessageType
		entity.LastMessageAt = &unix
		res = tx.Create(&entity)
	} else {
		updates := map[string]interface{}{
			"news_count":        gorm.Expr("news_count + ?", 1),
			"sender_id":         b.SenderID,
			"last_message":      b.LastMessage,
			"last_message_type": b.LastMessageType,
			"last_message_at":   &unix,
		}
		if b.Tips != nil {
			updates["tips"] = *b.Tips
		}

		res = tx.Model(&model.Conversation{}).Where(&model.Conversation{IDModel: model.IDModel{ID: entity.ID}}).Updates(updates)
	}

	if res.RowsAffected == 0 {
		return false
	}

	// 修改发送人会话
	res = tx.Where(&model.Conversation{
		UserID:     b.SenderID,
		ObjectType: enum.ObjectTypeUser,
		ObjectID:   b.ObjectID,
	}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&senderEntity)

	// 创建/修改 发送者会话
	if res.RowsAffected == 0 {
		senderEntity.UserID = b.SenderID
		senderEntity.ObjectType = enum.ObjectTypeUser
		senderEntity.ObjectID = b.ObjectID
		senderEntity.SenderID = b.SenderID
		senderEntity.LastMessage = b.LastMessage
		senderEntity.LastMessageType = b.LastMessageType
		senderEntity.LastMessageAt = &unix
		res = tx.Create(&senderEntity)
	} else {
		updates := map[string]interface{}{
			"sender_id":         b.SenderID,
			"last_message":      b.LastMessage,
			"last_message_type": b.LastMessageType,
			"last_message_at":   &unix,
		}

		res = tx.Model(&model.Conversation{}).Where(&model.Conversation{IDModel: model.IDModel{ID: senderEntity.ID}}).Updates(updates)
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true

}

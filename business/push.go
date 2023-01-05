package business

import (
	messageProto "contact/api/qvbilam/message/v1"
	"contact/enum"
	"contact/global"
	"context"
	"encoding/json"
)

type PushBusiness struct {
}

func (b *PushBusiness) GroupJoin(groupId int64, user interface{}) error {
	type e struct {
		User interface{} `json:"user"`
	}
	ext := e{}
	ext.User = user
	extra, _ := json.Marshal(ext)

	_, err := global.MessageServerClient.CreateGroupTipMessage(context.Background(), &messageProto.CreateGroupRequest{
		UserId:  0,
		GroupId: groupId,
		Message: &messageProto.MessageRequest{
			Code:    enum.PushGroupJoinCode,
			Content: "{user}, 加入群聊, {click}",
			Extra:   string(extra),
		},
	})
	return err
}

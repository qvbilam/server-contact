package api

import (
	proto "contact/api/qvbilam/contact/v1"
	"contact/business"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConversationServer struct {
	proto.UnimplementedConversationServer
}

func (s *ConversationServer) Get(ctx context.Context, request *proto.GetConversationRequest) (*proto.ConversationsResponse, error) {
	b := business.ConversationBusiness{UserID: request.UserId}
	total, conversations := b.List()
	newsCount := b.NewCount()

	var cs []*proto.ConversationResponse
	for _, c := range conversations {
		var obj *proto.ObjectResponse
		if c.Object != nil {
			obj.Id = c.Object.ID
			obj.Name = c.Object.Name
			obj.Avatar = c.Object.Avatar
			obj.Remark = c.Object.Remark
			obj.IsDND = c.Object.IsDND
		}
		cs = append(cs, &proto.ConversationResponse{
			Id:              c.ID,
			UserId:          c.UserID,
			ObjectType:      c.ObjectType,
			ObjectId:        c.ObjectID,
			Object:          obj,
			NewsCount:       c.NewsCount,
			Tips:            c.Tips,
			LastMessage:     c.LastMessage,
			LastMessageTime: c.LastMessageAt.Unix(),
		})
	}

	return &proto.ConversationsResponse{
		Total:         total,
		Conversations: cs,
		NewsCount:     newsCount,
	}, nil
}

func (s *ConversationServer) Create(ctx context.Context, request *proto.UpdateConversationRequest) (*emptypb.Empty, error) {
	b := business.ConversationBusiness{
		UserID:          request.UserId,
		ObjectType:      request.ObjectType,
		ObjectID:        request.ObjectId,
		Tips:            &request.Tips,
		LastMessage:     request.LastMessage,
		LastMessageTime: request.LastTime,
	}
	_ = b.Create()
	return &emptypb.Empty{}, nil
}

func (s *ConversationServer) Delete(ctx context.Context, request *proto.UpdateConversationRequest) (*emptypb.Empty, error) {
	b := business.ConversationBusiness{ID: request.Id, UserID: request.UserId}
	_ = b.Delete()
	return &emptypb.Empty{}, nil
}

func (s *ConversationServer) Read(ctx context.Context, request *proto.UpdateConversationRequest) (*emptypb.Empty, error) {
	b := business.ConversationBusiness{
		UserID:     request.UserId,
		ObjectType: request.ObjectType,
		ObjectID:   request.ObjectId,
	}
	_ = b.Read()
	return &emptypb.Empty{}, nil
}

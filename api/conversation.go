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
		cs = append(cs, &proto.ConversationResponse{
			Id:         c.ID,
			UserId:     c.UserID,
			ObjectType: c.ObjectType,
			ObjectId:   c.ObjectID,
			Object: &proto.ObjectResponse{
				Id:     c.Object.ID,
				Name:   c.Object.Name,
				Avatar: c.Object.Avatar,
				Remark: c.Object.Remark,
				IsDND:  c.Object.IsDND,
			},
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

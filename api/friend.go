package api

import (
	proto "contact/api/qvbilam/contact/v1"
	"contact/business"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FriendServer struct {
	proto.UnimplementedFriendServer
}

func (s *FriendServer) Apply(ctx context.Context, request *proto.UpdateFriendApplyRequest) (*emptypb.Empty, error) {
	b := business.FriendApplyBusiness{
		UserID:      request.UserId,
		ApplyUserID: request.ApplyUserId,
		Content:     request.Content,
	}
	if err := b.Apply(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FriendServer) ApplyAgree(ctx context.Context, request *proto.UpdateFriendApplyRequest) (*emptypb.Empty, error) {
	b := business.FriendApplyBusiness{
		ID:          request.Id,
		ApplyUserID: request.UserId,
	}
	if err := b.Agree(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *FriendServer) ApplyReject(ctx context.Context, request *proto.UpdateFriendApplyRequest) (*emptypb.Empty, error) {
	b := business.FriendApplyBusiness{
		ID:          request.Id,
		ApplyUserID: request.UserId,
	}
	if err := b.Reject(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *FriendServer) GetApply(ctx context.Context, request *proto.UpdateFriendApplyRequest) (*proto.FriendAppliesResponse, error) {
	b := business.FriendApplyBusiness{UserID: request.UserId}
	res := &proto.FriendAppliesResponse{}
	total, users := b.Users()
	res.Total = total
	var usersRes []*proto.FriendApplyResponse
	for _, u := range users {
		usersRes = append(usersRes, &proto.FriendApplyResponse{
			Id:          u.ID,
			UserId:      u.UserID,
			ApplyUserId: u.ApplyUserID,
			Content:     u.Content,
			Status:      u.Status,
			User:        nil,
		})
	}
	res.Users = usersRes
	return res, nil
}

package api

import (
	proto "contact/api/qvbilam/contact/v1"
	userProto "contact/api/qvbilam/user/v1"
	"contact/business"
	"contact/global"
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
	total, applies := b.Users()

	res := &proto.FriendAppliesResponse{}
	res.Total = total

	var userIds []int64
	var usersRes []*proto.FriendApplyResponse
	for _, u := range applies {
		userIds = append(userIds, u.ApplyUserID)
	}

	userMap := make(map[int64]*userProto.UserResponse)
	users, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{
		Id: userIds,
	})
	if err != nil {
		return nil, err
	}
	for _, u := range users.Users {
		userMap[u.Id] = u
	}

	for _, u := range applies {
		usersRes = append(usersRes, &proto.FriendApplyResponse{
			Id:          u.ID,
			UserId:      u.UserID,
			Content:     u.Content,
			ApplyUserId: u.ApplyUserID,
			ApplyUser:   userMap[u.ApplyUserID],
			Status:      u.Status,
		})
	}

	res.Users = usersRes
	return res, nil
}

func (s *FriendServer) Get(ctx context.Context, request *proto.SearchFriendRequest) (*proto.FriendsResponse, error) {
	b := business.FriendBusiness{
		UserID:        request.UserId,
		FriendUserIds: request.FriendIds,
		Keyword:       request.Keyword,
	}
	total, friends := b.Users()

	res := &proto.FriendsResponse{}
	res.Total = total
	var usersRes []*proto.FriendResponse

	var userIds []int64
	for _, u := range friends {
		userIds = append(userIds, u.FriendUserID)
	}

	userMap := make(map[int64]*userProto.UserResponse)
	users, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{
		Id: userIds,
	})
	if err != nil {
		return nil, err
	}
	for _, u := range users.Users {
		userMap[u.Id] = u
	}

	for _, u := range friends {
		usersRes = append(usersRes, &proto.FriendResponse{
			Id:     u.ID,
			UserId: u.UserID,
			Remark: u.Remark,
			Friend: userMap[u.FriendUserID],
		})
	}

	res.Friends = usersRes
	return res, nil
}

func (s *FriendServer) Update(ctx context.Context, request *proto.UpdateFriendRequest) (*emptypb.Empty, error) {
	b := business.FriendBusiness{
		ID:     request.Id,
		UserID: request.UserId,
		Remark: &request.Remark,
	}
	if err := b.Update(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FriendServer) Delete(ctx context.Context, request *proto.UpdateFriendRequest) (*emptypb.Empty, error) {
	b := business.FriendBusiness{
		ID:     request.Id,
		UserID: request.UserId,
	}
	if err := b.Delete(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

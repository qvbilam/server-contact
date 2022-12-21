package api

import (
	proto "contact/api/qvbilam/contact/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FriendServer struct {
	proto.UnimplementedFriendServer
}

func (s *FriendServer) Create(ctx context.Context, request *proto.UpdateFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

func (s *FriendServer) Update(ctx context.Context, request *proto.UpdateFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

func (s *FriendServer) Delete(ctx context.Context, request *proto.UpdateFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

func (s *FriendServer) Get(ctx context.Context, request *proto.SearchFriendRequest) (*proto.FriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

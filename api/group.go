package api

import (
	proto "contact/api/qvbilam/contact/v1"
	userProto "contact/api/qvbilam/user/v1"
	"contact/business"
	"contact/enum"
	"contact/global"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GroupServer struct {
	proto.UnimplementedGroupServer
}

func (s *GroupServer) Create(ctx context.Context, request *proto.UpdateGroupRequest) (*proto.GroupResponse, error) {
	if request.AllowMemberCount <= 0 {
		request.AllowMemberCount = 1000
	}
	b := business.GroupBusiness{
		Code:             &request.Code,
		UserId:           &request.UserId,
		Name:             request.Name,
		Avatar:           request.Avatar,
		Cover:            request.Cover,
		Introduce:        request.Introduce,
		MemberCount:      request.MemberCount,
		AllowMemberCount: request.AllowMemberCount,
		IsGlobalBanned:   &request.IsGlobalBanned,
		BannedEndAt:      request.BannedEndTime,
	}
	groupID, err := b.Create()
	if err != nil {
		return nil, err
	}
	return &proto.GroupResponse{Id: groupID}, nil
}

// Update todo
func (s *GroupServer) Update(ctx context.Context, request *proto.UpdateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

// Delete todo
func (s *GroupServer) Delete(ctx context.Context, request *proto.UpdateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

// Get todo
func (s *GroupServer) Get(ctx context.Context, request *proto.SearchGroupRequest) (*proto.GroupsResponse, error) {
	b := business.GroupBusiness{}
	models := b.List()
	var members []*proto.GroupResponse
	for _, model := range models {
		members = append(members, &proto.GroupResponse{
			Id:               model.ID,
			UserId:           model.UserID,
			Code:             model.Code,
			Name:             model.Name,
			Avatar:           model.Avatar,
			Cover:            model.Cover,
			Introduce:        model.Introduce,
			MemberCount:      model.MemberCount,
			AllowMemberCount: model.AllowMemberCount,
		})
	}

	return &proto.GroupsResponse{
		Total:  int64(len(members)),
		Groups: members,
	}, nil
}

func (s *GroupServer) Mine(ctx context.Context, request *proto.SearchGroupRequest) (*proto.GroupsResponse, error) {
	b := business.GroupBusiness{
		UserId: &request.UserId,
	}

	user, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{
		Id: request.UserId,
	})
	if err != nil {
		return nil, err
	}

	total, models := b.Mine()
	var members []*proto.GroupResponse
	for _, model := range models {
		members = append(members, &proto.GroupResponse{
			Id:               model.ID,
			UserId:           model.UserID,
			Code:             model.Code,
			Name:             model.Name,
			Avatar:           model.Avatar,
			Cover:            model.Cover,
			Introduce:        model.Introduce,
			MemberCount:      model.MemberCount,
			AllowMemberCount: model.AllowMemberCount,
			Member: &proto.GroupMemberResponse{
				Id:          model.Member.ID,
				User:        user,
				Nickname:    model.Member.Nickname,
				Role:        model.Member.Role,
				Remark:      model.Member.Remark,
				IsDnd:       model.Member.IsDnd,
				IsBanned:    model.Member.IsBanned,
				CreatedTime: model.Member.CreatedAt.Unix(),
			},
		})
	}

	return &proto.GroupsResponse{
		Total:  total,
		Groups: members,
	}, nil
}

func (s *GroupServer) Member(ctx context.Context, request *proto.SearchGroupMemberRequest) (*proto.GroupMemberResponse, error) {
	user, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{
		Id: request.UserId,
	})
	if err != nil {
		return nil, err
	}

	b := business.GroupMemberBusiness{
		GroupID: &request.GroupId,
		UserID:  &request.UserId,
	}
	member, err := b.Member()
	if err != nil {
		return nil, err
	}

	return &proto.GroupMemberResponse{
		User:        user,
		Nickname:    member.Nickname,
		Role:        member.Role,
		Level:       member.Level,
		Remark:      member.Remark,
		IsDnd:       member.IsDnd,
		IsBanned:    member.IsBanned,
		CreatedTime: member.CreatedAt.Unix(),
	}, nil
}

func (s *GroupServer) Members(ctx context.Context, request *proto.SearchGroupMemberRequest) (*proto.GroupMembersResponse, error) {
	res := &proto.GroupMembersResponse{
		Total:   0,
		Members: nil,
	}

	// 获取群信息
	gb := business.GroupBusiness{ID: &request.GroupId}
	groupEntity := gb.Detail()

	// 获取群成员
	b := business.GroupMemberBusiness{
		GroupID: &request.GroupId,
	}
	models, total := b.Members()
	if total == 0 {
		return res, nil
	}

	// 获取用户
	var userIds []int64
	for _, entity := range models {
		userIds = append(userIds, entity.UserID)
	}
	userMap := make(map[int64]*userProto.UserResponse)
	users, _ := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{Id: userIds})
	for _, user := range users.Users {
		userMap[user.Id] = user
	}

	// 定义成员响应数据
	var members []*proto.GroupMemberResponse
	for _, entity := range models {
		member := &proto.GroupMemberResponse{
			Id: entity.ID,
			Group: &proto.GroupResponse{
				Id:             groupEntity.ID,
				Code:           groupEntity.Code,
				Name:           groupEntity.Name,
				IsGlobalBanned: groupEntity.IsGlobalBanned,
			},
			User:     userMap[entity.UserID],
			Nickname: entity.Nickname,
			Role:     entity.Role,
			Level:    entity.Level,
			Remark:   entity.Remark,
			IsDnd:    entity.IsDnd,
			IsBanned: entity.IsBanned,
		}
		//member.CreatedTime = time.Parse(entity.CreatedAt, "")

		userIds = append(userIds, entity.UserID)
		members = append(members, member)
	}

	res.Total = int64(len(userIds))
	res.Members = members
	return res, nil
}

func (s *GroupServer) Join(ctx context.Context, request *proto.UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	user, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: request.UserId})
	if err != nil {
		return nil, err
	}

	role := int64(enum.GroupRoleMember)
	b := business.GroupMemberBusiness{
		GroupID: &request.GroupId,
		UserID:  &request.UserId,
		Role:    &role,
	}
	if _, err := b.Create(); err != nil {
		return nil, err
	}

	push := business.PushBusiness{}
	err = push.GroupJoin(request.GroupId, user)
	fmt.Printf("发送消息结果: %s\n", err)
	return &emptypb.Empty{}, nil
}

func (s *GroupServer) Quit(ctx context.Context, request *proto.UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	b := business.GroupMemberBusiness{
		GroupID: &request.GroupId,
		UserID:  &request.UserId,
	}
	if _, err := b.Delete(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// KickOut todo
func (s *GroupServer) KickOut(ctx context.Context, request *proto.UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "服务未实现")
}

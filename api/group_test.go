package api

import (
	proto "contact/api/qvbilam/contact/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func client() proto.GroupClient {
	host := "127.0.0.1"
	port := 9805
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(any(err))
	}
	client := proto.NewGroupClient(conn)
	return client
}

// 测试创建群
func TestGroupServer_Create(t *testing.T) {
	c := client()
	request := proto.UpdateGroupRequest{
		UserId: 3,
		Code:   534511019,
		Name:   "二滑大魔王交流群1群",
	}
	gId, err := c.Create(context.Background(), &request)
	if err != nil {
		fmt.Printf("创建失败, err: %s\n", err.Error())
		return
	}
	fmt.Printf("创建成功, 群ID: %s\n", gId)
	return
}

// 测试获取群成员
func TestGroupServer_Member(t *testing.T) {
	c := client()
	groupId := 2
	m, err := c.Member(context.Background(), &proto.SearchGroupMemberRequest{GroupId: int64(groupId)})
	if err != nil {
		fmt.Printf("获取失败, err: %s\n", err.Error())
		return
	}
	fmt.Printf("用户: %+v\n", m.Members)
	return
}

// 测试加入群
func TestGroupServer_Join(t *testing.T) {
	c := client()

	groupId := 2
	users := []*proto.UpdateGroupMemberRequest{
		{
			GroupId: int64(groupId),
			UserId:  4,
		},
		{
			GroupId: int64(groupId),
			UserId:  5,
		},
	}

	for _, user := range users {
		_, err := c.Join(context.Background(), &proto.UpdateGroupMemberRequest{
			GroupId: int64(groupId),
			UserId:  user.UserId,
		})
		if err != nil {
			fmt.Printf("用户: %d, 加群失败, err: %s\n", user.UserId, err.Error())
			continue
		}
		fmt.Printf("用户: %d, 加群成功\n", user.UserId)
	}

}

// 测试退出群
func TestGroupServer_Quit(t *testing.T) {
	c := client()

	groupId := 2
	users := []*proto.UpdateGroupMemberRequest{
		{
			GroupId: int64(groupId),
			UserId:  5,
		},
	}

	for _, user := range users {
		_, err := c.Quit(context.Background(), &proto.UpdateGroupMemberRequest{
			GroupId: int64(groupId),
			UserId:  user.UserId,
		})
		if err != nil {
			fmt.Printf("用户: %d, 退群失败, err: %s\n", user.UserId, err.Error())
			continue
		}
		fmt.Printf("用户: %d, 退群成功\n", user.UserId)
	}
}

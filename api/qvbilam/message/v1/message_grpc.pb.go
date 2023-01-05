// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: message.proto

package messageV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	CreateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*QueueResponse, error)
	UpdateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreatePrivateMessage(ctx context.Context, in *CreatePrivateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateRoomMessage(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupTxtMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupCmdMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupTipMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) CreateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*QueueResponse, error) {
	out := new(QueueResponse)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) UpdateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/UpdateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DeleteQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/DeleteQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreatePrivateMessage(ctx context.Context, in *CreatePrivateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreatePrivateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateRoomMessage(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateRoomMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateGroupMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateGroupMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateGroupTxtMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateGroupTxtMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateGroupCmdMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateGroupCmdMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateGroupTipMessage(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/messagePb.v1.Message/CreateGroupTipMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	CreateQueue(context.Context, *UpdateQueueRequest) (*QueueResponse, error)
	UpdateQueue(context.Context, *UpdateQueueRequest) (*emptypb.Empty, error)
	DeleteQueue(context.Context, *UpdateQueueRequest) (*emptypb.Empty, error)
	CreatePrivateMessage(context.Context, *CreatePrivateRequest) (*emptypb.Empty, error)
	CreateRoomMessage(context.Context, *CreateRoomRequest) (*emptypb.Empty, error)
	CreateGroupMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error)
	CreateGroupTxtMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error)
	CreateGroupCmdMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error)
	CreateGroupTipMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) CreateQueue(context.Context, *UpdateQueueRequest) (*QueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueue not implemented")
}
func (UnimplementedMessageServer) UpdateQueue(context.Context, *UpdateQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueue not implemented")
}
func (UnimplementedMessageServer) DeleteQueue(context.Context, *UpdateQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueue not implemented")
}
func (UnimplementedMessageServer) CreatePrivateMessage(context.Context, *CreatePrivateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateMessage not implemented")
}
func (UnimplementedMessageServer) CreateRoomMessage(context.Context, *CreateRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoomMessage not implemented")
}
func (UnimplementedMessageServer) CreateGroupMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupMessage not implemented")
}
func (UnimplementedMessageServer) CreateGroupTxtMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupTxtMessage not implemented")
}
func (UnimplementedMessageServer) CreateGroupCmdMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupCmdMessage not implemented")
}
func (UnimplementedMessageServer) CreateGroupTipMessage(context.Context, *CreateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupTipMessage not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_CreateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateQueue(ctx, req.(*UpdateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_UpdateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).UpdateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/UpdateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).UpdateQueue(ctx, req.(*UpdateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DeleteQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DeleteQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/DeleteQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DeleteQueue(ctx, req.(*UpdateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreatePrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreatePrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreatePrivateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreatePrivateMessage(ctx, req.(*CreatePrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateRoomMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateRoomMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateRoomMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateRoomMessage(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateGroupMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateGroupMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateGroupMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateGroupMessage(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateGroupTxtMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateGroupTxtMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateGroupTxtMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateGroupTxtMessage(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateGroupCmdMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateGroupCmdMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateGroupCmdMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateGroupCmdMessage(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateGroupTipMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateGroupTipMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messagePb.v1.Message/CreateGroupTipMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateGroupTipMessage(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messagePb.v1.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueue",
			Handler:    _Message_CreateQueue_Handler,
		},
		{
			MethodName: "UpdateQueue",
			Handler:    _Message_UpdateQueue_Handler,
		},
		{
			MethodName: "DeleteQueue",
			Handler:    _Message_DeleteQueue_Handler,
		},
		{
			MethodName: "CreatePrivateMessage",
			Handler:    _Message_CreatePrivateMessage_Handler,
		},
		{
			MethodName: "CreateRoomMessage",
			Handler:    _Message_CreateRoomMessage_Handler,
		},
		{
			MethodName: "CreateGroupMessage",
			Handler:    _Message_CreateGroupMessage_Handler,
		},
		{
			MethodName: "CreateGroupTxtMessage",
			Handler:    _Message_CreateGroupTxtMessage_Handler,
		},
		{
			MethodName: "CreateGroupCmdMessage",
			Handler:    _Message_CreateGroupCmdMessage_Handler,
		},
		{
			MethodName: "CreateGroupTipMessage",
			Handler:    _Message_CreateGroupTipMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: friend.proto

package contactV1

import (
	v1 "contact/api/qvbilam/user/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64            `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Remark string           `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Friend *v1.UserResponse `protobuf:"bytes,4,opt,name=friend,proto3" json:"friend,omitempty"`
}

func (x *FriendResponse) Reset() {
	*x = FriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendResponse) ProtoMessage() {}

func (x *FriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendResponse.ProtoReflect.Descriptor instead.
func (*FriendResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{0}
}

func (x *FriendResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FriendResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *FriendResponse) GetFriend() *v1.UserResponse {
	if x != nil {
		return x.Friend
	}
	return nil
}

type FriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Friends []*FriendResponse `protobuf:"bytes,2,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *FriendsResponse) Reset() {
	*x = FriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsResponse) ProtoMessage() {}

func (x *FriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsResponse.ProtoReflect.Descriptor instead.
func (*FriendsResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{1}
}

func (x *FriendsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FriendsResponse) GetFriends() []*FriendResponse {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UpdateFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	FriendUserId int64  `protobuf:"varint,3,opt,name=friendUserId,proto3" json:"friendUserId,omitempty"`
	Remark       string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UpdateFriendRequest) Reset() {
	*x = UpdateFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFriendRequest) ProtoMessage() {}

func (x *UpdateFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFriendRequest.ProtoReflect.Descriptor instead.
func (*UpdateFriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFriendRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFriendRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateFriendRequest) GetFriendUserId() int64 {
	if x != nil {
		return x.FriendUserId
	}
	return 0
}

func (x *UpdateFriendRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type SearchFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FriendIds []int64 `protobuf:"varint,2,rep,packed,name=friendIds,proto3" json:"friendIds,omitempty"`
	Keyword   string  `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *SearchFriendRequest) Reset() {
	*x = SearchFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFriendRequest) ProtoMessage() {}

func (x *SearchFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFriendRequest.ProtoReflect.Descriptor instead.
func (*SearchFriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{3}
}

func (x *SearchFriendRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SearchFriendRequest) GetFriendIds() []int64 {
	if x != nil {
		return x.FriendIds
	}
	return nil
}

func (x *SearchFriendRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type UpdateFriendApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ApplyUserId int64  `protobuf:"varint,3,opt,name=applyUserId,proto3" json:"applyUserId,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status      int64  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateFriendApplyRequest) Reset() {
	*x = UpdateFriendApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFriendApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFriendApplyRequest) ProtoMessage() {}

func (x *UpdateFriendApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFriendApplyRequest.ProtoReflect.Descriptor instead.
func (*UpdateFriendApplyRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFriendApplyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFriendApplyRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateFriendApplyRequest) GetApplyUserId() int64 {
	if x != nil {
		return x.ApplyUserId
	}
	return 0
}

func (x *UpdateFriendApplyRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateFriendApplyRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FriendApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64            `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ApplyUserId int64            `protobuf:"varint,3,opt,name=applyUserId,proto3" json:"applyUserId,omitempty"`
	Content     string           `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status      int64            `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	ApplyUser   *v1.UserResponse `protobuf:"bytes,6,opt,name=applyUser,proto3" json:"applyUser,omitempty"`
}

func (x *FriendApplyResponse) Reset() {
	*x = FriendApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendApplyResponse) ProtoMessage() {}

func (x *FriendApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendApplyResponse.ProtoReflect.Descriptor instead.
func (*FriendApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{5}
}

func (x *FriendApplyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendApplyResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FriendApplyResponse) GetApplyUserId() int64 {
	if x != nil {
		return x.ApplyUserId
	}
	return 0
}

func (x *FriendApplyResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FriendApplyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FriendApplyResponse) GetApplyUser() *v1.UserResponse {
	if x != nil {
		return x.ApplyUser
	}
	return nil
}

type FriendAppliesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Users []*FriendApplyResponse `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *FriendAppliesResponse) Reset() {
	*x = FriendAppliesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendAppliesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendAppliesResponse) ProtoMessage() {}

func (x *FriendAppliesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendAppliesResponse.ProtoReflect.Descriptor instead.
func (*FriendAppliesResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{6}
}

func (x *FriendAppliesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FriendAppliesResponse) GetUsers() []*FriendApplyResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_friend_proto protoreflect.FileDescriptor

var file_friend_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x71,
	0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2f, 0x0a, 0x06,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x22, 0x5f, 0x0a,
	0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x79,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x65, 0x0a, 0x13, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x96, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x13, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a,
	0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x15, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x9a, 0x04, 0x0a,
	0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x21,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x0a,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x67, 0x72, 0x65, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x0b, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_proto_rawDescOnce sync.Once
	file_friend_proto_rawDescData = file_friend_proto_rawDesc
)

func file_friend_proto_rawDescGZIP() []byte {
	file_friend_proto_rawDescOnce.Do(func() {
		file_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_proto_rawDescData)
	})
	return file_friend_proto_rawDescData
}

var file_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_friend_proto_goTypes = []interface{}{
	(*FriendResponse)(nil),           // 0: contactPb.v1.FriendResponse
	(*FriendsResponse)(nil),          // 1: contactPb.v1.FriendsResponse
	(*UpdateFriendRequest)(nil),      // 2: contactPb.v1.UpdateFriendRequest
	(*SearchFriendRequest)(nil),      // 3: contactPb.v1.SearchFriendRequest
	(*UpdateFriendApplyRequest)(nil), // 4: contactPb.v1.UpdateFriendApplyRequest
	(*FriendApplyResponse)(nil),      // 5: contactPb.v1.FriendApplyResponse
	(*FriendAppliesResponse)(nil),    // 6: contactPb.v1.FriendAppliesResponse
	(*v1.UserResponse)(nil),          // 7: userPb.v1.UserResponse
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_friend_proto_depIdxs = []int32{
	7,  // 0: contactPb.v1.FriendResponse.friend:type_name -> userPb.v1.UserResponse
	0,  // 1: contactPb.v1.FriendsResponse.friends:type_name -> contactPb.v1.FriendResponse
	7,  // 2: contactPb.v1.FriendApplyResponse.applyUser:type_name -> userPb.v1.UserResponse
	5,  // 3: contactPb.v1.FriendAppliesResponse.users:type_name -> contactPb.v1.FriendApplyResponse
	3,  // 4: contactPb.v1.friend.Get:input_type -> contactPb.v1.SearchFriendRequest
	2,  // 5: contactPb.v1.friend.Update:input_type -> contactPb.v1.UpdateFriendRequest
	2,  // 6: contactPb.v1.friend.Delete:input_type -> contactPb.v1.UpdateFriendRequest
	4,  // 7: contactPb.v1.friend.GetApply:input_type -> contactPb.v1.UpdateFriendApplyRequest
	4,  // 8: contactPb.v1.friend.Apply:input_type -> contactPb.v1.UpdateFriendApplyRequest
	4,  // 9: contactPb.v1.friend.ApplyAgree:input_type -> contactPb.v1.UpdateFriendApplyRequest
	4,  // 10: contactPb.v1.friend.ApplyReject:input_type -> contactPb.v1.UpdateFriendApplyRequest
	1,  // 11: contactPb.v1.friend.Get:output_type -> contactPb.v1.FriendsResponse
	8,  // 12: contactPb.v1.friend.Update:output_type -> google.protobuf.Empty
	8,  // 13: contactPb.v1.friend.Delete:output_type -> google.protobuf.Empty
	6,  // 14: contactPb.v1.friend.GetApply:output_type -> contactPb.v1.FriendAppliesResponse
	8,  // 15: contactPb.v1.friend.Apply:output_type -> google.protobuf.Empty
	8,  // 16: contactPb.v1.friend.ApplyAgree:output_type -> google.protobuf.Empty
	8,  // 17: contactPb.v1.friend.ApplyReject:output_type -> google.protobuf.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_friend_proto_init() }
func file_friend_proto_init() {
	if File_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFriendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFriendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFriendApplyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendApplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendAppliesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_proto_goTypes,
		DependencyIndexes: file_friend_proto_depIdxs,
		MessageInfos:      file_friend_proto_msgTypes,
	}.Build()
	File_friend_proto = out.File
	file_friend_proto_rawDesc = nil
	file_friend_proto_goTypes = nil
	file_friend_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/usersearch.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RealName             string   `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Bio                  string   `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	FollowingsCount      uint64   `protobuf:"varint,5,opt,name=followings_count,json=followingsCount,proto3" json:"followings_count,omitempty"`
	FollowersCount       uint64   `protobuf:"varint,6,opt,name=followers_count,json=followersCount,proto3" json:"followers_count,omitempty"`
	FollowingsUsers      []*User  `protobuf:"bytes,7,rep,name=followings_users,json=followingsUsers,proto3" json:"followings_users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_fec7f30762a98884, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetFollowingsCount() uint64 {
	if m != nil {
		return m.FollowingsCount
	}
	return 0
}

func (m *User) GetFollowersCount() uint64 {
	if m != nil {
		return m.FollowersCount
	}
	return 0
}

func (m *User) GetFollowingsUsers() []*User {
	if m != nil {
		return m.FollowingsUsers
	}
	return nil
}

type UserSearchResponse struct {
	UserList             []*User  `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSearchResponse) Reset()         { *m = UserSearchResponse{} }
func (m *UserSearchResponse) String() string { return proto.CompactTextString(m) }
func (*UserSearchResponse) ProtoMessage()    {}
func (*UserSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_fec7f30762a98884, []int{1}
}
func (m *UserSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchResponse.Unmarshal(m, b)
}
func (m *UserSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchResponse.Marshal(b, m, deterministic)
}
func (dst *UserSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchResponse.Merge(dst, src)
}
func (m *UserSearchResponse) XXX_Size() int {
	return xxx_messageInfo_UserSearchResponse.Size(m)
}
func (m *UserSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSearchResponse proto.InternalMessageInfo

func (m *UserSearchResponse) GetUserList() []*User {
	if m != nil {
		return m.UserList
	}
	return nil
}

type UserSearchrequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSearchrequest) Reset()         { *m = UserSearchrequest{} }
func (m *UserSearchrequest) String() string { return proto.CompactTextString(m) }
func (*UserSearchrequest) ProtoMessage()    {}
func (*UserSearchrequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersearch_fec7f30762a98884, []int{2}
}
func (m *UserSearchrequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchrequest.Unmarshal(m, b)
}
func (m *UserSearchrequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchrequest.Marshal(b, m, deterministic)
}
func (dst *UserSearchrequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchrequest.Merge(dst, src)
}
func (m *UserSearchrequest) XXX_Size() int {
	return xxx_messageInfo_UserSearchrequest.Size(m)
}
func (m *UserSearchrequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchrequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserSearchrequest proto.InternalMessageInfo

func (m *UserSearchrequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserSearchResponse)(nil), "proto.UserSearchResponse")
	proto.RegisterType((*UserSearchrequest)(nil), "proto.UserSearchrequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSearchServiceClient is the client API for UserSearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSearchServiceClient interface {
	GetAllUsersLikeUsername(ctx context.Context, in *UserSearchrequest, opts ...grpc.CallOption) (*UserSearchResponse, error)
	GetUserWithUsername(ctx context.Context, in *UserSearchrequest, opts ...grpc.CallOption) (*User, error)
}

type userSearchServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserSearchServiceClient(cc *grpc.ClientConn) UserSearchServiceClient {
	return &userSearchServiceClient{cc}
}

func (c *userSearchServiceClient) GetAllUsersLikeUsername(ctx context.Context, in *UserSearchrequest, opts ...grpc.CallOption) (*UserSearchResponse, error) {
	out := new(UserSearchResponse)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetAllUsersLikeUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSearchServiceClient) GetUserWithUsername(ctx context.Context, in *UserSearchrequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserSearchService/GetUserWithUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSearchServiceServer is the server API for UserSearchService service.
type UserSearchServiceServer interface {
	GetAllUsersLikeUsername(context.Context, *UserSearchrequest) (*UserSearchResponse, error)
	GetUserWithUsername(context.Context, *UserSearchrequest) (*User, error)
}

func RegisterUserSearchServiceServer(s *grpc.Server, srv UserSearchServiceServer) {
	s.RegisterService(&_UserSearchService_serviceDesc, srv)
}

func _UserSearchService_GetAllUsersLikeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearchrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).GetAllUsersLikeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/GetAllUsersLikeUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).GetAllUsersLikeUsername(ctx, req.(*UserSearchrequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSearchService_GetUserWithUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearchrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServiceServer).GetUserWithUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSearchService/GetUserWithUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServiceServer).GetUserWithUsername(ctx, req.(*UserSearchrequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserSearchService",
	HandlerType: (*UserSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsersLikeUsername",
			Handler:    _UserSearchService_GetAllUsersLikeUsername_Handler,
		},
		{
			MethodName: "GetUserWithUsername",
			Handler:    _UserSearchService_GetUserWithUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/usersearch.proto",
}

func init() {
	proto.RegisterFile("api/proto/usersearch.proto", fileDescriptor_usersearch_fec7f30762a98884)
}

var fileDescriptor_usersearch_fec7f30762a98884 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x99, 0xfc, 0x51, 0x3e, 0xa1, 0x6c, 0x3a, 0x34, 0x19, 0x81, 0x78, 0x69, 0xbb, 0x68,
	0x18, 0x74, 0x8a, 0x20, 0x3a, 0x78, 0x11, 0x0f, 0x8a, 0x74, 0x5c, 0xc6, 0xe5, 0x95, 0x43, 0xe3,
	0xce, 0x36, 0x33, 0x6b, 0x7f, 0x4f, 0x7f, 0x68, 0x10, 0x6f, 0xd6, 0x74, 0x45, 0x88, 0x4e, 0xfb,
	0xf6, 0xf3, 0x7d, 0x3f, 0xe6, 0xfb, 0x85, 0xb6, 0xcc, 0x54, 0x3f, 0xb3, 0xc6, 0x9b, 0x7e, 0xee,
	0xd0, 0x3a, 0x94, 0x36, 0x59, 0xf4, 0x02, 0xe0, 0xb5, 0xf0, 0xe9, 0x7e, 0x33, 0xa8, 0xce, 0x1c,
	0x5a, 0x7e, 0x09, 0x0d, 0xea, 0x89, 0x53, 0xb9, 0x44, 0xc1, 0x3a, 0x2c, 0x6a, 0x4c, 0x8e, 0x08,
	0x8c, 0xe5, 0x12, 0x49, 0xb4, 0x28, 0x75, 0x21, 0x1e, 0x14, 0x22, 0x81, 0x20, 0xb6, 0xa0, 0x32,
	0x57, 0x46, 0x54, 0x02, 0xa6, 0x92, 0x5f, 0x01, 0xc8, 0x95, 0xf4, 0xd2, 0xc6, 0xb9, 0xd5, 0xa2,
	0x1a, 0x84, 0x46, 0x41, 0x66, 0x56, 0xf3, 0x1b, 0x68, 0xbd, 0x1a, 0xad, 0xcd, 0xa7, 0x4a, 0xdf,
	0x5c, 0x9c, 0x98, 0x3c, 0xf5, 0xa2, 0xd6, 0x61, 0x51, 0x75, 0x72, 0xb2, 0xe5, 0xcf, 0x84, 0xf9,
	0x35, 0xac, 0x11, 0xda, 0xdf, 0xce, 0x7a, 0xe8, 0x3c, 0xde, 0xe0, 0xa2, 0xf1, 0x7e, 0x67, 0x67,
	0x70, 0x2b, 0x0e, 0x3b, 0x95, 0xa8, 0x39, 0x68, 0x16, 0x86, 0x7b, 0xe4, 0xb2, 0x7c, 0x80, 0xfe,
	0x5d, 0xf7, 0x11, 0x38, 0x15, 0xd3, 0x10, 0xcd, 0x04, 0x5d, 0x66, 0x52, 0x87, 0x3c, 0x5a, 0x87,
	0xa1, 0x95, 0xf3, 0x82, 0xed, 0xaf, 0x09, 0xc9, 0x8c, 0x94, 0xf3, 0xdd, 0x5b, 0x38, 0xdd, 0xce,
	0x5b, 0xfc, 0xc8, 0xd1, 0xf9, 0x3f, 0xb3, 0x1c, 0x7c, 0xb1, 0xf2, 0xc8, 0x14, 0xed, 0x4a, 0x25,
	0xc8, 0xc7, 0x70, 0x3e, 0x44, 0xff, 0xa4, 0x75, 0x78, 0xd6, 0x48, 0xbd, 0x23, 0x15, 0xb4, 0x80,
	0x8b, 0xd2, 0xe5, 0x9d, 0x3b, 0xed, 0x8b, 0x3d, 0x65, 0xe3, 0xe0, 0x01, 0xce, 0x86, 0xe8, 0x49,
	0x78, 0x51, 0x7e, 0xf1, 0x8f, 0x5d, 0x65, 0x7f, 0xf3, 0x7a, 0xa8, 0xef, 0x7e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x9f, 0xf9, 0xd2, 0x41, 0x02, 0x00, 0x00,
}

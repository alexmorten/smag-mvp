// Code generated by protoc-gen-go. DO NOT EDIT.
// source: renewingAddress.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RenewedElasticResult struct {
	IsRenewed            bool     `protobuf:"varint,1,opt,name=isRenewed,proto3" json:"isRenewed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewedElasticResult) Reset()         { *m = RenewedElasticResult{} }
func (m *RenewedElasticResult) String() string { return proto.CompactTextString(m) }
func (*RenewedElasticResult) ProtoMessage()    {}
func (*RenewedElasticResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_55c265b1376e8420, []int{0}
}

func (m *RenewedElasticResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewedElasticResult.Unmarshal(m, b)
}
func (m *RenewedElasticResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewedElasticResult.Marshal(b, m, deterministic)
}
func (m *RenewedElasticResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewedElasticResult.Merge(m, src)
}
func (m *RenewedElasticResult) XXX_Size() int {
	return xxx_messageInfo_RenewedElasticResult.Size(m)
}
func (m *RenewedElasticResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewedElasticResult.DiscardUnknown(m)
}

var xxx_messageInfo_RenewedElasticResult proto.InternalMessageInfo

func (m *RenewedElasticResult) GetIsRenewed() bool {
	if m != nil {
		return m.IsRenewed
	}
	return false
}

type RenewingElasticIp struct {
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Pod                  string   `protobuf:"bytes,3,opt,name=pod,proto3" json:"pod,omitempty"`
	PodIp                string   `protobuf:"bytes,4,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewingElasticIp) Reset()         { *m = RenewingElasticIp{} }
func (m *RenewingElasticIp) String() string { return proto.CompactTextString(m) }
func (*RenewingElasticIp) ProtoMessage()    {}
func (*RenewingElasticIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_55c265b1376e8420, []int{1}
}

func (m *RenewingElasticIp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewingElasticIp.Unmarshal(m, b)
}
func (m *RenewingElasticIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewingElasticIp.Marshal(b, m, deterministic)
}
func (m *RenewingElasticIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewingElasticIp.Merge(m, src)
}
func (m *RenewingElasticIp) XXX_Size() int {
	return xxx_messageInfo_RenewingElasticIp.Size(m)
}
func (m *RenewingElasticIp) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewingElasticIp.DiscardUnknown(m)
}

var xxx_messageInfo_RenewingElasticIp proto.InternalMessageInfo

func (m *RenewingElasticIp) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *RenewingElasticIp) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *RenewingElasticIp) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *RenewingElasticIp) GetPodIp() string {
	if m != nil {
		return m.PodIp
	}
	return ""
}

func init() {
	proto.RegisterType((*RenewedElasticResult)(nil), "proto.renewedElasticResult")
	proto.RegisterType((*RenewingElasticIp)(nil), "proto.RenewingElasticIp")
}

func init() { proto.RegisterFile("renewingAddress.proto", fileDescriptor_55c265b1376e8420) }

var fileDescriptor_55c265b1376e8420 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3f, 0x4b, 0xc7, 0x30,
	0x10, 0x86, 0xad, 0xfd, 0x83, 0xbd, 0x41, 0xea, 0x61, 0x21, 0xa8, 0x88, 0x74, 0x72, 0xea, 0xa0,
	0x7e, 0x01, 0x07, 0x87, 0xac, 0x71, 0x16, 0xa9, 0xcd, 0x21, 0x81, 0x92, 0x1c, 0x49, 0xd4, 0xaf,
	0x2f, 0x5e, 0x8b, 0x15, 0xfc, 0x4d, 0x39, 0x9e, 0x87, 0xcb, 0xbd, 0x2f, 0xf4, 0x91, 0x3c, 0x7d,
	0x39, 0xff, 0xfe, 0x68, 0x6d, 0xa4, 0x94, 0x46, 0x8e, 0x21, 0x07, 0xac, 0xe5, 0x19, 0x1e, 0xe0,
	0x5c, 0x3c, 0xd9, 0xa7, 0x65, 0x4a, 0xd9, 0xcd, 0x86, 0xd2, 0xc7, 0x92, 0xf1, 0x0a, 0x5a, 0x97,
	0xcc, 0x6a, 0x54, 0x71, 0x53, 0xdc, 0x9e, 0x98, 0x1d, 0x0c, 0x0c, 0x67, 0x66, 0xfb, 0x75, 0x5b,
	0xd3, 0x8c, 0xd7, 0x00, 0xce, 0xa7, 0x3c, 0xf9, 0x99, 0xf4, 0xba, 0xd3, 0x9a, 0x3f, 0x04, 0x11,
	0x2a, 0x1f, 0x2c, 0xa9, 0x63, 0x31, 0x32, 0x63, 0x07, 0x25, 0x07, 0xab, 0x4a, 0x41, 0x3f, 0x23,
	0xf6, 0xd0, 0x70, 0xb0, 0xaf, 0x8e, 0x55, 0x25, 0xb0, 0xe6, 0x60, 0x35, 0xdf, 0xbd, 0x40, 0xf7,
	0x7b, 0xe9, 0x99, 0xe2, 0xa7, 0x9b, 0x09, 0x35, 0x9c, 0x4a, 0xf6, 0x3d, 0x82, 0x5a, 0xcb, 0x8d,
	0xff, 0xc2, 0x5d, 0x5c, 0x6e, 0xe6, 0x50, 0xd9, 0xe1, 0xe8, 0xad, 0x11, 0x7b, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x77, 0xd5, 0x13, 0x1b, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ElasticIpServiceClient is the client API for ElasticIpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElasticIpServiceClient interface {
	RenewElasticIp(ctx context.Context, in *RenewingElasticIp, opts ...grpc.CallOption) (*RenewedElasticResult, error)
}

type elasticIpServiceClient struct {
	cc *grpc.ClientConn
}

func NewElasticIpServiceClient(cc *grpc.ClientConn) ElasticIpServiceClient {
	return &elasticIpServiceClient{cc}
}

func (c *elasticIpServiceClient) RenewElasticIp(ctx context.Context, in *RenewingElasticIp, opts ...grpc.CallOption) (*RenewedElasticResult, error) {
	out := new(RenewedElasticResult)
	err := c.cc.Invoke(ctx, "/proto.ElasticIpService/renewElasticIp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticIpServiceServer is the server API for ElasticIpService service.
type ElasticIpServiceServer interface {
	RenewElasticIp(context.Context, *RenewingElasticIp) (*RenewedElasticResult, error)
}

// UnimplementedElasticIpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedElasticIpServiceServer struct {
}

func (*UnimplementedElasticIpServiceServer) RenewElasticIp(ctx context.Context, req *RenewingElasticIp) (*RenewedElasticResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewElasticIp not implemented")
}

func RegisterElasticIpServiceServer(s *grpc.Server, srv ElasticIpServiceServer) {
	s.RegisterService(&_ElasticIpService_serviceDesc, srv)
}

func _ElasticIpService_RenewElasticIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewingElasticIp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticIpServiceServer).RenewElasticIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ElasticIpService/RenewElasticIp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticIpServiceServer).RenewElasticIp(ctx, req.(*RenewingElasticIp))
	}
	return interceptor(ctx, in, info, handler)
}

var _ElasticIpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ElasticIpService",
	HandlerType: (*ElasticIpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "renewElasticIp",
			Handler:    _ElasticIpService_RenewElasticIp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "renewingAddress.proto",
}

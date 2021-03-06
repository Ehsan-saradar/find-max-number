// Code generated by protoc-gen-go. DO NOT EDIT.
// source: math.proto

package math_grpc

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

type FindMaxNumberRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxNumberRequest) Reset()         { *m = FindMaxNumberRequest{} }
func (m *FindMaxNumberRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxNumberRequest) ProtoMessage()    {}
func (*FindMaxNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{0}
}

func (m *FindMaxNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxNumberRequest.Unmarshal(m, b)
}
func (m *FindMaxNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxNumberRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxNumberRequest.Merge(m, src)
}
func (m *FindMaxNumberRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxNumberRequest.Size(m)
}
func (m *FindMaxNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxNumberRequest proto.InternalMessageInfo

func (m *FindMaxNumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *FindMaxNumberRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type FindMaxNumberResponse struct {
	MaxNumber            int32    `protobuf:"varint,1,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxNumberResponse) Reset()         { *m = FindMaxNumberResponse{} }
func (m *FindMaxNumberResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxNumberResponse) ProtoMessage()    {}
func (*FindMaxNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{1}
}

func (m *FindMaxNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxNumberResponse.Unmarshal(m, b)
}
func (m *FindMaxNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxNumberResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxNumberResponse.Merge(m, src)
}
func (m *FindMaxNumberResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxNumberResponse.Size(m)
}
func (m *FindMaxNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxNumberResponse proto.InternalMessageInfo

func (m *FindMaxNumberResponse) GetMaxNumber() int32 {
	if m != nil {
		return m.MaxNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*FindMaxNumberRequest)(nil), "math_grpc.FindMaxNumberRequest")
	proto.RegisterType((*FindMaxNumberResponse)(nil), "math_grpc.FindMaxNumberResponse")
}

func init() { proto.RegisterFile("math.proto", fileDescriptor_f139a3799a86a974) }

var fileDescriptor_f139a3799a86a974 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xe3, 0xd3, 0x8b, 0x0a, 0x92, 0x95,
	0x7c, 0xb8, 0x44, 0xdc, 0x32, 0xf3, 0x52, 0x7c, 0x13, 0x2b, 0xfc, 0x4a, 0x73, 0x93, 0x52, 0x8b,
	0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0xc0, 0x02, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x0c, 0x17, 0x67, 0x71, 0x66, 0x7a, 0x5e, 0x62,
	0x49, 0x69, 0x51, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x42, 0x40, 0xc9, 0x94, 0x4b,
	0x14, 0xcd, 0xb4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0xb6, 0x5c, 0x98, 0x20, 0xd4, 0x44,
	0x84, 0x80, 0x51, 0x3a, 0x17, 0xb7, 0x6f, 0x62, 0x49, 0x86, 0x47, 0x62, 0x5e, 0x4a, 0x4e, 0x6a,
	0x91, 0x50, 0x04, 0x17, 0x2f, 0x8a, 0x29, 0x42, 0xf2, 0x7a, 0x70, 0x07, 0xeb, 0x61, 0x73, 0xad,
	0x94, 0x02, 0x6e, 0x05, 0x10, 0x07, 0x28, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xfd,
	0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0x19, 0xb0, 0x7d, 0x0d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathHandlerClient is the client API for MathHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathHandlerClient interface {
	FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (MathHandler_FindMaxNumberClient, error)
}

type mathHandlerClient struct {
	cc *grpc.ClientConn
}

func NewMathHandlerClient(cc *grpc.ClientConn) MathHandlerClient {
	return &mathHandlerClient{cc}
}

func (c *mathHandlerClient) FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (MathHandler_FindMaxNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MathHandler_serviceDesc.Streams[0], "/math_grpc.MathHandler/FindMaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathHandlerFindMaxNumberClient{stream}
	return x, nil
}

type MathHandler_FindMaxNumberClient interface {
	Send(*FindMaxNumberRequest) error
	Recv() (*FindMaxNumberResponse, error)
	grpc.ClientStream
}

type mathHandlerFindMaxNumberClient struct {
	grpc.ClientStream
}

func (x *mathHandlerFindMaxNumberClient) Send(m *FindMaxNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mathHandlerFindMaxNumberClient) Recv() (*FindMaxNumberResponse, error) {
	m := new(FindMaxNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MathHandlerServer is the server API for MathHandler service.
type MathHandlerServer interface {
	FindMaxNumber(MathHandler_FindMaxNumberServer) error
}

// UnimplementedMathHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedMathHandlerServer struct {
}

func (*UnimplementedMathHandlerServer) FindMaxNumber(srv MathHandler_FindMaxNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaxNumber not implemented")
}

func RegisterMathHandlerServer(s *grpc.Server, srv MathHandlerServer) {
	s.RegisterService(&_MathHandler_serviceDesc, srv)
}

func _MathHandler_FindMaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MathHandlerServer).FindMaxNumber(&mathHandlerFindMaxNumberServer{stream})
}

type MathHandler_FindMaxNumberServer interface {
	Send(*FindMaxNumberResponse) error
	Recv() (*FindMaxNumberRequest, error)
	grpc.ServerStream
}

type mathHandlerFindMaxNumberServer struct {
	grpc.ServerStream
}

func (x *mathHandlerFindMaxNumberServer) Send(m *FindMaxNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mathHandlerFindMaxNumberServer) Recv() (*FindMaxNumberRequest, error) {
	m := new(FindMaxNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MathHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math_grpc.MathHandler",
	HandlerType: (*MathHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaxNumber",
			Handler:       _MathHandler_FindMaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "math.proto",
}

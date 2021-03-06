// Code generated by protoc-gen-go.
// source: gate.proto
// DO NOT EDIT!

package pb

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

type LoginRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Uid   int64  `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type LoginRsp struct {
	Uid    int64  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Result int32  `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
}

func (m *LoginRsp) Reset()                    { *m = LoginRsp{} }
func (m *LoginRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()               {}
func (*LoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *LoginRsp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRsp) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type RawMsg struct {
	MsgId   string `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	MsgData []byte `protobuf:"bytes,2,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (m *RawMsg) Reset()                    { *m = RawMsg{} }
func (m *RawMsg) String() string            { return proto.CompactTextString(m) }
func (*RawMsg) ProtoMessage()               {}
func (*RawMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RawMsg) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *RawMsg) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type PushMsg struct {
	Msg *RawMsg `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *PushMsg) Reset()                    { *m = PushMsg{} }
func (m *PushMsg) String() string            { return proto.CompactTextString(m) }
func (*PushMsg) ProtoMessage()               {}
func (*PushMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *PushMsg) GetMsg() *RawMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

type HeartBeatMsg struct {
	Ticker int64 `protobuf:"varint,1,opt,name=ticker" json:"ticker,omitempty"`
}

func (m *HeartBeatMsg) Reset()                    { *m = HeartBeatMsg{} }
func (m *HeartBeatMsg) String() string            { return proto.CompactTextString(m) }
func (*HeartBeatMsg) ProtoMessage()               {}
func (*HeartBeatMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *HeartBeatMsg) GetTicker() int64 {
	if m != nil {
		return m.Ticker
	}
	return 0
}

type KickAgentReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *KickAgentReq) Reset()                    { *m = KickAgentReq{} }
func (m *KickAgentReq) String() string            { return proto.CompactTextString(m) }
func (*KickAgentReq) ProtoMessage()               {}
func (*KickAgentReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *KickAgentReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type PushClientReq struct {
	Uid  int64   `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Data *RawMsg `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *PushClientReq) Reset()                    { *m = PushClientReq{} }
func (m *PushClientReq) String() string            { return proto.CompactTextString(m) }
func (*PushClientReq) ProtoMessage()               {}
func (*PushClientReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PushClientReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PushClientReq) GetData() *RawMsg {
	if m != nil {
		return m.Data
	}
	return nil
}

type BroadcastClientReq struct {
	Data *RawMsg `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *BroadcastClientReq) Reset()                    { *m = BroadcastClientReq{} }
func (m *BroadcastClientReq) String() string            { return proto.CompactTextString(m) }
func (*BroadcastClientReq) ProtoMessage()               {}
func (*BroadcastClientReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BroadcastClientReq) GetData() *RawMsg {
	if m != nil {
		return m.Data
	}
	return nil
}

type Rsp struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *Rsp) Reset()                    { *m = Rsp{} }
func (m *Rsp) String() string            { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()               {}
func (*Rsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *Rsp) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginRsp)(nil), "pb.LoginRsp")
	proto.RegisterType((*RawMsg)(nil), "pb.RawMsg")
	proto.RegisterType((*PushMsg)(nil), "pb.PushMsg")
	proto.RegisterType((*HeartBeatMsg)(nil), "pb.HeartBeatMsg")
	proto.RegisterType((*KickAgentReq)(nil), "pb.KickAgentReq")
	proto.RegisterType((*PushClientReq)(nil), "pb.PushClientReq")
	proto.RegisterType((*BroadcastClientReq)(nil), "pb.BroadcastClientReq")
	proto.RegisterType((*Rsp)(nil), "pb.Rsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoginService service

type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (LoginService_LoginClient, error)
	HeartBeat(ctx context.Context, in *HeartBeatMsg, opts ...grpc.CallOption) (*HeartBeatMsg, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (LoginService_LoginClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_LoginService_serviceDesc.Streams[0], c.cc, "/pb.LoginService/Login", opts...)
	if err != nil {
		return nil, err
	}
	x := &loginServiceLoginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoginService_LoginClient interface {
	Recv() (*PushMsg, error)
	grpc.ClientStream
}

type loginServiceLoginClient struct {
	grpc.ClientStream
}

func (x *loginServiceLoginClient) Recv() (*PushMsg, error) {
	m := new(PushMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loginServiceClient) HeartBeat(ctx context.Context, in *HeartBeatMsg, opts ...grpc.CallOption) (*HeartBeatMsg, error) {
	out := new(HeartBeatMsg)
	err := grpc.Invoke(ctx, "/pb.LoginService/HeartBeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginService service

type LoginServiceServer interface {
	Login(*LoginRequest, LoginService_LoginServer) error
	HeartBeat(context.Context, *HeartBeatMsg) (*HeartBeatMsg, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoginServiceServer).Login(m, &loginServiceLoginServer{stream})
}

type LoginService_LoginServer interface {
	Send(*PushMsg) error
	grpc.ServerStream
}

type loginServiceLoginServer struct {
	grpc.ServerStream
}

func (x *loginServiceLoginServer) Send(m *PushMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _LoginService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoginService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).HeartBeat(ctx, req.(*HeartBeatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _LoginService_HeartBeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Login",
			Handler:       _LoginService_Login_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gate.proto",
}

// Client API for GateService service

type GateServiceClient interface {
	KickAgent(ctx context.Context, in *KickAgentReq, opts ...grpc.CallOption) (*Rsp, error)
	PushClient(ctx context.Context, in *PushClientReq, opts ...grpc.CallOption) (*Rsp, error)
	BroadcastClient(ctx context.Context, in *BroadcastClientReq, opts ...grpc.CallOption) (*Rsp, error)
}

type gateServiceClient struct {
	cc *grpc.ClientConn
}

func NewGateServiceClient(cc *grpc.ClientConn) GateServiceClient {
	return &gateServiceClient{cc}
}

func (c *gateServiceClient) KickAgent(ctx context.Context, in *KickAgentReq, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := grpc.Invoke(ctx, "/pb.GateService/KickAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateServiceClient) PushClient(ctx context.Context, in *PushClientReq, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := grpc.Invoke(ctx, "/pb.GateService/PushClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateServiceClient) BroadcastClient(ctx context.Context, in *BroadcastClientReq, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := grpc.Invoke(ctx, "/pb.GateService/BroadcastClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GateService service

type GateServiceServer interface {
	KickAgent(context.Context, *KickAgentReq) (*Rsp, error)
	PushClient(context.Context, *PushClientReq) (*Rsp, error)
	BroadcastClient(context.Context, *BroadcastClientReq) (*Rsp, error)
}

func RegisterGateServiceServer(s *grpc.Server, srv GateServiceServer) {
	s.RegisterService(&_GateService_serviceDesc, srv)
}

func _GateService_KickAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickAgentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).KickAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GateService/KickAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).KickAgent(ctx, req.(*KickAgentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateService_PushClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).PushClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GateService/PushClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).PushClient(ctx, req.(*PushClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateService_BroadcastClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).BroadcastClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GateService/BroadcastClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).BroadcastClient(ctx, req.(*BroadcastClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GateService",
	HandlerType: (*GateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KickAgent",
			Handler:    _GateService_KickAgent_Handler,
		},
		{
			MethodName: "PushClient",
			Handler:    _GateService_PushClient_Handler,
		},
		{
			MethodName: "BroadcastClient",
			Handler:    _GateService_BroadcastClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gate.proto",
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xdb, 0x6e, 0xe2, 0x30,
	0x10, 0x86, 0x31, 0x81, 0xb0, 0x19, 0xb2, 0x5a, 0xd6, 0x5a, 0xa1, 0x08, 0xed, 0xae, 0x22, 0x5f,
	0xec, 0x46, 0x5c, 0xa0, 0x96, 0xa2, 0xaa, 0xb7, 0xd0, 0x4a, 0xa5, 0x6a, 0x2b, 0x55, 0xee, 0x13,
	0x98, 0xc4, 0x4a, 0x23, 0xc8, 0x81, 0xd8, 0x69, 0x9f, 0xa5, 0x6f, 0x5b, 0xd9, 0x39, 0x90, 0x16,
	0x71, 0xe7, 0x7f, 0x66, 0xfe, 0x99, 0xf1, 0x67, 0x03, 0x84, 0x4c, 0xf2, 0x59, 0x96, 0xa7, 0x32,
	0xc5, 0xdd, 0x6c, 0x43, 0x2e, 0xc1, 0x7e, 0x48, 0xc3, 0x28, 0xa1, 0x7c, 0x5f, 0x70, 0x21, 0xf1,
	0x2f, 0xe8, 0xcb, 0x74, 0xcb, 0x13, 0x07, 0xb9, 0xc8, 0xb3, 0x68, 0x29, 0xf0, 0x08, 0x8c, 0x22,
	0x0a, 0x9c, 0xae, 0x8b, 0x3c, 0x83, 0xaa, 0x23, 0x59, 0xc3, 0xb7, 0xd2, 0x27, 0xb2, 0x3a, 0x8b,
	0x9a, 0x2c, 0xc6, 0xd0, 0x4b, 0x58, 0xcc, 0xb5, 0xc1, 0xa2, 0xfa, 0x8c, 0xc7, 0x60, 0xe6, 0x5c,
	0x14, 0x3b, 0xe9, 0x18, 0x2e, 0xf2, 0xfa, 0xb4, 0x52, 0xe4, 0x0a, 0x4c, 0xca, 0xde, 0x1e, 0x45,
	0xa8, 0x66, 0xc7, 0x22, 0xbc, 0x0b, 0xea, 0xd9, 0x5a, 0x60, 0x07, 0x06, 0xb1, 0x08, 0x6f, 0x98,
	0x64, 0xba, 0x9d, 0x4d, 0x6b, 0x49, 0xfe, 0xc3, 0xe0, 0xa9, 0x10, 0x2f, 0xca, 0xfa, 0x1b, 0x8c,
	0x58, 0x84, 0xda, 0x38, 0x9c, 0xc3, 0x2c, 0xdb, 0xcc, 0xca, 0x9e, 0x54, 0x85, 0xc9, 0x3f, 0xb0,
	0xd7, 0x9c, 0xe5, 0x72, 0xc5, 0x99, 0x54, 0xd5, 0x63, 0x30, 0x65, 0xe4, 0x6f, 0x79, 0x5e, 0xed,
	0x5c, 0x29, 0xe2, 0x82, 0x7d, 0x1f, 0xf9, 0xdb, 0x65, 0xc8, 0x13, 0x49, 0xf9, 0xfe, 0xf8, 0x62,
	0x64, 0x09, 0xdf, 0xd5, 0xc8, 0xeb, 0x5d, 0x74, 0xaa, 0x04, 0xff, 0x85, 0x5e, 0x50, 0x2f, 0xfb,
	0x79, 0x17, 0x1d, 0x27, 0x0b, 0xc0, 0xab, 0x3c, 0x65, 0x81, 0xcf, 0x84, 0x3c, 0xf4, 0xa9, 0x5d,
	0xe8, 0x84, 0xeb, 0x0f, 0x18, 0x0a, 0xf5, 0x01, 0x22, 0x6a, 0x43, 0x9c, 0xc7, 0xd5, 0x33, 0x3e,
	0xf3, 0xfc, 0x35, 0xf2, 0x39, 0x9e, 0x42, 0x5f, 0x6b, 0x3c, 0x52, 0x9d, 0xda, 0x2f, 0x3c, 0x19,
	0xaa, 0x48, 0xc5, 0x8d, 0x74, 0xce, 0x10, 0x3e, 0x07, 0xab, 0xa1, 0x53, 0xd6, 0xb7, 0x61, 0x4d,
	0x8e, 0x22, 0xa4, 0x33, 0x7f, 0x47, 0x30, 0xbc, 0x65, 0x92, 0xd7, 0xe3, 0x3c, 0xb0, 0x1a, 0x70,
	0x65, 0x8b, 0x36, 0xc7, 0xc9, 0x40, 0x5f, 0x47, 0x64, 0xa4, 0x83, 0xa7, 0x00, 0x07, 0x80, 0xf8,
	0x67, 0xbd, 0x4b, 0x03, 0xa2, 0x5d, 0xbb, 0x80, 0x1f, 0x5f, 0x48, 0xe1, 0xb1, 0xca, 0x1e, 0xe3,
	0x6b, 0xb9, 0x36, 0xa6, 0xfe, 0xdc, 0x17, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x85, 0xc7,
	0xf6, 0xea, 0x02, 0x00, 0x00,
}

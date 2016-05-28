// Code generated by protoc-gen-go.
// source: bfdatafeed.proto
// DO NOT EDIT!

/*
Package bftrader is a generated protocol buffer package.

package bftrader.bfdatafeed;

It is generated from these files:
	bfdatafeed.proto

It has these top-level messages:
	BfBarData
	BfGetTickReq
	BfGetBarReq
	BfDatafeedGetContractReq
*/
package bftrader

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bftrader1 "."

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
const _ = proto.ProtoPackageIsVersion1

// Bar周期类型
type BfBarPeriod int32

const (
	BfBarPeriod_PERIOD_UNKNOWN BfBarPeriod = 0
	BfBarPeriod_PERIOD_S1      BfBarPeriod = 1
	BfBarPeriod_PERIOD_S3      BfBarPeriod = 2
	BfBarPeriod_PERIOD_S5      BfBarPeriod = 3
	BfBarPeriod_PERIOD_S10     BfBarPeriod = 4
	BfBarPeriod_PERIOD_S15     BfBarPeriod = 5
	BfBarPeriod_PERIOD_S30     BfBarPeriod = 6
	BfBarPeriod_PERIOD_M1      BfBarPeriod = 7
	BfBarPeriod_PERIOD_M3      BfBarPeriod = 8
	BfBarPeriod_PERIOD_M5      BfBarPeriod = 9
	BfBarPeriod_PERIOD_M10     BfBarPeriod = 10
	BfBarPeriod_PERIOD_M15     BfBarPeriod = 11
	BfBarPeriod_PERIOD_M30     BfBarPeriod = 12
	BfBarPeriod_PERIOD_H1      BfBarPeriod = 13
	BfBarPeriod_PERIOD_D1      BfBarPeriod = 14
	BfBarPeriod_PERIOD_W1      BfBarPeriod = 15
)

var BfBarPeriod_name = map[int32]string{
	0:  "PERIOD_UNKNOWN",
	1:  "PERIOD_S1",
	2:  "PERIOD_S3",
	3:  "PERIOD_S5",
	4:  "PERIOD_S10",
	5:  "PERIOD_S15",
	6:  "PERIOD_S30",
	7:  "PERIOD_M1",
	8:  "PERIOD_M3",
	9:  "PERIOD_M5",
	10: "PERIOD_M10",
	11: "PERIOD_M15",
	12: "PERIOD_M30",
	13: "PERIOD_H1",
	14: "PERIOD_D1",
	15: "PERIOD_W1",
}
var BfBarPeriod_value = map[string]int32{
	"PERIOD_UNKNOWN": 0,
	"PERIOD_S1":      1,
	"PERIOD_S3":      2,
	"PERIOD_S5":      3,
	"PERIOD_S10":     4,
	"PERIOD_S15":     5,
	"PERIOD_S30":     6,
	"PERIOD_M1":      7,
	"PERIOD_M3":      8,
	"PERIOD_M5":      9,
	"PERIOD_M10":     10,
	"PERIOD_M15":     11,
	"PERIOD_M30":     12,
	"PERIOD_H1":      13,
	"PERIOD_D1":      14,
	"PERIOD_W1":      15,
}

func (x BfBarPeriod) String() string {
	return proto.EnumName(BfBarPeriod_name, int32(x))
}
func (BfBarPeriod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Bar行情数据类
type BfBarData struct {
	// 代码相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 周期
	Period BfBarPeriod `protobuf:"varint,3,opt,name=period,enum=bftrader.BfBarPeriod" json:"period,omitempty"`
	// 成交数据
	ActionDate   string  `protobuf:"bytes,4,opt,name=actionDate" json:"actionDate,omitempty"`
	BarTime      string  `protobuf:"bytes,5,opt,name=barTime" json:"barTime,omitempty"`
	Volume       int32   `protobuf:"varint,6,opt,name=volume" json:"volume,omitempty"`
	OpenInterest float64 `protobuf:"fixed64,7,opt,name=openInterest" json:"openInterest,omitempty"`
	LastVolume   int32   `protobuf:"varint,8,opt,name=lastVolume" json:"lastVolume,omitempty"`
	// OHLC
	OpenPrice  float64 `protobuf:"fixed64,9,opt,name=openPrice" json:"openPrice,omitempty"`
	HighPrice  float64 `protobuf:"fixed64,10,opt,name=highPrice" json:"highPrice,omitempty"`
	LowPrice   float64 `protobuf:"fixed64,11,opt,name=lowPrice" json:"lowPrice,omitempty"`
	ClosePrice float64 `protobuf:"fixed64,12,opt,name=closePrice" json:"closePrice,omitempty"`
}

func (m *BfBarData) Reset()                    { *m = BfBarData{} }
func (m *BfBarData) String() string            { return proto.CompactTextString(m) }
func (*BfBarData) ProtoMessage()               {}
func (*BfBarData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BfGetTickReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	ToDate   string `protobuf:"bytes,3,opt,name=toDate" json:"toDate,omitempty"`
	ToTime   string `protobuf:"bytes,4,opt,name=toTime" json:"toTime,omitempty"`
	Count    int32  `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
}

func (m *BfGetTickReq) Reset()                    { *m = BfGetTickReq{} }
func (m *BfGetTickReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetTickReq) ProtoMessage()               {}
func (*BfGetTickReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BfGetBarReq struct {
	// 代码编号相关
	Symbol   string      `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string      `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Period   BfBarPeriod `protobuf:"varint,3,opt,name=period,enum=bftrader.BfBarPeriod" json:"period,omitempty"`
	ToDate   string      `protobuf:"bytes,4,opt,name=toDate" json:"toDate,omitempty"`
	ToTime   string      `protobuf:"bytes,5,opt,name=toTime" json:"toTime,omitempty"`
	Count    int32       `protobuf:"varint,6,opt,name=count" json:"count,omitempty"`
}

func (m *BfGetBarReq) Reset()                    { *m = BfGetBarReq{} }
func (m *BfGetBarReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetBarReq) ProtoMessage()               {}
func (*BfGetBarReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BfDatafeedGetContractReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
}

func (m *BfDatafeedGetContractReq) Reset()                    { *m = BfDatafeedGetContractReq{} }
func (m *BfDatafeedGetContractReq) String() string            { return proto.CompactTextString(m) }
func (*BfDatafeedGetContractReq) ProtoMessage()               {}
func (*BfDatafeedGetContractReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*BfBarData)(nil), "bftrader.BfBarData")
	proto.RegisterType((*BfGetTickReq)(nil), "bftrader.BfGetTickReq")
	proto.RegisterType((*BfGetBarReq)(nil), "bftrader.BfGetBarReq")
	proto.RegisterType((*BfDatafeedGetContractReq)(nil), "bftrader.BfDatafeedGetContractReq")
	proto.RegisterEnum("bftrader.BfBarPeriod", BfBarPeriod_name, BfBarPeriod_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for BfDatafeedService service

type BfDatafeedServiceClient interface {
	// 活跃检测
	Ping(ctx context.Context, in *bftrader1.BfPingData, opts ...grpc.CallOption) (*bftrader1.BfPingData, error)
	// 保存tick
	InsertTick(ctx context.Context, in *bftrader1.BfTickData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error)
	// 保存bar
	InsertBar(ctx context.Context, in *BfBarData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error)
	// 保存contract
	InsertContract(ctx context.Context, in *bftrader1.BfContractData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error)
	// 获取tick
	GetTick(ctx context.Context, in *BfGetTickReq, opts ...grpc.CallOption) (BfDatafeedService_GetTickClient, error)
	// 获取bar
	GetBar(ctx context.Context, in *BfGetBarReq, opts ...grpc.CallOption) (BfDatafeedService_GetBarClient, error)
	// 获取contract
	GetContract(ctx context.Context, in *BfDatafeedGetContractReq, opts ...grpc.CallOption) (BfDatafeedService_GetContractClient, error)
}

type bfDatafeedServiceClient struct {
	cc *grpc.ClientConn
}

func NewBfDatafeedServiceClient(cc *grpc.ClientConn) BfDatafeedServiceClient {
	return &bfDatafeedServiceClient{cc}
}

func (c *bfDatafeedServiceClient) Ping(ctx context.Context, in *bftrader1.BfPingData, opts ...grpc.CallOption) (*bftrader1.BfPingData, error) {
	out := new(bftrader1.BfPingData)
	err := grpc.Invoke(ctx, "/bftrader.BfDatafeedService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertTick(ctx context.Context, in *bftrader1.BfTickData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error) {
	out := new(bftrader1.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.BfDatafeedService/InsertTick", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertBar(ctx context.Context, in *BfBarData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error) {
	out := new(bftrader1.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.BfDatafeedService/InsertBar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertContract(ctx context.Context, in *bftrader1.BfContractData, opts ...grpc.CallOption) (*bftrader1.BfVoid, error) {
	out := new(bftrader1.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.BfDatafeedService/InsertContract", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) GetTick(ctx context.Context, in *BfGetTickReq, opts ...grpc.CallOption) (BfDatafeedService_GetTickClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[0], c.cc, "/bftrader.BfDatafeedService/GetTick", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetTickClient interface {
	Recv() (*bftrader1.BfTickData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetTickClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetTickClient) Recv() (*bftrader1.BfTickData, error) {
	m := new(bftrader1.BfTickData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfDatafeedServiceClient) GetBar(ctx context.Context, in *BfGetBarReq, opts ...grpc.CallOption) (BfDatafeedService_GetBarClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[1], c.cc, "/bftrader.BfDatafeedService/GetBar", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetBarClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetBarClient interface {
	Recv() (*BfBarData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetBarClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetBarClient) Recv() (*BfBarData, error) {
	m := new(BfBarData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfDatafeedServiceClient) GetContract(ctx context.Context, in *BfDatafeedGetContractReq, opts ...grpc.CallOption) (BfDatafeedService_GetContractClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[2], c.cc, "/bftrader.BfDatafeedService/GetContract", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetContractClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetContractClient interface {
	Recv() (*bftrader1.BfContractData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetContractClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetContractClient) Recv() (*bftrader1.BfContractData, error) {
	m := new(bftrader1.BfContractData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BfDatafeedService service

type BfDatafeedServiceServer interface {
	// 活跃检测
	Ping(context.Context, *bftrader1.BfPingData) (*bftrader1.BfPingData, error)
	// 保存tick
	InsertTick(context.Context, *bftrader1.BfTickData) (*bftrader1.BfVoid, error)
	// 保存bar
	InsertBar(context.Context, *BfBarData) (*bftrader1.BfVoid, error)
	// 保存contract
	InsertContract(context.Context, *bftrader1.BfContractData) (*bftrader1.BfVoid, error)
	// 获取tick
	GetTick(*BfGetTickReq, BfDatafeedService_GetTickServer) error
	// 获取bar
	GetBar(*BfGetBarReq, BfDatafeedService_GetBarServer) error
	// 获取contract
	GetContract(*BfDatafeedGetContractReq, BfDatafeedService_GetContractServer) error
}

func RegisterBfDatafeedServiceServer(s *grpc.Server, srv BfDatafeedServiceServer) {
	s.RegisterService(&_BfDatafeedService_serviceDesc, srv)
}

func _BfDatafeedService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader1.BfPingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertTick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader1.BfTickData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertTick(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BfBarData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertBar(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader1.BfContractData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertContract(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_GetTick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BfGetTickReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetTick(m, &bfDatafeedServiceGetTickServer{stream})
}

type BfDatafeedService_GetTickServer interface {
	Send(*bftrader1.BfTickData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetTickServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetTickServer) Send(m *bftrader1.BfTickData) error {
	return x.ServerStream.SendMsg(m)
}

func _BfDatafeedService_GetBar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BfGetBarReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetBar(m, &bfDatafeedServiceGetBarServer{stream})
}

type BfDatafeedService_GetBarServer interface {
	Send(*BfBarData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetBarServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetBarServer) Send(m *BfBarData) error {
	return x.ServerStream.SendMsg(m)
}

func _BfDatafeedService_GetContract_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BfDatafeedGetContractReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetContract(m, &bfDatafeedServiceGetContractServer{stream})
}

type BfDatafeedService_GetContractServer interface {
	Send(*bftrader1.BfContractData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetContractServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetContractServer) Send(m *bftrader1.BfContractData) error {
	return x.ServerStream.SendMsg(m)
}

var _BfDatafeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftrader.BfDatafeedService",
	HandlerType: (*BfDatafeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BfDatafeedService_Ping_Handler,
		},
		{
			MethodName: "InsertTick",
			Handler:    _BfDatafeedService_InsertTick_Handler,
		},
		{
			MethodName: "InsertBar",
			Handler:    _BfDatafeedService_InsertBar_Handler,
		},
		{
			MethodName: "InsertContract",
			Handler:    _BfDatafeedService_InsertContract_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTick",
			Handler:       _BfDatafeedService_GetTick_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBar",
			Handler:       _BfDatafeedService_GetBar_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetContract",
			Handler:       _BfDatafeedService_GetContract_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0x8d, 0x01, 0x1b, 0x3c, 0x10, 0xb2, 0xbf, 0xfd, 0x25, 0x91, 0x65, 0x55, 0x55, 0xe5, 0x53,
	0x55, 0xa9, 0x11, 0x84, 0x50, 0x55, 0xaa, 0xd4, 0x03, 0xa5, 0x6a, 0x51, 0x15, 0x40, 0x4e, 0x9a,
	0x1c, 0x2b, 0xdb, 0xac, 0xc1, 0x2a, 0xf1, 0x52, 0xb3, 0x49, 0x9a, 0x6f, 0x90, 0x4f, 0xd4, 0x6b,
	0xbf, 0x58, 0x0f, 0xdd, 0x3f, 0x06, 0x96, 0x08, 0x0e, 0xe5, 0xf8, 0xde, 0x9b, 0x37, 0x33, 0x3b,
	0x3b, 0xbb, 0x80, 0xc2, 0x78, 0x14, 0xb0, 0x20, 0x26, 0x64, 0x74, 0x32, 0xcb, 0x28, 0xa3, 0xb8,
	0x12, 0xc6, 0x2c, 0x0b, 0x46, 0x24, 0x73, 0x0f, 0xc2, 0x78, 0x1c, 0x30, 0x72, 0x1f, 0x3c, 0x28,
	0xc9, 0xfb, 0x53, 0x00, 0xbb, 0x13, 0x77, 0x82, 0xac, 0xcb, 0x2d, 0xf8, 0x18, 0xac, 0xf9, 0xc3,
	0x4d, 0x48, 0xa7, 0x8e, 0xf1, 0xc2, 0x78, 0x69, 0xfb, 0x39, 0xc2, 0x2e, 0x54, 0xc8, 0xcf, 0x68,
	0x12, 0xa4, 0x63, 0xe2, 0x14, 0xa4, 0xb2, 0xc4, 0xf8, 0x35, 0x58, 0x33, 0x92, 0x25, 0x74, 0xe4,
	0x14, 0xb9, 0x52, 0x3f, 0x3d, 0x3a, 0x59, 0x54, 0x3b, 0x91, 0x89, 0x87, 0x52, 0xf4, 0xf3, 0x20,
	0xfc, 0x1c, 0x20, 0x88, 0x58, 0x42, 0x53, 0x5e, 0x90, 0x38, 0x25, 0x99, 0x4c, 0x63, 0xb0, 0x03,
	0xe5, 0x30, 0xc8, 0x2e, 0x93, 0x1b, 0xe2, 0x98, 0x52, 0x5c, 0x40, 0xd1, 0xdc, 0x1d, 0x9d, 0xde,
	0x72, 0xc1, 0xe2, 0x82, 0xe9, 0xe7, 0x08, 0x7b, 0x50, 0xa3, 0x33, 0x92, 0xf6, 0x52, 0x46, 0x32,
	0x32, 0x67, 0x4e, 0x99, 0xab, 0x86, 0xbf, 0xc6, 0x89, 0xaa, 0xd3, 0x60, 0xce, 0xae, 0x94, 0xbf,
	0x22, 0xfd, 0x1a, 0x83, 0x9f, 0x81, 0x2d, 0xe2, 0x87, 0x59, 0x12, 0x11, 0xc7, 0x96, 0x09, 0x56,
	0x84, 0x50, 0x27, 0xc9, 0x78, 0xa2, 0x54, 0x50, 0xea, 0x92, 0x10, 0xc3, 0x99, 0xd2, 0x7b, 0x25,
	0x56, 0xa5, 0xb8, 0xc4, 0xa2, 0x6e, 0x34, 0xa5, 0x73, 0xa2, 0xd4, 0x9a, 0x54, 0x35, 0xc6, 0x7b,
	0x34, 0xa0, 0xd6, 0x89, 0x3f, 0x11, 0x76, 0x99, 0x44, 0xdf, 0x7d, 0xf2, 0x63, 0xa7, 0x1b, 0xe0,
	0x1e, 0x46, 0xe5, 0x38, 0x8b, 0xca, 0xa3, 0x90, 0xe2, 0xe5, 0x24, 0x4b, 0x0b, 0x5e, 0x0e, 0xf2,
	0x10, 0xcc, 0x88, 0xde, 0xa6, 0x4c, 0x0e, 0xd8, 0xf4, 0x15, 0xf0, 0x7e, 0x19, 0x50, 0x95, 0xad,
	0xf0, 0x3b, 0xdb, 0xb5, 0x93, 0x7f, 0xdc, 0x85, 0x55, 0xe3, 0xa5, 0x2d, 0x8d, 0x9b, 0x9b, 0x1b,
	0xb7, 0xf4, 0xc6, 0xfb, 0xe0, 0x74, 0xe2, 0x6e, 0xbe, 0xf1, 0xbc, 0xff, 0x0f, 0x34, 0xe5, 0x35,
	0x23, 0xb6, 0xe3, 0x21, 0x5e, 0x3d, 0x16, 0xc4, 0x20, 0x96, 0xdd, 0x62, 0x0c, 0xf5, 0xe1, 0x47,
	0xbf, 0x37, 0xe8, 0x7e, 0xfb, 0xda, 0xff, 0xd2, 0x1f, 0x5c, 0xf7, 0xd1, 0x1e, 0xde, 0x07, 0x3b,
	0xe7, 0x2e, 0x9a, 0xc8, 0xd0, 0x61, 0x0b, 0x15, 0x74, 0xd8, 0x46, 0x45, 0x5c, 0x07, 0x58, 0x06,
	0x37, 0x50, 0x69, 0x0d, 0xb7, 0x91, 0xa9, 0xe3, 0x56, 0x03, 0x59, 0x9a, 0xfd, 0xbc, 0x89, 0xca,
	0x3a, 0x6c, 0xa1, 0x8a, 0x0e, 0xdb, 0xc8, 0xd6, 0xcc, 0xe7, 0x3c, 0x39, 0xac, 0xe1, 0x36, 0xaa,
	0xea, 0x98, 0x27, 0xaf, 0x69, 0xf6, 0xcf, 0x4d, 0xb4, 0xaf, 0xc1, 0x6e, 0x13, 0xd5, 0x35, 0x78,
	0xdd, 0x44, 0x07, 0xa7, 0xbf, 0x8b, 0xf0, 0xdf, 0x6a, 0xb6, 0x17, 0x24, 0xbb, 0x13, 0x4b, 0x7d,
	0x06, 0xa5, 0x61, 0x92, 0x8e, 0xf1, 0xa1, 0x7e, 0xbb, 0x82, 0x11, 0x81, 0xee, 0x46, 0xd6, 0xdb,
	0xc3, 0x6f, 0x00, 0x7a, 0xe9, 0x9c, 0x64, 0x72, 0xd5, 0xd7, 0xbd, 0x82, 0x91, 0x5e, 0xa4, 0xb3,
	0x57, 0x34, 0x19, 0x71, 0xdf, 0x19, 0xd8, 0xca, 0xc7, 0x6f, 0x04, 0xff, 0xff, 0x64, 0xa1, 0xb6,
	0xba, 0xde, 0x43, 0x5d, 0xb9, 0x16, 0xdb, 0x80, 0x1d, 0x3d, 0x6a, 0xc1, 0x6e, 0xf5, 0xbf, 0x83,
	0x72, 0xfe, 0x2a, 0xf1, 0xb1, 0x2e, 0xaf, 0x9e, 0xaa, 0xbb, 0xf1, 0x08, 0xde, 0x5e, 0xc3, 0xc0,
	0x6f, 0xc1, 0x52, 0xef, 0x08, 0x1f, 0x3d, 0xf1, 0xaa, 0xb7, 0xe5, 0x6e, 0x3a, 0x86, 0x74, 0x0e,
	0xa0, 0xaa, 0x6d, 0x30, 0xf6, 0xf4, 0xb8, 0xcd, 0x2b, 0xee, 0x6e, 0x3d, 0x97, 0x48, 0x18, 0x5a,
	0xf2, 0x9b, 0x6f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x06, 0x15, 0xc3, 0x15, 0x06, 0x00,
	0x00,
}

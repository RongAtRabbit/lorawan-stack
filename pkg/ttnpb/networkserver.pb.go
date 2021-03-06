// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/networkserver.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GsNsClient is the client API for GsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GsNsClient interface {
	HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*types.Empty, error)
}

type gsNsClient struct {
	cc *grpc.ClientConn
}

func NewGsNsClient(cc *grpc.ClientConn) GsNsClient {
	return &gsNsClient{cc}
}

func (c *gsNsClient) HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GsNs/HandleUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsNsServer is the server API for GsNs service.
type GsNsServer interface {
	HandleUplink(context.Context, *UplinkMessage) (*types.Empty, error)
}

func RegisterGsNsServer(s *grpc.Server, srv GsNsServer) {
	s.RegisterService(&_GsNs_serviceDesc, srv)
}

func _GsNs_HandleUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsNsServer).HandleUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GsNs/HandleUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsNsServer).HandleUplink(ctx, req.(*UplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GsNs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GsNs",
	HandlerType: (*GsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUplink",
			Handler:    _GsNs_HandleUplink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

// AsNsClient is the client API for AsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AsNsClient interface {
	LinkApplication(ctx context.Context, opts ...grpc.CallOption) (AsNs_LinkApplicationClient, error)
	DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error)
	DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error)
}

type asNsClient struct {
	cc *grpc.ClientConn
}

func NewAsNsClient(cc *grpc.ClientConn) AsNsClient {
	return &asNsClient{cc}
}

func (c *asNsClient) LinkApplication(ctx context.Context, opts ...grpc.CallOption) (AsNs_LinkApplicationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AsNs_serviceDesc.Streams[0], "/ttn.lorawan.v3.AsNs/LinkApplication", opts...)
	if err != nil {
		return nil, err
	}
	x := &asNsLinkApplicationClient{stream}
	return x, nil
}

type AsNs_LinkApplicationClient interface {
	Send(*types.Empty) error
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type asNsLinkApplicationClient struct {
	grpc.ClientStream
}

func (x *asNsLinkApplicationClient) Send(m *types.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *asNsLinkApplicationClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *asNsClient) DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueueReplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueuePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error) {
	out := new(ApplicationDownlinks)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsNs/DownlinkQueueList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsNsServer is the server API for AsNs service.
type AsNsServer interface {
	LinkApplication(AsNs_LinkApplicationServer) error
	DownlinkQueueReplace(context.Context, *DownlinkQueueRequest) (*types.Empty, error)
	DownlinkQueuePush(context.Context, *DownlinkQueueRequest) (*types.Empty, error)
	DownlinkQueueList(context.Context, *EndDeviceIdentifiers) (*ApplicationDownlinks, error)
}

func RegisterAsNsServer(s *grpc.Server, srv AsNsServer) {
	s.RegisterService(&_AsNs_serviceDesc, srv)
}

func _AsNs_LinkApplication_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AsNsServer).LinkApplication(&asNsLinkApplicationServer{stream})
}

type AsNs_LinkApplicationServer interface {
	Send(*ApplicationUp) error
	Recv() (*types.Empty, error)
	grpc.ServerStream
}

type asNsLinkApplicationServer struct {
	grpc.ServerStream
}

func (x *asNsLinkApplicationServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *asNsLinkApplicationServer) Recv() (*types.Empty, error) {
	m := new(types.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AsNs_DownlinkQueueReplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueueReplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueuePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueuePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsNs/DownlinkQueueList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueList(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _AsNs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.AsNs",
	HandlerType: (*AsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownlinkQueueReplace",
			Handler:    _AsNs_DownlinkQueueReplace_Handler,
		},
		{
			MethodName: "DownlinkQueuePush",
			Handler:    _AsNs_DownlinkQueuePush_Handler,
		},
		{
			MethodName: "DownlinkQueueList",
			Handler:    _AsNs_DownlinkQueueList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LinkApplication",
			Handler:       _AsNs_LinkApplication_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

// NsEndDeviceRegistryClient is the client API for NsEndDeviceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsEndDeviceRegistryClient interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Set creates or updates the device.
	Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type nsEndDeviceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNsEndDeviceRegistryClient(cc *grpc.ClientConn) NsEndDeviceRegistryClient {
	return &nsEndDeviceRegistryClient{cc}
}

func (c *nsEndDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsEndDeviceRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsEndDeviceRegistryServer is the server API for NsEndDeviceRegistry service.
type NsEndDeviceRegistryServer interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// Set creates or updates the device.
	Set(context.Context, *SetEndDeviceRequest) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*types.Empty, error)
}

func RegisterNsEndDeviceRegistryServer(s *grpc.Server, srv NsEndDeviceRegistryServer) {
	s.RegisterService(&_NsEndDeviceRegistry_serviceDesc, srv)
}

func _NsEndDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, req.(*SetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsEndDeviceRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsEndDeviceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsEndDeviceRegistry",
	HandlerType: (*NsEndDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NsEndDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _NsEndDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NsEndDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/networkserver.proto",
}

func init() {
	proto.RegisterFile("lorawan-stack/api/networkserver.proto", fileDescriptor_networkserver_9c56cf1de73aa617)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/networkserver.proto", fileDescriptor_networkserver_9c56cf1de73aa617)
}

var fileDescriptor_networkserver_9c56cf1de73aa617 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x31, 0x48, 0x1c, 0x4f,
	0x14, 0xc6, 0x67, 0x54, 0x2c, 0x96, 0x3f, 0xff, 0x90, 0x4d, 0x08, 0xe4, 0x92, 0x3c, 0xc2, 0x69,
	0x20, 0x48, 0xdc, 0x0d, 0xda, 0xd9, 0x19, 0xee, 0xb8, 0x08, 0x2a, 0x51, 0x63, 0x63, 0x0a, 0xd9,
	0xbb, 0x7b, 0xee, 0x0d, 0xb7, 0xce, 0x6c, 0x76, 0xe6, 0x14, 0x09, 0x82, 0xa4, 0xb2, 0x0c, 0x84,
	0x40, 0xca, 0x90, 0x4a, 0x48, 0x23, 0x56, 0x96, 0x96, 0x96, 0x42, 0x1a, 0xab, 0xc4, 0x9d, 0x4d,
	0x61, 0x69, 0x29, 0xa9, 0xc2, 0xed, 0xae, 0xde, 0x79, 0xeb, 0x85, 0x1c, 0x49, 0xb7, 0x33, 0xef,
	0xfb, 0xbe, 0xf9, 0xbd, 0x99, 0xd9, 0x31, 0x1e, 0x79, 0x22, 0x70, 0xd6, 0x1d, 0x3e, 0x2a, 0x95,
	0x53, 0xa9, 0xdb, 0x8e, 0xcf, 0x6c, 0x8e, 0x6a, 0x5d, 0x04, 0x75, 0x89, 0xc1, 0x1a, 0x06, 0x96,
	0x1f, 0x08, 0x25, 0xcc, 0xff, 0x95, 0xe2, 0x56, 0x2a, 0xb5, 0xd6, 0xc6, 0x73, 0xa3, 0x2e, 0x53,
	0xb5, 0x46, 0xd9, 0xaa, 0x88, 0x55, 0xdb, 0x15, 0xae, 0xb0, 0x63, 0x59, 0xb9, 0xb1, 0x12, 0x8f,
	0xe2, 0x41, 0xfc, 0x95, 0xd8, 0x73, 0xf7, 0x5d, 0x21, 0x5c, 0x0f, 0xe3, 0x78, 0x87, 0x73, 0xa1,
	0x1c, 0xc5, 0x04, 0x97, 0x69, 0xf5, 0x5e, 0x5a, 0xbd, 0xcc, 0xc0, 0x55, 0x5f, 0x6d, 0xa4, 0xc5,
	0x7c, 0x16, 0x10, 0x79, 0x75, 0xb9, 0x8a, 0x6b, 0xac, 0x82, 0xa9, 0x66, 0x28, 0xab, 0x61, 0x55,
	0xe4, 0x8a, 0xad, 0x30, 0x0c, 0x2e, 0x56, 0x79, 0x98, 0x15, 0xad, 0xa2, 0x94, 0x8e, 0x8b, 0xa9,
	0x62, 0x6c, 0xc6, 0x18, 0x28, 0xc9, 0x59, 0x69, 0x16, 0x8d, 0xff, 0x9e, 0x3b, 0xbc, 0xea, 0xe1,
	0xa2, 0xef, 0x31, 0x5e, 0x37, 0x1f, 0x58, 0x57, 0xbb, 0xb7, 0x92, 0xf9, 0x99, 0xc4, 0x9d, 0xbb,
	0x63, 0x25, 0xfc, 0xd6, 0x05, 0xbf, 0x55, 0x6c, 0xf2, 0x8f, 0x7d, 0xeb, 0x33, 0x06, 0x26, 0x9b,
	0x79, 0xd3, 0xc6, 0x8d, 0x69, 0xc6, 0xeb, 0x93, 0xbe, 0xef, 0xb1, 0x4a, 0xdc, 0xb9, 0xd9, 0xc5,
	0x93, 0xcb, 0x2c, 0xd5, 0x66, 0x5a, 0xf4, 0x1f, 0xd3, 0xa7, 0xd4, 0x7c, 0x69, 0xdc, 0x2e, 0x88,
	0x75, 0xde, 0x24, 0x98, 0x6b, 0x60, 0x03, 0xe7, 0xd1, 0xf7, 0x9c, 0x0a, 0x9a, 0xc3, 0x9d, 0xd6,
	0x0e, 0xd5, 0xeb, 0x06, 0x4a, 0xd5, 0x0d, 0xd6, 0x9c, 0x33, 0x6e, 0x5e, 0xd1, 0xbf, 0x68, 0xc8,
	0xda, 0x5f, 0x46, 0x2e, 0x77, 0x44, 0x4e, 0x33, 0xa9, 0xb2, 0x91, 0x45, 0x5e, 0x2d, 0xc4, 0x67,
	0x39, 0xd5, 0x3a, 0xb1, 0xdc, 0xf0, 0x6f, 0xb6, 0xe1, 0x22, 0x53, 0x8e, 0x7d, 0x1f, 0x30, 0x6e,
	0xcd, 0xca, 0xcb, 0x80, 0x79, 0x74, 0x99, 0x54, 0xc1, 0x86, 0xb9, 0x47, 0x8d, 0xfe, 0x12, 0x2a,
	0x73, 0xa8, 0x33, 0xa5, 0x84, 0xaa, 0x4d, 0x9d, 0xd0, 0xdf, 0xed, 0x0a, 0x94, 0xaf, 0xbf, 0xfd,
	0xfa, 0xe3, 0x7d, 0x1f, 0x9a, 0x15, 0x9b, 0x4b, 0xdb, 0x69, 0x11, 0x48, 0xfb, 0x4d, 0xeb, 0x0a,
	0x2e, 0xb3, 0xaa, 0xb4, 0xda, 0x8a, 0xd7, 0x8c, 0x37, 0xed, 0x44, 0x9a, 0xf5, 0x5d, 0x7e, 0x6e,
	0x9a, 0x3f, 0xa9, 0xd1, 0xbf, 0x70, 0x1d, 0xf4, 0x42, 0x6f, 0xd0, 0x7b, 0x34, 0xa6, 0xfe, 0x42,
	0x73, 0xaf, 0xb2, 0xd8, 0xe9, 0x5f, 0xd3, 0x13, 0x72, 0x9b, 0xa7, 0x85, 0x3b, 0x41, 0x47, 0x96,
	0xa6, 0xf2, 0x85, 0x7f, 0xb1, 0xc2, 0x04, 0x1d, 0x31, 0x3f, 0x50, 0x63, 0xb0, 0x80, 0x1e, 0x2a,
	0xfc, 0xc3, 0x0b, 0xd2, 0xe5, 0xce, 0xe5, 0x67, 0xe2, 0xe6, 0x4b, 0x23, 0xc5, 0x2c, 0x59, 0x8f,
	0x0d, 0x37, 0xe7, 0x9e, 0x7d, 0xa6, 0x87, 0x21, 0xd0, 0xa3, 0x10, 0xe8, 0x71, 0x08, 0xe4, 0x24,
	0x04, 0x72, 0x1a, 0x02, 0x39, 0x0b, 0x81, 0x9c, 0x87, 0x40, 0xb7, 0x34, 0xd0, 0x6d, 0x0d, 0x64,
	0x47, 0x03, 0xdd, 0xd5, 0x40, 0xf6, 0x35, 0x90, 0x03, 0x0d, 0xe4, 0x50, 0x03, 0x3d, 0xd2, 0x40,
	0x8f, 0x35, 0x90, 0x13, 0x0d, 0xf4, 0x54, 0x03, 0x39, 0xd3, 0x40, 0xcf, 0x35, 0x90, 0xad, 0x08,
	0xc8, 0x76, 0x04, 0xf4, 0x5d, 0x04, 0xe4, 0x63, 0x04, 0xf4, 0x53, 0x04, 0x64, 0x27, 0x02, 0xb2,
	0x1b, 0x01, 0xdd, 0x8f, 0x80, 0x1e, 0x44, 0x40, 0x97, 0x9e, 0xb8, 0xc2, 0x52, 0x35, 0x54, 0x35,
	0xc6, 0x5d, 0x69, 0xa5, 0xef, 0xb1, 0x7d, 0xf5, 0xfd, 0xf2, 0xeb, 0xae, 0xad, 0x14, 0xf7, 0xcb,
	0xe5, 0xc1, 0x78, 0x0f, 0xc6, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xf3, 0xf0, 0xbc, 0xcb,
	0x05, 0x00, 0x00,
}

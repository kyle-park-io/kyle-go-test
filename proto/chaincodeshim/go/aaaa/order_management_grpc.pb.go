// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: order_management.proto

package aaaa

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChaincodeSupportClient is the client API for ChaincodeSupport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc grpc.ClientConnInterface
}

func NewChaincodeSupportClient(cc grpc.ClientConnInterface) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChaincodeSupport_ServiceDesc.Streams[0], "/ecommerce.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaincodeSupportServer is the server API for ChaincodeSupport service.
// All implementations should embed UnimplementedChaincodeSupportServer
// for forward compatibility
type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

// UnimplementedChaincodeSupportServer should be embedded to have forward compatible implementations.
type UnimplementedChaincodeSupportServer struct {
}

func (UnimplementedChaincodeSupportServer) Register(ChaincodeSupport_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}

// UnsafeChaincodeSupportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaincodeSupportServer will
// result in compilation errors.
type UnsafeChaincodeSupportServer interface {
	mustEmbedUnimplementedChaincodeSupportServer()
}

func RegisterChaincodeSupportServer(s grpc.ServiceRegistrar, srv ChaincodeSupportServer) {
	s.RegisterService(&ChaincodeSupport_ServiceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaincodeSupport_ServiceDesc is the grpc.ServiceDesc for ChaincodeSupport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaincodeSupport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order_management.proto",
}
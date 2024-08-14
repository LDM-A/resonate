// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: internal/app/subsystems/api/grpc/api/callback.proto

package api

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

const (
	Callbacks_CreateCallback_FullMethodName = "/callback.Callbacks/CreateCallback"
)

// CallbacksClient is the client API for Callbacks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallbacksClient interface {
	CreateCallback(ctx context.Context, in *CreateCallbackRequest, opts ...grpc.CallOption) (*CreateCallbackResponse, error)
}

type callbacksClient struct {
	cc grpc.ClientConnInterface
}

func NewCallbacksClient(cc grpc.ClientConnInterface) CallbacksClient {
	return &callbacksClient{cc}
}

func (c *callbacksClient) CreateCallback(ctx context.Context, in *CreateCallbackRequest, opts ...grpc.CallOption) (*CreateCallbackResponse, error) {
	out := new(CreateCallbackResponse)
	err := c.cc.Invoke(ctx, Callbacks_CreateCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallbacksServer is the server API for Callbacks service.
// All implementations must embed UnimplementedCallbacksServer
// for forward compatibility
type CallbacksServer interface {
	CreateCallback(context.Context, *CreateCallbackRequest) (*CreateCallbackResponse, error)
	mustEmbedUnimplementedCallbacksServer()
}

// UnimplementedCallbacksServer must be embedded to have forward compatible implementations.
type UnimplementedCallbacksServer struct {
}

func (UnimplementedCallbacksServer) CreateCallback(context.Context, *CreateCallbackRequest) (*CreateCallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCallback not implemented")
}
func (UnimplementedCallbacksServer) mustEmbedUnimplementedCallbacksServer() {}

// UnsafeCallbacksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallbacksServer will
// result in compilation errors.
type UnsafeCallbacksServer interface {
	mustEmbedUnimplementedCallbacksServer()
}

func RegisterCallbacksServer(s grpc.ServiceRegistrar, srv CallbacksServer) {
	s.RegisterService(&Callbacks_ServiceDesc, srv)
}

func _Callbacks_CreateCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbacksServer).CreateCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Callbacks_CreateCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbacksServer).CreateCallback(ctx, req.(*CreateCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Callbacks_ServiceDesc is the grpc.ServiceDesc for Callbacks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Callbacks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "callback.Callbacks",
	HandlerType: (*CallbacksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCallback",
			Handler:    _Callbacks_CreateCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/subsystems/api/grpc/api/callback.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pixkey.proto

package pb

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

// PixServiceControllerClient is the client API for PixServiceController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PixServiceControllerClient interface {
	RegisterPixKey(ctx context.Context, in *PixKeyRegistration, opts ...grpc.CallOption) (*PixKeyCreatedResult, error)
	FindPixKey(ctx context.Context, in *PixKey, opts ...grpc.CallOption) (*PixKeyInfo, error)
}

type pixServiceControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPixServiceControllerClient(cc grpc.ClientConnInterface) PixServiceControllerClient {
	return &pixServiceControllerClient{cc}
}

func (c *pixServiceControllerClient) RegisterPixKey(ctx context.Context, in *PixKeyRegistration, opts ...grpc.CallOption) (*PixKeyCreatedResult, error) {
	out := new(PixKeyCreatedResult)
	err := c.cc.Invoke(ctx, "/codePix.PixServiceController/RegisterPixKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pixServiceControllerClient) FindPixKey(ctx context.Context, in *PixKey, opts ...grpc.CallOption) (*PixKeyInfo, error) {
	out := new(PixKeyInfo)
	err := c.cc.Invoke(ctx, "/codePix.PixServiceController/FindPixKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PixServiceControllerServer is the server API for PixServiceController service.
// All implementations must embed UnimplementedPixServiceControllerServer
// for forward compatibility
type PixServiceControllerServer interface {
	RegisterPixKey(context.Context, *PixKeyRegistration) (*PixKeyCreatedResult, error)
	FindPixKey(context.Context, *PixKey) (*PixKeyInfo, error)
	mustEmbedUnimplementedPixServiceControllerServer()
}

// UnimplementedPixServiceControllerServer must be embedded to have forward compatible implementations.
type UnimplementedPixServiceControllerServer struct {
}

func (UnimplementedPixServiceControllerServer) RegisterPixKey(context.Context, *PixKeyRegistration) (*PixKeyCreatedResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPixKey not implemented")
}
func (UnimplementedPixServiceControllerServer) FindPixKey(context.Context, *PixKey) (*PixKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPixKey not implemented")
}
func (UnimplementedPixServiceControllerServer) mustEmbedUnimplementedPixServiceControllerServer() {}

// UnsafePixServiceControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PixServiceControllerServer will
// result in compilation errors.
type UnsafePixServiceControllerServer interface {
	mustEmbedUnimplementedPixServiceControllerServer()
}

func RegisterPixServiceControllerServer(s grpc.ServiceRegistrar, srv PixServiceControllerServer) {
	s.RegisterService(&PixServiceController_ServiceDesc, srv)
}

func _PixServiceController_RegisterPixKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PixKeyRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceControllerServer).RegisterPixKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codePix.PixServiceController/RegisterPixKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceControllerServer).RegisterPixKey(ctx, req.(*PixKeyRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _PixServiceController_FindPixKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PixKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceControllerServer).FindPixKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codePix.PixServiceController/FindPixKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceControllerServer).FindPixKey(ctx, req.(*PixKey))
	}
	return interceptor(ctx, in, info, handler)
}

// PixServiceController_ServiceDesc is the grpc.ServiceDesc for PixServiceController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PixServiceController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codePix.PixServiceController",
	HandlerType: (*PixServiceControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPixKey",
			Handler:    _PixServiceController_RegisterPixKey_Handler,
		},
		{
			MethodName: "FindPixKey",
			Handler:    _PixServiceController_FindPixKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pixkey.proto",
}

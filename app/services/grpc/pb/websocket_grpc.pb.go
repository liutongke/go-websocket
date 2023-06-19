// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: websocket.proto

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	SendUserMsgToLocal(ctx context.Context, in *SendUserMsgToLocalRequest, opts ...grpc.CallOption) (*SendUserResponse, error)
	SendGroupMsgToLocal(ctx context.Context, in *SendGroupMsgToLocalRequest, opts ...grpc.CallOption) (*SendGroupResponse, error)
	SendAllMsgToLocal(ctx context.Context, in *SendAllMsgToLocalRequest, opts ...grpc.CallOption) (*SendAllResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SendUserMsgToLocal(ctx context.Context, in *SendUserMsgToLocalRequest, opts ...grpc.CallOption) (*SendUserResponse, error) {
	out := new(SendUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Greeter/SendUserMsgToLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendGroupMsgToLocal(ctx context.Context, in *SendGroupMsgToLocalRequest, opts ...grpc.CallOption) (*SendGroupResponse, error) {
	out := new(SendGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.Greeter/SendGroupMsgToLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendAllMsgToLocal(ctx context.Context, in *SendAllMsgToLocalRequest, opts ...grpc.CallOption) (*SendAllResponse, error) {
	out := new(SendAllResponse)
	err := c.cc.Invoke(ctx, "/pb.Greeter/SendAllMsgToLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	SendUserMsgToLocal(context.Context, *SendUserMsgToLocalRequest) (*SendUserResponse, error)
	SendGroupMsgToLocal(context.Context, *SendGroupMsgToLocalRequest) (*SendGroupResponse, error)
	SendAllMsgToLocal(context.Context, *SendAllMsgToLocalRequest) (*SendAllResponse, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SendUserMsgToLocal(context.Context, *SendUserMsgToLocalRequest) (*SendUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserMsgToLocal not implemented")
}
func (UnimplementedGreeterServer) SendGroupMsgToLocal(context.Context, *SendGroupMsgToLocalRequest) (*SendGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupMsgToLocal not implemented")
}
func (UnimplementedGreeterServer) SendAllMsgToLocal(context.Context, *SendAllMsgToLocalRequest) (*SendAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAllMsgToLocal not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SendUserMsgToLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserMsgToLocalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendUserMsgToLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Greeter/SendUserMsgToLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendUserMsgToLocal(ctx, req.(*SendUserMsgToLocalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendGroupMsgToLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupMsgToLocalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendGroupMsgToLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Greeter/SendGroupMsgToLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendGroupMsgToLocal(ctx, req.(*SendGroupMsgToLocalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendAllMsgToLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAllMsgToLocalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendAllMsgToLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Greeter/SendAllMsgToLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendAllMsgToLocal(ctx, req.(*SendAllMsgToLocalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserMsgToLocal",
			Handler:    _Greeter_SendUserMsgToLocal_Handler,
		},
		{
			MethodName: "SendGroupMsgToLocal",
			Handler:    _Greeter_SendGroupMsgToLocal_Handler,
		},
		{
			MethodName: "SendAllMsgToLocal",
			Handler:    _Greeter_SendAllMsgToLocal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "websocket.proto",
}

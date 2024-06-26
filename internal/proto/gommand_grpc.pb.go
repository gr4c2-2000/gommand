// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/proto/gommand.proto

package proto

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
	Gommand_CommandInfo_FullMethodName = "/proto.Gommand/CommandInfo"
	Gommand_ExecCommand_FullMethodName = "/proto.Gommand/ExecCommand"
	Gommand_CommandList_FullMethodName = "/proto.Gommand/CommandList"
)

// GommandClient is the client API for Gommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GommandClient interface {
	CommandInfo(ctx context.Context, in *Input, opts ...grpc.CallOption) (*CommandInfoResult, error)
	ExecCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResult, error)
	CommandList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CommandListResult, error)
}

type gommandClient struct {
	cc grpc.ClientConnInterface
}

func NewGommandClient(cc grpc.ClientConnInterface) GommandClient {
	return &gommandClient{cc}
}

func (c *gommandClient) CommandInfo(ctx context.Context, in *Input, opts ...grpc.CallOption) (*CommandInfoResult, error) {
	out := new(CommandInfoResult)
	err := c.cc.Invoke(ctx, Gommand_CommandInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gommandClient) ExecCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResult, error) {
	out := new(CommandResult)
	err := c.cc.Invoke(ctx, Gommand_ExecCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gommandClient) CommandList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CommandListResult, error) {
	out := new(CommandListResult)
	err := c.cc.Invoke(ctx, Gommand_CommandList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GommandServer is the server API for Gommand service.
// All implementations must embed UnimplementedGommandServer
// for forward compatibility
type GommandServer interface {
	CommandInfo(context.Context, *Input) (*CommandInfoResult, error)
	ExecCommand(context.Context, *Command) (*CommandResult, error)
	CommandList(context.Context, *Empty) (*CommandListResult, error)
	mustEmbedUnimplementedGommandServer()
}

// UnimplementedGommandServer must be embedded to have forward compatible implementations.
type UnimplementedGommandServer struct {
}

func (UnimplementedGommandServer) CommandInfo(context.Context, *Input) (*CommandInfoResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandInfo not implemented")
}
func (UnimplementedGommandServer) ExecCommand(context.Context, *Command) (*CommandResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecCommand not implemented")
}
func (UnimplementedGommandServer) CommandList(context.Context, *Empty) (*CommandListResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandList not implemented")
}
func (UnimplementedGommandServer) mustEmbedUnimplementedGommandServer() {}

// UnsafeGommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GommandServer will
// result in compilation errors.
type UnsafeGommandServer interface {
	mustEmbedUnimplementedGommandServer()
}

func RegisterGommandServer(s grpc.ServiceRegistrar, srv GommandServer) {
	s.RegisterService(&Gommand_ServiceDesc, srv)
}

func _Gommand_CommandInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GommandServer).CommandInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gommand_CommandInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GommandServer).CommandInfo(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gommand_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GommandServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gommand_ExecCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GommandServer).ExecCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gommand_CommandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GommandServer).CommandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gommand_CommandList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GommandServer).CommandList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Gommand_ServiceDesc is the grpc.ServiceDesc for Gommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Gommand",
	HandlerType: (*GommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommandInfo",
			Handler:    _Gommand_CommandInfo_Handler,
		},
		{
			MethodName: "ExecCommand",
			Handler:    _Gommand_ExecCommand_Handler,
		},
		{
			MethodName: "CommandList",
			Handler:    _Gommand_CommandList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/gommand.proto",
}

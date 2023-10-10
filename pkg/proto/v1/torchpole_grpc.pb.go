// Copyright 2022 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/combizent/torchpole.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: v1/torchpole.proto

package v1

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
	TorchPole_ListUser_FullMethodName = "/v1.TorchPole/ListUser"
)

// TorchPoleClient is the client API for TorchPole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TorchPoleClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type torchPoleClient struct {
	cc grpc.ClientConnInterface
}

func NewTorchPoleClient(cc grpc.ClientConnInterface) TorchPoleClient {
	return &torchPoleClient{cc}
}

func (c *torchPoleClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, TorchPole_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TorchPoleServer is the server API for TorchPole service.
// All implementations must embed UnimplementedTorchPoleServer
// for forward compatibility
type TorchPoleServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedTorchPoleServer()
}

// UnimplementedTorchPoleServer must be embedded to have forward compatible implementations.
type UnimplementedTorchPoleServer struct {
}

func (UnimplementedTorchPoleServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedTorchPoleServer) mustEmbedUnimplementedTorchPoleServer() {}

// UnsafeTorchPoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TorchPoleServer will
// result in compilation errors.
type UnsafeTorchPoleServer interface {
	mustEmbedUnimplementedTorchPoleServer()
}

func RegisterTorchPoleServer(s grpc.ServiceRegistrar, srv TorchPoleServer) {
	s.RegisterService(&TorchPole_ServiceDesc, srv)
}

func _TorchPole_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorchPoleServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TorchPole_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorchPoleServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TorchPole_ServiceDesc is the grpc.ServiceDesc for TorchPole service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TorchPole_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TorchPole",
	HandlerType: (*TorchPoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _TorchPole_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/torchpole.proto",
}

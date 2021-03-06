// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// BattleClient is the client API for Battle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BattleClient interface {
	// 开始战斗
	Start(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleReply, error)
	NStart(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleReply, error)
}

type battleClient struct {
	cc grpc.ClientConnInterface
}

func NewBattleClient(cc grpc.ClientConnInterface) BattleClient {
	return &battleClient{cc}
}

func (c *battleClient) Start(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleReply, error) {
	out := new(BattleReply)
	err := c.cc.Invoke(ctx, "/hello.v1.Battle/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battleClient) NStart(ctx context.Context, in *BattleRequest, opts ...grpc.CallOption) (*BattleReply, error) {
	out := new(BattleReply)
	err := c.cc.Invoke(ctx, "/hello.v1.Battle/NStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BattleServer is the server API for Battle service.
// All implementations must embed UnimplementedBattleServer
// for forward compatibility
type BattleServer interface {
	// 开始战斗
	Start(context.Context, *BattleRequest) (*BattleReply, error)
	NStart(context.Context, *BattleRequest) (*BattleReply, error)
	mustEmbedUnimplementedBattleServer()
}

// UnimplementedBattleServer must be embedded to have forward compatible implementations.
type UnimplementedBattleServer struct {
}

func (UnimplementedBattleServer) Start(context.Context, *BattleRequest) (*BattleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedBattleServer) NStart(context.Context, *BattleRequest) (*BattleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NStart not implemented")
}
func (UnimplementedBattleServer) mustEmbedUnimplementedBattleServer() {}

// UnsafeBattleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BattleServer will
// result in compilation errors.
type UnsafeBattleServer interface {
	mustEmbedUnimplementedBattleServer()
}

func RegisterBattleServer(s grpc.ServiceRegistrar, srv BattleServer) {
	s.RegisterService(&Battle_ServiceDesc, srv)
}

func _Battle_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BattleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Battle/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServer).Start(ctx, req.(*BattleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battle_NStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BattleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServer).NStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Battle/NStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServer).NStart(ctx, req.(*BattleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Battle_ServiceDesc is the grpc.ServiceDesc for Battle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Battle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.v1.Battle",
	HandlerType: (*BattleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Battle_Start_Handler,
		},
		{
			MethodName: "NStart",
			Handler:    _Battle_NStart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/Batlle.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: riot/account/v1/account.proto

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

// RiotAccountServiceClient is the client API for RiotAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiotAccountServiceClient interface {
	Puuid(ctx context.Context, in *RiotAccountReq, opts ...grpc.CallOption) (*RiotAccount, error)
	RiotId(ctx context.Context, in *RiotAccountReq, opts ...grpc.CallOption) (*RiotAccount, error)
}

type riotAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRiotAccountServiceClient(cc grpc.ClientConnInterface) RiotAccountServiceClient {
	return &riotAccountServiceClient{cc}
}

func (c *riotAccountServiceClient) Puuid(ctx context.Context, in *RiotAccountReq, opts ...grpc.CallOption) (*RiotAccount, error) {
	out := new(RiotAccount)
	err := c.cc.Invoke(ctx, "/riot.account.v1.RiotAccountService/Puuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAccountServiceClient) RiotId(ctx context.Context, in *RiotAccountReq, opts ...grpc.CallOption) (*RiotAccount, error) {
	out := new(RiotAccount)
	err := c.cc.Invoke(ctx, "/riot.account.v1.RiotAccountService/RiotId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiotAccountServiceServer is the server API for RiotAccountService service.
// All implementations must embed UnimplementedRiotAccountServiceServer
// for forward compatibility
type RiotAccountServiceServer interface {
	Puuid(context.Context, *RiotAccountReq) (*RiotAccount, error)
	RiotId(context.Context, *RiotAccountReq) (*RiotAccount, error)
	mustEmbedUnimplementedRiotAccountServiceServer()
}

// UnimplementedRiotAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRiotAccountServiceServer struct {
}

func (UnimplementedRiotAccountServiceServer) Puuid(context.Context, *RiotAccountReq) (*RiotAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Puuid not implemented")
}
func (UnimplementedRiotAccountServiceServer) RiotId(context.Context, *RiotAccountReq) (*RiotAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RiotId not implemented")
}
func (UnimplementedRiotAccountServiceServer) mustEmbedUnimplementedRiotAccountServiceServer() {}

// UnsafeRiotAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RiotAccountServiceServer will
// result in compilation errors.
type UnsafeRiotAccountServiceServer interface {
	mustEmbedUnimplementedRiotAccountServiceServer()
}

func RegisterRiotAccountServiceServer(s grpc.ServiceRegistrar, srv RiotAccountServiceServer) {
	s.RegisterService(&RiotAccountService_ServiceDesc, srv)
}

func _RiotAccountService_Puuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiotAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAccountServiceServer).Puuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riot.account.v1.RiotAccountService/Puuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAccountServiceServer).Puuid(ctx, req.(*RiotAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAccountService_RiotId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiotAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAccountServiceServer).RiotId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riot.account.v1.RiotAccountService/RiotId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAccountServiceServer).RiotId(ctx, req.(*RiotAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RiotAccountService_ServiceDesc is the grpc.ServiceDesc for RiotAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RiotAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "riot.account.v1.RiotAccountService",
	HandlerType: (*RiotAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Puuid",
			Handler:    _RiotAccountService_Puuid_Handler,
		},
		{
			MethodName: "RiotId",
			Handler:    _RiotAccountService_RiotId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "riot/account/v1/account.proto",
}
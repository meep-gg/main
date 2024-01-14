// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: scylla/lol/champion/v1/champion.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChampionServiceClient is the client API for ChampionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChampionServiceClient interface {
	GetChampion(ctx context.Context, in *GetChampionRequest, opts ...grpc.CallOption) (*Champion, error)
	UpsertChampion(ctx context.Context, in *UpsertChampionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteChampion(ctx context.Context, in *DeleteChampionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type championServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChampionServiceClient(cc grpc.ClientConnInterface) ChampionServiceClient {
	return &championServiceClient{cc}
}

func (c *championServiceClient) GetChampion(ctx context.Context, in *GetChampionRequest, opts ...grpc.CallOption) (*Champion, error) {
	out := new(Champion)
	err := c.cc.Invoke(ctx, "/scylla.lol.champion.v1.ChampionService/GetChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *championServiceClient) UpsertChampion(ctx context.Context, in *UpsertChampionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scylla.lol.champion.v1.ChampionService/UpsertChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *championServiceClient) DeleteChampion(ctx context.Context, in *DeleteChampionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scylla.lol.champion.v1.ChampionService/DeleteChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChampionServiceServer is the server API for ChampionService service.
// All implementations must embed UnimplementedChampionServiceServer
// for forward compatibility
type ChampionServiceServer interface {
	GetChampion(context.Context, *GetChampionRequest) (*Champion, error)
	UpsertChampion(context.Context, *UpsertChampionRequest) (*emptypb.Empty, error)
	DeleteChampion(context.Context, *DeleteChampionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChampionServiceServer()
}

// UnimplementedChampionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChampionServiceServer struct {
}

func (UnimplementedChampionServiceServer) GetChampion(context.Context, *GetChampionRequest) (*Champion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampion not implemented")
}
func (UnimplementedChampionServiceServer) UpsertChampion(context.Context, *UpsertChampionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertChampion not implemented")
}
func (UnimplementedChampionServiceServer) DeleteChampion(context.Context, *DeleteChampionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChampion not implemented")
}
func (UnimplementedChampionServiceServer) mustEmbedUnimplementedChampionServiceServer() {}

// UnsafeChampionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChampionServiceServer will
// result in compilation errors.
type UnsafeChampionServiceServer interface {
	mustEmbedUnimplementedChampionServiceServer()
}

func RegisterChampionServiceServer(s grpc.ServiceRegistrar, srv ChampionServiceServer) {
	s.RegisterService(&ChampionService_ServiceDesc, srv)
}

func _ChampionService_GetChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChampionServiceServer).GetChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.champion.v1.ChampionService/GetChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChampionServiceServer).GetChampion(ctx, req.(*GetChampionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChampionService_UpsertChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertChampionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChampionServiceServer).UpsertChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.champion.v1.ChampionService/UpsertChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChampionServiceServer).UpsertChampion(ctx, req.(*UpsertChampionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChampionService_DeleteChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChampionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChampionServiceServer).DeleteChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.champion.v1.ChampionService/DeleteChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChampionServiceServer).DeleteChampion(ctx, req.(*DeleteChampionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChampionService_ServiceDesc is the grpc.ServiceDesc for ChampionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChampionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scylla.lol.champion.v1.ChampionService",
	HandlerType: (*ChampionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChampion",
			Handler:    _ChampionService_GetChampion_Handler,
		},
		{
			MethodName: "UpsertChampion",
			Handler:    _ChampionService_UpsertChampion_Handler,
		},
		{
			MethodName: "DeleteChampion",
			Handler:    _ChampionService_DeleteChampion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scylla/lol/champion/v1/champion.proto",
}

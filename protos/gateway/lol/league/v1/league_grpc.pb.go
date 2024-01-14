// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: gateway/lol/league/v1/league.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "meep.gg/protos/scylla/lol/league/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayLeagueServiceClient is the client API for GatewayLeagueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayLeagueServiceClient interface {
	GetLeagues(ctx context.Context, in *GatewayLeagueReq, opts ...grpc.CallOption) (*v1.ScyllaLeagueEntries, error)
	UpdateLeagues(ctx context.Context, in *GatewayLeagueReq, opts ...grpc.CallOption) (*v1.ScyllaLeagueEntries, error)
	UpdateChallenger(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateGrandmaster(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateMaster(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateEntry(ctx context.Context, in *GatewayUpdateEntryReq, opts ...grpc.CallOption) (*GatewayUpdateEntryRes, error)
}

type gatewayLeagueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayLeagueServiceClient(cc grpc.ClientConnInterface) GatewayLeagueServiceClient {
	return &gatewayLeagueServiceClient{cc}
}

func (c *gatewayLeagueServiceClient) GetLeagues(ctx context.Context, in *GatewayLeagueReq, opts ...grpc.CallOption) (*v1.ScyllaLeagueEntries, error) {
	out := new(v1.ScyllaLeagueEntries)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/GetLeagues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayLeagueServiceClient) UpdateLeagues(ctx context.Context, in *GatewayLeagueReq, opts ...grpc.CallOption) (*v1.ScyllaLeagueEntries, error) {
	out := new(v1.ScyllaLeagueEntries)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/UpdateLeagues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayLeagueServiceClient) UpdateChallenger(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/UpdateChallenger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayLeagueServiceClient) UpdateGrandmaster(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/UpdateGrandmaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayLeagueServiceClient) UpdateMaster(ctx context.Context, in *GatewayApexUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/UpdateMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayLeagueServiceClient) UpdateEntry(ctx context.Context, in *GatewayUpdateEntryReq, opts ...grpc.CallOption) (*GatewayUpdateEntryRes, error) {
	out := new(GatewayUpdateEntryRes)
	err := c.cc.Invoke(ctx, "/gateway.lol.league.v1.GatewayLeagueService/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayLeagueServiceServer is the server API for GatewayLeagueService service.
// All implementations must embed UnimplementedGatewayLeagueServiceServer
// for forward compatibility
type GatewayLeagueServiceServer interface {
	GetLeagues(context.Context, *GatewayLeagueReq) (*v1.ScyllaLeagueEntries, error)
	UpdateLeagues(context.Context, *GatewayLeagueReq) (*v1.ScyllaLeagueEntries, error)
	UpdateChallenger(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error)
	UpdateGrandmaster(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error)
	UpdateMaster(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error)
	UpdateEntry(context.Context, *GatewayUpdateEntryReq) (*GatewayUpdateEntryRes, error)
	mustEmbedUnimplementedGatewayLeagueServiceServer()
}

// UnimplementedGatewayLeagueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayLeagueServiceServer struct {
}

func (UnimplementedGatewayLeagueServiceServer) GetLeagues(context.Context, *GatewayLeagueReq) (*v1.ScyllaLeagueEntries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagues not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) UpdateLeagues(context.Context, *GatewayLeagueReq) (*v1.ScyllaLeagueEntries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeagues not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) UpdateChallenger(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChallenger not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) UpdateGrandmaster(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGrandmaster not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) UpdateMaster(context.Context, *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaster not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) UpdateEntry(context.Context, *GatewayUpdateEntryReq) (*GatewayUpdateEntryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (UnimplementedGatewayLeagueServiceServer) mustEmbedUnimplementedGatewayLeagueServiceServer() {}

// UnsafeGatewayLeagueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayLeagueServiceServer will
// result in compilation errors.
type UnsafeGatewayLeagueServiceServer interface {
	mustEmbedUnimplementedGatewayLeagueServiceServer()
}

func RegisterGatewayLeagueServiceServer(s grpc.ServiceRegistrar, srv GatewayLeagueServiceServer) {
	s.RegisterService(&GatewayLeagueService_ServiceDesc, srv)
}

func _GatewayLeagueService_GetLeagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayLeagueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).GetLeagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/GetLeagues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).GetLeagues(ctx, req.(*GatewayLeagueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayLeagueService_UpdateLeagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayLeagueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).UpdateLeagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/UpdateLeagues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).UpdateLeagues(ctx, req.(*GatewayLeagueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayLeagueService_UpdateChallenger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayApexUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).UpdateChallenger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/UpdateChallenger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).UpdateChallenger(ctx, req.(*GatewayApexUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayLeagueService_UpdateGrandmaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayApexUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).UpdateGrandmaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/UpdateGrandmaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).UpdateGrandmaster(ctx, req.(*GatewayApexUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayLeagueService_UpdateMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayApexUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).UpdateMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/UpdateMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).UpdateMaster(ctx, req.(*GatewayApexUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayLeagueService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayUpdateEntryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayLeagueServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.lol.league.v1.GatewayLeagueService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayLeagueServiceServer).UpdateEntry(ctx, req.(*GatewayUpdateEntryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayLeagueService_ServiceDesc is the grpc.ServiceDesc for GatewayLeagueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayLeagueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.lol.league.v1.GatewayLeagueService",
	HandlerType: (*GatewayLeagueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeagues",
			Handler:    _GatewayLeagueService_GetLeagues_Handler,
		},
		{
			MethodName: "UpdateLeagues",
			Handler:    _GatewayLeagueService_UpdateLeagues_Handler,
		},
		{
			MethodName: "UpdateChallenger",
			Handler:    _GatewayLeagueService_UpdateChallenger_Handler,
		},
		{
			MethodName: "UpdateGrandmaster",
			Handler:    _GatewayLeagueService_UpdateGrandmaster_Handler,
		},
		{
			MethodName: "UpdateMaster",
			Handler:    _GatewayLeagueService_UpdateMaster_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _GatewayLeagueService_UpdateEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/lol/league/v1/league.proto",
}
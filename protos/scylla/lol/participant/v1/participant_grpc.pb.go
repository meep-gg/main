// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: scylla/lol/participant/v1/participant.proto

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

// ScyllaParticipantServiceClient is the client API for ScyllaParticipantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScyllaParticipantServiceClient interface {
	// Participant
	GetParticipant(ctx context.Context, in *ScyllaParticipantReq, opts ...grpc.CallOption) (*ScyllaParticipant, error)
	GetParticipants(ctx context.Context, in *ScyllaGetParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error)
	GetMatchParticipants(ctx context.Context, in *ScyllaMatchParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error)
	PaginateParticipant(ctx context.Context, in *ScyllaPaginateParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error)
	UpsertParticipant(ctx context.Context, in *ScyllaUpsertParticipantReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteParticipant(ctx context.Context, in *ScyllaParticipantReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type scyllaParticipantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScyllaParticipantServiceClient(cc grpc.ClientConnInterface) ScyllaParticipantServiceClient {
	return &scyllaParticipantServiceClient{cc}
}

func (c *scyllaParticipantServiceClient) GetParticipant(ctx context.Context, in *ScyllaParticipantReq, opts ...grpc.CallOption) (*ScyllaParticipant, error) {
	out := new(ScyllaParticipant)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/GetParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scyllaParticipantServiceClient) GetParticipants(ctx context.Context, in *ScyllaGetParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error) {
	out := new(ScyllaParticipants)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/GetParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scyllaParticipantServiceClient) GetMatchParticipants(ctx context.Context, in *ScyllaMatchParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error) {
	out := new(ScyllaParticipants)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/GetMatchParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scyllaParticipantServiceClient) PaginateParticipant(ctx context.Context, in *ScyllaPaginateParticipantsReq, opts ...grpc.CallOption) (*ScyllaParticipants, error) {
	out := new(ScyllaParticipants)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/PaginateParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scyllaParticipantServiceClient) UpsertParticipant(ctx context.Context, in *ScyllaUpsertParticipantReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/UpsertParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scyllaParticipantServiceClient) DeleteParticipant(ctx context.Context, in *ScyllaParticipantReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scylla.lol.participant.v1.ScyllaParticipantService/DeleteParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScyllaParticipantServiceServer is the server API for ScyllaParticipantService service.
// All implementations must embed UnimplementedScyllaParticipantServiceServer
// for forward compatibility
type ScyllaParticipantServiceServer interface {
	// Participant
	GetParticipant(context.Context, *ScyllaParticipantReq) (*ScyllaParticipant, error)
	GetParticipants(context.Context, *ScyllaGetParticipantsReq) (*ScyllaParticipants, error)
	GetMatchParticipants(context.Context, *ScyllaMatchParticipantsReq) (*ScyllaParticipants, error)
	PaginateParticipant(context.Context, *ScyllaPaginateParticipantsReq) (*ScyllaParticipants, error)
	UpsertParticipant(context.Context, *ScyllaUpsertParticipantReq) (*emptypb.Empty, error)
	DeleteParticipant(context.Context, *ScyllaParticipantReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedScyllaParticipantServiceServer()
}

// UnimplementedScyllaParticipantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScyllaParticipantServiceServer struct {
}

func (UnimplementedScyllaParticipantServiceServer) GetParticipant(context.Context, *ScyllaParticipantReq) (*ScyllaParticipant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipant not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) GetParticipants(context.Context, *ScyllaGetParticipantsReq) (*ScyllaParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipants not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) GetMatchParticipants(context.Context, *ScyllaMatchParticipantsReq) (*ScyllaParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchParticipants not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) PaginateParticipant(context.Context, *ScyllaPaginateParticipantsReq) (*ScyllaParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaginateParticipant not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) UpsertParticipant(context.Context, *ScyllaUpsertParticipantReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertParticipant not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) DeleteParticipant(context.Context, *ScyllaParticipantReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParticipant not implemented")
}
func (UnimplementedScyllaParticipantServiceServer) mustEmbedUnimplementedScyllaParticipantServiceServer() {
}

// UnsafeScyllaParticipantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScyllaParticipantServiceServer will
// result in compilation errors.
type UnsafeScyllaParticipantServiceServer interface {
	mustEmbedUnimplementedScyllaParticipantServiceServer()
}

func RegisterScyllaParticipantServiceServer(s grpc.ServiceRegistrar, srv ScyllaParticipantServiceServer) {
	s.RegisterService(&ScyllaParticipantService_ServiceDesc, srv)
}

func _ScyllaParticipantService_GetParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaParticipantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).GetParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/GetParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).GetParticipant(ctx, req.(*ScyllaParticipantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScyllaParticipantService_GetParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaGetParticipantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).GetParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/GetParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).GetParticipants(ctx, req.(*ScyllaGetParticipantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScyllaParticipantService_GetMatchParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaMatchParticipantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).GetMatchParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/GetMatchParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).GetMatchParticipants(ctx, req.(*ScyllaMatchParticipantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScyllaParticipantService_PaginateParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaPaginateParticipantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).PaginateParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/PaginateParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).PaginateParticipant(ctx, req.(*ScyllaPaginateParticipantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScyllaParticipantService_UpsertParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaUpsertParticipantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).UpsertParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/UpsertParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).UpsertParticipant(ctx, req.(*ScyllaUpsertParticipantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScyllaParticipantService_DeleteParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScyllaParticipantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScyllaParticipantServiceServer).DeleteParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scylla.lol.participant.v1.ScyllaParticipantService/DeleteParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScyllaParticipantServiceServer).DeleteParticipant(ctx, req.(*ScyllaParticipantReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ScyllaParticipantService_ServiceDesc is the grpc.ServiceDesc for ScyllaParticipantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScyllaParticipantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scylla.lol.participant.v1.ScyllaParticipantService",
	HandlerType: (*ScyllaParticipantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParticipant",
			Handler:    _ScyllaParticipantService_GetParticipant_Handler,
		},
		{
			MethodName: "GetParticipants",
			Handler:    _ScyllaParticipantService_GetParticipants_Handler,
		},
		{
			MethodName: "GetMatchParticipants",
			Handler:    _ScyllaParticipantService_GetMatchParticipants_Handler,
		},
		{
			MethodName: "PaginateParticipant",
			Handler:    _ScyllaParticipantService_PaginateParticipant_Handler,
		},
		{
			MethodName: "UpsertParticipant",
			Handler:    _ScyllaParticipantService_UpsertParticipant_Handler,
		},
		{
			MethodName: "DeleteParticipant",
			Handler:    _ScyllaParticipantService_DeleteParticipant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scylla/lol/participant/v1/participant.proto",
}
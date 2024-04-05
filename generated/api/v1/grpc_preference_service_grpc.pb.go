// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/grpc_preference_service.proto

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
	GRPCPreferencesService_Get_FullMethodName = "/v1.GRPCPreferencesService/Get"
)

// GRPCPreferencesServiceClient is the client API for GRPCPreferencesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCPreferencesServiceClient interface {
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Preferences, error)
}

type gRPCPreferencesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCPreferencesServiceClient(cc grpc.ClientConnInterface) GRPCPreferencesServiceClient {
	return &gRPCPreferencesServiceClient{cc}
}

func (c *gRPCPreferencesServiceClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Preferences, error) {
	out := new(Preferences)
	err := c.cc.Invoke(ctx, GRPCPreferencesService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCPreferencesServiceServer is the server API for GRPCPreferencesService service.
// All implementations should embed UnimplementedGRPCPreferencesServiceServer
// for forward compatibility
type GRPCPreferencesServiceServer interface {
	Get(context.Context, *Empty) (*Preferences, error)
}

// UnimplementedGRPCPreferencesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGRPCPreferencesServiceServer struct {
}

func (UnimplementedGRPCPreferencesServiceServer) Get(context.Context, *Empty) (*Preferences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeGRPCPreferencesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCPreferencesServiceServer will
// result in compilation errors.
type UnsafeGRPCPreferencesServiceServer interface {
	mustEmbedUnimplementedGRPCPreferencesServiceServer()
}

func RegisterGRPCPreferencesServiceServer(s grpc.ServiceRegistrar, srv GRPCPreferencesServiceServer) {
	s.RegisterService(&GRPCPreferencesService_ServiceDesc, srv)
}

func _GRPCPreferencesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCPreferencesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCPreferencesService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCPreferencesServiceServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCPreferencesService_ServiceDesc is the grpc.ServiceDesc for GRPCPreferencesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCPreferencesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GRPCPreferencesService",
	HandlerType: (*GRPCPreferencesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GRPCPreferencesService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/grpc_preference_service.proto",
}
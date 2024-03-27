// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/process_service.proto

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
	ProcessService_CountProcesses_FullMethodName                            = "/v1.ProcessService/CountProcesses"
	ProcessService_GetProcessesByDeployment_FullMethodName                  = "/v1.ProcessService/GetProcessesByDeployment"
	ProcessService_GetGroupedProcessByDeployment_FullMethodName             = "/v1.ProcessService/GetGroupedProcessByDeployment"
	ProcessService_GetGroupedProcessByDeploymentAndContainer_FullMethodName = "/v1.ProcessService/GetGroupedProcessByDeploymentAndContainer"
)

// ProcessServiceClient is the client API for ProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessServiceClient interface {
	// CountProcesses returns the count of processes.
	CountProcesses(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountProcessesResponse, error)
	// GetProcessesByDeployment returns the processes executed in the given deployment.
	GetProcessesByDeployment(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetProcessesResponse, error)
	// GetGroupedProcessByDeployment returns all the processes executed grouped by deployment.
	GetGroupedProcessByDeployment(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetGroupedProcessesResponse, error)
	// GetGroupedProcessByDeploymentAndContainer returns all the processes executed grouped by deployment and container.
	GetGroupedProcessByDeploymentAndContainer(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetGroupedProcessesWithContainerResponse, error)
}

type processServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessServiceClient(cc grpc.ClientConnInterface) ProcessServiceClient {
	return &processServiceClient{cc}
}

func (c *processServiceClient) CountProcesses(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountProcessesResponse, error) {
	out := new(CountProcessesResponse)
	err := c.cc.Invoke(ctx, ProcessService_CountProcesses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processServiceClient) GetProcessesByDeployment(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetProcessesResponse, error) {
	out := new(GetProcessesResponse)
	err := c.cc.Invoke(ctx, ProcessService_GetProcessesByDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processServiceClient) GetGroupedProcessByDeployment(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetGroupedProcessesResponse, error) {
	out := new(GetGroupedProcessesResponse)
	err := c.cc.Invoke(ctx, ProcessService_GetGroupedProcessByDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processServiceClient) GetGroupedProcessByDeploymentAndContainer(ctx context.Context, in *GetProcessesByDeploymentRequest, opts ...grpc.CallOption) (*GetGroupedProcessesWithContainerResponse, error) {
	out := new(GetGroupedProcessesWithContainerResponse)
	err := c.cc.Invoke(ctx, ProcessService_GetGroupedProcessByDeploymentAndContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServiceServer is the server API for ProcessService service.
// All implementations should embed UnimplementedProcessServiceServer
// for forward compatibility
type ProcessServiceServer interface {
	// CountProcesses returns the count of processes.
	CountProcesses(context.Context, *RawQuery) (*CountProcessesResponse, error)
	// GetProcessesByDeployment returns the processes executed in the given deployment.
	GetProcessesByDeployment(context.Context, *GetProcessesByDeploymentRequest) (*GetProcessesResponse, error)
	// GetGroupedProcessByDeployment returns all the processes executed grouped by deployment.
	GetGroupedProcessByDeployment(context.Context, *GetProcessesByDeploymentRequest) (*GetGroupedProcessesResponse, error)
	// GetGroupedProcessByDeploymentAndContainer returns all the processes executed grouped by deployment and container.
	GetGroupedProcessByDeploymentAndContainer(context.Context, *GetProcessesByDeploymentRequest) (*GetGroupedProcessesWithContainerResponse, error)
}

// UnimplementedProcessServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProcessServiceServer struct {
}

func (UnimplementedProcessServiceServer) CountProcesses(context.Context, *RawQuery) (*CountProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountProcesses not implemented")
}
func (UnimplementedProcessServiceServer) GetProcessesByDeployment(context.Context, *GetProcessesByDeploymentRequest) (*GetProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProcessesByDeployment not implemented")
}
func (UnimplementedProcessServiceServer) GetGroupedProcessByDeployment(context.Context, *GetProcessesByDeploymentRequest) (*GetGroupedProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupedProcessByDeployment not implemented")
}
func (UnimplementedProcessServiceServer) GetGroupedProcessByDeploymentAndContainer(context.Context, *GetProcessesByDeploymentRequest) (*GetGroupedProcessesWithContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupedProcessByDeploymentAndContainer not implemented")
}

// UnsafeProcessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessServiceServer will
// result in compilation errors.
type UnsafeProcessServiceServer interface {
	mustEmbedUnimplementedProcessServiceServer()
}

func RegisterProcessServiceServer(s grpc.ServiceRegistrar, srv ProcessServiceServer) {
	s.RegisterService(&ProcessService_ServiceDesc, srv)
}

func _ProcessService_CountProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).CountProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessService_CountProcesses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).CountProcesses(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessService_GetProcessesByDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessesByDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).GetProcessesByDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessService_GetProcessesByDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).GetProcessesByDeployment(ctx, req.(*GetProcessesByDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessService_GetGroupedProcessByDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessesByDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).GetGroupedProcessByDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessService_GetGroupedProcessByDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).GetGroupedProcessByDeployment(ctx, req.(*GetProcessesByDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessService_GetGroupedProcessByDeploymentAndContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessesByDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).GetGroupedProcessByDeploymentAndContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessService_GetGroupedProcessByDeploymentAndContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).GetGroupedProcessByDeploymentAndContainer(ctx, req.(*GetProcessesByDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessService_ServiceDesc is the grpc.ServiceDesc for ProcessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProcessService",
	HandlerType: (*ProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountProcesses",
			Handler:    _ProcessService_CountProcesses_Handler,
		},
		{
			MethodName: "GetProcessesByDeployment",
			Handler:    _ProcessService_GetProcessesByDeployment_Handler,
		},
		{
			MethodName: "GetGroupedProcessByDeployment",
			Handler:    _ProcessService_GetGroupedProcessByDeployment_Handler,
		},
		{
			MethodName: "GetGroupedProcessByDeploymentAndContainer",
			Handler:    _ProcessService_GetGroupedProcessByDeploymentAndContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/process_service.proto",
}

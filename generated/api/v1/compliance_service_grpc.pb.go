// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/compliance_service.proto

package v1

import (
	context "context"
	storage "github.com/stackrox/rox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ComplianceService_GetStandards_FullMethodName                   = "/v1.ComplianceService/GetStandards"
	ComplianceService_GetStandard_FullMethodName                    = "/v1.ComplianceService/GetStandard"
	ComplianceService_GetRunResults_FullMethodName                  = "/v1.ComplianceService/GetRunResults"
	ComplianceService_GetAggregatedResults_FullMethodName           = "/v1.ComplianceService/GetAggregatedResults"
	ComplianceService_UpdateComplianceStandardConfig_FullMethodName = "/v1.ComplianceService/UpdateComplianceStandardConfig"
)

// ComplianceServiceClient is the client API for ComplianceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceServiceClient interface {
	GetStandards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetComplianceStandardsResponse, error)
	GetStandard(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetComplianceStandardResponse, error)
	GetRunResults(ctx context.Context, in *GetComplianceRunResultsRequest, opts ...grpc.CallOption) (*GetComplianceRunResultsResponse, error)
	GetAggregatedResults(ctx context.Context, in *ComplianceAggregationRequest, opts ...grpc.CallOption) (*storage.ComplianceAggregation_Response, error)
	UpdateComplianceStandardConfig(ctx context.Context, in *UpdateComplianceRequest, opts ...grpc.CallOption) (*Empty, error)
}

type complianceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceServiceClient(cc grpc.ClientConnInterface) ComplianceServiceClient {
	return &complianceServiceClient{cc}
}

func (c *complianceServiceClient) GetStandards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetComplianceStandardsResponse, error) {
	out := new(GetComplianceStandardsResponse)
	err := c.cc.Invoke(ctx, ComplianceService_GetStandards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceServiceClient) GetStandard(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetComplianceStandardResponse, error) {
	out := new(GetComplianceStandardResponse)
	err := c.cc.Invoke(ctx, ComplianceService_GetStandard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceServiceClient) GetRunResults(ctx context.Context, in *GetComplianceRunResultsRequest, opts ...grpc.CallOption) (*GetComplianceRunResultsResponse, error) {
	out := new(GetComplianceRunResultsResponse)
	err := c.cc.Invoke(ctx, ComplianceService_GetRunResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceServiceClient) GetAggregatedResults(ctx context.Context, in *ComplianceAggregationRequest, opts ...grpc.CallOption) (*storage.ComplianceAggregation_Response, error) {
	out := new(storage.ComplianceAggregation_Response)
	err := c.cc.Invoke(ctx, ComplianceService_GetAggregatedResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceServiceClient) UpdateComplianceStandardConfig(ctx context.Context, in *UpdateComplianceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ComplianceService_UpdateComplianceStandardConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceServiceServer is the server API for ComplianceService service.
// All implementations should embed UnimplementedComplianceServiceServer
// for forward compatibility
type ComplianceServiceServer interface {
	GetStandards(context.Context, *Empty) (*GetComplianceStandardsResponse, error)
	GetStandard(context.Context, *ResourceByID) (*GetComplianceStandardResponse, error)
	GetRunResults(context.Context, *GetComplianceRunResultsRequest) (*GetComplianceRunResultsResponse, error)
	GetAggregatedResults(context.Context, *ComplianceAggregationRequest) (*storage.ComplianceAggregation_Response, error)
	UpdateComplianceStandardConfig(context.Context, *UpdateComplianceRequest) (*Empty, error)
}

// UnimplementedComplianceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedComplianceServiceServer struct {
}

func (UnimplementedComplianceServiceServer) GetStandards(context.Context, *Empty) (*GetComplianceStandardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandards not implemented")
}
func (UnimplementedComplianceServiceServer) GetStandard(context.Context, *ResourceByID) (*GetComplianceStandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandard not implemented")
}
func (UnimplementedComplianceServiceServer) GetRunResults(context.Context, *GetComplianceRunResultsRequest) (*GetComplianceRunResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunResults not implemented")
}
func (UnimplementedComplianceServiceServer) GetAggregatedResults(context.Context, *ComplianceAggregationRequest) (*storage.ComplianceAggregation_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregatedResults not implemented")
}
func (UnimplementedComplianceServiceServer) UpdateComplianceStandardConfig(context.Context, *UpdateComplianceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComplianceStandardConfig not implemented")
}

// UnsafeComplianceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceServiceServer will
// result in compilation errors.
type UnsafeComplianceServiceServer interface {
	mustEmbedUnimplementedComplianceServiceServer()
}

func RegisterComplianceServiceServer(s grpc.ServiceRegistrar, srv ComplianceServiceServer) {
	s.RegisterService(&ComplianceService_ServiceDesc, srv)
}

func _ComplianceService_GetStandards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceServiceServer).GetStandards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceService_GetStandards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceServiceServer).GetStandards(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceService_GetStandard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceServiceServer).GetStandard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceService_GetStandard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceServiceServer).GetStandard(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceService_GetRunResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComplianceRunResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceServiceServer).GetRunResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceService_GetRunResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceServiceServer).GetRunResults(ctx, req.(*GetComplianceRunResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceService_GetAggregatedResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplianceAggregationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceServiceServer).GetAggregatedResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceService_GetAggregatedResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceServiceServer).GetAggregatedResults(ctx, req.(*ComplianceAggregationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceService_UpdateComplianceStandardConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateComplianceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceServiceServer).UpdateComplianceStandardConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceService_UpdateComplianceStandardConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceServiceServer).UpdateComplianceStandardConfig(ctx, req.(*UpdateComplianceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplianceService_ServiceDesc is the grpc.ServiceDesc for ComplianceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ComplianceService",
	HandlerType: (*ComplianceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStandards",
			Handler:    _ComplianceService_GetStandards_Handler,
		},
		{
			MethodName: "GetStandard",
			Handler:    _ComplianceService_GetStandard_Handler,
		},
		{
			MethodName: "GetRunResults",
			Handler:    _ComplianceService_GetRunResults_Handler,
		},
		{
			MethodName: "GetAggregatedResults",
			Handler:    _ComplianceService_GetAggregatedResults_Handler,
		},
		{
			MethodName: "UpdateComplianceStandardConfig",
			Handler:    _ComplianceService_UpdateComplianceStandardConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/compliance_service.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Api_GetRocketPoolNodes_FullMethodName = "/pb.Api/GetRocketPoolNodes"
	Api_GetOdaoNodes_FullMethodName       = "/pb.Api/GetOdaoNodes"
	Api_GetSoloValidators_FullMethodName  = "/pb.Api/GetSoloValidators"
	Api_ValidateEIP1271_FullMethodName    = "/pb.Api/ValidateEIP1271"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	GetRocketPoolNodes(ctx context.Context, in *RocketPoolNodesRequest, opts ...grpc.CallOption) (*RocketPoolNodes, error)
	GetOdaoNodes(ctx context.Context, in *OdaoNodesRequest, opts ...grpc.CallOption) (*OdaoNodes, error)
	GetSoloValidators(ctx context.Context, in *SoloValidatorsRequest, opts ...grpc.CallOption) (*SoloValidators, error)
	ValidateEIP1271(ctx context.Context, in *ValidateEIP1271Request, opts ...grpc.CallOption) (*ValidateEIP1271Response, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetRocketPoolNodes(ctx context.Context, in *RocketPoolNodesRequest, opts ...grpc.CallOption) (*RocketPoolNodes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RocketPoolNodes)
	err := c.cc.Invoke(ctx, Api_GetRocketPoolNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOdaoNodes(ctx context.Context, in *OdaoNodesRequest, opts ...grpc.CallOption) (*OdaoNodes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OdaoNodes)
	err := c.cc.Invoke(ctx, Api_GetOdaoNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetSoloValidators(ctx context.Context, in *SoloValidatorsRequest, opts ...grpc.CallOption) (*SoloValidators, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SoloValidators)
	err := c.cc.Invoke(ctx, Api_GetSoloValidators_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ValidateEIP1271(ctx context.Context, in *ValidateEIP1271Request, opts ...grpc.CallOption) (*ValidateEIP1271Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateEIP1271Response)
	err := c.cc.Invoke(ctx, Api_ValidateEIP1271_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility.
type ApiServer interface {
	GetRocketPoolNodes(context.Context, *RocketPoolNodesRequest) (*RocketPoolNodes, error)
	GetOdaoNodes(context.Context, *OdaoNodesRequest) (*OdaoNodes, error)
	GetSoloValidators(context.Context, *SoloValidatorsRequest) (*SoloValidators, error)
	ValidateEIP1271(context.Context, *ValidateEIP1271Request) (*ValidateEIP1271Response, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiServer struct{}

func (UnimplementedApiServer) GetRocketPoolNodes(context.Context, *RocketPoolNodesRequest) (*RocketPoolNodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRocketPoolNodes not implemented")
}
func (UnimplementedApiServer) GetOdaoNodes(context.Context, *OdaoNodesRequest) (*OdaoNodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOdaoNodes not implemented")
}
func (UnimplementedApiServer) GetSoloValidators(context.Context, *SoloValidatorsRequest) (*SoloValidators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoloValidators not implemented")
}
func (UnimplementedApiServer) ValidateEIP1271(context.Context, *ValidateEIP1271Request) (*ValidateEIP1271Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateEIP1271 not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}
func (UnimplementedApiServer) testEmbeddedByValue()             {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	// If the following call pancis, it indicates UnimplementedApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_GetRocketPoolNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RocketPoolNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetRocketPoolNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetRocketPoolNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetRocketPoolNodes(ctx, req.(*RocketPoolNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOdaoNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OdaoNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOdaoNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetOdaoNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOdaoNodes(ctx, req.(*OdaoNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetSoloValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoloValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetSoloValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetSoloValidators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetSoloValidators(ctx, req.(*SoloValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ValidateEIP1271_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateEIP1271Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ValidateEIP1271(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ValidateEIP1271_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ValidateEIP1271(ctx, req.(*ValidateEIP1271Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRocketPoolNodes",
			Handler:    _Api_GetRocketPoolNodes_Handler,
		},
		{
			MethodName: "GetOdaoNodes",
			Handler:    _Api_GetOdaoNodes_Handler,
		},
		{
			MethodName: "GetSoloValidators",
			Handler:    _Api_GetSoloValidators_Handler,
		},
		{
			MethodName: "ValidateEIP1271",
			Handler:    _Api_ValidateEIP1271_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

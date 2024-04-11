// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/networkservices/v1/dep.proto

package networkservicespb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DepServiceClient is the client API for DepService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepServiceClient interface {
	// Lists `LbTrafficExtension` resources in a given project and location.
	ListLbTrafficExtensions(ctx context.Context, in *ListLbTrafficExtensionsRequest, opts ...grpc.CallOption) (*ListLbTrafficExtensionsResponse, error)
	// Gets details of the specified `LbTrafficExtension` resource.
	GetLbTrafficExtension(ctx context.Context, in *GetLbTrafficExtensionRequest, opts ...grpc.CallOption) (*LbTrafficExtension, error)
	// Creates a new `LbTrafficExtension` resource in a given project and
	// location.
	CreateLbTrafficExtension(ctx context.Context, in *CreateLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the parameters of the specified `LbTrafficExtension` resource.
	UpdateLbTrafficExtension(ctx context.Context, in *UpdateLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes the specified `LbTrafficExtension` resource.
	DeleteLbTrafficExtension(ctx context.Context, in *DeleteLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists `LbRouteExtension` resources in a given project and location.
	ListLbRouteExtensions(ctx context.Context, in *ListLbRouteExtensionsRequest, opts ...grpc.CallOption) (*ListLbRouteExtensionsResponse, error)
	// Gets details of the specified `LbRouteExtension` resource.
	GetLbRouteExtension(ctx context.Context, in *GetLbRouteExtensionRequest, opts ...grpc.CallOption) (*LbRouteExtension, error)
	// Creates a new `LbRouteExtension` resource in a given project and location.
	CreateLbRouteExtension(ctx context.Context, in *CreateLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the parameters of the specified `LbRouteExtension` resource.
	UpdateLbRouteExtension(ctx context.Context, in *UpdateLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes the specified `LbRouteExtension` resource.
	DeleteLbRouteExtension(ctx context.Context, in *DeleteLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type depServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepServiceClient(cc grpc.ClientConnInterface) DepServiceClient {
	return &depServiceClient{cc}
}

func (c *depServiceClient) ListLbTrafficExtensions(ctx context.Context, in *ListLbTrafficExtensionsRequest, opts ...grpc.CallOption) (*ListLbTrafficExtensionsResponse, error) {
	out := new(ListLbTrafficExtensionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/ListLbTrafficExtensions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) GetLbTrafficExtension(ctx context.Context, in *GetLbTrafficExtensionRequest, opts ...grpc.CallOption) (*LbTrafficExtension, error) {
	out := new(LbTrafficExtension)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/GetLbTrafficExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) CreateLbTrafficExtension(ctx context.Context, in *CreateLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/CreateLbTrafficExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) UpdateLbTrafficExtension(ctx context.Context, in *UpdateLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/UpdateLbTrafficExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) DeleteLbTrafficExtension(ctx context.Context, in *DeleteLbTrafficExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/DeleteLbTrafficExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) ListLbRouteExtensions(ctx context.Context, in *ListLbRouteExtensionsRequest, opts ...grpc.CallOption) (*ListLbRouteExtensionsResponse, error) {
	out := new(ListLbRouteExtensionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/ListLbRouteExtensions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) GetLbRouteExtension(ctx context.Context, in *GetLbRouteExtensionRequest, opts ...grpc.CallOption) (*LbRouteExtension, error) {
	out := new(LbRouteExtension)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/GetLbRouteExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) CreateLbRouteExtension(ctx context.Context, in *CreateLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/CreateLbRouteExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) UpdateLbRouteExtension(ctx context.Context, in *UpdateLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/UpdateLbRouteExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depServiceClient) DeleteLbRouteExtension(ctx context.Context, in *DeleteLbRouteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.networkservices.v1.DepService/DeleteLbRouteExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepServiceServer is the server API for DepService service.
// All implementations must embed UnimplementedDepServiceServer
// for forward compatibility
type DepServiceServer interface {
	// Lists `LbTrafficExtension` resources in a given project and location.
	ListLbTrafficExtensions(context.Context, *ListLbTrafficExtensionsRequest) (*ListLbTrafficExtensionsResponse, error)
	// Gets details of the specified `LbTrafficExtension` resource.
	GetLbTrafficExtension(context.Context, *GetLbTrafficExtensionRequest) (*LbTrafficExtension, error)
	// Creates a new `LbTrafficExtension` resource in a given project and
	// location.
	CreateLbTrafficExtension(context.Context, *CreateLbTrafficExtensionRequest) (*longrunningpb.Operation, error)
	// Updates the parameters of the specified `LbTrafficExtension` resource.
	UpdateLbTrafficExtension(context.Context, *UpdateLbTrafficExtensionRequest) (*longrunningpb.Operation, error)
	// Deletes the specified `LbTrafficExtension` resource.
	DeleteLbTrafficExtension(context.Context, *DeleteLbTrafficExtensionRequest) (*longrunningpb.Operation, error)
	// Lists `LbRouteExtension` resources in a given project and location.
	ListLbRouteExtensions(context.Context, *ListLbRouteExtensionsRequest) (*ListLbRouteExtensionsResponse, error)
	// Gets details of the specified `LbRouteExtension` resource.
	GetLbRouteExtension(context.Context, *GetLbRouteExtensionRequest) (*LbRouteExtension, error)
	// Creates a new `LbRouteExtension` resource in a given project and location.
	CreateLbRouteExtension(context.Context, *CreateLbRouteExtensionRequest) (*longrunningpb.Operation, error)
	// Updates the parameters of the specified `LbRouteExtension` resource.
	UpdateLbRouteExtension(context.Context, *UpdateLbRouteExtensionRequest) (*longrunningpb.Operation, error)
	// Deletes the specified `LbRouteExtension` resource.
	DeleteLbRouteExtension(context.Context, *DeleteLbRouteExtensionRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedDepServiceServer()
}

// UnimplementedDepServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepServiceServer struct {
}

func (UnimplementedDepServiceServer) ListLbTrafficExtensions(context.Context, *ListLbTrafficExtensionsRequest) (*ListLbTrafficExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLbTrafficExtensions not implemented")
}
func (UnimplementedDepServiceServer) GetLbTrafficExtension(context.Context, *GetLbTrafficExtensionRequest) (*LbTrafficExtension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLbTrafficExtension not implemented")
}
func (UnimplementedDepServiceServer) CreateLbTrafficExtension(context.Context, *CreateLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLbTrafficExtension not implemented")
}
func (UnimplementedDepServiceServer) UpdateLbTrafficExtension(context.Context, *UpdateLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLbTrafficExtension not implemented")
}
func (UnimplementedDepServiceServer) DeleteLbTrafficExtension(context.Context, *DeleteLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLbTrafficExtension not implemented")
}
func (UnimplementedDepServiceServer) ListLbRouteExtensions(context.Context, *ListLbRouteExtensionsRequest) (*ListLbRouteExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLbRouteExtensions not implemented")
}
func (UnimplementedDepServiceServer) GetLbRouteExtension(context.Context, *GetLbRouteExtensionRequest) (*LbRouteExtension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLbRouteExtension not implemented")
}
func (UnimplementedDepServiceServer) CreateLbRouteExtension(context.Context, *CreateLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLbRouteExtension not implemented")
}
func (UnimplementedDepServiceServer) UpdateLbRouteExtension(context.Context, *UpdateLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLbRouteExtension not implemented")
}
func (UnimplementedDepServiceServer) DeleteLbRouteExtension(context.Context, *DeleteLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLbRouteExtension not implemented")
}
func (UnimplementedDepServiceServer) mustEmbedUnimplementedDepServiceServer() {}

// UnsafeDepServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepServiceServer will
// result in compilation errors.
type UnsafeDepServiceServer interface {
	mustEmbedUnimplementedDepServiceServer()
}

func RegisterDepServiceServer(s grpc.ServiceRegistrar, srv DepServiceServer) {
	s.RegisterService(&DepService_ServiceDesc, srv)
}

func _DepService_ListLbTrafficExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLbTrafficExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).ListLbTrafficExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/ListLbTrafficExtensions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).ListLbTrafficExtensions(ctx, req.(*ListLbTrafficExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_GetLbTrafficExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLbTrafficExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).GetLbTrafficExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/GetLbTrafficExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).GetLbTrafficExtension(ctx, req.(*GetLbTrafficExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_CreateLbTrafficExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLbTrafficExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).CreateLbTrafficExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/CreateLbTrafficExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).CreateLbTrafficExtension(ctx, req.(*CreateLbTrafficExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_UpdateLbTrafficExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLbTrafficExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).UpdateLbTrafficExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/UpdateLbTrafficExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).UpdateLbTrafficExtension(ctx, req.(*UpdateLbTrafficExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_DeleteLbTrafficExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLbTrafficExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).DeleteLbTrafficExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/DeleteLbTrafficExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).DeleteLbTrafficExtension(ctx, req.(*DeleteLbTrafficExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_ListLbRouteExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLbRouteExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).ListLbRouteExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/ListLbRouteExtensions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).ListLbRouteExtensions(ctx, req.(*ListLbRouteExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_GetLbRouteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLbRouteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).GetLbRouteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/GetLbRouteExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).GetLbRouteExtension(ctx, req.(*GetLbRouteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_CreateLbRouteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLbRouteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).CreateLbRouteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/CreateLbRouteExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).CreateLbRouteExtension(ctx, req.(*CreateLbRouteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_UpdateLbRouteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLbRouteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).UpdateLbRouteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/UpdateLbRouteExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).UpdateLbRouteExtension(ctx, req.(*UpdateLbRouteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepService_DeleteLbRouteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLbRouteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServiceServer).DeleteLbRouteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.networkservices.v1.DepService/DeleteLbRouteExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServiceServer).DeleteLbRouteExtension(ctx, req.(*DeleteLbRouteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepService_ServiceDesc is the grpc.ServiceDesc for DepService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.networkservices.v1.DepService",
	HandlerType: (*DepServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLbTrafficExtensions",
			Handler:    _DepService_ListLbTrafficExtensions_Handler,
		},
		{
			MethodName: "GetLbTrafficExtension",
			Handler:    _DepService_GetLbTrafficExtension_Handler,
		},
		{
			MethodName: "CreateLbTrafficExtension",
			Handler:    _DepService_CreateLbTrafficExtension_Handler,
		},
		{
			MethodName: "UpdateLbTrafficExtension",
			Handler:    _DepService_UpdateLbTrafficExtension_Handler,
		},
		{
			MethodName: "DeleteLbTrafficExtension",
			Handler:    _DepService_DeleteLbTrafficExtension_Handler,
		},
		{
			MethodName: "ListLbRouteExtensions",
			Handler:    _DepService_ListLbRouteExtensions_Handler,
		},
		{
			MethodName: "GetLbRouteExtension",
			Handler:    _DepService_GetLbRouteExtension_Handler,
		},
		{
			MethodName: "CreateLbRouteExtension",
			Handler:    _DepService_CreateLbRouteExtension_Handler,
		},
		{
			MethodName: "UpdateLbRouteExtension",
			Handler:    _DepService_UpdateLbRouteExtension_Handler,
		},
		{
			MethodName: "DeleteLbRouteExtension",
			Handler:    _DepService_DeleteLbRouteExtension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/networkservices/v1/dep.proto",
}

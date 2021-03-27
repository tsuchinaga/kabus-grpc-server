// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kabuspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// KabusServiceClient is the client API for KabusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KabusServiceClient interface {
	GetRegisteredSymbols(ctx context.Context, in *GetRegisteredSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error)
	RegisterSymbols(ctx context.Context, in *RegisterSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error)
	UnregisterSymbols(ctx context.Context, in *UnregisterSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error)
	UnregisterAllSymbols(ctx context.Context, in *UnregisterAllSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error)
}

type kabusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKabusServiceClient(cc grpc.ClientConnInterface) KabusServiceClient {
	return &kabusServiceClient{cc}
}

func (c *kabusServiceClient) GetRegisteredSymbols(ctx context.Context, in *GetRegisteredSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error) {
	out := new(RegisteredSymbols)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetRegisteredSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) RegisterSymbols(ctx context.Context, in *RegisterSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error) {
	out := new(RegisteredSymbols)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/RegisterSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) UnregisterSymbols(ctx context.Context, in *UnregisterSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error) {
	out := new(RegisteredSymbols)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/UnregisterSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) UnregisterAllSymbols(ctx context.Context, in *UnregisterAllSymbolsRequest, opts ...grpc.CallOption) (*RegisteredSymbols, error) {
	out := new(RegisteredSymbols)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/UnregisterAllSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KabusServiceServer is the server API for KabusService service.
// All implementations must embed UnimplementedKabusServiceServer
// for forward compatibility
type KabusServiceServer interface {
	GetRegisteredSymbols(context.Context, *GetRegisteredSymbolsRequest) (*RegisteredSymbols, error)
	RegisterSymbols(context.Context, *RegisterSymbolsRequest) (*RegisteredSymbols, error)
	UnregisterSymbols(context.Context, *UnregisterSymbolsRequest) (*RegisteredSymbols, error)
	UnregisterAllSymbols(context.Context, *UnregisterAllSymbolsRequest) (*RegisteredSymbols, error)
	mustEmbedUnimplementedKabusServiceServer()
}

// UnimplementedKabusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKabusServiceServer struct {
}

func (UnimplementedKabusServiceServer) GetRegisteredSymbols(context.Context, *GetRegisteredSymbolsRequest) (*RegisteredSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredSymbols not implemented")
}
func (UnimplementedKabusServiceServer) RegisterSymbols(context.Context, *RegisterSymbolsRequest) (*RegisteredSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSymbols not implemented")
}
func (UnimplementedKabusServiceServer) UnregisterSymbols(context.Context, *UnregisterSymbolsRequest) (*RegisteredSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterSymbols not implemented")
}
func (UnimplementedKabusServiceServer) UnregisterAllSymbols(context.Context, *UnregisterAllSymbolsRequest) (*RegisteredSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterAllSymbols not implemented")
}
func (UnimplementedKabusServiceServer) mustEmbedUnimplementedKabusServiceServer() {}

// UnsafeKabusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KabusServiceServer will
// result in compilation errors.
type UnsafeKabusServiceServer interface {
	mustEmbedUnimplementedKabusServiceServer()
}

func RegisterKabusServiceServer(s grpc.ServiceRegistrar, srv KabusServiceServer) {
	s.RegisterService(&KabusService_ServiceDesc, srv)
}

func _KabusService_GetRegisteredSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetRegisteredSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetRegisteredSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetRegisteredSymbols(ctx, req.(*GetRegisteredSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_RegisterSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).RegisterSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/RegisterSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).RegisterSymbols(ctx, req.(*RegisterSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_UnregisterSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).UnregisterSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/UnregisterSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).UnregisterSymbols(ctx, req.(*UnregisterSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_UnregisterAllSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterAllSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).UnregisterAllSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/UnregisterAllSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).UnregisterAllSymbols(ctx, req.(*UnregisterAllSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KabusService_ServiceDesc is the grpc.ServiceDesc for KabusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KabusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kabuspb.KabusService",
	HandlerType: (*KabusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegisteredSymbols",
			Handler:    _KabusService_GetRegisteredSymbols_Handler,
		},
		{
			MethodName: "RegisterSymbols",
			Handler:    _KabusService_RegisterSymbols_Handler,
		},
		{
			MethodName: "UnregisterSymbols",
			Handler:    _KabusService_UnregisterSymbols_Handler,
		},
		{
			MethodName: "UnregisterAllSymbols",
			Handler:    _KabusService_UnregisterAllSymbols_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kabuspb/kabus.proto",
}

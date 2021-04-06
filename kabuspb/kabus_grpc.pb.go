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
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*Token, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*Token, error)
	GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*Board, error)
	GetSymbol(ctx context.Context, in *GetSymbolRequest, opts ...grpc.CallOption) (*Symbol, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*Orders, error)
	GetPositions(ctx context.Context, in *GetPositionsRequest, opts ...grpc.CallOption) (*Positions, error)
	GetFutureSymbolCodeInfo(ctx context.Context, in *GetFutureSymbolCodeInfoRequest, opts ...grpc.CallOption) (*SymbolCodeInfo, error)
	GetOptionSymbolCodeInfo(ctx context.Context, in *GetOptionSymbolCodeInfoRequest, opts ...grpc.CallOption) (*SymbolCodeInfo, error)
	GetPriceRanking(ctx context.Context, in *GetPriceRankingRequest, opts ...grpc.CallOption) (*PriceRanking, error)
	GetTickRanking(ctx context.Context, in *GetTickRankingRequest, opts ...grpc.CallOption) (*TickRanking, error)
	GetVolumeRanking(ctx context.Context, in *GetVolumeRankingRequest, opts ...grpc.CallOption) (*VolumeRanking, error)
	GetValueRanking(ctx context.Context, in *GetValueRankingRequest, opts ...grpc.CallOption) (*ValueRanking, error)
	GetMarginRanking(ctx context.Context, in *GetMarginRankingRequest, opts ...grpc.CallOption) (*MarginRanking, error)
	GetIndustryRanking(ctx context.Context, in *GetIndustryRankingRequest, opts ...grpc.CallOption) (*IndustryRanking, error)
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

func (c *kabusServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*Board, error) {
	out := new(Board)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetSymbol(ctx context.Context, in *GetSymbolRequest, opts ...grpc.CallOption) (*Symbol, error) {
	out := new(Symbol)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetSymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetPositions(ctx context.Context, in *GetPositionsRequest, opts ...grpc.CallOption) (*Positions, error) {
	out := new(Positions)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetFutureSymbolCodeInfo(ctx context.Context, in *GetFutureSymbolCodeInfoRequest, opts ...grpc.CallOption) (*SymbolCodeInfo, error) {
	out := new(SymbolCodeInfo)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetFutureSymbolCodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetOptionSymbolCodeInfo(ctx context.Context, in *GetOptionSymbolCodeInfoRequest, opts ...grpc.CallOption) (*SymbolCodeInfo, error) {
	out := new(SymbolCodeInfo)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetOptionSymbolCodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetPriceRanking(ctx context.Context, in *GetPriceRankingRequest, opts ...grpc.CallOption) (*PriceRanking, error) {
	out := new(PriceRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetPriceRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetTickRanking(ctx context.Context, in *GetTickRankingRequest, opts ...grpc.CallOption) (*TickRanking, error) {
	out := new(TickRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetTickRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetVolumeRanking(ctx context.Context, in *GetVolumeRankingRequest, opts ...grpc.CallOption) (*VolumeRanking, error) {
	out := new(VolumeRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetVolumeRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetValueRanking(ctx context.Context, in *GetValueRankingRequest, opts ...grpc.CallOption) (*ValueRanking, error) {
	out := new(ValueRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetValueRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetMarginRanking(ctx context.Context, in *GetMarginRankingRequest, opts ...grpc.CallOption) (*MarginRanking, error) {
	out := new(MarginRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetMarginRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabusServiceClient) GetIndustryRanking(ctx context.Context, in *GetIndustryRankingRequest, opts ...grpc.CallOption) (*IndustryRanking, error) {
	out := new(IndustryRanking)
	err := c.cc.Invoke(ctx, "/kabuspb.KabusService/GetIndustryRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	GetToken(context.Context, *GetTokenRequest) (*Token, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*Token, error)
	GetBoard(context.Context, *GetBoardRequest) (*Board, error)
	GetSymbol(context.Context, *GetSymbolRequest) (*Symbol, error)
	GetOrders(context.Context, *GetOrdersRequest) (*Orders, error)
	GetPositions(context.Context, *GetPositionsRequest) (*Positions, error)
	GetFutureSymbolCodeInfo(context.Context, *GetFutureSymbolCodeInfoRequest) (*SymbolCodeInfo, error)
	GetOptionSymbolCodeInfo(context.Context, *GetOptionSymbolCodeInfoRequest) (*SymbolCodeInfo, error)
	GetPriceRanking(context.Context, *GetPriceRankingRequest) (*PriceRanking, error)
	GetTickRanking(context.Context, *GetTickRankingRequest) (*TickRanking, error)
	GetVolumeRanking(context.Context, *GetVolumeRankingRequest) (*VolumeRanking, error)
	GetValueRanking(context.Context, *GetValueRankingRequest) (*ValueRanking, error)
	GetMarginRanking(context.Context, *GetMarginRankingRequest) (*MarginRanking, error)
	GetIndustryRanking(context.Context, *GetIndustryRankingRequest) (*IndustryRanking, error)
	GetRegisteredSymbols(context.Context, *GetRegisteredSymbolsRequest) (*RegisteredSymbols, error)
	RegisterSymbols(context.Context, *RegisterSymbolsRequest) (*RegisteredSymbols, error)
	UnregisterSymbols(context.Context, *UnregisterSymbolsRequest) (*RegisteredSymbols, error)
	UnregisterAllSymbols(context.Context, *UnregisterAllSymbolsRequest) (*RegisteredSymbols, error)
	mustEmbedUnimplementedKabusServiceServer()
}

// UnimplementedKabusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKabusServiceServer struct {
}

func (UnimplementedKabusServiceServer) GetToken(context.Context, *GetTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedKabusServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedKabusServiceServer) GetBoard(context.Context, *GetBoardRequest) (*Board, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedKabusServiceServer) GetSymbol(context.Context, *GetSymbolRequest) (*Symbol, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbol not implemented")
}
func (UnimplementedKabusServiceServer) GetOrders(context.Context, *GetOrdersRequest) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedKabusServiceServer) GetPositions(context.Context, *GetPositionsRequest) (*Positions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedKabusServiceServer) GetFutureSymbolCodeInfo(context.Context, *GetFutureSymbolCodeInfoRequest) (*SymbolCodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFutureSymbolCodeInfo not implemented")
}
func (UnimplementedKabusServiceServer) GetOptionSymbolCodeInfo(context.Context, *GetOptionSymbolCodeInfoRequest) (*SymbolCodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionSymbolCodeInfo not implemented")
}
func (UnimplementedKabusServiceServer) GetPriceRanking(context.Context, *GetPriceRankingRequest) (*PriceRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceRanking not implemented")
}
func (UnimplementedKabusServiceServer) GetTickRanking(context.Context, *GetTickRankingRequest) (*TickRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickRanking not implemented")
}
func (UnimplementedKabusServiceServer) GetVolumeRanking(context.Context, *GetVolumeRankingRequest) (*VolumeRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolumeRanking not implemented")
}
func (UnimplementedKabusServiceServer) GetValueRanking(context.Context, *GetValueRankingRequest) (*ValueRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueRanking not implemented")
}
func (UnimplementedKabusServiceServer) GetMarginRanking(context.Context, *GetMarginRankingRequest) (*MarginRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginRanking not implemented")
}
func (UnimplementedKabusServiceServer) GetIndustryRanking(context.Context, *GetIndustryRankingRequest) (*IndustryRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndustryRanking not implemented")
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

func _KabusService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetBoard(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetSymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetSymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetSymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetSymbol(ctx, req.(*GetSymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetPositions(ctx, req.(*GetPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetFutureSymbolCodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFutureSymbolCodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetFutureSymbolCodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetFutureSymbolCodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetFutureSymbolCodeInfo(ctx, req.(*GetFutureSymbolCodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetOptionSymbolCodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOptionSymbolCodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetOptionSymbolCodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetOptionSymbolCodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetOptionSymbolCodeInfo(ctx, req.(*GetOptionSymbolCodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetPriceRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetPriceRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetPriceRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetPriceRanking(ctx, req.(*GetPriceRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetTickRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTickRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetTickRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetTickRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetTickRanking(ctx, req.(*GetTickRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetVolumeRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetVolumeRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetVolumeRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetVolumeRanking(ctx, req.(*GetVolumeRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetValueRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetValueRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetValueRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetValueRanking(ctx, req.(*GetValueRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetMarginRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarginRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetMarginRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetMarginRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetMarginRanking(ctx, req.(*GetMarginRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabusService_GetIndustryRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndustryRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabusServiceServer).GetIndustryRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kabuspb.KabusService/GetIndustryRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabusServiceServer).GetIndustryRanking(ctx, req.(*GetIndustryRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "GetToken",
			Handler:    _KabusService_GetToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _KabusService_RefreshToken_Handler,
		},
		{
			MethodName: "GetBoard",
			Handler:    _KabusService_GetBoard_Handler,
		},
		{
			MethodName: "GetSymbol",
			Handler:    _KabusService_GetSymbol_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _KabusService_GetOrders_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _KabusService_GetPositions_Handler,
		},
		{
			MethodName: "GetFutureSymbolCodeInfo",
			Handler:    _KabusService_GetFutureSymbolCodeInfo_Handler,
		},
		{
			MethodName: "GetOptionSymbolCodeInfo",
			Handler:    _KabusService_GetOptionSymbolCodeInfo_Handler,
		},
		{
			MethodName: "GetPriceRanking",
			Handler:    _KabusService_GetPriceRanking_Handler,
		},
		{
			MethodName: "GetTickRanking",
			Handler:    _KabusService_GetTickRanking_Handler,
		},
		{
			MethodName: "GetVolumeRanking",
			Handler:    _KabusService_GetVolumeRanking_Handler,
		},
		{
			MethodName: "GetValueRanking",
			Handler:    _KabusService_GetValueRanking_Handler,
		},
		{
			MethodName: "GetMarginRanking",
			Handler:    _KabusService_GetMarginRanking_Handler,
		},
		{
			MethodName: "GetIndustryRanking",
			Handler:    _KabusService_GetIndustryRanking_Handler,
		},
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

package server

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
)

func NewServer(
	security repositories.Security,
	tokenService services.TokenService,
	registerSymbolService services.RegisterSymbolService,
	boardStreamService services.BoardStreamService) kabuspb.KabusServiceServer {
	return &server{security: security, tokenService: tokenService, registerSymbolService: registerSymbolService, boardStreamService: boardStreamService}
}

type server struct {
	kabuspb.UnimplementedKabusServiceServer
	security              repositories.Security
	tokenService          services.TokenService
	registerSymbolService services.RegisterSymbolService
	boardStreamService    services.BoardStreamService
}

func (s *server) SendStockOrder(ctx context.Context, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderStock(ctx, token, req)
}

func (s *server) SendMarginOrder(ctx context.Context, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderMargin(ctx, token, req)
}

func (s *server) SendFutureOrder(ctx context.Context, req *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderFuture(ctx, token, req)
}

func (s *server) SendOptionOrder(ctx context.Context, req *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderOption(ctx, token, req)
}

func (s *server) CancelOrder(ctx context.Context, req *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.CancelOrder(ctx, token, req)
}

func (s *server) GetStockWallet(ctx context.Context, req *kabuspb.GetStockWalletRequest) (*kabuspb.StockWallet, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.GetStockWallet(ctx, token, req)
}

func (s *server) GetMarginWallet(ctx context.Context, req *kabuspb.GetMarginWalletRequest) (*kabuspb.MarginWallet, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.GetMarginWallet(ctx, token, req)
}

func (s *server) GetFutureWallet(ctx context.Context, req *kabuspb.GetFutureWalletRequest) (*kabuspb.FutureWallet, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.GetFutureWallet(ctx, token, req)
}

func (s *server) GetOptionWallet(ctx context.Context, req *kabuspb.GetOptionWalletRequest) (*kabuspb.OptionWallet, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.GetOptionWallet(ctx, token, req)
}

func (s *server) GetBoard(ctx context.Context, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Board(ctx, token, req)
}

func (s *server) GetOrders(ctx context.Context, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Orders(ctx, token, req)
}

func (s *server) GetPositions(ctx context.Context, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Positions(ctx, token, req)
}

func (s *server) GetFutureSymbolCodeInfo(ctx context.Context, req *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.SymbolNameFuture(ctx, token, req)
}

func (s *server) GetOptionSymbolCodeInfo(ctx context.Context, req *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.SymbolNameOption(ctx, token, req)
}

func (s *server) GetRegisteredSymbols(_ context.Context, req *kabuspb.GetRegisteredSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) RegisterSymbols(ctx context.Context, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := s.security.RegisterSymbols(ctx, token, req); err != nil {
		return nil, err
	}

	s.registerSymbolService.Add(req.RequesterName, req.Symbols)

	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) UnregisterSymbols(ctx context.Context, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := s.security.UnregisterSymbols(ctx, token, req); err != nil {
		return nil, err
	}

	s.registerSymbolService.Remove(req.RequesterName, req.Symbols)

	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) UnregisterAllSymbols(ctx context.Context, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := s.security.UnregisterAll(ctx, token, req); err != nil {
		return nil, err
	}

	s.registerSymbolService.Remove(req.RequesterName, s.registerSymbolService.Get(req.RequesterName))
	return &kabuspb.RegisteredSymbols{
		Symbols: []*kabuspb.RegisterSymbol{},
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) GetSymbol(ctx context.Context, req *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Symbol(ctx, token, req)
}

func (s *server) GetPriceRanking(ctx context.Context, req *kabuspb.GetPriceRankingRequest) (*kabuspb.PriceRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.PriceRanking(ctx, token, req)
}

func (s *server) GetTickRanking(ctx context.Context, req *kabuspb.GetTickRankingRequest) (*kabuspb.TickRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.TickRanking(ctx, token, req)
}

func (s *server) GetVolumeRanking(ctx context.Context, req *kabuspb.GetVolumeRankingRequest) (*kabuspb.VolumeRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.VolumeRanking(ctx, token, req)
}

func (s *server) GetValueRanking(ctx context.Context, req *kabuspb.GetValueRankingRequest) (*kabuspb.ValueRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.ValueRanking(ctx, token, req)
}

func (s *server) GetMarginRanking(ctx context.Context, req *kabuspb.GetMarginRankingRequest) (*kabuspb.MarginRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.MarginRanking(ctx, token, req)
}

func (s *server) GetIndustryRanking(ctx context.Context, req *kabuspb.GetIndustryRankingRequest) (*kabuspb.IndustryRanking, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.IndustryRanking(ctx, token, req)
}

func (s *server) GetExchange(ctx context.Context, req *kabuspb.GetExchangeRequest) (*kabuspb.ExchangeInfo, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Exchange(ctx, token, req)
}

func (s *server) GetRegulation(ctx context.Context, req *kabuspb.GetRegulationRequest) (*kabuspb.Regulation, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Regulation(ctx, token, req)
}

func (s *server) GetPrimaryExchange(ctx context.Context, req *kabuspb.GetPrimaryExchangeRequest) (*kabuspb.PrimaryExchange, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.PrimaryExchange(ctx, token, req)
}

func (s *server) GetSoftLimit(ctx context.Context, req *kabuspb.GetSoftLimitRequest) (*kabuspb.SoftLimit, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.SoftLimit(ctx, token, req)
}

func (s *server) GetBoardsStreaming(_ *kabuspb.GetBoardsStreamingRequest, stream kabuspb.KabusService_GetBoardsStreamingServer) error {
	return s.boardStreamService.Connect(stream)
}

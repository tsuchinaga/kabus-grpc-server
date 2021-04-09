package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
)

func NewServer(security repositories.Security, tokenService services.TokenService, registerSymbolService services.RegisterSymbolService, setting repositories.Setting) kabuspb.KabusServiceServer {
	return &server{security: security, tokenService: tokenService, registerSymbolService: registerSymbolService, setting: setting}
}

type server struct {
	kabuspb.UnimplementedKabusServiceServer
	security              repositories.Security
	tokenService          services.TokenService
	registerSymbolService services.RegisterSymbolService
	setting               repositories.Setting
}

func (s *server) GetToken(context.Context, *kabuspb.GetTokenRequest) (*kabuspb.Token, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}
	return &kabuspb.Token{Token: token, ExpiredAt: timestamppb.New(s.tokenService.GetExpiredAt())}, nil
}

func (s *server) RefreshToken(context.Context, *kabuspb.RefreshTokenRequest) (*kabuspb.Token, error) {
	token, err := s.tokenService.Refresh(context.Background())
	if err != nil {
		return nil, err
	}
	return &kabuspb.Token{Token: token, ExpiredAt: timestamppb.New(s.tokenService.GetExpiredAt())}, nil
}

func (s *server) SendStockOrder(ctx context.Context, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderStock(ctx, token, req, s.setting.Password())
}

func (s *server) SendMarginOrder(ctx context.Context, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderMargin(ctx, token, req, s.setting.Password())
}

func (s *server) SendFutureOrder(ctx context.Context, req *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderFuture(ctx, token, req, s.setting.Password())
}

func (s *server) SendOptionOrder(ctx context.Context, req *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error) {
	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return s.security.SendOrderOption(ctx, token, req, s.setting.Password())
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

func (s *server) GetRegisteredSymbols(context.Context, *kabuspb.GetRegisteredSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return &kabuspb.RegisteredSymbols{Symbols: s.registerSymbolService.Get()}, nil
}

func (s *server) RegisterSymbols(ctx context.Context, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.RegisterSymbols(ctx, token, req)
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Set(res.Symbols)
	return res, nil
}

func (s *server) UnregisterSymbols(ctx context.Context, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.UnregisterSymbols(ctx, token, req)
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Set(res.Symbols)
	return res, err
}

func (s *server) UnregisterAllSymbols(ctx context.Context, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.UnregisterAll(ctx, token, req)
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Set(res.Symbols)
	return res, err
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

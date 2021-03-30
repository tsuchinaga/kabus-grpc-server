package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
)

func NewServer(security repositories.Security, tokenService services.TokenService, registerSymbolService services.RegisterSymbolService) kabuspb.KabusServiceServer {
	return &server{security: security, tokenService: tokenService, registerSymbolService: registerSymbolService}
}

type server struct {
	kabuspb.UnimplementedKabusServiceServer
	security              repositories.Security
	tokenService          services.TokenService
	registerSymbolService services.RegisterSymbolService
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

func (s *server) GetBoard(ctx context.Context, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.security.Board(ctx, token, req)
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

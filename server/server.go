package server

import (
	"context"

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

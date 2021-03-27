package repositories

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

type Security interface {
	Token(ctx context.Context, password string) (string, error)
	RegisterSymbols(ctx context.Context, token string, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterSymbols(ctx context.Context, token string, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterAll(ctx context.Context, token string, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
}

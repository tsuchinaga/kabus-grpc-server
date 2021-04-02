package repositories

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

type Security interface {
	Board(ctx context.Context, token string, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error)
	Symbol(ctx context.Context, token string, req *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error)
	Orders(ctx context.Context, token string, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error)
	SymbolNameFuture(ctx context.Context, token string, req *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error)
	SymbolNameOption(ctx context.Context, token string, req *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error)
	Token(ctx context.Context, password string) (string, error)
	RegisterSymbols(ctx context.Context, token string, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterSymbols(ctx context.Context, token string, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterAll(ctx context.Context, token string, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
}

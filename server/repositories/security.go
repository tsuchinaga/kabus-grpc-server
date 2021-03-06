package repositories

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

type Security interface {
	Board(ctx context.Context, token string, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error)
	Symbol(ctx context.Context, token string, req *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error)
	Orders(ctx context.Context, token string, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error)
	Positions(ctx context.Context, token string, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error)
	SymbolNameFuture(ctx context.Context, token string, req *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error)
	SymbolNameOption(ctx context.Context, token string, req *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error)
	Token(ctx context.Context, password string) (string, error)
	RegisterSymbols(ctx context.Context, token string, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterSymbols(ctx context.Context, token string, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	UnregisterAll(ctx context.Context, token string, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error)
	PriceRanking(ctx context.Context, token string, req *kabuspb.GetPriceRankingRequest) (*kabuspb.PriceRanking, error)
	TickRanking(ctx context.Context, token string, req *kabuspb.GetTickRankingRequest) (*kabuspb.TickRanking, error)
	VolumeRanking(ctx context.Context, token string, req *kabuspb.GetVolumeRankingRequest) (*kabuspb.VolumeRanking, error)
	ValueRanking(ctx context.Context, token string, req *kabuspb.GetValueRankingRequest) (*kabuspb.ValueRanking, error)
	MarginRanking(ctx context.Context, token string, req *kabuspb.GetMarginRankingRequest) (*kabuspb.MarginRanking, error)
	IndustryRanking(ctx context.Context, token string, req *kabuspb.GetIndustryRankingRequest) (*kabuspb.IndustryRanking, error)
	SendOrderStock(ctx context.Context, token string, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error)
	SendOrderMargin(ctx context.Context, token string, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error)
	SendOrderFuture(ctx context.Context, token string, req *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error)
	SendOrderOption(ctx context.Context, token string, req *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error)
	CancelOrder(ctx context.Context, token string, req *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error)
	GetStockWallet(ctx context.Context, token string, req *kabuspb.GetStockWalletRequest) (*kabuspb.StockWallet, error)
	GetMarginWallet(ctx context.Context, token string, req *kabuspb.GetMarginWalletRequest) (*kabuspb.MarginWallet, error)
	GetFutureWallet(ctx context.Context, token string, req *kabuspb.GetFutureWalletRequest) (*kabuspb.FutureWallet, error)
	GetOptionWallet(ctx context.Context, token string, req *kabuspb.GetOptionWalletRequest) (*kabuspb.OptionWallet, error)
	Exchange(ctx context.Context, token string, req *kabuspb.GetExchangeRequest) (*kabuspb.ExchangeInfo, error)
	Regulation(ctx context.Context, token string, req *kabuspb.GetRegulationRequest) (*kabuspb.Regulation, error)
	PrimaryExchange(ctx context.Context, token string, req *kabuspb.GetPrimaryExchangeRequest) (*kabuspb.PrimaryExchange, error)
	SoftLimit(ctx context.Context, token string, req *kabuspb.GetSoftLimitRequest) (*kabuspb.SoftLimit, error)
	MarginPremium(ctx context.Context, token string, req *kabuspb.GetMarginPremiumRequest) (*kabuspb.MarginPremium, error)
	IsMissMatchApiKeyError(err error) bool
}

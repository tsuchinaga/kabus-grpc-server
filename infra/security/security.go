package security

import (
	"context"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func NewSecurity(restClient kabus.RESTClient) repositories.Security {
	return &security{restClient: restClient}
}

type security struct {
	restClient kabus.RESTClient
}

func (s *security) SymbolNameFuture(ctx context.Context, token string, req *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	res, err := s.restClient.SymbolNameFutureWithContext(ctx, token, kabus.SymbolNameFutureRequest{
		FutureCode: toFutureCode(req.FutureCode),
		DerivMonth: toYmNum(req.DerivativeMonth),
	})
	if err != nil {
		return nil, err
	}
	return &kabuspb.SymbolCodeInfo{Code: res.Symbol, Name: res.SymbolName}, nil
}

func (s *security) SymbolNameOption(ctx context.Context, token string, req *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	res, err := s.restClient.SymbolNameOptionWithContext(ctx, token, kabus.SymbolNameOptionRequest{
		DerivMonth:  toYmNum(req.DerivativeMonth),
		PutOrCall:   toPutOrCall(req.CallOrPut),
		StrikePrice: int(req.StrikePrice),
	})
	if err != nil {
		return nil, err
	}
	return &kabuspb.SymbolCodeInfo{Code: res.Symbol, Name: res.SymbolName}, nil
}

func (s *security) Token(ctx context.Context, password string) (string, error) {
	token, err := s.restClient.TokenWithContext(ctx, kabus.TokenRequest{APIPassword: password})
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func (s *security) RegisterSymbols(ctx context.Context, token string, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	symbols := make([]kabus.RegisterSymbol, len(req.Symbols))
	for i, symbol := range req.Symbols {
		symbols[i] = kabus.RegisterSymbol{Symbol: symbol.Symbol, Exchange: toExchange(symbol.Exchange)}
	}
	res, err := s.restClient.RegisterWithContext(ctx, token, kabus.RegisterRequest{Symbols: symbols})
	if err != nil {
		return nil, err
	}

	resSymbols := make([]*kabuspb.RegisterSymbol, len(res.RegisterList))
	for i, symbol := range res.RegisterList {
		resSymbols[i] = &kabuspb.RegisterSymbol{Symbol: symbol.Symbol, Exchange: fromExchange(symbol.Exchange)}
	}
	return &kabuspb.RegisteredSymbols{Symbols: resSymbols}, nil
}

func (s *security) UnregisterSymbols(ctx context.Context, token string, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	symbols := make([]kabus.UnregisterSymbol, len(req.Symbols))
	for i, symbol := range req.Symbols {
		symbols[i] = kabus.UnregisterSymbol{Symbol: symbol.Symbol, Exchange: toExchange(symbol.Exchange)}
	}
	res, err := s.restClient.UnregisterWithContext(ctx, token, kabus.UnregisterRequest{Symbols: symbols})
	if err != nil {
		return nil, err
	}

	resSymbols := make([]*kabuspb.RegisterSymbol, len(res.RegisterList))
	for i, symbol := range res.RegisterList {
		resSymbols[i] = &kabuspb.RegisterSymbol{Symbol: symbol.Symbol, Exchange: fromExchange(symbol.Exchange)}
	}
	return &kabuspb.RegisteredSymbols{Symbols: resSymbols}, nil
}

func (s *security) UnregisterAll(ctx context.Context, token string, _ *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	res, err := s.restClient.UnregisterAllWithContext(ctx, token)
	if err != nil {
		return nil, err
	}

	resSymbols := make([]*kabuspb.RegisterSymbol, len(res.RegistList))
	for i, symbol := range res.RegistList {
		resSymbols[i] = &kabuspb.RegisterSymbol{Symbol: symbol.Symbol, Exchange: fromExchange(symbol.Exchange)}
	}
	return &kabuspb.RegisteredSymbols{Symbols: resSymbols}, nil
}

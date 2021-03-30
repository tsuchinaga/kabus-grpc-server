package security

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

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

func (s *security) Board(ctx context.Context, token string, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error) {
	res, err := s.restClient.BoardWithContext(ctx, token, kabus.BoardRequest{Symbol: req.SymbolCode, Exchange: toExchange(req.Exchange)})
	if err != nil {
		return nil, err
	}
	return &kabuspb.Board{
		SymbolCode:               res.Symbol,
		SymbolName:               res.SymbolName,
		Exchange:                 fromExchange(res.Exchange),
		ExchangeName:             res.ExchangeName,
		CurrentPrice:             res.CurrentPrice,
		CurrentPriceTime:         timestamppb.New(res.CurrentPriceTime),
		CurrentPriceChangeStatus: fromCurrentPriceChangeStatus(res.CurrentPriceChangeStatus),
		CurrentPriceStatus:       fromCurrentPriceStatus(res.CurrentPriceStatus),
		CalculationPrice:         res.CalcPrice,
		PreviousClose:            res.PreviousClose,
		PreviousCloseTime:        timestamppb.New(res.PreviousCloseTime),
		ChangePreviousClose:      res.ChangePreviousClose,
		ChangePreviousClosePer:   res.ChangePreviousClosePer,
		OpeningPrice:             res.OpeningPrice,
		OpeningPriceTime:         timestamppb.New(res.OpeningPriceTime),
		HighPrice:                res.HighPrice,
		HighPriceTime:            timestamppb.New(res.HighPriceTime),
		LowPrice:                 res.LowPrice,
		LowPriceTime:             timestamppb.New(res.LowPriceTime),
		TradingVolume:            res.TradingVolume,
		TradingVolumeTime:        timestamppb.New(res.TradingVolumeTime),
		Vwap:                     res.VWAP,
		TradingValue:             res.TradingValue,
		BidQuantity:              res.BidQty,
		BidPrice:                 res.BidPrice,
		BidTime:                  timestamppb.New(res.BidTime),
		BidSign:                  fromBidAskSign(res.BidSign),
		MarketOrderSellQuantity:  res.MarketOrderSellQty,
		Sell1:                    fromFirstBoardSign(res.Sell1),
		Sell2:                    fromBoardSign(res.Sell2),
		Sell3:                    fromBoardSign(res.Sell3),
		Sell4:                    fromBoardSign(res.Sell4),
		Sell5:                    fromBoardSign(res.Sell5),
		Sell6:                    fromBoardSign(res.Sell6),
		Sell7:                    fromBoardSign(res.Sell7),
		Sell8:                    fromBoardSign(res.Sell8),
		Sell9:                    fromBoardSign(res.Sell9),
		Sell10:                   fromBoardSign(res.Sell10),
		AskQuantity:              res.AskQty,
		AskPrice:                 res.AskPrice,
		AskTime:                  timestamppb.New(res.AskTime),
		AskSign:                  fromBidAskSign(res.AskSign),
		MarketOrderBuyQuantity:   res.MarketOrderBuyQty,
		Buy1:                     fromFirstBoardSign(res.Buy1),
		Buy2:                     fromBoardSign(res.Buy2),
		Buy3:                     fromBoardSign(res.Buy3),
		Buy4:                     fromBoardSign(res.Buy4),
		Buy5:                     fromBoardSign(res.Buy5),
		Buy6:                     fromBoardSign(res.Buy6),
		Buy7:                     fromBoardSign(res.Buy7),
		Buy8:                     fromBoardSign(res.Buy8),
		Buy9:                     fromBoardSign(res.Buy9),
		Buy10:                    fromBoardSign(res.Buy10),
		OverSellQuantity:         res.OverSellQty,
		UnderBuyQuantity:         res.UnderBuyQty,
		TotalMarketValue:         res.TotalMarketValue,
		ClearingPrice:            res.ClearingPrice,
		ImpliedVolatility:        res.IV,
		Gamma:                    res.Gamma,
		Theta:                    res.Theta,
		Vega:                     res.Vega,
		Delta:                    res.Delta,
	}, nil
}

func (s *security) Symbol(ctx context.Context, token string, req *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error) {
	res, err := s.restClient.SymbolWithContext(ctx, token, kabus.SymbolRequest{Symbol: req.SymbolCode, Exchange: toExchange(req.Exchange)})
	if err != nil {
		return nil, err
	}
	return &kabuspb.Symbol{
		Code:               res.Symbol,
		Name:               res.SymbolName,
		DisplayName:        res.DisplayName,
		Exchange:           fromExchange(res.Exchange),
		ExchangeName:       res.ExchangeName,
		IndustryCategory:   res.BisCategory,
		TotalMarketValue:   res.TotalMarketValue,
		TotalStocks:        res.TotalStocks,
		TradingUnit:        res.TradingUnit,
		FiscalYearEndBasic: timestamppb.New(res.FiscalYearEndBasic.Time),
		PriceRangeGroup:    fromPriceRangeGroup(res.PriceRangeGroup),
		KabucomMarginBuy:   res.KCMarginBuy,
		KabucomMarginSell:  res.KCMarginSell,
		MarginBuy:          res.MarginBuy,
		MarginSell:         res.MarginSell,
		UpperLimit:         res.UpperLimit,
		LowerLimit:         res.LowerLimit,
		Underlyer:          fromUnderlyer(res.Underlyer),
		DerivativeMonth:    timestamppb.New(res.DerivMonth.Time),
		TradeStart:         timestamppb.New(res.TradeStart.Time),
		TradeEnd:           timestamppb.New(res.TradeEnd.Time),
		StrikePrice:        res.StrikePrice,
		CallOrPut:          fromPutOrCallNum(res.PutOrCall),
		ClearingPrice:      res.ClearingPrice,
	}, nil
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

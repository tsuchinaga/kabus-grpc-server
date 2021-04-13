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

func (s *security) Orders(ctx context.Context, token string, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	res, err := s.restClient.OrdersWithContext(ctx, token, kabus.OrdersRequest{
		Product:          toProduct(req.Product),
		ID:               req.Id,
		UpdateTime:       req.UpdateTime.AsTime(),
		IsGetOrderDetail: toIsGetOrderDetail(req.GetDetails),
		Symbol:           req.Symbol,
		State:            toOrderState(req.State),
		Side:             toSide(req.Side),
		CashMargin:       toCashMargin(req.TradeType),
	})
	if err != nil {
		return nil, err
	}

	return fromOrders(res), nil
}

func (s *security) Positions(ctx context.Context, token string, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	res, err := s.restClient.PositionsWithContext(ctx, token, kabus.PositionsRequest{
		Product: toProduct(req.Product),
		Symbol:  req.SymbolCode,
	})
	if err != nil {
		return nil, err
	}

	return fromPositions(res), nil
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

	resSymbols := make([]*kabuspb.RegisterSymbol, len(res.RegisterList))
	for i, symbol := range res.RegisterList {
		resSymbols[i] = &kabuspb.RegisterSymbol{Symbol: symbol.Symbol, Exchange: fromExchange(symbol.Exchange)}
	}
	return &kabuspb.RegisteredSymbols{Symbols: resSymbols}, nil
}

func (s *security) PriceRanking(ctx context.Context, token string, req *kabuspb.GetPriceRankingRequest) (*kabuspb.PriceRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             toRankingTypeFromPriceRankingType(req.RankingType),
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToPriceRanking(res), nil
}

func (s *security) TickRanking(ctx context.Context, token string, req *kabuspb.GetTickRankingRequest) (*kabuspb.TickRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             kabus.RankingTypeTickCount,
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToTickRanking(res), nil
}

func (s *security) VolumeRanking(ctx context.Context, token string, req *kabuspb.GetVolumeRankingRequest) (*kabuspb.VolumeRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             kabus.RankingTypeVolumeRapidIncrease,
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToVolumeRanking(res), nil
}

func (s *security) ValueRanking(ctx context.Context, token string, req *kabuspb.GetValueRankingRequest) (*kabuspb.ValueRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             kabus.RankingTypeValueRapidIncrease,
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToValueRanking(res), nil
}

func (s *security) MarginRanking(ctx context.Context, token string, req *kabuspb.GetMarginRankingRequest) (*kabuspb.MarginRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             toRankingTypeFromMarginRankingType(req.RankingType),
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToMarginRanking(res), nil
}

func (s *security) IndustryRanking(ctx context.Context, token string, req *kabuspb.GetIndustryRankingRequest) (*kabuspb.IndustryRanking, error) {
	res, err := s.restClient.RankingWithContext(ctx, token, kabus.RankingRequest{
		Type:             toRankingTypeFromIndustryRankingType(req.RankingType),
		ExchangeDivision: toExchangeDivision(req.ExchangeDivision),
	})
	if err != nil {
		return nil, err
	}
	return fromRankingToIndustryRanking(res), nil
}

func (s *security) SendOrderStock(ctx context.Context, token string, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.restClient.SendOrderStockWithContext(ctx, token, toSendOrderStockRequestFromSendStockOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return &kabuspb.OrderResponse{ResultCode: int32(res.Result), OrderId: res.OrderID}, nil
}

func (s *security) SendOrderMargin(ctx context.Context, token string, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.restClient.SendOrderStockWithContext(ctx, token, toSendOrderStockRequestFromSendMarginOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return &kabuspb.OrderResponse{ResultCode: int32(res.Result), OrderId: res.OrderID}, nil
}

func (s *security) SendOrderFuture(ctx context.Context, token string, req *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.restClient.SendOrderFutureWithContext(ctx, token, toSendOrderFutureRequest(req))
	if err != nil {
		return nil, err
	}
	return &kabuspb.OrderResponse{ResultCode: int32(res.Result), OrderId: res.OrderID}, nil
}

func (s *security) SendOrderOption(ctx context.Context, token string, req *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.restClient.SendOrderOptionWithContext(ctx, token, toSendOrderOptionRequest(req))
	if err != nil {
		return nil, err
	}
	return &kabuspb.OrderResponse{ResultCode: int32(res.Result), OrderId: res.OrderID}, nil
}

func (s *security) CancelOrder(ctx context.Context, token string, req *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
	res, err := s.restClient.CancelOrderWithContext(ctx, token, toCancelOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return &kabuspb.OrderResponse{ResultCode: int32(res.Result), OrderId: res.OrderID}, nil
}

func (s *security) GetStockWallet(ctx context.Context, token string, req *kabuspb.GetStockWalletRequest) (*kabuspb.StockWallet, error) {
	var (
		res *kabus.WalletCashResponse
		err error
	)
	if req.SymbolCode == "" {
		res, err = s.restClient.WalletCashWithContext(ctx, token)
	} else {
		res, err = s.restClient.WalletCashSymbolWithContext(ctx, token, toWalletCashSymbolRequest(req))
	}

	if err != nil {
		return nil, err
	}
	return &kabuspb.StockWallet{StockAccountWallet: res.StockAccountWallet}, nil
}
func (s *security) GetMarginWallet(ctx context.Context, token string, req *kabuspb.GetMarginWalletRequest) (*kabuspb.MarginWallet, error) {
	var (
		res *kabus.WalletMarginResponse
		err error
	)
	if req.SymbolCode == "" {
		res, err = s.restClient.WalletMarginWithContext(ctx, token)
	} else {
		res, err = s.restClient.WalletMarginSymbolWithContext(ctx, token, toWalletMarginSymbolRequest(req))
	}

	if err != nil {
		return nil, err
	}
	return &kabuspb.MarginWallet{
		MarginAccountWallet:          res.MarginAccountWallet,
		DepositKeepRate:              res.DepositkeepRate,
		ConsignmentDepositRate:       res.ConsignmentDepositRate,
		CashOfConsignmentDepositRate: res.CashOfConsignmentDepositRate,
	}, nil
}
func (s *security) GetFutureWallet(ctx context.Context, token string, req *kabuspb.GetFutureWalletRequest) (*kabuspb.FutureWallet, error) {
	var (
		res *kabus.WalletFutureResponse
		err error
	)
	if req.SymbolCode == "" {
		res, err = s.restClient.WalletFutureWithContext(ctx, token)
	} else {
		res, err = s.restClient.WalletFutureSymbolWithContext(ctx, token, toWalletFutureSymbolRequest(req))
	}

	if err != nil {
		return nil, err
	}
	return &kabuspb.FutureWallet{FutureTradeLimit: res.FutureTradeLimit, MarginRequirement: res.MarginRequirement}, nil
}
func (s *security) GetOptionWallet(ctx context.Context, token string, req *kabuspb.GetOptionWalletRequest) (*kabuspb.OptionWallet, error) {
	var (
		res *kabus.WalletOptionResponse
		err error
	)
	if req.SymbolCode == "" {
		res, err = s.restClient.WalletOptionWithContext(ctx, token)
	} else {
		res, err = s.restClient.WalletOptionSymbolWithContext(ctx, token, toWalletOptionSymbolRequest(req))
	}

	if err != nil {
		return nil, err
	}
	return &kabuspb.OptionWallet{
		OptionBuyTradeLimit:  res.OptionBuyTradeLimit,
		OptionSellTradeLimit: res.OptionSellTradeLimit,
		MarginRequirement:    res.MarginRequirement,
	}, nil
}

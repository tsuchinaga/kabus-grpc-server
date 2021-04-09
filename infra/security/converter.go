package security

import (
	"time"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toExchange(exchange kabuspb.Exchange) kabus.Exchange {
	switch exchange {
	case kabuspb.Exchange_EXCHANGE_TOUSHOU:
		return kabus.ExchangeToushou
	case kabuspb.Exchange_EXCHANGE_MEISHOU:
		return kabus.ExchangeMeishou
	case kabuspb.Exchange_EXCHANGE_FUKUSHOU:
		return kabus.ExchangeFukushou
	case kabuspb.Exchange_EXCHANGE_SATSUSHOU:
		return kabus.ExchangeSatsushou
	case kabuspb.Exchange_EXCHANGE_ALL_SESSION:
		return kabus.ExchangeAll
	case kabuspb.Exchange_EXCHANGE_DAY_SESSION:
		return kabus.ExchangeDaytime
	case kabuspb.Exchange_EXCHANGE_NIGHT_SESSION:
		return kabus.ExchangeEvening
	default:
		return kabus.ExchangeUnspecified
	}
}

func fromExchange(exchange kabus.Exchange) kabuspb.Exchange {
	switch exchange {
	case kabus.ExchangeToushou:
		return kabuspb.Exchange_EXCHANGE_TOUSHOU
	case kabus.ExchangeMeishou:
		return kabuspb.Exchange_EXCHANGE_MEISHOU
	case kabus.ExchangeFukushou:
		return kabuspb.Exchange_EXCHANGE_FUKUSHOU
	case kabus.ExchangeSatsushou:
		return kabuspb.Exchange_EXCHANGE_SATSUSHOU
	case kabus.ExchangeAll:
		return kabuspb.Exchange_EXCHANGE_ALL_SESSION
	case kabus.ExchangeDaytime:
		return kabuspb.Exchange_EXCHANGE_DAY_SESSION
	case kabus.ExchangeEvening:
		return kabuspb.Exchange_EXCHANGE_NIGHT_SESSION
	default:
		return kabuspb.Exchange_EXCHANGE_UNSPECIFIED
	}
}

func toFutureCode(futureCode kabuspb.FutureCode) kabus.FutureCode {
	switch futureCode {
	case kabuspb.FutureCode_FUTURE_CODE_NK225:
		return kabus.FutureCodeNK225
	case kabuspb.FutureCode_FUTURE_CODE_NK225_MINI:
		return kabus.FutureCodeNK225Mini
	case kabuspb.FutureCode_FUTURE_CODE_TOPIX:
		return kabus.FutureCodeTOPIX
	case kabuspb.FutureCode_FUTURE_CODE_TOPIX_MINI:
		return kabus.FutureCodeTOPIXMini
	case kabuspb.FutureCode_FUTURE_CODE_MOTHERS:
		return kabus.FutureCodeMOTHERS
	case kabuspb.FutureCode_FUTURE_CODE_JPX400:
		return kabus.FutureCodeJPX400
	case kabuspb.FutureCode_FUTURE_CODE_DOW:
		return kabus.FutureCodeDOW
	case kabuspb.FutureCode_FUTURE_CODE_VI:
		return kabus.FutureCodeVI
	case kabuspb.FutureCode_FUTURE_CODE_CORE30:
		return kabus.FutureCodeCore30
	case kabuspb.FutureCode_FUTURE_CODE_REIT:
		return kabus.FutureCodeREIT
	}
	return kabus.FutureCodeUnspecified
}

func toYmNum(timestamp *timestamppb.Timestamp) kabus.YmNUM {
	t := timestamp.AsTime().In(time.Local)
	if t.IsZero() {
		return kabus.YmNUMToday
	}
	return kabus.NewYmNUM(t)
}

func toPutOrCall(callPut kabuspb.CallPut) kabus.PutOrCall {
	switch callPut {
	case kabuspb.CallPut_CALL_PUT_CALL:
		return kabus.PutOrCallCall
	case kabuspb.CallPut_CALL_PUT_PUT:
		return kabus.PutOrCallPut
	}
	return kabus.PutOrCallUnspecified
}

func fromCurrentPriceChangeStatus(status kabus.CurrentPriceChangeStatus) string {
	switch status {
	case kabus.CurrentPriceChangeStatusNoEffect:
		return "0000"
	case kabus.CurrentPriceChangeStatusNoChange:
		return "0056"
	case kabus.CurrentPriceChangeStatusUp:
		return "0057"
	case kabus.CurrentPriceChangeStatusDown:
		return "0058"
	case kabus.CurrentPriceChangeStatusOpenPriceAfterBreak:
		return "0059"
	case kabus.CurrentPriceChangeStatusTradingSessionClose:
		return "0060"
	case kabus.CurrentPriceChangeStatusClose:
		return "0061"
	case kabus.CurrentPriceChangeStatusBreakClose:
		return "0062"
	case kabus.CurrentPriceChangeStatusDownClose:
		return "0063"
	case kabus.CurrentPriceChangeStatusTarnOverClose:
		return "0064"
	case kabus.CurrentPriceChangeStatusSpecialQuoteClose:
		return "0066"
	case kabus.CurrentPriceChangeStatusReservationClose:
		return "0067"
	case kabus.CurrentPriceChangeStatusStopClose:
		return "0068"
	case kabus.CurrentPriceChangeCircuitBreakerClose:
		return "0069"
	case kabus.CurrentPriceChangeDynamicCircuitBreakerClose:
		return "0431"
	}
	return ""
}

func fromCurrentPriceStatus(status kabus.CurrentPriceStatus) int32 {
	switch status {
	case kabus.CurrentPriceStatusCurrentPrice:
		return 1
	case kabus.CurrentPriceStatusNoContinuousTicks:
		return 2
	case kabus.CurrentPriceStatusItayose:
		return 3
	case kabus.CurrentPriceStatusSystemError:
		return 4
	case kabus.CurrentPriceStatusPause:
		return 5
	case kabus.CurrentPriceStatusStopTrading:
		return 6
	case kabus.CurrentPriceStatusRestart:
		return 7
	case kabus.CurrentPriceStatusClosePrice:
		return 8
	case kabus.CurrentPriceStatusSystemStop:
		return 9
	case kabus.CurrentPriceStatusRoughQuote:
		return 10
	case kabus.CurrentPriceStatusReference:
		return 11
	case kabus.CurrentPriceStatusInCircuitBreak:
		return 12
	case kabus.CurrentPriceStatusRestoration:
		return 13
	case kabus.CurrentPriceStatusReleaseCircuitBreak:
		return 14
	case kabus.CurrentPriceStatusReleasePause:
		return 15
	case kabus.CurrentPriceStatusInReservation:
		return 16
	case kabus.CurrentPriceStatusReleaseReservation:
		return 17
	case kabus.CurrentPriceStatusFileError:
		return 18
	case kabus.CurrentPriceStatusReleaseFileError:
		return 19
	case kabus.CurrentPriceStatusSpreadStrategy:
		return 20
	case kabus.CurrentPriceStatusInDynamicCircuitBreak:
		return 21
	case kabus.CurrentPriceStatusReleaseDynamicCircuitBreak:
		return 22
	case kabus.CurrentPriceStatusContractedInItayose:
		return 23
	}
	return 0
}

func fromBidAskSign(sign kabus.BidAskSign) string {
	switch sign {
	case kabus.BidAskSignNoEffect:
		return "0000"
	case kabus.BidAskSignGeneral:
		return "0101"
	case kabus.BidAskSignSpecial:
		return "0102"
	case kabus.BidAskSignAttention:
		return "0103"
	case kabus.BidAskSignBeforeOpen:
		return "0107"
	case kabus.BidAskSignSpecialBeforeStop:
		return "0108"
	case kabus.BidAskSignAfterClose:
		return "0109"
	case kabus.BidAskSignNotExistsContractPoint:
		return "0116"
	case kabus.BidAskSignExistsContractPoint:
		return "0117"
	case kabus.BidAskSignContinuous:
		return "0118"
	case kabus.BidAskSignContinuousBeforeStop:
		return "0119"
	case kabus.BidAskSignMoving:
		return "0120"
	}
	return ""
}

func fromFirstBoardSign(firstBoardSign kabus.FirstBoardSign) *kabuspb.FirstQuote {
	return &kabuspb.FirstQuote{
		Time:     timestamppb.New(firstBoardSign.Time),
		Sign:     fromBidAskSign(firstBoardSign.Sign),
		Price:    firstBoardSign.Price,
		Quantity: firstBoardSign.Qty,
	}
}

func fromBoardSign(boardSign kabus.BoardSign) *kabuspb.Quote {
	return &kabuspb.Quote{
		Price:    boardSign.Price,
		Quantity: boardSign.Qty,
	}
}

func fromPutOrCallNum(putOrCall kabus.PutOrCallNum) kabuspb.CallPut {
	switch putOrCall {
	case kabus.PutOrCallNumCall:
		return kabuspb.CallPut_CALL_PUT_CALL
	case kabus.PutOrCallNumPut:
		return kabuspb.CallPut_CALL_PUT_PUT
	}
	return kabuspb.CallPut_CALL_PUT_UNSPECIFIED
}

func fromUnderlyer(underlyer kabus.Underlyer) string {
	switch underlyer {
	case kabus.UnderlyerNK225:
		return "NK225"
	case kabus.UnderlyerNK300:
		return "NK300"
	case kabus.UnderlyerMOTHERS:
		return "MOTHERS"
	case kabus.UnderlyerJPX400:
		return "JPX400"
	case kabus.UnderlyerTOPIX:
		return "TOPIX"
	case kabus.UnderlyerNKVI:
		return "NKVI"
	case kabus.UnderlyerDJIA:
		return "DJIA"
	case kabus.UnderlyerTSEREITINDEX:
		return "TSEREITINDEX"
	case kabus.UnderlyerTOPIXCORE30:
		return "TOPIXCORE30"
	}
	return ""
}

func fromPriceRangeGroup(priceRangeGroup kabus.PriceRangeGroup) string {
	switch priceRangeGroup {
	case kabus.PriceRangeGroup10000:
		return "10000"
	case kabus.PriceRangeGroup10003:
		return "10003"
	case kabus.PriceRangeGroup10118:
		return "10118"
	case kabus.PriceRangeGroup10119:
		return "10119"
	case kabus.PriceRangeGroup10318:
		return "10318"
	case kabus.PriceRangeGroup10706:
		return "10706"
	case kabus.PriceRangeGroup10718:
		return "10718"
	case kabus.PriceRangeGroup12122:
		return "12122"
	case kabus.PriceRangeGroup14473:
		return "14473"
	case kabus.PriceRangeGroup14515:
		return "14515"
	case kabus.PriceRangeGroup15411:
		return "15411"
	case kabus.PriceRangeGroup15569:
		return "15569"
	case kabus.PriceRangeGroup17163:
		return "17163"
	}
	return ""
}

func toOrderState(status kabuspb.OrderState) kabus.OrderState {
	switch status {
	case kabuspb.OrderState_ORDER_STATE_WAIT:
		return kabus.OrderStateWait
	case kabuspb.OrderState_ORDER_STATE_PROCESSING:
		return kabus.OrderStateProcessing
	case kabuspb.OrderState_ORDER_STATE_PROCESSED:
		return kabus.OrderStateProcessed
	case kabuspb.OrderState_ORDER_STATE_IN_MODIFY:
		return kabus.OrderStateInCancel
	case kabuspb.OrderState_ORDER_STATE_DONE:
		return kabus.OrderStateDone
	}
	return kabus.OrderStateUnspecified
}

func toSide(side kabuspb.Side) kabus.Side {
	switch side {
	case kabuspb.Side_SIDE_BUY:
		return kabus.SideBuy
	case kabuspb.Side_SIDE_SELL:
		return kabus.SideSell
	}
	return kabus.SideUnspecified
}

func fromOrders(orders *kabus.OrdersResponse) *kabuspb.Orders {
	res := &kabuspb.Orders{Orders: make([]*kabuspb.Order, len(*orders))}
	for i, order := range *orders {
		res.Orders[i] = &kabuspb.Order{
			Id:                 order.ID,
			State:              fromState(order.State),
			OrderState:         fromOrderState(order.OrderState),
			OrderType:          fromOrdType(order.OrdType),
			ReceiveTime:        timestamppb.New(order.RecvTime),
			SymbolCode:         order.Symbol,
			SymbolName:         order.SymbolName,
			Exchange:           fromOrderExchange(order.Exchange),
			ExchangeName:       order.ExchangeName,
			TimeInForce:        fromTimeInForce(order.TimeInForce),
			Price:              order.Price,
			OrderQuantity:      order.OrderQty,
			CumulativeQuantity: order.CumQty,
			Side:               fromSide(order.Side),
			TradeType:          fromCashMargin(order.CashMargin),
			AccountType:        fromAccountType(order.AccountType),
			DeliveryType:       fromDelivType(order.DelivType),
			ExpireDay:          timestamppb.New(order.ExpireDay.Time),
			MarginTradeType:    fromMarginTradeType(order.MarginTradeType),
			Details:            fromOrderDetails(order.Details),
		}
	}
	return res
}

func fromOrderDetails(details []kabus.OrderDetail) []*kabuspb.OrderDetail {
	res := make([]*kabuspb.OrderDetail, len(details))
	for i, detail := range details {
		res[i] = &kabuspb.OrderDetail{
			SequenceNumber: int32(detail.SeqNum),
			Id:             detail.ID,
			RecordType:     fromRecType(detail.RecType),
			ExchangeId:     detail.ExchangeID,
			State:          fromOrderDetailState(detail.State),
			TransactTime:   timestamppb.New(detail.TransactTime),
			OrderType:      fromOrdType(detail.OrdType),
			Price:          detail.Price,
			Quantity:       detail.Qty,
			ExecutionId:    detail.ExecutionID,
			ExecutionDay:   timestamppb.New(detail.ExecutionDay),
			DeliveryDay:    timestamppb.New(detail.DelivDay.Time),
			Commission:     detail.Commission,
			CommissionTax:  detail.CommissionTax,
		}
	}
	return res
}

func fromState(state kabus.State) kabuspb.State {
	switch state {
	case kabus.StateWait:
		return kabuspb.State_STATE_WAIT
	case kabus.StateProcessing:
		return kabuspb.State_STATE_PROCESSING
	case kabus.StateProcessed:
		return kabuspb.State_STATE_PROCESSED
	case kabus.StateInCancel:
		return kabuspb.State_STATE_IN_MODIFY
	case kabus.StateDone:
		return kabuspb.State_STATE_DONE
	}
	return kabuspb.State_STATE_UNSPECIFIED
}

func fromOrderState(orderState kabus.OrderState) kabuspb.OrderState {
	switch orderState {
	case kabus.OrderStateWait:
		return kabuspb.OrderState_ORDER_STATE_WAIT
	case kabus.OrderStateProcessing:
		return kabuspb.OrderState_ORDER_STATE_PROCESSING
	case kabus.OrderStateProcessed:
		return kabuspb.OrderState_ORDER_STATE_PROCESSED
	case kabus.OrderStateInCancel:
		return kabuspb.OrderState_ORDER_STATE_IN_MODIFY
	case kabus.OrderStateDone:
		return kabuspb.OrderState_ORDER_STATE_DONE
	}
	return kabuspb.OrderState_ORDER_STATE_UNSPECIFIED
}

func fromOrdType(ordType kabus.OrdType) kabuspb.OrderType {
	switch ordType {
	case kabus.OrdTypeInTrading:
		return kabuspb.OrderType_ORDER_TYPE_ZARABA
	case kabus.OrdTypeOpen:
		return kabuspb.OrderType_ORDER_TYPE_OPEN
	case kabus.OrdTypeClose:
		return kabuspb.OrderType_ORDER_TYPE_CLOSE
	case kabus.OrdTypeNoContracted:
		return kabuspb.OrderType_ORDER_TYPE_FUNARI
	case kabus.OrdTypeMarketToLimit:
		return kabuspb.OrderType_ORDER_TYPE_MTLO
	case kabus.OrdTypeIOC:
		return kabuspb.OrderType_ORDER_TYPE_IOC
	}
	return kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED
}

func fromOrderExchange(exchange kabus.OrderExchange) kabuspb.OrderExchange {
	switch exchange {
	case kabus.OrderExchangeToushou:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_TOUSHOU
	case kabus.OrderExchangeMeishou:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_MEISHOU
	case kabus.OrderExchangeFukushou:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_FUKUSHOU
	case kabus.OrderExchangeSatsushou:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_SATSUSHOU
	case kabus.OrderExchangeSOR:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_SOR
	case kabus.OrderExchangeAll:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_ALL_SESSION
	case kabus.OrderExchangeDaytime:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_DAY_SESSION
	case kabus.OrderExchangeEvening:
		return kabuspb.OrderExchange_ORDER_EXCHANGE_NIGHT_SESSION
	}
	return kabuspb.OrderExchange_ORDER_EXCHANGE_UNSPECIFIED
}

func fromSide(side kabus.Side) kabuspb.Side {
	switch side {
	case kabus.SideBuy:
		return kabuspb.Side_SIDE_BUY
	case kabus.SideSell:
		return kabuspb.Side_SIDE_SELL
	}
	return kabuspb.Side_SIDE_UNSPECIFIED
}

func fromAccountType(accountType kabus.AccountType) kabuspb.AccountType {
	switch accountType {
	case kabus.AccountTypeGeneral:
		return kabuspb.AccountType_ACCOUNT_TYPE_GENERAL
	case kabus.AccountTypeSpecific:
		return kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC
	case kabus.AccountTypeCorporation:
		return kabuspb.AccountType_ACCOUNT_TYPE_CORPORATION
	}
	return kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED
}

func fromDelivType(delivType kabus.DelivType) kabuspb.DeliveryType {
	switch delivType {
	case kabus.DelivTypeAuto:
		return kabuspb.DeliveryType_DELIVERY_TYPE_AUTO
	case kabus.DelivTypeCash:
		return kabuspb.DeliveryType_DELIVERY_TYPE_CASH
	}
	return kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED
}

func fromRecType(recType kabus.RecType) kabuspb.RecordType {
	switch recType {
	case kabus.RecTypeReceived:
		return kabuspb.RecordType_RECORD_TYPE_RECEIVE
	case kabus.RecTypeCarried:
		return kabuspb.RecordType_RECORD_TYPE_CARRIED
	case kabus.RecTypeExpired:
		return kabuspb.RecordType_RECORD_TYPE_EXPIRED
	case kabus.RecTypeOrdered:
		return kabuspb.RecordType_RECORD_TYPE_ORDERED
	case kabus.RecTypeModified:
		return kabuspb.RecordType_RECORD_TYPE_MODIFIED
	case kabus.RecTypeCanceled:
		return kabuspb.RecordType_RECORD_TYPE_CANCELED
	case kabus.RecTypeRevocation:
		return kabuspb.RecordType_RECORD_TYPE_REVOCATION
	case kabus.RecTypeContracted:
		return kabuspb.RecordType_RECORD_TYPE_CONTRACTED
	}
	return kabuspb.RecordType_RECORD_TYPE_UNSPECIFIED
}

func fromOrderDetailState(orderDetailState kabus.OrderDetailState) kabuspb.OrderDetailState {
	switch orderDetailState {
	case kabus.OrderDetailStateWait:
		return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_WAIT
	case kabus.OrderDetailStateProcessing:
		return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSING
	case kabus.OrderDetailStateProcessed:
		return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED
	case kabus.OrderDetailStateError:
		return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_ERROR
	case kabus.OrderDetailStateDeleted:
		return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_DELETED
	}
	return kabuspb.OrderDetailState_ORDER_DETAIL_STATE_UNSPECIFIED
}

func toCashMargin(tradeType kabuspb.TradeType) kabus.CashMargin {
	switch tradeType {
	case kabuspb.TradeType_TRADE_TYPE_ENTRY:
		return kabus.CashMarginMarginEntry
	case kabuspb.TradeType_TRADE_TYPE_EXIT:
		return kabus.CashMarginMarginExit
	}
	return kabus.CashMarginUnspecified
}

func fromCashMargin(cashMargin kabus.CashMargin) kabuspb.TradeType {
	switch cashMargin {
	case kabus.CashMarginMarginEntry:
		return kabuspb.TradeType_TRADE_TYPE_ENTRY
	case kabus.CashMarginMarginExit:
		return kabuspb.TradeType_TRADE_TYPE_EXIT
	}
	return kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED
}

func fromMarginTradeType(marginTradeType kabus.MarginTradeType) kabuspb.MarginTradeType {
	switch marginTradeType {
	case kabus.MarginTradeTypeSystem:
		return kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM
	case kabus.MarginTradeTypeGeneralLong:
		return kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_LONG
	case kabus.MarginTradeTypeGeneralShort:
		return kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_SHORT
	}
	return kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED
}

func fromTimeInForce(timeInForce kabus.TimeInForce) kabuspb.TimeInForce {
	switch timeInForce {
	case kabus.TimeInForceFAS:
		return kabuspb.TimeInForce_TIME_IN_FORCE_FAS
	case kabus.TimeInForceFAK:
		return kabuspb.TimeInForce_TIME_IN_FORCE_FAK
	case kabus.TimeInForceFOK:
		return kabuspb.TimeInForce_TIME_IN_FORCE_FOK
	}
	return kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED
}

func toProduct(product kabuspb.Product) kabus.Product {
	switch product {
	case kabuspb.Product_PRODUCT_STOCK:
		return kabus.ProductCash
	case kabuspb.Product_PRODUCT_MARGIN:
		return kabus.ProductMargin
	case kabuspb.Product_PRODUCT_FUTURE:
		return kabus.ProductFuture
	case kabuspb.Product_PRODUCT_OPTION:
		return kabus.ProductOption
	}
	return kabus.ProductAll
}

func toIsGetOrderDetail(getDetails bool) kabus.IsGetOrderDetail {
	if !getDetails {
		return kabus.IsGetOrderDetailFalse
	} else {
		return kabus.IsGetOrderDetailTrue
	}
}

func fromPositions(positions *kabus.PositionsResponse) *kabuspb.Positions {
	res := &kabuspb.Positions{Positions: make([]*kabuspb.Position, len(*positions))}
	for i, position := range *positions {
		res.Positions[i] = &kabuspb.Position{
			ExecutionId:     position.ExecutionID,
			AccountType:     fromAccountType(position.AccountType),
			SymbolCode:      position.Symbol,
			SymbolName:      position.SymbolName,
			Exchange:        fromExchange(position.Exchange),
			ExchangeName:    position.ExchangeName,
			SecurityType:    fromSecurityType(position.SecurityType),
			ExecutionDay:    timestamppb.New(position.ExecutionDay.Time),
			Price:           position.Price,
			LeavesQuantity:  position.LeavesQty,
			HoldQuantity:    position.HoldQty,
			Side:            fromSide(position.Side),
			Expenses:        position.Expenses,
			Commission:      position.Commission,
			CommissionTax:   position.CommissionTax,
			ExpireDay:       timestamppb.New(position.ExpireDay.Time),
			MarginTradeType: fromMarginTradeType(position.MarginTradeType),
			CurrentPrice:    position.CurrentPrice,
			Valuation:       position.Valuation,
			ProfitLoss:      position.ProfitLoss,
			ProfitLossRate:  position.ProfitLossRate,
		}
	}
	return res
}

func fromSecurityType(securityType kabus.SecurityType) kabuspb.SecurityType {
	switch securityType {
	case kabus.SecurityTypeStock:
		return kabuspb.SecurityType_SECURITY_TYPE_STOCK
	case kabus.SecurityTypeNK225:
		return kabuspb.SecurityType_SECURITY_TYPE_NK225
	case kabus.SecurityTypeNK225Mini:
		return kabuspb.SecurityType_SECURITY_TYPE_NK225_MINI
	case kabus.SecurityTypeJPX400:
		return kabuspb.SecurityType_SECURITY_TYPE_JPX400
	case kabus.SecurityTypeTOPIX:
		return kabuspb.SecurityType_SECURITY_TYPE_TOPIX
	case kabus.SecurityTypeTOPIXMini:
		return kabuspb.SecurityType_SECURITY_TYPE_TOPIX_MINI
	case kabus.SecurityTypeMothers:
		return kabuspb.SecurityType_SECURITY_TYPE_MOTHERS
	case kabus.SecurityTypeREIT:
		return kabuspb.SecurityType_SECURITY_TYPE_REIT
	case kabus.SecurityTypeDOW:
		return kabuspb.SecurityType_SECURITY_TYPE_DOW
	case kabus.SecurityTypeVI:
		return kabuspb.SecurityType_SECURITY_TYPE_VI
	case kabus.SecurityTypeCORE30:
		return kabuspb.SecurityType_SECURITY_TYPE_CODE30
	}
	return kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED
}

func toRankingTypeFromPriceRankingType(rankingType kabuspb.PriceRankingType) kabus.RankingType {
	switch rankingType {
	case kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE:
		return kabus.RankingTypePriceIncreaseRate
	case kabuspb.PriceRankingType_PRICE_RANKING_TYPE_DECREASE_RATE:
		return kabus.RankingTypePriceDecreaseRate
	case kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VOLUME:
		return kabus.RankingTypeVolume
	case kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VALUE:
		return kabus.RankingTypeValue
	}
	return kabus.RankingTypeUnspecified
}

func fromRankingTypeToPriceRankingType(rankingType kabus.RankingType) kabuspb.PriceRankingType {
	switch rankingType {
	case kabus.RankingTypePriceIncreaseRate:
		return kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE
	case kabus.RankingTypePriceDecreaseRate:
		return kabuspb.PriceRankingType_PRICE_RANKING_TYPE_DECREASE_RATE
	case kabus.RankingTypeVolume:
		return kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VOLUME
	case kabus.RankingTypeValue:
		return kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VALUE
	}
	return kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED
}

func toRankingTypeFromMarginRankingType(rankingType kabuspb.MarginRankingType) kabus.RankingType {
	switch rankingType {
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_INCREASE:
		return kabus.RankingTypeMarginSellBalanceIncrease
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_DECREASE:
		return kabus.RankingTypeMarginSellBalanceDecrease
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_INCREASE:
		return kabus.RankingTypeMarginBuyBalanceIncrease
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_DECREASE:
		return kabus.RankingTypeMarginBuyBalanceDecrease
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION:
		return kabus.RankingTypeMarginHighMagnification
	case kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_LOW_MAGNIFICATION:
		return kabus.RankingTypeMarginLowMagnification
	}
	return kabus.RankingTypeUnspecified
}

func fromRankingTypeToMarginRankingType(rankingType kabus.RankingType) kabuspb.MarginRankingType {
	switch rankingType {
	case kabus.RankingTypeMarginSellBalanceIncrease:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_INCREASE
	case kabus.RankingTypeMarginSellBalanceDecrease:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_DECREASE
	case kabus.RankingTypeMarginBuyBalanceIncrease:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_INCREASE
	case kabus.RankingTypeMarginBuyBalanceDecrease:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_DECREASE
	case kabus.RankingTypeMarginHighMagnification:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION
	case kabus.RankingTypeMarginLowMagnification:
		return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_LOW_MAGNIFICATION
	}
	return kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED
}

func toRankingTypeFromIndustryRankingType(rankingType kabuspb.IndustryRankingType) kabus.RankingType {
	switch rankingType {
	case kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE:
		return kabus.RankingTypePriceIncreaseRateByCategory
	case kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_DECREASE_RATE:
		return kabus.RankingTypePriceDecreaseRateByCategory
	}
	return kabus.RankingTypeUnspecified
}

func fromRankingTypeToIndustryRankingType(rankingType kabus.RankingType) kabuspb.IndustryRankingType {
	switch rankingType {
	case kabus.RankingTypePriceIncreaseRateByCategory:
		return kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE
	case kabus.RankingTypePriceDecreaseRateByCategory:
		return kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_DECREASE_RATE
	}
	return kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED
}

func toExchangeDivision(exchangeDivision kabuspb.ExchangeDivision) kabus.ExchangeDivision {
	switch exchangeDivision {
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL:
		return kabus.ExchangeDivisionALL
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_ALL:
		return kabus.ExchangeDivisionToushou
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_1:
		return kabus.ExchangeDivisionToushou1
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_2:
		return kabus.ExchangeDivisionToushou2
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MOTHERS:
		return kabus.ExchangeDivisionMothers
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_JASDAQ:
		return kabus.ExchangeDivisionJASDAQ
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MEISHOU:
		return kabus.ExchangeDivisionMeishou
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_FUKUSHOU:
		return kabus.ExchangeDivisionFukushou
	case kabuspb.ExchangeDivision_EXCHANGE_DIVISION_SATSUSHOU:
		return kabus.ExchangeDivisionSatsushou
	}
	return kabus.ExchangeDivisionUnspecified
}

func fromExchangeDivision(exchangeDivision kabus.ExchangeDivision) kabuspb.ExchangeDivision {
	switch exchangeDivision {
	case kabus.ExchangeDivisionALL:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL
	case kabus.ExchangeDivisionToushou:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_ALL
	case kabus.ExchangeDivisionToushou1:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_1
	case kabus.ExchangeDivisionToushou2:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_2
	case kabus.ExchangeDivisionMothers:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MOTHERS
	case kabus.ExchangeDivisionJASDAQ:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_JASDAQ
	case kabus.ExchangeDivisionMeishou:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MEISHOU
	case kabus.ExchangeDivisionFukushou:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_FUKUSHOU
	case kabus.ExchangeDivisionSatsushou:
		return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_SATSUSHOU
	}
	return kabuspb.ExchangeDivision_EXCHANGE_DIVISION_UNSPECIFIED
}

func fromTrend(trend kabus.RankingTrend) kabuspb.RankingTrend {
	switch trend {
	case kabus.RankingTrendNoData:
		return kabuspb.RankingTrend_RANKING_TREND_NO_DATA
	case kabus.RankingTrendRiseOver20:
		return kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20
	case kabus.RankingTrendRise:
		return kabuspb.RankingTrend_RANKING_TREND_RISE
	case kabus.RankingTrendUnchanged:
		return kabuspb.RankingTrend_RANKING_TREND_NO_CHANGE
	case kabus.RankingTrendDescent:
		return kabuspb.RankingTrend_RANKING_TREND_DESCENT
	case kabus.RankingTrendDescentOver20:
		return kabuspb.RankingTrend_RANKING_TREND_DESCENT_OVER_20
	}
	return kabuspb.RankingTrend_RANKING_TREND_UNSPECIFIED
}

func fromRankingToPriceRanking(ranking *kabus.RankingResponse) *kabuspb.PriceRanking {
	res := &kabuspb.PriceRanking{
		Type:             fromRankingTypeToPriceRankingType(ranking.Type),
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.PriceRankingInfo{},
	}

	if ranking.PriceRanking != nil {
		res.Ranking = make([]*kabuspb.PriceRankingInfo, len(ranking.PriceRanking))
		for i, rank := range ranking.PriceRanking {
			res.Ranking[i] = fromPriceRanking(rank)
		}
	}

	return res
}

func fromPriceRanking(rank kabus.PriceRanking) *kabuspb.PriceRankingInfo {
	return &kabuspb.PriceRankingInfo{
		No:               int32(rank.No),
		Trend:            fromTrend(rank.Trend),
		AverageRanking:   rank.AverageRanking,
		SymbolCode:       rank.Symbol,
		SymbolName:       rank.SymbolName,
		CurrentPrice:     rank.CurrentPrice,
		ChangeRatio:      rank.ChangeRatio,
		ChangePercentage: rank.ChangePercentage,
		CurrentPriceTime: timestamppb.New(rank.CurrentPriceTime.Time),
		TradingVolume:    rank.TradingVolume,
		Turnover:         rank.Turnover,
		ExchangeName:     rank.ExchangeName,
		IndustryName:     rank.CategoryName,
	}
}

func fromRankingToTickRanking(ranking *kabus.RankingResponse) *kabuspb.TickRanking {
	res := &kabuspb.TickRanking{
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.TickRankingInfo{},
	}

	if ranking.TickRanking != nil {
		res.Ranking = make([]*kabuspb.TickRankingInfo, len(ranking.TickRanking))
		for i, rank := range ranking.TickRanking {
			res.Ranking[i] = fromTickRanking(rank)
		}
	}

	return res
}

func fromTickRanking(rank kabus.TickRanking) *kabuspb.TickRankingInfo {
	return &kabuspb.TickRankingInfo{
		No:               int32(rank.No),
		Trend:            fromTrend(rank.Trend),
		AverageRanking:   rank.AverageRanking,
		SymbolCode:       rank.Symbol,
		SymbolName:       rank.SymbolName,
		CurrentPrice:     rank.CurrentPrice,
		ChangeRatio:      rank.ChangeRatio,
		TickCount:        int32(rank.TickCount),
		UpCount:          int32(rank.UpCount),
		DownCount:        int32(rank.DownCount),
		ChangePercentage: rank.ChangePercentage,
		TradingVolume:    rank.TradingVolume,
		Turnover:         rank.Turnover,
		ExchangeName:     rank.ExchangeName,
		IndustryName:     rank.CategoryName,
	}
}

func fromRankingToVolumeRanking(ranking *kabus.RankingResponse) *kabuspb.VolumeRanking {
	res := &kabuspb.VolumeRanking{
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.VolumeRankingInfo{},
	}

	if ranking.VolumeRapidRanking != nil {
		res.Ranking = make([]*kabuspb.VolumeRankingInfo, len(ranking.VolumeRapidRanking))
		for i, rank := range ranking.VolumeRapidRanking {
			res.Ranking[i] = fromVolumeRapidRanking(rank)
		}
	}

	return res
}

func fromVolumeRapidRanking(rank kabus.VolumeRapidRanking) *kabuspb.VolumeRankingInfo {
	return &kabuspb.VolumeRankingInfo{
		No:                   int32(rank.No),
		Trend:                fromTrend(rank.Trend),
		AverageRanking:       rank.AverageRanking,
		SymbolCode:           rank.Symbol,
		SymbolName:           rank.SymbolName,
		CurrentPrice:         rank.CurrentPrice,
		ChangeRatio:          rank.ChangeRatio,
		RapidTradePercentage: rank.RapidTradePercentage,
		TradingVolume:        rank.TradingVolume,
		CurrentPriceTime:     timestamppb.New(rank.CurrentPriceTime.Time),
		ChangePercentage:     rank.ChangePercentage,
		ExchangeName:         rank.ExchangeName,
		IndustryName:         rank.CategoryName,
	}
}

func fromRankingToValueRanking(ranking *kabus.RankingResponse) *kabuspb.ValueRanking {
	res := &kabuspb.ValueRanking{
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.ValueRankingInfo{},
	}

	if ranking.ValueRapidRanking != nil {
		res.Ranking = make([]*kabuspb.ValueRankingInfo, len(ranking.ValueRapidRanking))
		for i, rank := range ranking.ValueRapidRanking {
			res.Ranking[i] = fromValueRapidRanking(rank)
		}
	}

	return res
}

func fromValueRapidRanking(rank kabus.ValueRapidRanking) *kabuspb.ValueRankingInfo {
	return &kabuspb.ValueRankingInfo{
		No:                     int32(rank.No),
		Trend:                  fromTrend(rank.Trend),
		AverageRanking:         rank.AverageRanking,
		SymbolCode:             rank.Symbol,
		SymbolName:             rank.SymbolName,
		CurrentPrice:           rank.CurrentPrice,
		ChangeRatio:            rank.ChangeRatio,
		RapidPaymentPercentage: rank.RapidPaymentPercentage,
		Turnover:               rank.Turnover,
		CurrentPriceTime:       timestamppb.New(rank.CurrentPriceTime.Time),
		ChangePercentage:       rank.ChangePercentage,
		ExchangeName:           rank.ExchangeName,
		IndustryName:           rank.CategoryName,
	}
}

func fromRankingToMarginRanking(ranking *kabus.RankingResponse) *kabuspb.MarginRanking {
	res := &kabuspb.MarginRanking{
		Type:             fromRankingTypeToMarginRankingType(ranking.Type),
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.MarginRankingInfo{},
	}

	if ranking.MarginRanking != nil {
		res.Ranking = make([]*kabuspb.MarginRankingInfo, len(ranking.MarginRanking))
		for i, rank := range ranking.MarginRanking {
			res.Ranking[i] = fromMarginRanking(rank)
		}
	}

	return res
}

func fromMarginRanking(rank kabus.MarginRanking) *kabuspb.MarginRankingInfo {
	return &kabuspb.MarginRankingInfo{
		No:                         int32(rank.No),
		SymbolCode:                 rank.Symbol,
		SymbolName:                 rank.SymbolName,
		SellRapidPaymentPercentage: rank.SellRapidPaymentPercentage,
		SellLastWeekRatio:          rank.SellLastWeekRatio,
		BuyRapidPaymentPercentage:  rank.BuyRapidPaymentPercentage,
		BuyLastWeekRatio:           rank.BuyLastWeekRatio,
		Ratio:                      rank.Ratio,
		ExchangeName:               rank.ExchangeName,
		IndustryName:               rank.CategoryName,
	}
}

func fromRankingToIndustryRanking(ranking *kabus.RankingResponse) *kabuspb.IndustryRanking {
	res := &kabuspb.IndustryRanking{
		Type:             fromRankingTypeToIndustryRankingType(ranking.Type),
		ExchangeDivision: fromExchangeDivision(ranking.ExchangeDivision),
		Ranking:          []*kabuspb.IndustryRankingInfo{},
	}

	if ranking.CategoryPriceRanking != nil {
		res.Ranking = make([]*kabuspb.IndustryRankingInfo, len(ranking.CategoryPriceRanking))
		for i, rank := range ranking.CategoryPriceRanking {
			res.Ranking[i] = fromCategoryRanking(rank)
		}
	}

	return res
}

func fromCategoryRanking(rank kabus.CategoryPriceRanking) *kabuspb.IndustryRankingInfo {
	return &kabuspb.IndustryRankingInfo{
		No:               int32(rank.No),
		Trend:            fromTrend(rank.Trend),
		AverageRanking:   rank.AverageRanking,
		IndustryCode:     rank.Category,
		IndustryName:     rank.CategoryName,
		CurrentPrice:     rank.CurrentPrice,
		ChangeRatio:      rank.ChangeRatio,
		CurrentPriceTime: timestamppb.New(rank.CurrentPriceTime.Time),
		ChangePercentage: rank.ChangePercentage,
	}
}

func toSendOrderStockRequestFromSendStockOrderRequest(req *kabuspb.SendStockOrderRequest, password string) kabus.SendOrderStockRequest {
	delivType := kabus.DelivTypeUnspecified
	if req.Side == kabuspb.Side_SIDE_BUY {
		delivType = toDelivType(req.DeliveryType)
	}

	fundType := kabus.FundTypeUnspecified
	if req.Side == kabuspb.Side_SIDE_BUY {
		fundType = toFundType(req.FundType)
	}

	return kabus.SendOrderStockRequest{
		Password:        password,
		Symbol:          req.SymbolCode,
		Exchange:        toStockExchange(req.Exchange),
		SecurityType:    kabus.SecurityTypeStock, // 株式 固定値
		Side:            toSide(req.Side),
		CashMargin:      kabus.CashMarginCash,             // 現物 固定値
		MarginTradeType: kabus.MarginTradeTypeUnspecified, // 未指定 固定値
		DelivType:       delivType,
		FundType:        fundType,
		AccountType:     toAccountType(req.AccountType),
		Qty:             int(req.Quantity),
		Price:           req.Price,
		ExpireDay:       toExpireDay(req.ExpireDay),
		FrontOrderType:  toStockFrontOrderType(req.OrderType),
	}
}

func toSendOrderStockRequestFromSendMarginOrderRequest(req *kabuspb.SendMarginOrderRequest, password string) kabus.SendOrderStockRequest {
	delivType := kabus.DelivTypeUnspecified
	if req.TradeType == kabuspb.TradeType_TRADE_TYPE_EXIT {
		delivType = toDelivType(req.DeliveryType)
	}

	return kabus.SendOrderStockRequest{
		Password:        password,
		Symbol:          req.SymbolCode,
		Exchange:        toStockExchange(req.Exchange),
		SecurityType:    kabus.SecurityTypeStock, // 株式 固定値
		Side:            toSide(req.Side),
		CashMargin:      toCashMargin(req.TradeType),
		MarginTradeType: toMarginTradeType(req.MarginTradeType),
		DelivType:       delivType,
		AccountType:     toAccountType(req.AccountType),
		Qty:             int(req.Quantity),
		ClosePositions:  toClosePositions(req.ClosePositions),
		Price:           req.Price,
		ExpireDay:       toExpireDay(req.ExpireDay),
		FrontOrderType:  toStockFrontOrderType(req.OrderType),
	}
}

func toSendOrderFutureRequest(req *kabuspb.SendFutureOrderRequest, password string) kabus.SendOrderFutureRequest {
	return kabus.SendOrderFutureRequest{
		Password:           password,
		Symbol:             req.SymbolCode,
		Exchange:           toFutureExchange(req.Exchange),
		TradeType:          toTradeType(req.TradeType),
		TimeInForce:        toTimeInForce(req.TimeInForce),
		Side:               toSide(req.Side),
		Qty:                int(req.Quantity),
		ClosePositionOrder: 0,
		ClosePositions:     toClosePositions(req.ClosePositions),
		FrontOrderType:     toFutureFrontOrderType(req.OrderType),
		Price:              req.Price,
		ExpireDay:          toExpireDay(req.ExpireDay),
	}
}

func toSendOrderOptionRequest(req *kabuspb.SendOptionOrderRequest, password string) kabus.SendOrderOptionRequest {
	return kabus.SendOrderOptionRequest{
		Password:           password,
		Symbol:             req.SymbolCode,
		Exchange:           toOptionExchange(req.Exchange),
		TradeType:          toTradeType(req.TradeType),
		TimeInForce:        toTimeInForce(req.TimeInForce),
		Side:               toSide(req.Side),
		Qty:                int(req.Quantity),
		ClosePositionOrder: 0,
		ClosePositions:     toClosePositions(req.ClosePositions),
		FrontOrderType:     toOptionFrontOrderType(req.OrderType),
		Price:              req.Price,
		ExpireDay:          toExpireDay(req.ExpireDay),
	}
}

func toStockExchange(exchange kabuspb.StockExchange) kabus.StockExchange {
	switch exchange {
	case kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU:
		return kabus.StockExchangeToushou
	case kabuspb.StockExchange_STOCK_EXCHANGE_MEISHOU:
		return kabus.StockExchangeMeishou
	case kabuspb.StockExchange_STOCK_EXCHANGE_FUKUSHOU:
		return kabus.StockExchangeFukushou
	case kabuspb.StockExchange_STOCK_EXCHANGE_SATSUSHOU:
		return kabus.StockExchangeSatsushou
	}
	return kabus.StockExchangeUnspecified
}

func toMarginTradeType(tradeType kabuspb.MarginTradeType) kabus.MarginTradeType {
	switch tradeType {
	case kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM:
		return kabus.MarginTradeTypeSystem
	case kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_LONG:
		return kabus.MarginTradeTypeGeneralLong
	case kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_SHORT:
		return kabus.MarginTradeTypeGeneralShort
	}
	return kabus.MarginTradeTypeUnspecified
}

func toDelivType(deliveryType kabuspb.DeliveryType) kabus.DelivType {
	switch deliveryType {
	case kabuspb.DeliveryType_DELIVERY_TYPE_AUTO:
		return kabus.DelivTypeAuto
	case kabuspb.DeliveryType_DELIVERY_TYPE_CASH:
		return kabus.DelivTypeCash
	}
	return kabus.DelivTypeUnspecified
}

func toFundType(fundType kabuspb.FundType) kabus.FundType {
	switch fundType {
	case kabuspb.FundType_FUND_TYPE_PROTECTED:
		return kabus.FundTypeProtected
	case kabuspb.FundType_FUND_TYPE_SUBSTITUTE_MARGIN:
		return kabus.FundTypeTransferMargin
	case kabuspb.FundType_FUND_TYPE_MARGIN_TRADING:
		return kabus.FundTypeMarginTrading
	}
	return kabus.FundTypeUnspecified
}

func toAccountType(accountType kabuspb.AccountType) kabus.AccountType {
	switch accountType {
	case kabuspb.AccountType_ACCOUNT_TYPE_GENERAL:
		return kabus.AccountTypeGeneral
	case kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC:
		return kabus.AccountTypeSpecific
	case kabuspb.AccountType_ACCOUNT_TYPE_CORPORATION:
		return kabus.AccountTypeCorporation
	}
	return kabus.AccountTypeUnspecified
}

func toStockFrontOrderType(orderType kabuspb.StockOrderType) kabus.StockFrontOrderType {
	switch orderType {
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO:
		return kabus.StockFrontOrderTypeMarket
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOOM:
		return kabus.StockFrontOrderTypeMOOM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOOA:
		return kabus.StockFrontOrderTypeMOOA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOCM:
		return kabus.StockFrontOrderTypeMOCM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOCA:
		return kabus.StockFrontOrderTypeMOCA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_MO:
		return kabus.StockFrontOrderTypeIOCMarket
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO:
		return kabus.StockFrontOrderTypeLimit
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOOM:
		return kabus.StockFrontOrderTypeLOOM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOOA:
		return kabus.StockFrontOrderTypeLOOA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOCM:
		return kabus.StockFrontOrderTypeLOCM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOCA:
		return kabus.StockFrontOrderTypeLOCA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_M:
		return kabus.StockFrontOrderTypeFunariM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_A:
		return kabus.StockFrontOrderTypeFunariA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_LO:
		return kabus.StockFrontOrderTypeIOCLimit
	}
	return kabus.StockFrontOrderTypeUnspecified
}

func toExpireDay(expireDay *timestamppb.Timestamp) kabus.YmdNUM {
	if expireDay == nil || expireDay.AsTime().IsZero() {
		return kabus.YmdNUMToday
	}
	return kabus.NewYmdNUM(expireDay.AsTime().In(time.Local))
}

func toClosePositions(closePositions []*kabuspb.ClosePosition) []kabus.ClosePosition {
	if closePositions == nil {
		return nil
	}

	res := make([]kabus.ClosePosition, len(closePositions))

	for i, cp := range closePositions {
		res[i] = kabus.ClosePosition{HoldID: cp.ExecutionId, Qty: int(cp.Quantity)}
	}

	return res
}

func toFutureExchange(exchange kabuspb.FutureExchange) kabus.FutureExchange {
	switch exchange {
	case kabuspb.FutureExchange_FUTURE_EXCHANGE_ALL_SESSION:
		return kabus.FutureExchangeAll
	case kabuspb.FutureExchange_FUTURE_EXCHANGE_DAY_SESSION:
		return kabus.FutureExchangeDaytime
	case kabuspb.FutureExchange_FUTURE_EXCHANGE_NIGHT_SESSION:
		return kabus.FutureExchangeEvening
	}
	return kabus.FutureExchangeUnspecified
}

func toTradeType(tradeType kabuspb.TradeType) kabus.TradeType {
	switch tradeType {
	case kabuspb.TradeType_TRADE_TYPE_ENTRY:
		return kabus.TradeTypeEntry
	case kabuspb.TradeType_TRADE_TYPE_EXIT:
		return kabus.TradeTypeExit
	}
	return kabus.TradeTypeUnspecified
}

func toTimeInForce(timeInForce kabuspb.TimeInForce) kabus.TimeInForce {
	switch timeInForce {
	case kabuspb.TimeInForce_TIME_IN_FORCE_FAS:
		return kabus.TimeInForceFAS
	case kabuspb.TimeInForce_TIME_IN_FORCE_FAK:
		return kabus.TimeInForceFAK
	case kabuspb.TimeInForce_TIME_IN_FORCE_FOK:
		return kabus.TimeInForceFOK
	}
	return kabus.TimeInForceUnspecified
}

func toFutureFrontOrderType(orderType kabuspb.FutureOrderType) kabus.FutureFrontOrderType {
	switch orderType {
	case kabuspb.FutureOrderType_FUTURE_ORDER_TYPE_MO:
		return kabus.FutureFrontOrderTypeMarket
	case kabuspb.FutureOrderType_FUTURE_ORDER_TYPE_MOC:
		return kabus.FutureFrontOrderTypeMarketClose
	case kabuspb.FutureOrderType_FUTURE_ORDER_TYPE_LO:
		return kabus.FutureFrontOrderTypeLimit
	case kabuspb.FutureOrderType_FUTURE_ORDER_TYPE_LOC:
		return kabus.FutureFrontOrderTypeLimitClose
	}
	return kabus.FutureFrontOrderTypeUnspecified
}

func toOptionExchange(exchange kabuspb.OptionExchange) kabus.OptionExchange {
	switch exchange {
	case kabuspb.OptionExchange_OPTION_EXCHANGE_ALL_SESSION:
		return kabus.OptionExchangeAll
	case kabuspb.OptionExchange_OPTION_EXCHANGE_DAY_SESSION:
		return kabus.OptionExchangeDaytime
	case kabuspb.OptionExchange_OPTION_EXCHANGE_NIGHT_SESSION:
		return kabus.OptionExchangeEvening
	}
	return kabus.OptionExchangeUnspecified
}

func toOptionFrontOrderType(orderType kabuspb.OptionOrderType) kabus.OptionFrontOrderType {
	switch orderType {
	case kabuspb.OptionOrderType_OPTION_ORDER_TYPE_MO:
		return kabus.OptionFrontOrderTypeMarket
	case kabuspb.OptionOrderType_OPTION_ORDER_TYPE_MOC:
		return kabus.OptionFrontOrderTypeMarketClose
	case kabuspb.OptionOrderType_OPTION_ORDER_TYPE_LO:
		return kabus.OptionFrontOrderTypeLimit
	case kabuspb.OptionOrderType_OPTION_ORDER_TYPE_LOC:
		return kabus.OptionFrontOrderTypeLimitClose
	}
	return kabus.OptionFrontOrderTypeUnspecified
}

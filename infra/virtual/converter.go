package virtual

import (
	"strconv"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toStockOrderRequest(req *kabuspb.SendStockOrderRequest) *vs.StockOrderRequest {
	if req == nil {
		return nil
	}

	return &vs.StockOrderRequest{
		Side:               toSide(req.Side),
		ExecutionCondition: toStockExecutionCondition(req.OrderType),
		SymbolCode:         req.SymbolCode,
		Quantity:           req.Quantity,
		LimitPrice:         req.Price,
		ExpiredAt:          toExpiredAt(req.ExpireDay),
		StopCondition:      toStopCondition(req.StopOrder),
	}
}

func fromOrderResult(res *vs.OrderResult) *kabuspb.OrderResponse {
	if res == nil {
		return nil
	}

	return &kabuspb.OrderResponse{
		ResultCode: 0,
		OrderId:    res.OrderCode,
	}
}

func fromStockOrders(res []*vs.StockOrder) *kabuspb.Orders {
	orders := make([]*kabuspb.Order, len(res))
	if res == nil {
		return &kabuspb.Orders{Orders: orders}
	}

	for i, o := range res {
		orders[i] = &kabuspb.Order{
			Id:                 o.Code,
			State:              fromOrderStatusToState(o.OrderStatus),
			OrderState:         fromOrderStatusToOrderState(o.OrderStatus),
			OrderType:          fromStockExecutionCondition(o.ExecutionCondition),
			ReceiveTime:        timestamppb.New(o.OrderedAt),
			SymbolCode:         o.SymbolCode,
			SymbolName:         "",                                               // TODO 必要なら
			Exchange:           kabuspb.OrderExchange_ORDER_EXCHANGE_UNSPECIFIED, // TODO 必要なら
			ExchangeName:       "",                                               // TODO 必要なら
			TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
			Price:              o.LimitPrice,
			OrderQuantity:      o.OrderQuantity,
			CumulativeQuantity: o.ContractedQuantity,
			Side:               fromSide(o.Side),
			TradeType:          tradeTypeFromSide(o.Side),
			AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
			DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
			ExpireDay:          timestamppb.New(o.ExpiredAt),
			MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
			Details:            fromContracts(o.Contracts),
		}
	}

	return &kabuspb.Orders{
		Orders: orders,
	}
}

func fromStockPositions(res []*vs.StockPosition) *kabuspb.Positions {
	positions := make([]*kabuspb.Position, len(res))
	if res == nil {
		return &kabuspb.Positions{Positions: positions}
	}

	for i, p := range res {
		positions[i] = &kabuspb.Position{
			ExecutionId:     p.Code,
			AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
			SymbolCode:      p.SymbolCode,
			SymbolName:      "",                                    // TODO 必要なら
			Exchange:        kabuspb.Exchange_EXCHANGE_UNSPECIFIED, // TODO 必要なら
			ExchangeName:    "",                                    // TODO 必要なら
			SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED,
			ExecutionDay:    timestamppb.New(p.ContractedAt),
			Price:           p.Price,
			LeavesQuantity:  p.OwnedQuantity,
			HoldQuantity:    p.HoldQuantity,
			Side:            fromSide(p.Side),
			Expenses:        0,
			Commission:      0,
			CommissionTax:   0,
			ExpireDay:       nil,
			MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
			CurrentPrice:    0, // TODO 必要なら
			Valuation:       0, // TODO 必要なら
			ProfitLoss:      0, // TODO 必要なら
			ProfitLossRate:  0, // TODO 必要なら
		}
	}

	return &kabuspb.Positions{Positions: positions}
}

func toSide(side kabuspb.Side) vs.Side {
	switch side {
	case kabuspb.Side_SIDE_BUY:
		return vs.SideBuy
	case kabuspb.Side_SIDE_SELL:
		return vs.SideSell
	}
	return vs.SideUnspecified
}

func toStockExecutionCondition(orderType kabuspb.StockOrderType) vs.StockExecutionCondition {
	switch orderType {
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO:
		return vs.StockExecutionConditionMO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOMO:
		return vs.StockExecutionConditionMOMO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOAO:
		return vs.StockExecutionConditionMOAO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOMC:
		return vs.StockExecutionConditionMOMC
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOAC:
		return vs.StockExecutionConditionMOAC
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_MO:
		return vs.StockExecutionConditionIOCMO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO:
		return vs.StockExecutionConditionLO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOMO:
		return vs.StockExecutionConditionLOMO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOAO:
		return vs.StockExecutionConditionLOAO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOMC:
		return vs.StockExecutionConditionLOMC
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOAC:
		return vs.StockExecutionConditionLOAC
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_M:
		return vs.StockExecutionConditionFunariM
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_A:
		return vs.StockExecutionConditionFunariA
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_LO:
		return vs.StockExecutionConditionIOCLO
	case kabuspb.StockOrderType_STOCK_ORDER_TYPE_STOP:
		return vs.StockExecutionConditionStop
	}
	return vs.StockExecutionConditionUnspecified
}

func toStopCondition(order *kabuspb.StockStopOrder) *vs.StockStopCondition {
	if order == nil {
		return nil
	}

	return &vs.StockStopCondition{
		StopPrice:                  order.TriggerPrice,
		ComparisonOperator:         toComparisonOperator(order.UnderOver),
		ExecutionConditionAfterHit: toExecutionConditionAfterHit(order.AfterHitOrderType),
		LimitPriceAfterHit:         order.AfterHitPrice,
	}
}

func toComparisonOperator(uo kabuspb.UnderOver) vs.ComparisonOperator {
	switch uo {
	case kabuspb.UnderOver_UNDER_OVER_UNDER:
		return vs.ComparisonOperatorLE
	case kabuspb.UnderOver_UNDER_OVER_OVER:
		return vs.ComparisonOperatorGE
	}
	return vs.ComparisonOperatorUnspecified
}

func toExecutionConditionAfterHit(orderType kabuspb.StockAfterHitOrderType) vs.StockExecutionCondition {
	switch orderType {
	case kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_MO:
		return vs.StockExecutionConditionMO
	case kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_LO:
		return vs.StockExecutionConditionLO
	case kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_FUNARI:
		return vs.StockExecutionConditionFunariA
	}
	return vs.StockExecutionConditionUnspecified
}

func fromOrderStatusToState(status vs.OrderStatus) kabuspb.State {
	switch status {
	case vs.OrderStatusWait:
		return kabuspb.State_STATE_WAIT
	case vs.OrderStatusNew:
		return kabuspb.State_STATE_PROCESSING
	case vs.OrderStatusInOrder, vs.OrderStatusPart, vs.OrderStatusInCancel:
		return kabuspb.State_STATE_PROCESSED
	case vs.OrderStatusDone, vs.OrderStatusCanceled:
		return kabuspb.State_STATE_DONE
	}
	return kabuspb.State_STATE_UNSPECIFIED
}

func fromOrderStatusToOrderState(status vs.OrderStatus) kabuspb.OrderState {
	switch status {
	case vs.OrderStatusWait:
		return kabuspb.OrderState_ORDER_STATE_WAIT
	case vs.OrderStatusNew:
		return kabuspb.OrderState_ORDER_STATE_PROCESSING
	case vs.OrderStatusInOrder, vs.OrderStatusPart, vs.OrderStatusInCancel:
		return kabuspb.OrderState_ORDER_STATE_PROCESSED
	case vs.OrderStatusDone, vs.OrderStatusCanceled:
		return kabuspb.OrderState_ORDER_STATE_DONE
	}
	return kabuspb.OrderState_ORDER_STATE_UNSPECIFIED
}

func fromStockExecutionCondition(condition vs.StockExecutionCondition) kabuspb.OrderType {
	switch condition {
	case vs.StockExecutionConditionMO, vs.StockExecutionConditionLO, vs.StockExecutionConditionStop:
		return kabuspb.OrderType_ORDER_TYPE_ZARABA
	case vs.StockExecutionConditionMOMO, vs.StockExecutionConditionMOAO, vs.StockExecutionConditionLOMO, vs.StockExecutionConditionLOAO:
		return kabuspb.OrderType_ORDER_TYPE_OPEN
	case vs.StockExecutionConditionMOMC, vs.StockExecutionConditionMOAC, vs.StockExecutionConditionLOMC, vs.StockExecutionConditionLOAC:
		return kabuspb.OrderType_ORDER_TYPE_CLOSE
	case vs.StockExecutionConditionIOCMO, vs.StockExecutionConditionIOCLO:
		return kabuspb.OrderType_ORDER_TYPE_IOC
	case vs.StockExecutionConditionFunariM, vs.StockExecutionConditionFunariA:
		return kabuspb.OrderType_ORDER_TYPE_FUNARI
	}
	return kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED
}

func fromSide(side vs.Side) kabuspb.Side {
	switch side {
	case vs.SideBuy:
		return kabuspb.Side_SIDE_BUY
	case vs.SideSell:
		return kabuspb.Side_SIDE_SELL
	}
	return kabuspb.Side_SIDE_UNSPECIFIED
}

func tradeTypeFromSide(side vs.Side) kabuspb.TradeType {
	switch side {
	case vs.SideBuy:
		return kabuspb.TradeType_TRADE_TYPE_ENTRY
	case vs.SideSell:
		return kabuspb.TradeType_TRADE_TYPE_EXIT
	}
	return kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED
}

func fromContracts(contracts []*vs.Contract) []*kabuspb.OrderDetail {
	details := make([]*kabuspb.OrderDetail, 0)
	if contracts == nil {
		return details
	}

	for _, c := range contracts {
		if c == nil {
			continue
		}
		details = append(details, &kabuspb.OrderDetail{
			SequenceNumber: int32(len(details) + 1),
			Id:             strconv.FormatInt(int64(len(details)+1), 10),
			RecordType:     kabuspb.RecordType_RECORD_TYPE_CONTRACTED,
			ExchangeId:     "virtual-security",
			State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
			TransactTime:   timestamppb.New(c.ContractedAt),
			OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA, // TODO 必要なら仮想証券会社で約定ルールを持つようにする
			Price:          c.Price,
			Quantity:       c.Quantity,
			ExecutionId:    c.PositionCode,
			ExecutionDay:   timestamppb.New(c.ContractedAt),
			DeliveryDay:    nil,
			Commission:     0,
			CommissionTax:  0,
		})
	}
	return details
}

func toExpiredAt(timestamp *timestamppb.Timestamp) time.Time {
	if timestamp == nil {
		return time.Time{}
	}

	t := timestamp.AsTime().In(time.Local)
	if t.IsZero() {
		return time.Time{}
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func toRegisterPrice(req *kabuspb.Board) vs.RegisterPriceRequest {
	if req == nil {
		return vs.RegisterPriceRequest{}
	}

	return vs.RegisterPriceRequest{
		ExchangeType: toExchangeType(req.Exchange),
		SymbolCode:   req.SymbolCode,
		Price:        req.CurrentPrice,
		PriceTime:    req.CurrentPriceTime.AsTime().In(time.Local),
		Ask:          req.AskPrice,
		AskTime:      req.AskTime.AsTime().In(time.Local),
		Bid:          req.BidPrice,
		BidTime:      req.BidTime.AsTime().In(time.Local),
	}
}

func toExchangeType(exchange kabuspb.Exchange) vs.ExchangeType {
	switch exchange {
	case kabuspb.Exchange_EXCHANGE_TOUSHOU, kabuspb.Exchange_EXCHANGE_MEISHOU, kabuspb.Exchange_EXCHANGE_FUKUSHOU, kabuspb.Exchange_EXCHANGE_SATSUSHOU:
		return vs.ExchangeTypeStock
	case kabuspb.Exchange_EXCHANGE_ALL_SESSION, kabuspb.Exchange_EXCHANGE_DAY_SESSION, kabuspb.Exchange_EXCHANGE_NIGHT_SESSION:
		return vs.ExchangeTypeFuture
	}
	return vs.ExchangeTypeUnspecified
}

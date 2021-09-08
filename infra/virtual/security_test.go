package virtual

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
)

type testVirtualSecurity struct {
	vs.VirtualSecurity
	stockOrder1            *vs.OrderResult
	stockOrder2            error
	cancelStockOrder       error
	cancelStockOrderCount  int
	stockOrders1           []*vs.StockOrder
	stockOrders2           error
	stockOrdersCount       int
	stockPositions1        []*vs.StockPosition
	stockPositions2        error
	stockPositionsCount    int
	registerPrice1         error
	marginOrder1           *vs.OrderResult
	marginOrder2           error
	cancelMarginOrder      error
	cancelMarginOrderCount int
	marginOrders1          []*vs.MarginOrder
	marginOrders2          error
	marginOrdersCount      int
	marginPositions1       []*vs.MarginPosition
	marginPositions2       error
	marginPositionsCount   int
}

func (t *testVirtualSecurity) StockOrder(*vs.StockOrderRequest) (*vs.OrderResult, error) {
	return t.stockOrder1, t.stockOrder2
}
func (t *testVirtualSecurity) CancelStockOrder(*vs.CancelOrderRequest) error {
	t.cancelStockOrderCount++
	return t.cancelStockOrder
}
func (t *testVirtualSecurity) StockOrders() ([]*vs.StockOrder, error) {
	t.stockOrdersCount++
	return t.stockOrders1, t.stockOrders2
}
func (t *testVirtualSecurity) StockPositions() ([]*vs.StockPosition, error) {
	t.stockPositionsCount++
	return t.stockPositions1, t.stockPositions2
}
func (t *testVirtualSecurity) RegisterPrice(vs.RegisterPriceRequest) error {
	return t.registerPrice1
}
func (t *testVirtualSecurity) MarginOrder(*vs.MarginOrderRequest) (*vs.OrderResult, error) {
	return t.marginOrder1, t.marginOrder2
}
func (t *testVirtualSecurity) CancelMarginOrder(*vs.CancelOrderRequest) error {
	t.cancelMarginOrderCount++
	return t.cancelMarginOrder
}
func (t *testVirtualSecurity) MarginOrders() ([]*vs.MarginOrder, error) {
	t.marginOrdersCount++
	return t.marginOrders1, t.marginOrders2
}
func (t *testVirtualSecurity) MarginPositions() ([]*vs.MarginPosition, error) {
	t.marginPositionsCount++
	return t.marginPositions1, t.marginPositions2
}

func Test_NewSecurity(t *testing.T) {
	t.Parallel()
	sec := &testVirtualSecurity{}
	want := &security{virtual: sec}
	got := NewSecurity(sec)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_security_SendOrderStock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		security vs.VirtualSecurity
		want1    *kabuspb.OrderResponse
		want2    error
	}{
		{name: "errが返されたらerrを返す",
			security: &testVirtualSecurity{stockOrder2: vs.NilArgumentError},
			want1:    nil,
			want2:    vs.NilArgumentError,
		},
		{name: "errがなければ結果を返す",
			security: &testVirtualSecurity{stockOrder1: &vs.OrderResult{OrderCode: "sor-1"}},
			want1:    &kabuspb.OrderResponse{ResultCode: 0, OrderId: "sor-1"},
			want2:    nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.SendOrderStock(context.Background(), "no-token", nil)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_security_Orders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		security              *testVirtualSecurity
		arg3                  *kabuspb.GetOrdersRequest
		want1                 *kabuspb.Orders
		want2                 error
		wantStockOrdersCount  int
		wantMarginOrdersCount int
	}{
		{name: "存在しないproductが指定されたら空配列が返される",
			security: &testVirtualSecurity{stockOrders2: vs.NilArgumentError},
			arg3:     &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_UNSPECIFIED},
			want1:    &kabuspb.Orders{Orders: []*kabuspb.Order{}},
			want2:    nil,
		},
		{name: "productがallでerrがなければ全ての結果をつないで返す",
			security: &testVirtualSecurity{
				stockOrders1: []*vs.StockOrder{
					{
						Code:               "sor-1234",
						OrderStatus:        vs.OrderStatusDone,
						Side:               vs.SideBuy,
						ExecutionCondition: vs.StockExecutionConditionMO,
						SymbolCode:         "1234",
						OrderQuantity:      100,
						ContractedQuantity: 100,
						CanceledQuantity:   0,
						LimitPrice:         0,
						ExpiredAt:          time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local),
						StopCondition:      nil,
						OrderedAt:          time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local),
						CanceledAt:         time.Time{},
						Contracts:          []*vs.Contract{},
						Message:            "",
					},
				},
				marginOrders1: []*vs.MarginOrder{
					{
						Code:               "mor-1234",
						OrderStatus:        vs.OrderStatusDone,
						TradeType:          vs.TradeTypeEntry,
						Side:               vs.SideBuy,
						ExecutionCondition: vs.StockExecutionConditionMO,
						SymbolCode:         "1234",
						OrderQuantity:      100,
						ContractedQuantity: 100,
						CanceledQuantity:   0,
						LimitPrice:         0,
						ExpiredAt:          time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local),
						StopCondition:      nil,
						ExitPositionList:   nil,
						OrderedAt:          time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local),
						CanceledAt:         time.Time{},
						Contracts:          []*vs.Contract{},
						Message:            "",
					},
				}},
			arg3: &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_ALL},
			want1: &kabuspb.Orders{Orders: []*kabuspb.Order{
				{
					Id:                 "sor-1234",
					State:              kabuspb.State_STATE_DONE,
					OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
					OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
					ReceiveTime:        timestamppb.New(time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local)),
					SymbolCode:         "1234",
					SymbolName:         "",
					Exchange:           0,
					ExchangeName:       "",
					TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
					Price:              0,
					OrderQuantity:      100,
					CumulativeQuantity: 100,
					Side:               kabuspb.Side_SIDE_BUY,
					TradeType:          kabuspb.TradeType_TRADE_TYPE_ENTRY,
					AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
					ExpireDay:          timestamppb.New(time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local)),
					MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
					Details:            []*kabuspb.OrderDetail{},
				},
				{
					Id:                 "mor-1234",
					State:              kabuspb.State_STATE_DONE,
					OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
					OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
					ReceiveTime:        timestamppb.New(time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local)),
					SymbolCode:         "1234",
					SymbolName:         "",
					Exchange:           0,
					ExchangeName:       "",
					TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
					Price:              0,
					OrderQuantity:      100,
					CumulativeQuantity: 100,
					Side:               kabuspb.Side_SIDE_BUY,
					TradeType:          kabuspb.TradeType_TRADE_TYPE_ENTRY,
					AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
					ExpireDay:          timestamppb.New(time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local)),
					MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
					Details:            []*kabuspb.OrderDetail{},
				},
			}},
			want2:                 nil,
			wantStockOrdersCount:  1,
			wantMarginOrdersCount: 1,
		},
		{name: "productがstockでerrが返されたらerrを返す",
			security:             &testVirtualSecurity{stockOrders2: vs.NilArgumentError},
			arg3:                 &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_STOCK},
			want1:                nil,
			want2:                vs.NilArgumentError,
			wantStockOrdersCount: 1,
		},
		{name: "productがstockでerrがなければ結果を返す",
			security: &testVirtualSecurity{stockOrders1: []*vs.StockOrder{
				{
					Code:               "sor-1234",
					OrderStatus:        vs.OrderStatusDone,
					Side:               vs.SideBuy,
					ExecutionCondition: vs.StockExecutionConditionMO,
					SymbolCode:         "1234",
					OrderQuantity:      100,
					ContractedQuantity: 100,
					CanceledQuantity:   0,
					LimitPrice:         0,
					ExpiredAt:          time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local),
					StopCondition:      nil,
					OrderedAt:          time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local),
					CanceledAt:         time.Time{},
					Contracts:          []*vs.Contract{},
					Message:            "",
				},
			}},
			arg3: &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_STOCK},
			want1: &kabuspb.Orders{Orders: []*kabuspb.Order{
				{
					Id:                 "sor-1234",
					State:              kabuspb.State_STATE_DONE,
					OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
					OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
					ReceiveTime:        timestamppb.New(time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local)),
					SymbolCode:         "1234",
					SymbolName:         "",
					Exchange:           0,
					ExchangeName:       "",
					TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
					Price:              0,
					OrderQuantity:      100,
					CumulativeQuantity: 100,
					Side:               kabuspb.Side_SIDE_BUY,
					TradeType:          kabuspb.TradeType_TRADE_TYPE_ENTRY,
					AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
					ExpireDay:          timestamppb.New(time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local)),
					MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
					Details:            []*kabuspb.OrderDetail{},
				},
			}},
			want2:                nil,
			wantStockOrdersCount: 1,
		},
		{name: "productがmarginでerrが返されたらerrを返す",
			security:              &testVirtualSecurity{marginOrders2: vs.NilArgumentError},
			arg3:                  &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_MARGIN},
			want1:                 nil,
			want2:                 vs.NilArgumentError,
			wantMarginOrdersCount: 1,
		},
		{name: "productがmarginでerrがなければ結果を返す",
			security: &testVirtualSecurity{marginOrders1: []*vs.MarginOrder{
				{
					Code:               "mor-1234",
					OrderStatus:        vs.OrderStatusDone,
					TradeType:          vs.TradeTypeEntry,
					Side:               vs.SideBuy,
					ExecutionCondition: vs.StockExecutionConditionMO,
					SymbolCode:         "1234",
					OrderQuantity:      100,
					ContractedQuantity: 100,
					CanceledQuantity:   0,
					LimitPrice:         0,
					ExpiredAt:          time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local),
					StopCondition:      nil,
					ExitPositionList:   nil,
					OrderedAt:          time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local),
					CanceledAt:         time.Time{},
					Contracts:          []*vs.Contract{},
					Message:            "",
				},
			}},
			arg3: &kabuspb.GetOrdersRequest{Product: kabuspb.Product_PRODUCT_MARGIN},
			want1: &kabuspb.Orders{Orders: []*kabuspb.Order{
				{
					Id:                 "mor-1234",
					State:              kabuspb.State_STATE_DONE,
					OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
					OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
					ReceiveTime:        timestamppb.New(time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local)),
					SymbolCode:         "1234",
					SymbolName:         "",
					Exchange:           0,
					ExchangeName:       "",
					TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
					Price:              0,
					OrderQuantity:      100,
					CumulativeQuantity: 100,
					Side:               kabuspb.Side_SIDE_BUY,
					TradeType:          kabuspb.TradeType_TRADE_TYPE_ENTRY,
					AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
					ExpireDay:          timestamppb.New(time.Date(2021, 7, 16, 0, 0, 0, 0, time.Local)),
					MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
					Details:            []*kabuspb.OrderDetail{},
				},
			}},
			want2:                 nil,
			wantMarginOrdersCount: 1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.Orders(context.Background(), "no-token", test.arg3)
			if !reflect.DeepEqual(test.want1, got1) ||
				!errors.Is(got2, test.want2) ||
				!reflect.DeepEqual(test.wantStockOrdersCount, test.security.stockOrdersCount) ||
				!reflect.DeepEqual(test.wantMarginOrdersCount, test.security.marginOrdersCount) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					test.want1, test.want2, test.wantStockOrdersCount, test.wantMarginOrdersCount,
					got1, got2, test.security.stockOrdersCount, test.security.marginOrdersCount)
			}
		})
	}
}

func Test_security_Positions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                     string
		security                 *testVirtualSecurity
		arg3                     *kabuspb.GetPositionsRequest
		want1                    *kabuspb.Positions
		want2                    error
		wantStockPositionsCount  int
		wantMarginPositionsCount int
	}{
		{name: "productが未指定なら空配列を返す",
			security: &testVirtualSecurity{stockPositions2: vs.NilArgumentError},
			arg3:     &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_UNSPECIFIED},
			want1:    &kabuspb.Positions{Positions: []*kabuspb.Position{}},
			want2:    nil,
		},
		{name: "productがallでerrがなければそれぞれの結果を繋げて返す",
			security: &testVirtualSecurity{
				stockPositions1: []*vs.StockPosition{
					{
						Code:               "spo-1234",
						OrderCode:          "sor-234",
						SymbolCode:         "1234",
						Side:               vs.SideBuy,
						ContractedQuantity: 300,
						OwnedQuantity:      300,
						HoldQuantity:       100,
						Price:              1000,
						ContractedAt:       time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local),
					},
				},
				marginPositions1: []*vs.MarginPosition{
					{
						Code:               "mpo-1234",
						OrderCode:          "mor-234",
						SymbolCode:         "1234",
						Side:               vs.SideBuy,
						ContractedQuantity: 300,
						OwnedQuantity:      300,
						HoldQuantity:       100,
						Price:              1000,
						ContractedAt:       time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local),
					},
				}},
			arg3: &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_ALL},
			want1: &kabuspb.Positions{Positions: []*kabuspb.Position{
				{
					ExecutionId:     "spo-1234",
					AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					SymbolCode:      "1234",
					SymbolName:      "",
					Exchange:        0,
					ExchangeName:    "",
					SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED,
					ExecutionDay:    timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
					Price:           1000,
					LeavesQuantity:  300,
					HoldQuantity:    100,
					Side:            kabuspb.Side_SIDE_BUY,
					Expenses:        0,
					Commission:      0,
					CommissionTax:   0,
					ExpireDay:       nil,
					MarginTradeType: 0,
					CurrentPrice:    0,
					Valuation:       0,
					ProfitLoss:      0,
					ProfitLossRate:  0,
				},
				{
					ExecutionId:     "mpo-1234",
					AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					SymbolCode:      "1234",
					SymbolName:      "",
					Exchange:        0,
					ExchangeName:    "",
					SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED,
					ExecutionDay:    timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
					Price:           1000,
					LeavesQuantity:  300,
					HoldQuantity:    100,
					Side:            kabuspb.Side_SIDE_BUY,
					Expenses:        0,
					Commission:      0,
					CommissionTax:   0,
					ExpireDay:       nil,
					MarginTradeType: 0,
					CurrentPrice:    0,
					Valuation:       0,
					ProfitLoss:      0,
					ProfitLossRate:  0,
				},
			}},
			want2:                    nil,
			wantStockPositionsCount:  1,
			wantMarginPositionsCount: 1,
		},
		{name: "productがstockでerrが返されたらerrを返す",
			security:                &testVirtualSecurity{stockPositions2: vs.NilArgumentError},
			arg3:                    &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_STOCK},
			want1:                   nil,
			want2:                   vs.NilArgumentError,
			wantStockPositionsCount: 1,
		},
		{name: "productがstockでerrがなければ結果を返す",
			security: &testVirtualSecurity{stockPositions1: []*vs.StockPosition{
				{
					Code:               "spo-1234",
					OrderCode:          "sor-234",
					SymbolCode:         "1234",
					Side:               vs.SideBuy,
					ContractedQuantity: 300,
					OwnedQuantity:      300,
					HoldQuantity:       100,
					Price:              1000,
					ContractedAt:       time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local),
				},
			}},
			arg3: &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_STOCK},
			want1: &kabuspb.Positions{Positions: []*kabuspb.Position{
				{
					ExecutionId:     "spo-1234",
					AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					SymbolCode:      "1234",
					SymbolName:      "",
					Exchange:        0,
					ExchangeName:    "",
					SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED,
					ExecutionDay:    timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
					Price:           1000,
					LeavesQuantity:  300,
					HoldQuantity:    100,
					Side:            kabuspb.Side_SIDE_BUY,
					Expenses:        0,
					Commission:      0,
					CommissionTax:   0,
					ExpireDay:       nil,
					MarginTradeType: 0,
					CurrentPrice:    0,
					Valuation:       0,
					ProfitLoss:      0,
					ProfitLossRate:  0,
				},
			}},
			want2:                   nil,
			wantStockPositionsCount: 1,
		},
		{name: "productがmarginでerrが返されたらerrを返す",
			security:                 &testVirtualSecurity{marginPositions2: vs.NilArgumentError},
			arg3:                     &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_MARGIN},
			want1:                    nil,
			want2:                    vs.NilArgumentError,
			wantMarginPositionsCount: 1,
		},
		{name: "productがmarginでerrがなければ結果を返す",
			security: &testVirtualSecurity{marginPositions1: []*vs.MarginPosition{
				{
					Code:               "mpo-1234",
					OrderCode:          "mor-234",
					SymbolCode:         "1234",
					Side:               vs.SideBuy,
					ContractedQuantity: 300,
					OwnedQuantity:      300,
					HoldQuantity:       100,
					Price:              1000,
					ContractedAt:       time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local),
				},
			}},
			arg3: &kabuspb.GetPositionsRequest{Product: kabuspb.Product_PRODUCT_MARGIN},
			want1: &kabuspb.Positions{Positions: []*kabuspb.Position{
				{
					ExecutionId:     "mpo-1234",
					AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED,
					SymbolCode:      "1234",
					SymbolName:      "",
					Exchange:        0,
					ExchangeName:    "",
					SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED,
					ExecutionDay:    timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
					Price:           1000,
					LeavesQuantity:  300,
					HoldQuantity:    100,
					Side:            kabuspb.Side_SIDE_BUY,
					Expenses:        0,
					Commission:      0,
					CommissionTax:   0,
					ExpireDay:       nil,
					MarginTradeType: 0,
					CurrentPrice:    0,
					Valuation:       0,
					ProfitLoss:      0,
					ProfitLossRate:  0,
				},
			}},
			want2:                    nil,
			wantMarginPositionsCount: 1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.Positions(context.Background(), "no-token", test.arg3)
			if !reflect.DeepEqual(test.want1, got1) ||
				!errors.Is(got2, test.want2) ||
				!reflect.DeepEqual(test.wantStockPositionsCount, test.security.stockPositionsCount) ||
				!reflect.DeepEqual(test.wantMarginPositionsCount, test.security.marginPositionsCount) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					test.want1, test.want2, test.wantStockPositionsCount, test.wantMarginPositionsCount,
					got1, got2, test.security.stockPositionsCount, test.security.marginPositionsCount)
			}
		})
	}
}

func Test_security_SendPrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		registerPrice error
		arg           *kabuspb.Board
		hasError      bool
	}{
		{name: "nilなら何もせずにnilを返す",
			arg:      nil,
			hasError: false},
		{name: "registerPriceがerrを返したらerrを返す",
			registerPrice: errors.New("error message"),
			arg:           &kabuspb.Board{},
			hasError:      true},
		{name: "registerPriceがnilを返したらnilを返す",
			registerPrice: nil,
			arg:           &kabuspb.Board{},
			hasError:      false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: &testVirtualSecurity{registerPrice1: test.registerPrice}}
			got := security.SendPrice(context.Background(), test.arg)
			if (got != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.hasError, got)
			}
		})
	}
}

func Test_security_SendOrderMargin(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		security vs.VirtualSecurity
		want1    *kabuspb.OrderResponse
		want2    error
	}{
		{name: "errが返されたらerrを返す",
			security: &testVirtualSecurity{marginOrder2: vs.NilArgumentError},
			want1:    nil,
			want2:    vs.NilArgumentError,
		},
		{name: "errがなければ結果を返す",
			security: &testVirtualSecurity{marginOrder1: &vs.OrderResult{OrderCode: "sor-1"}},
			want1:    &kabuspb.OrderResponse{ResultCode: 0, OrderId: "sor-1"},
			want2:    nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.SendOrderMargin(context.Background(), "no-token", nil)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_security_CancelOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                       string
		virtual                    *testVirtualSecurity
		arg3                       *kabuspb.CancelOrderRequest
		want1                      *kabuspb.OrderResponse
		want2                      error
		wantCancelStockOrderCount  int
		wantCancelMarginOrderCount int
	}{
		{name: "注文コードが空文字なら何もしない",
			virtual: &testVirtualSecurity{},
			arg3:    &kabuspb.CancelOrderRequest{OrderId: ""},
			want1:   &kabuspb.OrderResponse{}},
		{name: "注文コードの先頭3文字が想定外の形なら何もしない",
			virtual: &testVirtualSecurity{},
			arg3:    &kabuspb.CancelOrderRequest{OrderId: "ORIGINAL-ORDER-CODE"},
			want1:   &kabuspb.OrderResponse{}},
		{name: "注文コードの先頭3文字がsorなら現物注文の取消を叩く",
			virtual:                   &testVirtualSecurity{cancelStockOrder: vs.NoDataError},
			arg3:                      &kabuspb.CancelOrderRequest{OrderId: "sor-uuid-001"},
			want1:                     &kabuspb.OrderResponse{},
			want2:                     vs.NoDataError,
			wantCancelStockOrderCount: 1},
		{name: "注文コードの先頭3文字がmorなら信用注文の取消を叩く",
			virtual:                    &testVirtualSecurity{cancelMarginOrder: vs.NoDataError},
			arg3:                       &kabuspb.CancelOrderRequest{OrderId: "mor-uuid-001"},
			want1:                      &kabuspb.OrderResponse{},
			want2:                      vs.NoDataError,
			wantCancelMarginOrderCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.virtual}
			got1, got2 := security.CancelOrder(context.Background(), "", test.arg3)
			if !reflect.DeepEqual(test.want1, got1) ||
				!errors.Is(got2, test.want2) ||
				test.wantCancelStockOrderCount != test.virtual.cancelStockOrderCount ||
				test.wantCancelMarginOrderCount != test.virtual.cancelMarginOrderCount {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					test.want1, test.want2, test.wantCancelStockOrderCount, test.wantCancelMarginOrderCount,
					got1, got2, test.virtual.cancelStockOrder, test.virtual.cancelMarginOrder)
			}
		})
	}
}

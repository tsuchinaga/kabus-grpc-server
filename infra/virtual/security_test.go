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
	stockOrder1     *vs.OrderResult
	stockOrder2     error
	stockOrders1    []*vs.StockOrder
	stockOrders2    error
	stockPositions1 []*vs.StockPosition
	stockPositions2 error
}

func (t *testVirtualSecurity) StockOrder(*vs.StockOrderRequest) (*vs.OrderResult, error) {
	return t.stockOrder1, t.stockOrder2
}
func (t *testVirtualSecurity) StockOrders() ([]*vs.StockOrder, error) {
	return t.stockOrders1, t.stockOrders2
}
func (t *testVirtualSecurity) StockPositions() ([]*vs.StockPosition, error) {
	return t.stockPositions1, t.stockPositions2
}

func Test_Name(t *testing.T) {
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
		name     string
		security vs.VirtualSecurity
		want1    *kabuspb.Orders
		want2    error
	}{
		{name: "errが返されたらerrを返す",
			security: &testVirtualSecurity{stockOrders2: vs.NilArgumentError},
			want1:    nil,
			want2:    vs.NilArgumentError,
		},
		{name: "errがなければ結果を返す",
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
			want2: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.Orders(context.Background(), "no-token", nil)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_security_Positions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		security vs.VirtualSecurity
		want1    *kabuspb.Positions
		want2    error
	}{
		{name: "errが返されたらerrを返す",
			security: &testVirtualSecurity{stockPositions2: vs.NilArgumentError},
			want1:    nil,
			want2:    vs.NilArgumentError,
		},
		{name: "errがなければ結果を返す",
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
			want2: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{virtual: test.security}
			got1, got2 := security.Positions(context.Background(), "no-token", nil)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

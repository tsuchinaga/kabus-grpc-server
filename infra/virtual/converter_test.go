package virtual

import (
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	vs "gitlab.com/tsuchinaga/kabus-virtual-security"
)

func Test_toSide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.Side
		want vs.Side
	}{
		{name: "未指定 を変換できる", arg: kabuspb.Side_SIDE_UNSPECIFIED, want: vs.SideUnspecified},
		{name: "買い を変換できる", arg: kabuspb.Side_SIDE_BUY, want: vs.SideBuy},
		{name: "売り を変換できる", arg: kabuspb.Side_SIDE_SELL, want: vs.SideSell},
		{name: "未定義 を変換できる", arg: kabuspb.Side(-1), want: vs.SideUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toSide(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toStockExecutionCondition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.StockOrderType
		want vs.StockExecutionCondition
	}{
		{name: "未指定 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_UNSPECIFIED, want: vs.StockExecutionConditionUnspecified},
		{name: "成行 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO, want: vs.StockExecutionConditionMO},
		{name: "寄成(前場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOMO, want: vs.StockExecutionConditionMOMO},
		{name: "寄成(後場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOAO, want: vs.StockExecutionConditionMOAO},
		{name: "引成(前場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOMC, want: vs.StockExecutionConditionMOMC},
		{name: "引成(後場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_MOAC, want: vs.StockExecutionConditionMOAC},
		{name: "IOC成行 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_MO, want: vs.StockExecutionConditionIOCMO},
		{name: "指値 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO, want: vs.StockExecutionConditionLO},
		{name: "寄指(前場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOMO, want: vs.StockExecutionConditionLOMO},
		{name: "寄指(後場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOAO, want: vs.StockExecutionConditionLOAO},
		{name: "引指(前場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOMC, want: vs.StockExecutionConditionLOMC},
		{name: "引指(後場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LOAC, want: vs.StockExecutionConditionLOAC},
		{name: "IOC指値 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_IOC_LO, want: vs.StockExecutionConditionIOCLO},
		{name: "不成(前場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_M, want: vs.StockExecutionConditionFunariM},
		{name: "不成(後場) を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_FUNARI_A, want: vs.StockExecutionConditionFunariA},
		{name: "逆指値 を変換できる", arg: kabuspb.StockOrderType_STOCK_ORDER_TYPE_STOP, want: vs.StockExecutionConditionStop},
		{name: "未定義 を変換できる", arg: kabuspb.StockOrderType(-1), want: vs.StockExecutionConditionUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toStockExecutionCondition(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toComparisonOperator(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.UnderOver
		want vs.ComparisonOperator
	}{
		{name: "未指定 を変換できる", arg: kabuspb.UnderOver_UNDER_OVER_UNSPECIFIED, want: vs.ComparisonOperatorUnspecified},
		{name: "以上 を変換できる", arg: kabuspb.UnderOver_UNDER_OVER_OVER, want: vs.ComparisonOperatorGE},
		{name: "以下 を変換できる", arg: kabuspb.UnderOver_UNDER_OVER_UNDER, want: vs.ComparisonOperatorLE},
		{name: "未定義 を変換できる", arg: kabuspb.UnderOver(-1), want: vs.ComparisonOperatorUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toComparisonOperator(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toExecutionConditionAfterHit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.StockAfterHitOrderType
		want vs.StockExecutionCondition
	}{
		{name: "未指定 を変換できる", arg: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_UNSPECIFIED, want: vs.StockExecutionConditionUnspecified},
		{name: "成行 を変換できる", arg: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_MO, want: vs.StockExecutionConditionMO},
		{name: "指値 を変換できる", arg: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_LO, want: vs.StockExecutionConditionLO},
		{name: "不成 を変換できる", arg: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_FUNARI, want: vs.StockExecutionConditionFunariA},
		{name: "未定義 を変換できる", arg: kabuspb.StockAfterHitOrderType(-1), want: vs.StockExecutionConditionUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExecutionConditionAfterHit(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toStopCondition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabuspb.StockStopOrder
		want *vs.StockStopCondition
	}{
		{name: "nilならnilを返す"},
		{name: "変換して返せる", arg: &kabuspb.StockStopOrder{
			TriggerType:       kabuspb.TriggerType_TRIGGER_TYPE_ORDER_SYMBOL,
			TriggerPrice:      1000,
			UnderOver:         kabuspb.UnderOver_UNDER_OVER_UNDER,
			AfterHitOrderType: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_LO,
			AfterHitPrice:     995,
		}, want: &vs.StockStopCondition{
			StopPrice:                  1000,
			ComparisonOperator:         vs.ComparisonOperatorLE,
			ExecutionConditionAfterHit: vs.StockExecutionConditionLO,
			LimitPriceAfterHit:         995,
		}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toStopCondition(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderStatusToState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.OrderStatus
		want kabuspb.State
	}{
		{name: "未指定 を変換できる", arg: vs.OrderStatusUnspecified, want: kabuspb.State_STATE_UNSPECIFIED},
		{name: "新規 を変換できる", arg: vs.OrderStatusNew, want: kabuspb.State_STATE_PROCESSING},
		{name: "待機 を変換できる", arg: vs.OrderStatusWait, want: kabuspb.State_STATE_WAIT},
		{name: "注文中 を変換できる", arg: vs.OrderStatusInOrder, want: kabuspb.State_STATE_PROCESSED},
		{name: "部分約定 を変換できる", arg: vs.OrderStatusPart, want: kabuspb.State_STATE_PROCESSED},
		{name: "全約定 を変換できる", arg: vs.OrderStatusDone, want: kabuspb.State_STATE_DONE},
		{name: "取消中 を変換できる", arg: vs.OrderStatusInCancel, want: kabuspb.State_STATE_PROCESSED},
		{name: "取消済み を変換できる", arg: vs.OrderStatusCanceled, want: kabuspb.State_STATE_DONE},
		{name: "未定義 を変換できる", arg: vs.OrderStatus("foo"), want: kabuspb.State_STATE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderStatusToState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderStatusToOrderState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.OrderStatus
		want kabuspb.OrderState
	}{
		{name: "未指定 を変換できる", arg: vs.OrderStatusUnspecified, want: kabuspb.OrderState_ORDER_STATE_UNSPECIFIED},
		{name: "新規 を変換できる", arg: vs.OrderStatusNew, want: kabuspb.OrderState_ORDER_STATE_PROCESSING},
		{name: "待機 を変換できる", arg: vs.OrderStatusWait, want: kabuspb.OrderState_ORDER_STATE_WAIT},
		{name: "注文中 を変換できる", arg: vs.OrderStatusInOrder, want: kabuspb.OrderState_ORDER_STATE_PROCESSED},
		{name: "部分約定 を変換できる", arg: vs.OrderStatusPart, want: kabuspb.OrderState_ORDER_STATE_PROCESSED},
		{name: "全約定 を変換できる", arg: vs.OrderStatusDone, want: kabuspb.OrderState_ORDER_STATE_DONE},
		{name: "取消中 を変換できる", arg: vs.OrderStatusInCancel, want: kabuspb.OrderState_ORDER_STATE_PROCESSED},
		{name: "取消済み を変換できる", arg: vs.OrderStatusCanceled, want: kabuspb.OrderState_ORDER_STATE_DONE},
		{name: "未定義 を変換できる", arg: vs.OrderStatus("foo"), want: kabuspb.OrderState_ORDER_STATE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderStatusToOrderState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromStockExecutionCondition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.StockExecutionCondition
		want kabuspb.OrderType
	}{
		{name: "未指定 を変換できる", arg: vs.StockExecutionConditionUnspecified, want: kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED},
		{name: "成行 を変換できる", arg: vs.StockExecutionConditionMO, want: kabuspb.OrderType_ORDER_TYPE_ZARABA},
		{name: "寄成(前場) を変換できる", arg: vs.StockExecutionConditionMOMO, want: kabuspb.OrderType_ORDER_TYPE_OPEN},
		{name: "寄成(後場) を変換できる", arg: vs.StockExecutionConditionMOAO, want: kabuspb.OrderType_ORDER_TYPE_OPEN},
		{name: "引成(前場) を変換できる", arg: vs.StockExecutionConditionMOMC, want: kabuspb.OrderType_ORDER_TYPE_CLOSE},
		{name: "引成(後場) を変換できる", arg: vs.StockExecutionConditionMOAC, want: kabuspb.OrderType_ORDER_TYPE_CLOSE},
		{name: "IOC成行 を変換できる", arg: vs.StockExecutionConditionIOCMO, want: kabuspb.OrderType_ORDER_TYPE_IOC},
		{name: "指値 を変換できる", arg: vs.StockExecutionConditionLO, want: kabuspb.OrderType_ORDER_TYPE_ZARABA},
		{name: "寄指(前場) を変換できる", arg: vs.StockExecutionConditionLOMO, want: kabuspb.OrderType_ORDER_TYPE_OPEN},
		{name: "寄指(後場) を変換できる", arg: vs.StockExecutionConditionLOAO, want: kabuspb.OrderType_ORDER_TYPE_OPEN},
		{name: "引指(前場) を変換できる", arg: vs.StockExecutionConditionLOMC, want: kabuspb.OrderType_ORDER_TYPE_CLOSE},
		{name: "引指(後場) を変換できる", arg: vs.StockExecutionConditionLOAC, want: kabuspb.OrderType_ORDER_TYPE_CLOSE},
		{name: "IOC指値 を変換できる", arg: vs.StockExecutionConditionIOCLO, want: kabuspb.OrderType_ORDER_TYPE_IOC},
		{name: "不成(前場) を変換できる", arg: vs.StockExecutionConditionFunariM, want: kabuspb.OrderType_ORDER_TYPE_FUNARI},
		{name: "不成(後場) を変換できる", arg: vs.StockExecutionConditionFunariA, want: kabuspb.OrderType_ORDER_TYPE_FUNARI},
		{name: "逆指値 を変換できる", arg: vs.StockExecutionConditionStop, want: kabuspb.OrderType_ORDER_TYPE_ZARABA},
		{name: "未定義 を変換できる", arg: vs.StockExecutionCondition("foo"), want: kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromStockExecutionCondition(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromSide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.Side
		want kabuspb.Side
	}{
		{name: "未指定 を変換できる", arg: vs.SideUnspecified, want: kabuspb.Side_SIDE_UNSPECIFIED},
		{name: "買い を変換できる", arg: vs.SideBuy, want: kabuspb.Side_SIDE_BUY},
		{name: "売り を変換できる", arg: vs.SideSell, want: kabuspb.Side_SIDE_SELL},
		{name: "未定義 を変換できる", arg: vs.Side("foo"), want: kabuspb.Side_SIDE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromSide(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_tradeTypeFromSide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.Side
		want kabuspb.TradeType
	}{
		{name: "未指定 を変換できる", arg: vs.SideUnspecified, want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
		{name: "買い を変換できる", arg: vs.SideBuy, want: kabuspb.TradeType_TRADE_TYPE_ENTRY},
		{name: "売り を変換できる", arg: vs.SideSell, want: kabuspb.TradeType_TRADE_TYPE_EXIT},
		{name: "未定義 を変換できる", arg: vs.Side("foo"), want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := tradeTypeFromSide(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromContracts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*vs.Contract
		want []*kabuspb.OrderDetail
	}{
		{name: "nilなら空配列を返す", arg: nil, want: []*kabuspb.OrderDetail{}},
		{name: "空配列なら空配列を返す", arg: []*vs.Contract{}, want: []*kabuspb.OrderDetail{}},
		{name: "nilだけの配列なら空配列を返す", arg: []*vs.Contract{nil, nil, nil}, want: []*kabuspb.OrderDetail{}},
		{name: "各約定データをマッピングして返す", arg: []*vs.Contract{
			{
				ContractCode: "sco-1234",
				OrderCode:    "sor-1111",
				PositionCode: "spo-100",
				Price:        1000,
				Quantity:     300,
				ContractedAt: time.Date(2021, 7, 16, 13, 0, 0, 0, time.Local),
			}, {
				ContractCode: "sco-2345",
				OrderCode:    "sor-2222",
				PositionCode: "spo-200",
				Price:        1500,
				Quantity:     100,
				ContractedAt: time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local),
			},
		}, want: []*kabuspb.OrderDetail{
			{
				SequenceNumber: 1,
				Id:             "1",
				RecordType:     kabuspb.RecordType_RECORD_TYPE_CONTRACTED,
				ExchangeId:     "virtual-security",
				State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
				TransactTime:   timestamppb.New(time.Date(2021, 7, 16, 13, 0, 0, 0, time.Local)),
				OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA,
				Price:          1000,
				Quantity:       300,
				ExecutionId:    "spo-100",
				ExecutionDay:   timestamppb.New(time.Date(2021, 7, 16, 13, 0, 0, 0, time.Local)),
				DeliveryDay:    nil,
				Commission:     0,
				CommissionTax:  0,
			}, {
				SequenceNumber: 2,
				Id:             "2",
				RecordType:     kabuspb.RecordType_RECORD_TYPE_CONTRACTED,
				ExchangeId:     "virtual-security",
				State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
				TransactTime:   timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
				OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA,
				Price:          1500,
				Quantity:       100,
				ExecutionId:    "spo-200",
				ExecutionDay:   timestamppb.New(time.Date(2021, 7, 16, 14, 0, 0, 0, time.Local)),
				DeliveryDay:    nil,
				Commission:     0,
				CommissionTax:  0,
			},
		}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromContracts(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toExpiredAt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *timestamppb.Timestamp
		want time.Time
	}{
		{name: "nilならゼロ値に変換", arg: nil, want: time.Time{}},
		{name: "ゼロ値ならゼロ値を返す", arg: timestamppb.New(time.Time{}), want: time.Time{}},
		{name: "有効な値なら年月日を返す",
			arg:  timestamppb.New(time.Date(2021, 7, 17, 16, 17, 0, 0, time.Local)),
			want: time.Date(2021, 7, 17, 0, 0, 0, 0, time.Local)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExpiredAt(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toStockOrderRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabuspb.SendStockOrderRequest
		want *vs.StockOrderRequest
	}{
		{name: "nilならnilを返す", arg: nil, want: nil},
		{name: "各要素を対応付けて返す",
			arg: &kabuspb.SendStockOrderRequest{
				Password:     "password",
				SymbolCode:   "1234",
				Exchange:     kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU,
				Side:         kabuspb.Side_SIDE_BUY,
				DeliveryType: kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
				FundType:     kabuspb.FundType_FUND_TYPE_MARGIN_TRADING,
				AccountType:  kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				Quantity:     100,
				OrderType:    kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO,
				Price:        1000,
				ExpireDay:    nil,
				StopOrder:    nil,
				IsVirtual:    true,
			},
			want: &vs.StockOrderRequest{
				Side:               vs.SideBuy,
				ExecutionCondition: vs.StockExecutionConditionLO,
				SymbolCode:         "1234",
				Quantity:           100,
				LimitPrice:         1000,
				ExpiredAt:          time.Time{},
				StopCondition:      nil,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toStockOrderRequest(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *vs.OrderResult
		want *kabuspb.OrderResponse
	}{
		{name: "nilならnilを返す", arg: nil, want: nil},
		{name: "各要素を対応付けて返す",
			arg:  &vs.OrderResult{OrderCode: "sor-1234"},
			want: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "sor-1234"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderResult(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromStockOrders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*vs.StockOrder
		want *kabuspb.Orders
	}{
		{name: "nilなら空配列を返す", arg: nil, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "空配列でも空配列を返す", arg: []*vs.StockOrder{}, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "各要素を対応付けて返す",
			arg: []*vs.StockOrder{
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
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{
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
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromStockOrders(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromStockPositions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*vs.StockPosition
		want *kabuspb.Positions
	}{
		{name: "nilなら空配列を返す", arg: nil, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "空配列でも空配列を返す", arg: []*vs.StockPosition{}, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "各要素を対応付けて返す",
			arg: []*vs.StockPosition{
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
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{
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
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromStockPositions(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

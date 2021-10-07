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

func Test_toStopConditionFromStockStopOrder(t *testing.T) {
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
			got := toStopConditionFromStockStopOrder(test.arg)
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
		{name: "要素がnilならスキップする", arg: []*vs.StockOrder{nil, nil, nil}, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
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
		{name: "各要素がnilならスキップする", arg: []*vs.StockPosition{nil, nil, nil}, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
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

func Test_toExchangeType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.Exchange
		want vs.ExchangeType
	}{
		{name: "未指定 を変換できる", arg: kabuspb.Exchange_EXCHANGE_UNSPECIFIED, want: vs.ExchangeTypeUnspecified},
		{name: "東証 を変換できる", arg: kabuspb.Exchange_EXCHANGE_TOUSHOU, want: vs.ExchangeTypeStock},
		{name: "名証 を変換できる", arg: kabuspb.Exchange_EXCHANGE_MEISHOU, want: vs.ExchangeTypeStock},
		{name: "福証 を変換できる", arg: kabuspb.Exchange_EXCHANGE_FUKUSHOU, want: vs.ExchangeTypeStock},
		{name: "札証 を変換できる", arg: kabuspb.Exchange_EXCHANGE_SATSUSHOU, want: vs.ExchangeTypeStock},
		{name: "日通し を変換できる", arg: kabuspb.Exchange_EXCHANGE_ALL_SESSION, want: vs.ExchangeTypeFuture},
		{name: "日中場 を変換できる", arg: kabuspb.Exchange_EXCHANGE_DAY_SESSION, want: vs.ExchangeTypeFuture},
		{name: "夕場 を変換できる", arg: kabuspb.Exchange_EXCHANGE_NIGHT_SESSION, want: vs.ExchangeTypeFuture},
		{name: "未定義 を変換できる", arg: kabuspb.Exchange(-1), want: vs.ExchangeTypeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExchangeType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toRegisterPrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabuspb.Board
		want vs.RegisterPriceRequest
	}{
		{name: "nilならゼロ値を返す", arg: nil, want: vs.RegisterPriceRequest{}},
		{name: "各項目をマッピングして返す",
			arg: &kabuspb.Board{
				SymbolCode:               "5401",
				SymbolName:               "新日鐵住金",
				Exchange:                 kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:             "東証１部",
				CurrentPrice:             2408,
				CurrentPriceTime:         timestamppb.New(time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local)),
				CurrentPriceChangeStatus: "0058",
				CurrentPriceStatus:       1,
				CalculationPrice:         343.7,
				PreviousClose:            1048,
				PreviousCloseTime:        timestamppb.New(time.Date(2020, 7, 21, 0, 0, 0, 0, time.Local)),
				ChangePreviousClose:      1360,
				ChangePreviousClosePer:   129.77,
				OpeningPrice:             2380,
				OpeningPriceTime:         timestamppb.New(time.Date(2020, 7, 22, 9, 0, 0, 0, time.Local)),
				HighPrice:                2418,
				HighPriceTime:            timestamppb.New(time.Date(2020, 7, 22, 13, 25, 47, 0, time.Local)),
				LowPrice:                 2370,
				LowPriceTime:             timestamppb.New(time.Date(2020, 7, 22, 10, 0, 4, 0, time.Local)),
				TradingVolume:            4571500,
				TradingVolumeTime:        timestamppb.New(time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local)),
				Vwap:                     2394.4262,
				TradingValue:             10946119350,
				BidQuantity:              100,
				BidPrice:                 2408.5,
				BidTime:                  timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)),
				BidSign:                  "0101",
				MarketOrderSellQuantity:  0,
				Sell1:                    &kabuspb.FirstQuote{Time: timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)), Sign: "0101", Price: 2408.5, Quantity: 100},
				Sell2:                    &kabuspb.Quote{Price: 2409, Quantity: 800},
				Sell3:                    &kabuspb.Quote{Price: 2409.5, Quantity: 2100},
				Sell4:                    &kabuspb.Quote{Price: 2410, Quantity: 800},
				Sell5:                    &kabuspb.Quote{Price: 2410.5, Quantity: 500},
				Sell6:                    &kabuspb.Quote{Price: 2411, Quantity: 8400},
				Sell7:                    &kabuspb.Quote{Price: 2411.5, Quantity: 1200},
				Sell8:                    &kabuspb.Quote{Price: 2412, Quantity: 27200},
				Sell9:                    &kabuspb.Quote{Price: 2412.5, Quantity: 400},
				Sell10:                   &kabuspb.Quote{Price: 2413, Quantity: 16400},
				AskQuantity:              200,
				AskPrice:                 2407.5,
				AskTime:                  timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)),
				AskSign:                  "0101",
				MarketOrderBuyQuantity:   0,
				Buy1:                     &kabuspb.FirstQuote{Time: timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)), Sign: "0101", Price: 2407.5, Quantity: 200},
				Buy2:                     &kabuspb.Quote{Price: 2407, Quantity: 400},
				Buy3:                     &kabuspb.Quote{Price: 2406.5, Quantity: 1000},
				Buy4:                     &kabuspb.Quote{Price: 2406, Quantity: 5800},
				Buy5:                     &kabuspb.Quote{Price: 2405.5, Quantity: 7500},
				Buy6:                     &kabuspb.Quote{Price: 2405, Quantity: 2200},
				Buy7:                     &kabuspb.Quote{Price: 2404.5, Quantity: 16700},
				Buy8:                     &kabuspb.Quote{Price: 2404, Quantity: 30100},
				Buy9:                     &kabuspb.Quote{Price: 2403.5, Quantity: 1300},
				Buy10:                    &kabuspb.Quote{Price: 2403, Quantity: 3000},
				OverSellQuantity:         974900,
				UnderBuyQuantity:         756000,
				TotalMarketValue:         3266254659361.4,
				ClearingPrice:            23000,
				ImpliedVolatility:        22.11,
				Gamma:                    0.000183,
				Theta:                    -6.5073,
				Vega:                     39.3109,
				Delta:                    0.4794,
			},
			want: vs.RegisterPriceRequest{
				ExchangeType: vs.ExchangeTypeStock,
				SymbolCode:   "5401",
				Price:        2408,
				PriceTime:    time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local),
				Ask:          2407.5,
				AskTime:      time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
				Bid:          2408.5,
				BidTime:      time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toRegisterPrice(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toTradeType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.TradeType
		want vs.TradeType
	}{
		{name: "未指定 を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED, want: vs.TradeTypeUnspecified},
		{name: "Entry を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_ENTRY, want: vs.TradeTypeEntry},
		{name: "Exit を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_EXIT, want: vs.TradeTypeExit},
		{name: "未定義 を変換できる", arg: kabuspb.TradeType(-1), want: vs.TradeTypeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toTradeType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toStopConditionFromMarginStopOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabuspb.MarginStopOrder
		want *vs.StockStopCondition
	}{
		{name: "nilならnilを返す"},
		{name: "変換して返せる",
			arg: &kabuspb.MarginStopOrder{
				TriggerType:       kabuspb.TriggerType_TRIGGER_TYPE_ORDER_SYMBOL,
				TriggerPrice:      1000,
				UnderOver:         kabuspb.UnderOver_UNDER_OVER_UNDER,
				AfterHitOrderType: kabuspb.StockAfterHitOrderType_STOCK_AFTER_HIT_ORDER_TYPE_LO,
				AfterHitPrice:     995,
			},
			want: &vs.StockStopCondition{
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
			got := toStopConditionFromMarginStopOrder(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toExitPositionList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*kabuspb.ClosePosition
		want []vs.ExitPosition
	}{
		{name: "nilならnilが返される", arg: nil, want: nil},
		{name: "配列が空なら空配列が返される", arg: []*kabuspb.ClosePosition{}, want: []vs.ExitPosition{}},
		{name: "配列内のnilはスキップされる", arg: []*kabuspb.ClosePosition{nil, nil, nil}, want: []vs.ExitPosition{}},
		{name: "配列内の各項目が変換されて返される",
			arg: []*kabuspb.ClosePosition{
				{ExecutionId: "mpo-uuid-01", Quantity: 100},
				{ExecutionId: "mpo-uuid-02", Quantity: 50},
				{ExecutionId: "mpo-uuid-03", Quantity: 200},
			},
			want: []vs.ExitPosition{
				{PositionCode: "mpo-uuid-01", Quantity: 100},
				{PositionCode: "mpo-uuid-02", Quantity: 50},
				{PositionCode: "mpo-uuid-03", Quantity: 200},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExitPositionList(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.arg, got)
			}
		})
	}
}

func Test_toMarginOrderRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabuspb.SendMarginOrderRequest
		want *vs.MarginOrderRequest
	}{
		{name: "nilならnilを返す", arg: nil, want: nil},
		{name: "各要素を対応付けて返す",
			arg: &kabuspb.SendMarginOrderRequest{
				Password:        "password",
				SymbolCode:      "1234",
				Exchange:        kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU,
				Side:            kabuspb.Side_SIDE_BUY,
				TradeType:       kabuspb.TradeType_TRADE_TYPE_EXIT,
				MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM,
				DeliveryType:    kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
				AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				Quantity:        300,
				ClosePositions: []*kabuspb.ClosePosition{
					{ExecutionId: "mpo-uuid-01", Quantity: 100},
					{ExecutionId: "mpo-uuid-02", Quantity: 200},
				},
				OrderType: kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO,
				Price:     1000,
				ExpireDay: nil,
				StopOrder: nil,
				IsVirtual: true,
			},
			want: &vs.MarginOrderRequest{
				TradeType:          vs.TradeTypeExit,
				Side:               vs.SideBuy,
				ExecutionCondition: vs.StockExecutionConditionLO,
				SymbolCode:         "1234",
				Quantity:           300,
				LimitPrice:         1000,
				ExpiredAt:          time.Time{},
				StopCondition:      nil,
				ExitPositionList: []vs.ExitPosition{
					{PositionCode: "mpo-uuid-01", Quantity: 100},
					{PositionCode: "mpo-uuid-02", Quantity: 200},
				},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toMarginOrderRequest(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromMarginOrders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*vs.MarginOrder
		want *kabuspb.Orders
	}{
		{name: "nilなら空配列を返す", arg: nil, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "空配列でも空配列を返す", arg: []*vs.MarginOrder{}, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "要素がnilならスキップする", arg: []*vs.MarginOrder{nil, nil, nil}, want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "各要素を対応付けて返す",
			arg: []*vs.MarginOrder{
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
					OrderedAt:          time.Date(2021, 7, 16, 9, 0, 0, 0, time.Local),
					CanceledAt:         time.Time{},
					Contracts:          []*vs.Contract{},
					Message:            "",
				},
			},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{
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
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromMarginOrders(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromTradeType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  vs.TradeType
		want kabuspb.TradeType
	}{
		{name: "未指定 を変換できる", arg: vs.TradeTypeUnspecified, want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
		{name: "買い を変換できる", arg: vs.TradeTypeEntry, want: kabuspb.TradeType_TRADE_TYPE_ENTRY},
		{name: "売り を変換できる", arg: vs.TradeTypeExit, want: kabuspb.TradeType_TRADE_TYPE_EXIT},
		{name: "未定義 を変換できる", arg: vs.TradeType("foo"), want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromTradeType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromMarginPositions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*vs.MarginPosition
		want *kabuspb.Positions
	}{
		{name: "nilなら空配列を返す", arg: nil, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "空配列でも空配列を返す", arg: []*vs.MarginPosition{}, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "各要素がnilならスキップする", arg: []*vs.MarginPosition{nil, nil, nil}, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "各要素を対応付けて返す",
			arg: []*vs.MarginPosition{
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
			got := fromMarginPositions(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

package security

import (
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_toExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.Exchange
		want kabus.Exchange
	}{
		{name: "東証の変換ができる", arg: kabuspb.Exchange_EXCHANGE_TOUSHOU, want: kabus.ExchangeToushou},
		{name: "名証の変換ができる", arg: kabuspb.Exchange_EXCHANGE_MEISHOU, want: kabus.ExchangeMeishou},
		{name: "福証の変換ができる", arg: kabuspb.Exchange_EXCHANGE_FUKUSHOU, want: kabus.ExchangeFukushou},
		{name: "札証の変換ができる", arg: kabuspb.Exchange_EXCHANGE_SATSUSHOU, want: kabus.ExchangeSatsushou},
		{name: "日通しの変換ができる", arg: kabuspb.Exchange_EXCHANGE_ALL_SESSION, want: kabus.ExchangeAll},
		{name: "日中の変換ができる", arg: kabuspb.Exchange_EXCHANGE_DAY_SESSION, want: kabus.ExchangeDaytime},
		{name: "夜間の変換ができる", arg: kabuspb.Exchange_EXCHANGE_NIGHT_SESSION, want: kabus.ExchangeEvening},
		{name: "指定なしの変換ができる", arg: kabuspb.Exchange_EXCHANGE_UNSPECIFIED, want: kabus.ExchangeUnspecified},
		{name: "未定義の変換ができる", arg: kabuspb.Exchange(-1), want: kabus.ExchangeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExchange(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.Exchange
		want kabuspb.Exchange
	}{
		{name: "東証の変換ができる", arg: kabus.ExchangeToushou, want: kabuspb.Exchange_EXCHANGE_TOUSHOU},
		{name: "名証の変換ができる", arg: kabus.ExchangeMeishou, want: kabuspb.Exchange_EXCHANGE_MEISHOU},
		{name: "福証の変換ができる", arg: kabus.ExchangeFukushou, want: kabuspb.Exchange_EXCHANGE_FUKUSHOU},
		{name: "札証の変換ができる", arg: kabus.ExchangeSatsushou, want: kabuspb.Exchange_EXCHANGE_SATSUSHOU},
		{name: "日通しの変換ができる", arg: kabus.ExchangeAll, want: kabuspb.Exchange_EXCHANGE_ALL_SESSION},
		{name: "日中の変換ができる", arg: kabus.ExchangeDaytime, want: kabuspb.Exchange_EXCHANGE_DAY_SESSION},
		{name: "夜間の変換ができる", arg: kabus.ExchangeEvening, want: kabuspb.Exchange_EXCHANGE_NIGHT_SESSION},
		{name: "指定なしの変換ができる", arg: kabus.ExchangeUnspecified, want: kabuspb.Exchange_EXCHANGE_UNSPECIFIED},
		{name: "未定義の変換ができる", arg: kabus.Exchange(-1), want: kabuspb.Exchange_EXCHANGE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromExchange(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toFutureCode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.FutureCode
		want kabus.FutureCode
	}{
		{name: "日経平均先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_NK225, want: kabus.FutureCodeNK225},
		{name: "日経225mini先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_NK225_MINI, want: kabus.FutureCodeNK225Mini},
		{name: "TOPIX先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_TOPIX, want: kabus.FutureCodeTOPIX},
		{name: "ミニTOPIX先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_TOPIX_MINI, want: kabus.FutureCodeTOPIXMini},
		{name: "東証マザーズ先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_MOTHERS, want: kabus.FutureCodeMOTHERS},
		{name: "JPX日経400先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_JPX400, want: kabus.FutureCodeJPX400},
		{name: "NYダウ先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_DOW, want: kabus.FutureCodeDOW},
		{name: "日経平均VI先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_VI, want: kabus.FutureCodeVI},
		{name: "TOPIX Core30先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_CORE30, want: kabus.FutureCodeCore30},
		{name: "東証REIT指数先物の変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_REIT, want: kabus.FutureCodeREIT},
		{name: "指定なしの変換ができる", arg: kabuspb.FutureCode_FUTURE_CODE_UNSPECIFIED, want: kabus.FutureCodeUnspecified},
		{name: "未定義の変換ができる", arg: kabuspb.FutureCode(-1), want: kabus.FutureCodeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toFutureCode(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toYmNum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *timestamppb.Timestamp
		want kabus.YmNUM
	}{
		{name: "ゼロ値なら0を意味するYmNUMが返される", arg: timestamppb.New(time.Time{}), want: kabus.YmNUMToday},
		{name: "Time型を持ったYmNUMが返される",
			arg:  timestamppb.New(time.Date(2021, 3, 28, 10, 27, 0, 0, time.Local)),
			want: kabus.NewYmNUM(time.Date(2021, 3, 28, 10, 27, 0, 0, time.Local))},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toYmNum(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toPutOrCall(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.CallPut
		want kabus.PutOrCall
	}{
		{name: "CALLの変換ができる", arg: kabuspb.CallPut_CALL_PUT_CALL, want: kabus.PutOrCallCall},
		{name: "PUTの変換ができる", arg: kabuspb.CallPut_CALL_PUT_PUT, want: kabus.PutOrCallPut},
		{name: "指定なしの変換ができる", arg: kabuspb.CallPut_CALL_PUT_UNSPECIFIED, want: kabus.PutOrCallUnspecified},
		{name: "未定義の変換ができる", arg: kabuspb.CallPut(-1), want: kabus.PutOrCallUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toPutOrCall(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromCurrentPriceChangeStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.CurrentPriceChangeStatus
		want string
	}{
		{name: "未指定 を変換できる", arg: kabus.CurrentPriceChangeStatusUnspecified, want: ""},
		{name: "事象なし を変換できる", arg: kabus.CurrentPriceChangeStatusNoEffect, want: "0000"},
		{name: "変わらず を変換できる", arg: kabus.CurrentPriceChangeStatusNoChange, want: "0056"},
		{name: "UP を変換できる", arg: kabus.CurrentPriceChangeStatusUp, want: "0057"},
		{name: "DOWN を変換できる", arg: kabus.CurrentPriceChangeStatusDown, want: "0058"},
		{name: "中断板寄り後の初値 を変換できる", arg: kabus.CurrentPriceChangeStatusOpenPriceAfterBreak, want: "0059"},
		{name: "ザラバ引け を変換できる", arg: kabus.CurrentPriceChangeStatusTradingSessionClose, want: "0060"},
		{name: "板寄り引け を変換できる", arg: kabus.CurrentPriceChangeStatusClose, want: "0061"},
		{name: "中断引け を変換できる", arg: kabus.CurrentPriceChangeStatusBreakClose, want: "0062"},
		{name: "ダウン引け を変換できる", arg: kabus.CurrentPriceChangeStatusDownClose, want: "0063"},
		{name: "逆転終値 を変換できる", arg: kabus.CurrentPriceChangeStatusTarnOverClose, want: "0064"},
		{name: "特別気配引け を変換できる", arg: kabus.CurrentPriceChangeStatusSpecialQuoteClose, want: "0066"},
		{name: "一時留保引け を変換できる", arg: kabus.CurrentPriceChangeStatusReservationClose, want: "0067"},
		{name: "売買停止引け を変換できる", arg: kabus.CurrentPriceChangeStatusStopClose, want: "0068"},
		{name: "サーキットブレーカ引け を変換できる", arg: kabus.CurrentPriceChangeCircuitBreakerClose, want: "0069"},
		{name: "ダイナミックサーキットブレーカ引け を変換できる", arg: kabus.CurrentPriceChangeDynamicCircuitBreakerClose, want: "0431"},
		{name: "未定義 を変換できる", arg: kabus.CurrentPriceChangeStatus("-1"), want: ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromCurrentPriceChangeStatus(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromCurrentPriceStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.CurrentPriceStatus
		want int32
	}{
		{name: "指定なし を変換できる", arg: kabus.CurrentPriceStatusUnspecified, want: 0},
		{name: "現値 を変換できる", arg: kabus.CurrentPriceStatusCurrentPrice, want: 1},
		{name: "不連続歩み を変換できる", arg: kabus.CurrentPriceStatusNoContinuousTicks, want: 2},
		{name: "板寄せ を変換できる", arg: kabus.CurrentPriceStatusItayose, want: 3},
		{name: "システム障害 を変換できる", arg: kabus.CurrentPriceStatusSystemError, want: 4},
		{name: "中断 を変換できる", arg: kabus.CurrentPriceStatusPause, want: 5},
		{name: "売買停止 を変換できる", arg: kabus.CurrentPriceStatusStopTrading, want: 6},
		{name: "売買停止・システム停止解除 を変換できる", arg: kabus.CurrentPriceStatusRestart, want: 7},
		{name: "終値 を変換できる", arg: kabus.CurrentPriceStatusClosePrice, want: 8},
		{name: "システム停止 を変換できる", arg: kabus.CurrentPriceStatusSystemStop, want: 9},
		{name: "概算値 を変換できる", arg: kabus.CurrentPriceStatusRoughQuote, want: 10},
		{name: "参考値 を変換できる", arg: kabus.CurrentPriceStatusReference, want: 11},
		{name: "サーキットブレイク実施中 を変換できる", arg: kabus.CurrentPriceStatusInCircuitBreak, want: 12},
		{name: "システム障害解除 を変換できる", arg: kabus.CurrentPriceStatusRestoration, want: 13},
		{name: "サーキットブレイク解除 を変換できる", arg: kabus.CurrentPriceStatusReleaseCircuitBreak, want: 14},
		{name: "中断解除 を変換できる", arg: kabus.CurrentPriceStatusReleasePause, want: 15},
		{name: "一時留保中 を変換できる", arg: kabus.CurrentPriceStatusInReservation, want: 16},
		{name: "一時留保解除 を変換できる", arg: kabus.CurrentPriceStatusReleaseReservation, want: 17},
		{name: "ファイル障害 を変換できる", arg: kabus.CurrentPriceStatusFileError, want: 18},
		{name: "ファイル障害解除 を変換できる", arg: kabus.CurrentPriceStatusReleaseFileError, want: 19},
		{name: "Spread/Strategy を変換できる", arg: kabus.CurrentPriceStatusSpreadStrategy, want: 20},
		{name: "ダイナミックサーキットブレイク発動 を変換できる", arg: kabus.CurrentPriceStatusInDynamicCircuitBreak, want: 21},
		{name: "ダイナミックサーキットブレイク解除 を変換できる", arg: kabus.CurrentPriceStatusReleaseDynamicCircuitBreak, want: 22},
		{name: "板寄せ約定 を変換できる", arg: kabus.CurrentPriceStatusContractedInItayose, want: 23},
		{name: "未定義 を変換できる", arg: kabus.CurrentPriceStatus(-1), want: 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromCurrentPriceStatus(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromBidAskSign(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.BidAskSign
		want string
	}{
		{name: "指定なし を変換できる", arg: kabus.BidAskSignUnspecified, want: ""},
		{name: "事象なし を変換できる", arg: kabus.BidAskSignNoEffect, want: "0000"},
		{name: "一般気配 を変換できる", arg: kabus.BidAskSignGeneral, want: "0101"},
		{name: "特別気配 を変換できる", arg: kabus.BidAskSignSpecial, want: "0102"},
		{name: "注意気配 を変換できる", arg: kabus.BidAskSignAttention, want: "0103"},
		{name: "寄前気配 を変換できる", arg: kabus.BidAskSignBeforeOpen, want: "0107"},
		{name: "停止前特別気配 を変換できる", arg: kabus.BidAskSignSpecialBeforeStop, want: "0108"},
		{name: "引け後気配 を変換できる", arg: kabus.BidAskSignAfterClose, want: "0109"},
		{name: "寄前気配約定成立ポイントなし を変換できる", arg: kabus.BidAskSignNotExistsContractPoint, want: "0116"},
		{name: "寄前気配約定成立ポイントあり を変換できる", arg: kabus.BidAskSignExistsContractPoint, want: "0117"},
		{name: "連続約定気配 を変換できる", arg: kabus.BidAskSignContinuous, want: "0118"},
		{name: "停止前の連続約定気配 を変換できる", arg: kabus.BidAskSignContinuousBeforeStop, want: "0119"},
		{name: "買い上がり売り下がり中 を変換できる", arg: kabus.BidAskSignMoving, want: "0120"},
		{name: "未定義 を変換できる", arg: kabus.BidAskSign("-1"), want: ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromBidAskSign(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromFirstBoardSign(t *testing.T) {
	t.Parallel()
	got := fromFirstBoardSign(kabus.FirstBoardSign{
		Time:  time.Date(2021, 3, 30, 22, 53, 0, 0, time.Local),
		Sign:  kabus.BidAskSignNoEffect,
		Price: 22500,
		Qty:   1,
	})
	want := &kabuspb.FirstQuote{
		Time:     timestamppb.New(time.Date(2021, 3, 30, 22, 53, 0, 0, time.Local)),
		Sign:     "0000",
		Price:    22500,
		Quantity: 1,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_fromBoardSign(t *testing.T) {
	t.Parallel()
	got := fromBoardSign(kabus.BoardSign{
		Price: 22500,
		Qty:   1,
	})
	want := &kabuspb.Quote{
		Price:    22500,
		Quantity: 1,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_fromPutOrCallNum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.PutOrCallNum
		want kabuspb.CallPut
	}{
		{name: "未指定 を変換できる", arg: kabus.PutOrCallNumUnspecified, want: kabuspb.CallPut_CALL_PUT_UNSPECIFIED},
		{name: "プット を変換できる", arg: kabus.PutOrCallNumPut, want: kabuspb.CallPut_CALL_PUT_PUT},
		{name: "コール を変換できる", arg: kabus.PutOrCallNumCall, want: kabuspb.CallPut_CALL_PUT_CALL},
		{name: "未定義 を変換できる", arg: kabus.PutOrCallNum(-1), want: kabuspb.CallPut_CALL_PUT_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromPutOrCallNum(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromUnderlyer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.Underlyer
		want string
	}{
		{name: "未指定 を変換できる", arg: kabus.UnderlyerUnspecified, want: ""},
		{name: "日経225 を変換できる", arg: kabus.UnderlyerNK225, want: "NK225"},
		{name: "日経300 を変換できる", arg: kabus.UnderlyerNK300, want: "NK300"},
		{name: "東証マザーズ を変換できる", arg: kabus.UnderlyerMOTHERS, want: "MOTHERS"},
		{name: "JPX日経400 を変換できる", arg: kabus.UnderlyerJPX400, want: "JPX400"},
		{name: "TOPIX を変換できる", arg: kabus.UnderlyerTOPIX, want: "TOPIX"},
		{name: "日経平均VI を変換できる", arg: kabus.UnderlyerNKVI, want: "NKVI"},
		{name: "NYダウ を変換できる", arg: kabus.UnderlyerDJIA, want: "DJIA"},
		{name: "東証REIT指数 を変換できる", arg: kabus.UnderlyerTSEREITINDEX, want: "TSEREITINDEX"},
		{name: "TOPIX Core30 を変換できる", arg: kabus.UnderlyerTOPIXCORE30, want: "TOPIXCORE30"},
		{name: "未定義 を変換できる", arg: kabus.Underlyer("-1"), want: ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromUnderlyer(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromPriceRangeGroup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.PriceRangeGroup
		want string
	}{
		{name: "未指定 を変換できる", arg: kabus.PriceRangeGroupUnspecified, want: ""},
		{name: "10000 を変換できる", arg: kabus.PriceRangeGroup10000, want: "10000"},
		{name: "10003 を変換できる", arg: kabus.PriceRangeGroup10003, want: "10003"},
		{name: "10118 を変換できる", arg: kabus.PriceRangeGroup10118, want: "10118"},
		{name: "10119 を変換できる", arg: kabus.PriceRangeGroup10119, want: "10119"},
		{name: "10318 を変換できる", arg: kabus.PriceRangeGroup10318, want: "10318"},
		{name: "10706 を変換できる", arg: kabus.PriceRangeGroup10706, want: "10706"},
		{name: "10718 を変換できる", arg: kabus.PriceRangeGroup10718, want: "10718"},
		{name: "12122 を変換できる", arg: kabus.PriceRangeGroup12122, want: "12122"},
		{name: "14473 を変換できる", arg: kabus.PriceRangeGroup14473, want: "14473"},
		{name: "14515 を変換できる", arg: kabus.PriceRangeGroup14515, want: "14515"},
		{name: "15411 を変換できる", arg: kabus.PriceRangeGroup15411, want: "15411"},
		{name: "15569 を変換できる", arg: kabus.PriceRangeGroup15569, want: "15569"},
		{name: "17163 を変換できる", arg: kabus.PriceRangeGroup17163, want: "17163"},
		{name: "未定義 を変換できる", arg: kabus.PriceRangeGroup("-1"), want: ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromPriceRangeGroup(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toOrderState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.OrderState
		want kabus.OrderState
	}{
		{name: "未指定 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_UNSPECIFIED, want: kabus.OrderStateUnspecified},
		{name: "待機 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_WAIT, want: kabus.OrderStateWait},
		{name: "処理中 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_PROCESSING, want: kabus.OrderStateProcessing},
		{name: "処理済 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_PROCESSED, want: kabus.OrderStateProcessed},
		{name: "訂正取消送信中 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_IN_MODIFY, want: kabus.OrderStateInCancel},
		{name: "終了 を変換できる", arg: kabuspb.OrderState_ORDER_STATE_DONE, want: kabus.OrderStateDone},
		{name: "未定義 を変換できる", arg: kabuspb.OrderState(-1), want: kabus.OrderStateUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toOrderState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toSide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.Side
		want kabus.Side
	}{
		{name: "未指定 を変換できる", arg: kabuspb.Side_SIDE_UNSPECIFIED, want: kabus.SideUnspecified},
		{name: "買い を変換できる", arg: kabuspb.Side_SIDE_BUY, want: kabus.SideBuy},
		{name: "売り を変換できる", arg: kabuspb.Side_SIDE_SELL, want: kabus.SideSell},
		{name: "未定義 を変換できる", arg: kabuspb.Side(-1), want: kabus.SideUnspecified},
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

func Test_fromOrders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.OrdersResponse
		want *kabuspb.Orders
	}{
		{name: "空配列なら空配列を返す",
			arg:  &kabus.OrdersResponse{},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{}}},
		{name: "渡す要素が1つなら要素が1つの配列が返される",
			arg: &kabus.OrdersResponse{{
				ID:              "20210331A02N36008375",
				State:           kabus.StateDone,
				OrderState:      kabus.OrderStateDone,
				OrdType:         kabus.OrdTypeInTrading,
				RecvTime:        time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local),
				Symbol:          "1475",
				SymbolName:      "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:        kabus.OrderExchangeToushou,
				ExchangeName:    "東証ETF/ETN",
				TimeInForce:     kabus.TimeInForceUnspecified,
				Price:           0,
				OrderQty:        1,
				CumQty:          1,
				Side:            kabus.SideBuy,
				CashMargin:      kabus.CashMarginUnspecified,
				AccountType:     kabus.AccountTypeSpecific,
				DelivType:       kabus.DelivTypeCash,
				ExpireDay:       kabus.YmdNUM{Time: time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)},
				MarginTradeType: kabus.MarginTradeTypeUnspecified,
				Details:         []kabus.OrderDetail{},
			}},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{{
				Id:                 "20210331A02N36008375",
				State:              kabuspb.State_STATE_DONE,
				OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
				OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
				ReceiveTime:        timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local)),
				SymbolCode:         "1475",
				SymbolName:         "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:           kabuspb.OrderExchange_ORDER_EXCHANGE_TOUSHOU,
				ExchangeName:       "東証ETF/ETN",
				TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
				Price:              0,
				OrderQuantity:      1,
				CumulativeQuantity: 1,
				Side:               kabuspb.Side_SIDE_BUY,
				TradeType:          kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED,
				AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
				ExpireDay:          timestamppb.New(time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)),
				MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
				Details:            []*kabuspb.OrderDetail{},
			}}}},
		{name: "渡す要素が2つなら要素が2つの配列が返される",
			arg: &kabus.OrdersResponse{{
				ID:              "20210331A02N36008375",
				State:           kabus.StateDone,
				OrderState:      kabus.OrderStateDone,
				OrdType:         kabus.OrdTypeInTrading,
				RecvTime:        time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local),
				Symbol:          "1475",
				SymbolName:      "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:        kabus.OrderExchangeToushou,
				ExchangeName:    "東証ETF/ETN",
				TimeInForce:     kabus.TimeInForceUnspecified,
				Price:           0,
				OrderQty:        1,
				CumQty:          1,
				Side:            kabus.SideBuy,
				CashMargin:      kabus.CashMarginUnspecified,
				AccountType:     kabus.AccountTypeSpecific,
				DelivType:       kabus.DelivTypeCash,
				ExpireDay:       kabus.YmdNUM{Time: time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)},
				MarginTradeType: kabus.MarginTradeTypeUnspecified,
				Details:         []kabus.OrderDetail{},
			}, {
				ID:              "20210331A02N36008399",
				State:           kabus.StateDone,
				OrderState:      kabus.OrderStateDone,
				OrdType:         kabus.OrdTypeInTrading,
				RecvTime:        time.Date(2021, 3, 31, 11, 28, 37, 291907000, time.Local),
				Symbol:          "1475",
				SymbolName:      "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:        kabus.OrderExchangeSOR,
				ExchangeName:    "SOR",
				TimeInForce:     kabus.TimeInForceUnspecified,
				Price:           0,
				OrderQty:        1,
				CumQty:          1,
				Side:            kabus.SideSell,
				CashMargin:      kabus.CashMarginUnspecified,
				AccountType:     kabus.AccountTypeSpecific,
				DelivType:       kabus.DelivTypeUnspecified,
				ExpireDay:       kabus.YmdNUM{Time: time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)},
				MarginTradeType: kabus.MarginTradeTypeUnspecified,
				Details:         []kabus.OrderDetail{},
			}},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{{
				Id:                 "20210331A02N36008375",
				State:              kabuspb.State_STATE_DONE,
				OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
				OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
				ReceiveTime:        timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local)),
				SymbolCode:         "1475",
				SymbolName:         "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:           kabuspb.OrderExchange_ORDER_EXCHANGE_TOUSHOU,
				ExchangeName:       "東証ETF/ETN",
				TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
				Price:              0,
				OrderQuantity:      1,
				CumulativeQuantity: 1,
				Side:               kabuspb.Side_SIDE_BUY,
				TradeType:          kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED,
				AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
				ExpireDay:          timestamppb.New(time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)),
				MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
				Details:            []*kabuspb.OrderDetail{},
			}, {
				Id:                 "20210331A02N36008399",
				State:              kabuspb.State_STATE_DONE,
				OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
				OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
				ReceiveTime:        timestamppb.New(time.Date(2021, 3, 31, 11, 28, 37, 291907000, time.Local)),
				SymbolCode:         "1475",
				SymbolName:         "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:           kabuspb.OrderExchange_ORDER_EXCHANGE_SOR,
				ExchangeName:       "SOR",
				TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
				Price:              0,
				OrderQuantity:      1,
				CumulativeQuantity: 1,
				Side:               kabuspb.Side_SIDE_SELL,
				TradeType:          kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED,
				AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED,
				ExpireDay:          timestamppb.New(time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)),
				MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
				Details:            []*kabuspb.OrderDetail{},
			}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrders(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderDetails(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []kabus.OrderDetail
		want []*kabuspb.OrderDetail
	}{
		{name: "空配列が渡されたら空配列が返される",
			arg:  []kabus.OrderDetail{},
			want: []*kabuspb.OrderDetail{}},
		{name: "渡す要素が1つなら要素が1つの配列が返される",
			arg: []kabus.OrderDetail{
				{
					SeqNum:        1,
					ID:            "20200715A02N04738436",
					RecType:       kabus.RecTypeReceived,
					ExchangeID:    "00000000-0000-0000-0000-00000000",
					State:         kabus.OrderDetailStateProcessed,
					TransactTime:  time.Date(2020, 7, 16, 18, 0, 51, 763683000, time.Local),
					OrdType:       kabus.OrdTypeInTrading,
					Price:         704.5,
					Qty:           1500,
					ExecutionID:   "",
					ExecutionDay:  time.Date(2020, 7, 2, 18, 2, 0, 0, time.Local),
					DelivDay:      kabus.NewYmdNUM(time.Date(2020, 7, 6, 0, 0, 0, 0, time.Local)),
					Commission:    0,
					CommissionTax: 0,
				},
			},
			want: []*kabuspb.OrderDetail{
				{
					SequenceNumber: 1,
					Id:             "20200715A02N04738436",
					RecordType:     kabuspb.RecordType_RECORD_TYPE_RECEIVE,
					ExchangeId:     "00000000-0000-0000-0000-00000000",
					State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
					TransactTime:   timestamppb.New(time.Date(2020, 7, 16, 18, 0, 51, 763683000, time.Local)),
					OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA,
					Price:          704.5,
					Quantity:       1500,
					ExecutionId:    "",
					ExecutionDay:   timestamppb.New(time.Date(2020, 7, 2, 18, 2, 0, 0, time.Local)),
					DeliveryDay:    timestamppb.New(time.Date(2020, 7, 6, 0, 0, 0, 0, time.Local)),
					Commission:     0,
					CommissionTax:  0,
				},
			}},
		{name: "渡す要素が3つなら要素が3つの配列が返される",
			arg: []kabus.OrderDetail{
				{
					SeqNum:        1,
					ID:            "20210331A02N36008375",
					RecType:       kabus.RecTypeReceived,
					ExchangeID:    "",
					State:         kabus.OrderDetailStateProcessed,
					TransactTime:  time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local),
					OrdType:       kabus.OrdTypeInTrading,
					Price:         0,
					Qty:           1,
					ExecutionID:   "",
					ExecutionDay:  time.Time{},
					DelivDay:      kabus.NewYmdNUM(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
					Commission:    0,
					CommissionTax: 0,
				}, {
					SeqNum:        4,
					ID:            "20210331B02N36008376",
					RecType:       kabus.RecTypeOrdered,
					ExchangeID:    "1F111300012175",
					State:         kabus.OrderDetailStateProcessed,
					TransactTime:  time.Date(2021, 3, 31, 11, 28, 19, 53576000, time.Local),
					OrdType:       kabus.OrdTypeInTrading,
					Price:         0,
					Qty:           1,
					ExecutionID:   "",
					ExecutionDay:  time.Time{},
					DelivDay:      kabus.NewYmdNUM(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
					Commission:    0,
					CommissionTax: 0,
				}, {
					SeqNum:        5,
					ID:            "20210331E02N36008377",
					RecType:       kabus.RecTypeContracted,
					ExchangeID:    "416",
					State:         kabus.OrderDetailStateProcessed,
					TransactTime:  time.Date(2021, 3, 31, 11, 28, 19, 535867000, time.Local),
					OrdType:       kabus.OrdTypeUnspecified,
					Price:         2018,
					Qty:           1,
					ExecutionID:   "E20210331022VE",
					ExecutionDay:  time.Date(2021, 3, 31, 11, 28, 19, 535867000, time.Local),
					DelivDay:      kabus.NewYmdNUM(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
					Commission:    0,
					CommissionTax: 0,
				},
			},
			want: []*kabuspb.OrderDetail{{
				SequenceNumber: 1,
				Id:             "20210331A02N36008375",
				RecordType:     kabuspb.RecordType_RECORD_TYPE_RECEIVE,
				ExchangeId:     "",
				State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
				TransactTime:   timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local)),
				OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA,
				Price:          0,
				Quantity:       1,
				ExecutionId:    "",
				ExecutionDay:   timestamppb.New(time.Time{}),
				DeliveryDay:    timestamppb.New(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
				Commission:     0,
				CommissionTax:  0,
			}, {
				SequenceNumber: 4,
				Id:             "20210331B02N36008376",
				RecordType:     kabuspb.RecordType_RECORD_TYPE_ORDERED,
				ExchangeId:     "1F111300012175",
				State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
				TransactTime:   timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 53576000, time.Local)),
				OrderType:      kabuspb.OrderType_ORDER_TYPE_ZARABA,
				Price:          0,
				Quantity:       1,
				ExecutionId:    "",
				ExecutionDay:   timestamppb.New(time.Time{}),
				DeliveryDay:    timestamppb.New(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
				Commission:     0,
				CommissionTax:  0,
			}, {
				SequenceNumber: 5,
				Id:             "20210331E02N36008377",
				RecordType:     kabuspb.RecordType_RECORD_TYPE_CONTRACTED,
				ExchangeId:     "416",
				State:          kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED,
				TransactTime:   timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 535867000, time.Local)),
				OrderType:      kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED,
				Price:          2018,
				Quantity:       1,
				ExecutionId:    "E20210331022VE",
				ExecutionDay:   timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 535867000, time.Local)),
				DeliveryDay:    timestamppb.New(time.Date(2021, 4, 2, 0, 0, 0, 0, time.Local)),
				Commission:     0,
				CommissionTax:  0,
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderDetails(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.State
		want kabuspb.State
	}{
		{name: "未指定 を変換できる", arg: kabus.StateUnspecified, want: kabuspb.State_STATE_UNSPECIFIED},
		{name: "待機 を変換できる", arg: kabus.StateWait, want: kabuspb.State_STATE_WAIT},
		{name: "処理中 を変換できる", arg: kabus.StateProcessing, want: kabuspb.State_STATE_PROCESSING},
		{name: "処理済 を変換できる", arg: kabus.StateProcessed, want: kabuspb.State_STATE_PROCESSED},
		{name: "訂正取消送信中 を変換できる", arg: kabus.StateInCancel, want: kabuspb.State_STATE_IN_MODIFY},
		{name: "終了 を変換できる", arg: kabus.StateDone, want: kabuspb.State_STATE_DONE},
		{name: "未定義 を変換できる", arg: kabus.State(-1), want: kabuspb.State_STATE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.OrderState
		want kabuspb.OrderState
	}{
		{name: "未指定 を変換できる", arg: kabus.OrderStateUnspecified, want: kabuspb.OrderState_ORDER_STATE_UNSPECIFIED},
		{name: "待機 を変換できる", arg: kabus.OrderStateWait, want: kabuspb.OrderState_ORDER_STATE_WAIT},
		{name: "処理中 を変換できる", arg: kabus.OrderStateProcessing, want: kabuspb.OrderState_ORDER_STATE_PROCESSING},
		{name: "処理済 を変換できる", arg: kabus.OrderStateProcessed, want: kabuspb.OrderState_ORDER_STATE_PROCESSED},
		{name: "訂正取消送信中 を変換できる", arg: kabus.OrderStateInCancel, want: kabuspb.OrderState_ORDER_STATE_IN_MODIFY},
		{name: "終了 を変換できる", arg: kabus.OrderStateDone, want: kabuspb.OrderState_ORDER_STATE_DONE},
		{name: "未定義 を変換できる", arg: kabus.OrderState(-1), want: kabuspb.OrderState_ORDER_STATE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrdType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.OrdType
		want kabuspb.OrderType
	}{
		{name: "未指定 を変換できる", arg: kabus.OrdTypeUnspecified, want: kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED},
		{name: "ザラバ を変換できる", arg: kabus.OrdTypeInTrading, want: kabuspb.OrderType_ORDER_TYPE_ZARABA},
		{name: "寄り を変換できる", arg: kabus.OrdTypeOpen, want: kabuspb.OrderType_ORDER_TYPE_OPEN},
		{name: "引け を変換できる", arg: kabus.OrdTypeClose, want: kabuspb.OrderType_ORDER_TYPE_CLOSE},
		{name: "不成 を変換できる", arg: kabus.OrdTypeNoContracted, want: kabuspb.OrderType_ORDER_TYPE_FUNARI},
		{name: "対当指値 を変換できる", arg: kabus.OrdTypeMarketToLimit, want: kabuspb.OrderType_ORDER_TYPE_MTLO},
		{name: "IOC を変換できる", arg: kabus.OrdTypeIOC, want: kabuspb.OrderType_ORDER_TYPE_IOC},
		{name: "未定義 を変換できる", arg: kabus.OrdType(-1), want: kabuspb.OrderType_ORDER_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrdType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.OrderExchange
		want kabuspb.OrderExchange
	}{
		{name: "未指定 を変換できる", arg: kabus.OrderExchangeUnspecified, want: kabuspb.OrderExchange_ORDER_EXCHANGE_UNSPECIFIED},
		{name: "東証 を変換できる", arg: kabus.OrderExchangeToushou, want: kabuspb.OrderExchange_ORDER_EXCHANGE_TOUSHOU},
		{name: "名証 を変換できる", arg: kabus.OrderExchangeMeishou, want: kabuspb.OrderExchange_ORDER_EXCHANGE_MEISHOU},
		{name: "福証 を変換できる", arg: kabus.OrderExchangeFukushou, want: kabuspb.OrderExchange_ORDER_EXCHANGE_FUKUSHOU},
		{name: "札証 を変換できる", arg: kabus.OrderExchangeSatsushou, want: kabuspb.OrderExchange_ORDER_EXCHANGE_SATSUSHOU},
		{name: "SOR を変換できる", arg: kabus.OrderExchangeSOR, want: kabuspb.OrderExchange_ORDER_EXCHANGE_SOR},
		{name: "日通し を変換できる", arg: kabus.OrderExchangeAll, want: kabuspb.OrderExchange_ORDER_EXCHANGE_ALL_SESSION},
		{name: "日中 を変換できる", arg: kabus.OrderExchangeDaytime, want: kabuspb.OrderExchange_ORDER_EXCHANGE_DAY_SESSION},
		{name: "夜間 を変換できる", arg: kabus.OrderExchangeEvening, want: kabuspb.OrderExchange_ORDER_EXCHANGE_NIGHT_SESSION},
		{name: "未定義 を変換できる", arg: kabus.OrderExchange(-1), want: kabuspb.OrderExchange_ORDER_EXCHANGE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderExchange(test.arg)
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
		arg  kabus.Side
		want kabuspb.Side
	}{
		{name: "未指定 を変換できる", arg: kabus.SideUnspecified, want: kabuspb.Side_SIDE_UNSPECIFIED},
		{name: "売 を変換できる", arg: kabus.SideSell, want: kabuspb.Side_SIDE_SELL},
		{name: "買 を変換できる", arg: kabus.SideBuy, want: kabuspb.Side_SIDE_BUY},
		{name: "未定義 を変換できる", arg: kabus.Side("-1"), want: kabuspb.Side_SIDE_UNSPECIFIED},
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

func Test_fromAccountType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.AccountType
		want kabuspb.AccountType
	}{
		{name: "未指定 を変換できる", arg: kabus.AccountTypeUnspecified, want: kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED},
		{name: "一般 を変換できる", arg: kabus.AccountTypeGeneral, want: kabuspb.AccountType_ACCOUNT_TYPE_GENERAL},
		{name: "特定 を変換できる", arg: kabus.AccountTypeSpecific, want: kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC},
		{name: "法人 を変換できる", arg: kabus.AccountTypeCorporation, want: kabuspb.AccountType_ACCOUNT_TYPE_CORPORATION},
		{name: "未定義 を変換できる", arg: kabus.AccountType(-1), want: kabuspb.AccountType_ACCOUNT_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromAccountType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromDelivType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.DelivType
		want kabuspb.DeliveryType
	}{
		{name: "未指定 を変換できる", arg: kabus.DelivTypeUnspecified, want: kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED},
		{name: "自動振替 を変換できる", arg: kabus.DelivTypeAuto, want: kabuspb.DeliveryType_DELIVERY_TYPE_AUTO},
		{name: "お預り金 を変換できる", arg: kabus.DelivTypeCash, want: kabuspb.DeliveryType_DELIVERY_TYPE_CASH},
		{name: "未定義 を変換できる", arg: kabus.DelivType(-1), want: kabuspb.DeliveryType_DELIVERY_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromDelivType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRecType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.RecType
		want kabuspb.RecordType
	}{
		{name: "未指定 を変換できる", arg: kabus.RecTypeUnspecified, want: kabuspb.RecordType_RECORD_TYPE_UNSPECIFIED},
		{name: "受付 を変換できる", arg: kabus.RecTypeReceived, want: kabuspb.RecordType_RECORD_TYPE_RECEIVE},
		{name: "繰越 を変換できる", arg: kabus.RecTypeCarried, want: kabuspb.RecordType_RECORD_TYPE_CARRIED},
		{name: "期限切れ を変換できる", arg: kabus.RecTypeExpired, want: kabuspb.RecordType_RECORD_TYPE_EXPIRED},
		{name: "発注 を変換できる", arg: kabus.RecTypeOrdered, want: kabuspb.RecordType_RECORD_TYPE_ORDERED},
		{name: "訂正 を変換できる", arg: kabus.RecTypeModified, want: kabuspb.RecordType_RECORD_TYPE_MODIFIED},
		{name: "取消 を変換できる", arg: kabus.RecTypeCanceled, want: kabuspb.RecordType_RECORD_TYPE_CANCELED},
		{name: "失効 を変換できる", arg: kabus.RecTypeRevocation, want: kabuspb.RecordType_RECORD_TYPE_REVOCATION},
		{name: "約定 を変換できる", arg: kabus.RecTypeContracted, want: kabuspb.RecordType_RECORD_TYPE_CONTRACTED},
		{name: "未定義 を変換できる", arg: kabus.RecType(-1), want: kabuspb.RecordType_RECORD_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRecType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromOrderDetailState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.OrderDetailState
		want kabuspb.OrderDetailState
	}{
		{name: "未指定 を変換できる", arg: kabus.OrderDetailStateUnspecified, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_UNSPECIFIED},
		{name: "待機 を変換できる", arg: kabus.OrderDetailStateWait, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_WAIT},
		{name: "処理中 を変換できる", arg: kabus.OrderDetailStateProcessing, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSING},
		{name: "処理済 を変換できる", arg: kabus.OrderDetailStateProcessed, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_PROCESSED},
		{name: "エラー を変換できる", arg: kabus.OrderDetailStateError, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_ERROR},
		{name: "削除済み を変換できる", arg: kabus.OrderDetailStateDeleted, want: kabuspb.OrderDetailState_ORDER_DETAIL_STATE_DELETED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromOrderDetailState(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toCashMargin(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.TradeType
		want kabus.CashMargin
	}{
		{name: "未指定 を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED, want: kabus.CashMarginUnspecified},
		{name: "新規 を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_ENTRY, want: kabus.CashMarginMarginEntry},
		{name: "返済 を変換できる", arg: kabuspb.TradeType_TRADE_TYPE_EXIT, want: kabus.CashMarginMarginExit},
		{name: "未定義 を変換できる", arg: kabuspb.TradeType(-1), want: kabus.CashMarginUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toCashMargin(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromCashMargin(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.CashMargin
		want kabuspb.TradeType
	}{
		{name: "未指定 を変換できる", arg: kabus.CashMarginUnspecified, want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
		{name: "現物 を変換できる", arg: kabus.CashMarginCash, want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
		{name: "信用新規 を変換できる", arg: kabus.CashMarginMarginEntry, want: kabuspb.TradeType_TRADE_TYPE_ENTRY},
		{name: "信用返済 を変換できる", arg: kabus.CashMarginMarginExit, want: kabuspb.TradeType_TRADE_TYPE_EXIT},
		{name: "未定義 を変換できる", arg: kabus.CashMargin(-1), want: kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromCashMargin(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromMarginTradeType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.MarginTradeType
		want kabuspb.MarginTradeType
	}{
		{name: "未指定 を変換できる", arg: kabus.MarginTradeTypeUnspecified, want: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED},
		{name: "制度信用 を変換できる", arg: kabus.MarginTradeTypeSystem, want: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM},
		{name: "一般信用 を変換できる", arg: kabus.MarginTradeTypeGeneralLong, want: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_LONG},
		{name: "一般信用(売短) を変換できる", arg: kabus.MarginTradeTypeGeneralShort, want: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_GENERAL_SHORT},
		{name: "未定義 を変換できる", arg: kabus.MarginTradeType(-1), want: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromMarginTradeType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromTimeInForce(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.TimeInForce
		want kabuspb.TimeInForce
	}{
		{name: "未指定 を変換できる", arg: kabus.TimeInForceUnspecified, want: kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED},
		{name: "FAS を変換できる", arg: kabus.TimeInForceFAS, want: kabuspb.TimeInForce_TIME_IN_FORCE_FAS},
		{name: "FAK を変換できる", arg: kabus.TimeInForceFAK, want: kabuspb.TimeInForce_TIME_IN_FORCE_FAK},
		{name: "FOK を変換できる", arg: kabus.TimeInForceFOK, want: kabuspb.TimeInForce_TIME_IN_FORCE_FOK},
		{name: "未定義 を変換できる", arg: kabus.TimeInForce(-1), want: kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromTimeInForce(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toProduct(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.Product
		want kabus.Product
	}{
		{name: "未指定 を変換できる", arg: kabuspb.Product_PRODUCT_UNSPECIFIED, want: kabus.ProductAll},
		{name: "すべて を変換できる", arg: kabuspb.Product_PRODUCT_ALL, want: kabus.ProductAll},
		{name: "現物 を変換できる", arg: kabuspb.Product_PRODUCT_STOCK, want: kabus.ProductCash},
		{name: "信用 を変換できる", arg: kabuspb.Product_PRODUCT_MARGIN, want: kabus.ProductMargin},
		{name: "先物 を変換できる", arg: kabuspb.Product_PRODUCT_FUTURE, want: kabus.ProductFuture},
		{name: "オプション を変換できる", arg: kabuspb.Product_PRODUCT_OPTION, want: kabus.ProductOption},
		{name: "未定義 を変換できる", arg: kabuspb.Product(-1), want: kabus.ProductAll},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toProduct(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toIsGetOrderDetail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  bool
		want kabus.IsGetOrderDetail
	}{
		{name: "true を変換できる", arg: true, want: kabus.IsGetOrderDetailTrue},
		{name: "false を変換できる", arg: false, want: kabus.IsGetOrderDetailFalse},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toIsGetOrderDetail(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromSecurityType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.SecurityType
		want kabuspb.SecurityType
	}{
		{name: "未指定 を変換できる", arg: kabus.SecurityTypeUnspecified, want: kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED},
		{name: "株式 を変換できる", arg: kabus.SecurityTypeStock, want: kabuspb.SecurityType_SECURITY_TYPE_STOCK},
		{name: "日経225先物 を変換できる", arg: kabus.SecurityTypeNK225, want: kabuspb.SecurityType_SECURITY_TYPE_NK225},
		{name: "日経225mini先物 を変換できる", arg: kabus.SecurityTypeNK225Mini, want: kabuspb.SecurityType_SECURITY_TYPE_NK225_MINI},
		{name: "JPX日経インデックス400先物 を変換できる", arg: kabus.SecurityTypeJPX400, want: kabuspb.SecurityType_SECURITY_TYPE_JPX400},
		{name: "TOPIX先物 を変換できる", arg: kabus.SecurityTypeTOPIX, want: kabuspb.SecurityType_SECURITY_TYPE_TOPIX},
		{name: "ミニTOPIX先物 を変換できる", arg: kabus.SecurityTypeTOPIXMini, want: kabuspb.SecurityType_SECURITY_TYPE_TOPIX_MINI},
		{name: "東証マザーズ指数先物 を変換できる", arg: kabus.SecurityTypeMothers, want: kabuspb.SecurityType_SECURITY_TYPE_MOTHERS},
		{name: "東証REIT指数先物 を変換できる", arg: kabus.SecurityTypeREIT, want: kabuspb.SecurityType_SECURITY_TYPE_REIT},
		{name: "NYダウ先物 を変換できる", arg: kabus.SecurityTypeDOW, want: kabuspb.SecurityType_SECURITY_TYPE_DOW},
		{name: "日経平均VI先物 を変換できる", arg: kabus.SecurityTypeVI, want: kabuspb.SecurityType_SECURITY_TYPE_VI},
		{name: "TOPIX Core30先物 を変換できる", arg: kabus.SecurityTypeCORE30, want: kabuspb.SecurityType_SECURITY_TYPE_CODE30},
		{name: "未定義 を変換できる", arg: kabus.SecurityType(-1), want: kabuspb.SecurityType_SECURITY_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromSecurityType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromPositions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.PositionsResponse
		want *kabuspb.Positions
	}{
		{name: "空配列が渡されたら空配列を返す", arg: &kabus.PositionsResponse{}, want: &kabuspb.Positions{Positions: []*kabuspb.Position{}}},
		{name: "要素が1つの配列が渡されたら要素が1つの配列を返す",
			arg: &kabus.PositionsResponse{{
				ExecutionID:     "20200715E02N04738464",
				AccountType:     kabus.AccountTypeSpecific,
				Symbol:          "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabus.ExchangeToushou,
				ExchangeName:    "東証１部",
				SecurityType:    kabus.SecurityTypeNK225,
				ExecutionDay:    kabus.NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQty:       500,
				HoldQty:         0,
				Side:            kabus.SideSell,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       kabus.NewYmdNUM(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabus.MarginTradeTypeSystem,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}},
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{{
				ExecutionId:     "20200715E02N04738464",
				AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SymbolCode:      "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:    "東証１部",
				SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_NK225,
				ExecutionDay:    timestamppb.New(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQuantity:  500,
				HoldQuantity:    0,
				Side:            kabuspb.Side_SIDE_SELL,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       timestamppb.New(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}}}},
		{name: "要素が2つの配列が渡されたら要素が2つの配列を返す",
			arg: &kabus.PositionsResponse{{
				ExecutionID:     "20200715E02N04738464",
				AccountType:     kabus.AccountTypeSpecific,
				Symbol:          "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabus.ExchangeToushou,
				ExchangeName:    "東証１部",
				SecurityType:    kabus.SecurityTypeNK225,
				ExecutionDay:    kabus.NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQty:       500,
				HoldQty:         0,
				Side:            kabus.SideSell,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       kabus.NewYmdNUM(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabus.MarginTradeTypeSystem,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}, {
				ExecutionID:     "20200715E02N04738464",
				AccountType:     kabus.AccountTypeSpecific,
				Symbol:          "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabus.ExchangeToushou,
				ExchangeName:    "東証１部",
				SecurityType:    kabus.SecurityTypeNK225,
				ExecutionDay:    kabus.NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQty:       500,
				HoldQty:         0,
				Side:            kabus.SideSell,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       kabus.NewYmdNUM(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabus.MarginTradeTypeSystem,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}},
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{{
				ExecutionId:     "20200715E02N04738464",
				AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SymbolCode:      "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:    "東証１部",
				SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_NK225,
				ExecutionDay:    timestamppb.New(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQuantity:  500,
				HoldQuantity:    0,
				Side:            kabuspb.Side_SIDE_SELL,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       timestamppb.New(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}, {
				ExecutionId:     "20200715E02N04738464",
				AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SymbolCode:      "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:    "東証１部",
				SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_NK225,
				ExecutionDay:    timestamppb.New(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQuantity:  500,
				HoldQuantity:    0,
				Side:            kabuspb.Side_SIDE_SELL,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       timestamppb.New(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromPositions(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toRankingTypeFromPriceRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.PriceRankingType
		want kabus.RankingType
	}{
		{name: "未指定 を変換できる", arg: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED, want: kabus.RankingTypeUnspecified},
		{name: "値上がり率 を変換できる", arg: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE, want: kabus.RankingTypePriceIncreaseRate},
		{name: "値下がり率 を変換できる", arg: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_DECREASE_RATE, want: kabus.RankingTypePriceDecreaseRate},
		{name: "売買高上位 を変換できる", arg: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VOLUME, want: kabus.RankingTypeVolume},
		{name: "売買代金上位 を変換できる", arg: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VALUE, want: kabus.RankingTypeValue},
		{name: "未定義 を変換できる", arg: kabuspb.PriceRankingType(-1), want: kabus.RankingTypeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toRankingTypeFromPriceRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingTypeToPriceRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.RankingType
		want kabuspb.PriceRankingType
	}{
		{name: "未指定 を変換できる", arg: kabus.RankingTypeUnspecified, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRate, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE},
		{name: "値下がり率 を変換できる", arg: kabus.RankingTypePriceDecreaseRate, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_DECREASE_RATE},
		{name: "売買高上位 を変換できる", arg: kabus.RankingTypeVolume, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VOLUME},
		{name: "売買代金 を変換できる", arg: kabus.RankingTypeValue, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_VALUE},
		{name: "TICK回数 を変換できる", arg: kabus.RankingTypeTickCount, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "売買高急増 を変換できる", arg: kabus.RankingTypeVolumeRapidIncrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "売買代金急増 を変換できる", arg: kabus.RankingTypeValueRapidIncrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用売残増 を変換できる", arg: kabus.RankingTypeMarginSellBalanceIncrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用売残減 を変換できる", arg: kabus.RankingTypeMarginSellBalanceDecrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用買残増 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceIncrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用買残減 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceDecrease, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用高倍率 を変換できる", arg: kabus.RankingTypeMarginHighMagnification, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "信用低倍率 を変換できる", arg: kabus.RankingTypeMarginLowMagnification, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "業種別値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRateByCategory, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "業種別値下がり率 を変換できる", arg: kabus.RankingTypePriceDecreaseRateByCategory, want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
		{name: "未定義 を変換できる", arg: kabus.RankingType("-1"), want: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingTypeToPriceRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toRankingTypeFromMarginRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.MarginRankingType
		want kabus.RankingType
	}{
		{name: "未指定 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED, want: kabus.RankingTypeUnspecified},
		{name: "信用売残増 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_INCREASE, want: kabus.RankingTypeMarginSellBalanceIncrease},
		{name: "信用売残減 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_DECREASE, want: kabus.RankingTypeMarginSellBalanceDecrease},
		{name: "信用買残増 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_INCREASE, want: kabus.RankingTypeMarginBuyBalanceIncrease},
		{name: "信用買残減 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_DECREASE, want: kabus.RankingTypeMarginBuyBalanceDecrease},
		{name: "信用高倍率 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION, want: kabus.RankingTypeMarginHighMagnification},
		{name: "信用低倍率 を変換できる", arg: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_LOW_MAGNIFICATION, want: kabus.RankingTypeMarginLowMagnification},
		{name: "未定義 を変換できる", arg: kabuspb.MarginRankingType(-1), want: kabus.RankingTypeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toRankingTypeFromMarginRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingTypeToMarginRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.RankingType
		want kabuspb.MarginRankingType
	}{
		{name: "未指定 を変換できる", arg: kabus.RankingTypeUnspecified, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRate, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "売買高上位 を変換できる", arg: kabus.RankingTypeVolume, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "売買代金 を変換できる", arg: kabus.RankingTypeValue, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "TICK回数 を変換できる", arg: kabus.RankingTypeTickCount, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "売買高急増 を変換できる", arg: kabus.RankingTypeVolumeRapidIncrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "売買代金急増 を変換できる", arg: kabus.RankingTypeValueRapidIncrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "信用売残増 を変換できる", arg: kabus.RankingTypeMarginSellBalanceIncrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_INCREASE},
		{name: "信用売残減 を変換できる", arg: kabus.RankingTypeMarginSellBalanceDecrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_SELL_BALANCE_DECREASE},
		{name: "信用買残増 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceIncrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_INCREASE},
		{name: "信用買残減 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceDecrease, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_BUY_BALANCE_DECREASE},
		{name: "信用高倍率 を変換できる", arg: kabus.RankingTypeMarginHighMagnification, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION},
		{name: "信用低倍率 を変換できる", arg: kabus.RankingTypeMarginLowMagnification, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_LOW_MAGNIFICATION},
		{name: "業種別値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRateByCategory, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "業種別値下がり率 を変換できる", arg: kabus.RankingTypePriceDecreaseRateByCategory, want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
		{name: "未定義 を変換できる", arg: kabus.RankingType("-1"), want: kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingTypeToMarginRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toRankingTypeFromIndustryRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.IndustryRankingType
		want kabus.RankingType
	}{
		{name: "未指定 を変換できる", arg: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED, want: kabus.RankingTypeUnspecified},
		{name: "値上がり率 を変換できる", arg: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE, want: kabus.RankingTypePriceIncreaseRateByCategory},
		{name: "値下がり率 を変換できる", arg: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_DECREASE_RATE, want: kabus.RankingTypePriceDecreaseRateByCategory},
		{name: "未定義 を変換できる", arg: kabuspb.IndustryRankingType(-1), want: kabus.RankingTypeUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toRankingTypeFromIndustryRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingTypeToIndustryRankingType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.RankingType
		want kabuspb.IndustryRankingType
	}{
		{name: "未指定 を変換できる", arg: kabus.RankingTypeUnspecified, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRate, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "売買高上位 を変換できる", arg: kabus.RankingTypeVolume, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "売買代金 を変換できる", arg: kabus.RankingTypeValue, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "TICK回数 を変換できる", arg: kabus.RankingTypeTickCount, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "売買高急増 を変換できる", arg: kabus.RankingTypeVolumeRapidIncrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "売買代金急増 を変換できる", arg: kabus.RankingTypeValueRapidIncrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用売残増 を変換できる", arg: kabus.RankingTypeMarginSellBalanceIncrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用売残減 を変換できる", arg: kabus.RankingTypeMarginSellBalanceDecrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用買残増 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceIncrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用買残減 を変換できる", arg: kabus.RankingTypeMarginBuyBalanceDecrease, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用高倍率 を変換できる", arg: kabus.RankingTypeMarginHighMagnification, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "信用低倍率 を変換できる", arg: kabus.RankingTypeMarginLowMagnification, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
		{name: "業種別値上がり率 を変換できる", arg: kabus.RankingTypePriceIncreaseRateByCategory, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE},
		{name: "業種別値下がり率 を変換できる", arg: kabus.RankingTypePriceDecreaseRateByCategory, want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_DECREASE_RATE},
		{name: "未定義 を変換できる", arg: kabus.RankingType("-1"), want: kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingTypeToIndustryRankingType(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toExchangeDivision(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabuspb.ExchangeDivision
		want kabus.ExchangeDivision
	}{
		{name: "未指定 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_UNSPECIFIED, want: kabus.ExchangeDivisionUnspecified},
		{name: "全市場 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL, want: kabus.ExchangeDivisionALL},
		{name: "東証全体 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_ALL, want: kabus.ExchangeDivisionToushou},
		{name: "東証一部 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_1, want: kabus.ExchangeDivisionToushou1},
		{name: "東証二部 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_2, want: kabus.ExchangeDivisionToushou2},
		{name: "東証マザーズ を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MOTHERS, want: kabus.ExchangeDivisionMothers},
		{name: "JASDAQ を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_JASDAQ, want: kabus.ExchangeDivisionJASDAQ},
		{name: "名証 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MEISHOU, want: kabus.ExchangeDivisionMeishou},
		{name: "福証 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_FUKUSHOU, want: kabus.ExchangeDivisionFukushou},
		{name: "札証 を変換できる", arg: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_SATSUSHOU, want: kabus.ExchangeDivisionSatsushou},
		{name: "未定義 を変換できる", arg: kabuspb.ExchangeDivision(-1), want: kabus.ExchangeDivisionUnspecified},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := toExchangeDivision(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromExchangeDivision(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.ExchangeDivision
		want kabuspb.ExchangeDivision
	}{
		{name: "未指定 を変換できる", arg: kabus.ExchangeDivisionUnspecified, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_UNSPECIFIED},
		{name: "全市場 を変換できる", arg: kabus.ExchangeDivisionALL, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
		{name: "東証全体 を変換できる", arg: kabus.ExchangeDivisionToushou, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_ALL},
		{name: "東証一部 を変換できる", arg: kabus.ExchangeDivisionToushou1, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_1},
		{name: "東証二部 を変換できる", arg: kabus.ExchangeDivisionToushou2, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_TOUSHOU_2},
		{name: "東証マザーズ を変換できる", arg: kabus.ExchangeDivisionMothers, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MOTHERS},
		{name: "JASDAQ を変換できる", arg: kabus.ExchangeDivisionJASDAQ, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_JASDAQ},
		{name: "名証 を変換できる", arg: kabus.ExchangeDivisionMeishou, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_MEISHOU},
		{name: "福証 を変換できる", arg: kabus.ExchangeDivisionFukushou, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_FUKUSHOU},
		{name: "札証 を変換できる", arg: kabus.ExchangeDivisionSatsushou, want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_SATSUSHOU},
		{name: "未定義 を変換できる", arg: kabus.ExchangeDivision("-1"), want: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromExchangeDivision(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromTrend(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  kabus.RankingTrend
		want kabuspb.RankingTrend
	}{
		{name: "未指定 を変換できる", arg: kabus.RankingTrendUnspecified, want: kabuspb.RankingTrend_RANKING_TREND_UNSPECIFIED},
		{name: "対象データ無し を変換できる", arg: kabus.RankingTrendNoData, want: kabuspb.RankingTrend_RANKING_TREND_NO_DATA},
		{name: "過去10営業日より20位以上上昇 を変換できる", arg: kabus.RankingTrendRiseOver20, want: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20},
		{name: "過去10営業日より1～19位上昇 を変換できる", arg: kabus.RankingTrendRise, want: kabuspb.RankingTrend_RANKING_TREND_RISE},
		{name: "過去10営業日と変わらず を変換できる", arg: kabus.RankingTrendUnchanged, want: kabuspb.RankingTrend_RANKING_TREND_NO_CHANGE},
		{name: "過去10営業日より1～19位下落 を変換できる", arg: kabus.RankingTrendDescent, want: kabuspb.RankingTrend_RANKING_TREND_DESCENT},
		{name: "過去10営業日より20位以上下落 を変換できる", arg: kabus.RankingTrendDescentOver20, want: kabuspb.RankingTrend_RANKING_TREND_DESCENT_OVER_20},
		{name: "未定義 を変換できる", arg: kabus.RankingTrend("-1"), want: kabuspb.RankingTrend_RANKING_TREND_UNSPECIFIED},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromTrend(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToPriceRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.PriceRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypePriceIncreaseRate, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.PriceRanking{
				Type:             kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.PriceRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRate,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.PriceRanking{
				Type:             kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.PriceRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRate,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking: []kabus.PriceRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1689", SymbolName: "ガスETF/ETF(C)", CurrentPrice: 2, ChangeRatio: 1, ChangePercentage: 100, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 5722.4, Turnover: 10.4136, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, ChangePercentage: 54.65, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 3117.5, Turnover: 3194.7121, ExchangeName: "東証JQS", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.PriceRanking{
				Type:             kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.PriceRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1689", SymbolName: "ガスETF/ETF(C)", CurrentPrice: 2, ChangeRatio: 1, ChangePercentage: 100, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), TradingVolume: 5722.4, Turnover: 10.4136, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, ChangePercentage: 54.65, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), TradingVolume: 3117.5, Turnover: 3194.7121, ExchangeName: "東証JQS", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToPriceRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToTickRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.TickRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypeTickCount, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.TickRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.TickRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeTickCount,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.TickRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.TickRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeTickCount,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				TickRanking: []kabus.TickRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 22, Symbol: "2929", SymbolName: "ﾌｧｰﾏﾌｰｽﾞ", CurrentPrice: 2748, ChangeRatio: 99, TickCount: 40579, UpCount: 12722, DownCount: 12798, ChangePercentage: 3.73, TradingVolume: 16086.8, Turnover: 43810.0498, ExchangeName: "東証２部", CategoryName: "食料品"},
					{No: 2, Trend: kabus.RankingTrendUnchanged, AverageRanking: 2, Symbol: "9984", SymbolName: "ｿﾌﾄﾊﾞﾝｸG", CurrentPrice: 8285, ChangeRatio: -309, TickCount: 32219, UpCount: 8655, DownCount: 8562, ChangePercentage: -3.59, TradingVolume: 16688.8, Turnover: 138143.1773, ExchangeName: "東証１部", CategoryName: "情報・通信業"},
				},
			},
			want: &kabuspb.TickRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.TickRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 22, SymbolCode: "2929", SymbolName: "ﾌｧｰﾏﾌｰｽﾞ", CurrentPrice: 2748, ChangeRatio: 99, TickCount: 40579, UpCount: 12722, DownCount: 12798, ChangePercentage: 3.73, TradingVolume: 16086.8, Turnover: 43810.0498, ExchangeName: "東証２部", IndustryName: "食料品"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_NO_CHANGE, AverageRanking: 2, SymbolCode: "9984", SymbolName: "ｿﾌﾄﾊﾞﾝｸG", CurrentPrice: 8285, ChangeRatio: -309, TickCount: 32219, UpCount: 8655, DownCount: 8562, ChangePercentage: -3.59, TradingVolume: 16688.8, Turnover: 138143.1773, ExchangeName: "東証１部", IndustryName: "情報・通信業"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToTickRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToVolumeRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.VolumeRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypeVolumeRapidIncrease, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.VolumeRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.VolumeRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeVolumeRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.VolumeRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.VolumeRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeVolumeRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				VolumeRapidRanking: []kabus.VolumeRapidRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidTradePercentage: 49900, TradingVolume: 1, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidTradePercentage: 28189.47, TradingVolume: 3117.5, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.VolumeRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.VolumeRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidTradePercentage: 49900, TradingVolume: 1, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)), ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidTradePercentage: 28189.47, TradingVolume: 3117.5, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 54.65, ExchangeName: "東証JQS", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToVolumeRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToValueRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.ValueRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypeValueRapidIncrease, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.ValueRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.ValueRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeValueRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.ValueRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.ValueRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeValueRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				ValueRapidRanking: []kabus.ValueRapidRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidPaymentPercentage: 55381.47, Turnover: 3194.7121, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidPaymentPercentage: 50159.4, Turnover: 7.75, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
				},
			},
			want: &kabuspb.ValueRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.ValueRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidPaymentPercentage: 55381.47, Turnover: 3194.7121, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 54.65, ExchangeName: "東証JQS", IndustryName: "電気機器"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidPaymentPercentage: 50159.4, Turnover: 7.75, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)), ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToValueRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToMarginRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.MarginRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypeMarginHighMagnification, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.MarginRanking{
				Type:             kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.MarginRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeMarginHighMagnification,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.MarginRanking{
				Type:             kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.MarginRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypeMarginHighMagnification,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				MarginRanking: []kabus.MarginRanking{
					{No: 1, Symbol: "3150", SymbolName: "グリムス", Ratio: 14467, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.5, BuyRapidPaymentPercentage: 1446.7, BuyLastWeekRatio: 139.7, ExchangeName: "東証１部", CategoryName: "卸売業"},
					{No: 2, Symbol: "6955", SymbolName: "ＦＤＫ", Ratio: 10536.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: -0.8, BuyRapidPaymentPercentage: 2107.3, BuyLastWeekRatio: 121.6, ExchangeName: "東証２部", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.MarginRanking{
				Type:             kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.MarginRankingInfo{
					{No: 1, SymbolCode: "3150", SymbolName: "グリムス", Ratio: 14467, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.5, BuyRapidPaymentPercentage: 1446.7, BuyLastWeekRatio: 139.7, ExchangeName: "東証１部", IndustryName: "卸売業"},
					{No: 2, SymbolCode: "6955", SymbolName: "ＦＤＫ", Ratio: 10536.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: -0.8, BuyRapidPaymentPercentage: 2107.3, BuyLastWeekRatio: 121.6, ExchangeName: "東証２部", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToMarginRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_fromRankingToIndustryRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  *kabus.RankingResponse
		want *kabuspb.IndustryRanking
	}{
		{name: "ランキングの配列がnilなら空の配列が返される",
			arg: &kabus.RankingResponse{Type: kabus.RankingTypePriceIncreaseRateByCategory, ExchangeDivision: kabus.ExchangeDivisionALL},
			want: &kabuspb.IndustryRanking{
				Type:             kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.IndustryRankingInfo{}}},
		{name: "ランキングの配列が空なら空の配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRateByCategory,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking:     []kabus.PriceRanking{},
			},
			want: &kabuspb.IndustryRanking{
				Type:             kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking:          []*kabuspb.IndustryRankingInfo{}}},
		{name: "ランキングの配列に要素が2つあるなら要素が2つの配列が返される",
			arg: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRateByCategory,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				CategoryPriceRanking: []kabus.CategoryPriceRanking{
					{No: 1, Trend: kabus.RankingTrendRise, AverageRanking: 18, Category: "343", CategoryName: "IS 空運", CurrentPrice: 170.97, ChangeRatio: 6.72, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 4.09},
					{No: 2, Trend: kabus.RankingTrendRise, AverageRanking: 16, Category: "341", CategoryName: "IS 陸運", CurrentPrice: 1895.49, ChangeRatio: 15.41, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.82},
				},
			},
			want: &kabuspb.IndustryRanking{
				Type:             kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.IndustryRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE, AverageRanking: 18, IndustryCode: "343", IndustryName: "IS 空運", CurrentPrice: 170.97, ChangeRatio: 6.72, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 4.09},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE, AverageRanking: 16, IndustryCode: "341", IndustryName: "IS 陸運", CurrentPrice: 1895.49, ChangeRatio: 15.41, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 0.82},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := fromRankingToIndustryRanking(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

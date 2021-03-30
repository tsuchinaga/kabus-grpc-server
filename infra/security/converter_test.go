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
		{name: "事象なし を変換できる", arg: kabus.CurrentPriceChangeStatusUnspecified, want: "0000"},
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

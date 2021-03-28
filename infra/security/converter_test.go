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

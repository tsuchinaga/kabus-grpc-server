package security

import (
	"reflect"
	"testing"

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
		{name: "東証の変換ができる", arg: kabuspb.Exchange_TOUSHOU, want: kabus.ExchangeToushou},
		{name: "名証の変換ができる", arg: kabuspb.Exchange_MEISHOU, want: kabus.ExchangeMeishou},
		{name: "福証の変換ができる", arg: kabuspb.Exchange_FUKUSHOU, want: kabus.ExchangeFukushou},
		{name: "札証の変換ができる", arg: kabuspb.Exchange_SATSUSHOU, want: kabus.ExchangeSatsushou},
		{name: "日通しの変換ができる", arg: kabuspb.Exchange_ALL_SESSION, want: kabus.ExchangeAll},
		{name: "日中の変換ができる", arg: kabuspb.Exchange_DAY_SESSION, want: kabus.ExchangeDaytime},
		{name: "夜間の変換ができる", arg: kabuspb.Exchange_NIGHT_SESSION, want: kabus.ExchangeEvening},
		{name: "指定なしの変換ができる", arg: kabuspb.Exchange_UNSPECIFIED, want: kabus.ExchangeUnspecified},
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
		{name: "東証の変換ができる", arg: kabus.ExchangeToushou, want: kabuspb.Exchange_TOUSHOU},
		{name: "名証の変換ができる", arg: kabus.ExchangeMeishou, want: kabuspb.Exchange_MEISHOU},
		{name: "福証の変換ができる", arg: kabus.ExchangeFukushou, want: kabuspb.Exchange_FUKUSHOU},
		{name: "札証の変換ができる", arg: kabus.ExchangeSatsushou, want: kabuspb.Exchange_SATSUSHOU},
		{name: "日通しの変換ができる", arg: kabus.ExchangeAll, want: kabuspb.Exchange_ALL_SESSION},
		{name: "日中の変換ができる", arg: kabus.ExchangeDaytime, want: kabuspb.Exchange_DAY_SESSION},
		{name: "夜間の変換ができる", arg: kabus.ExchangeEvening, want: kabuspb.Exchange_NIGHT_SESSION},
		{name: "指定なしの変換ができる", arg: kabus.ExchangeUnspecified, want: kabuspb.Exchange_UNSPECIFIED},
		{name: "未定義の変換ができる", arg: kabus.Exchange(-1), want: kabuspb.Exchange_UNSPECIFIED},
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

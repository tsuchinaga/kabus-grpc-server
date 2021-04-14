package stores

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_GetRegisterSymbolStore(t *testing.T) {
	t.Parallel()
	got := GetRegisterSymbolStore()
	want := &registerSymbol{store: []*kabuspb.RegisterSymbol{}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_registerSymbol_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store repositories.RegisterSymbolStore
		want  []*kabuspb.RegisterSymbol
	}{
		{name: "storeが空配列なら空配列を返す", store: &registerSymbol{store: []*kabuspb.RegisterSymbol{}}, want: []*kabuspb.RegisterSymbol{}},
		{name: "storeにデータがあればそれを返す",
			store: &registerSymbol{store: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, {SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:  []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, {SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.store.GetAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_SetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *registerSymbol
		arg   []*kabuspb.RegisterSymbol
		want  []*kabuspb.RegisterSymbol
	}{
		{name: "データのある配列で上書き出る",
			store: &registerSymbol{store: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, {SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			arg:   []*kabuspb.RegisterSymbol{{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want:  []*kabuspb.RegisterSymbol{{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
		{name: "空配列で上書きできる",
			store: &registerSymbol{store: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, {SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			arg:   []*kabuspb.RegisterSymbol{},
			want:  []*kabuspb.RegisterSymbol{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.SetAll(test.arg)
			got := test.store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

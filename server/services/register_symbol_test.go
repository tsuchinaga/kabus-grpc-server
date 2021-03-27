package services

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_NewRegisterSymbolService(t *testing.T) {
	t.Parallel()
	got := NewRegisterSymbolService(&testRegisterSymbolStore{})
	want := &registerSymbol{registerSymbolStore: &testRegisterSymbolStore{}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_registerSymbol_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []*kabuspb.RegisterSymbol
		want   []*kabuspb.RegisterSymbol
	}{
		{name: "storeのgetAllが空なら空を返す",
			getAll: []*kabuspb.RegisterSymbol{},
			want:   []*kabuspb.RegisterSymbol{}},
		{name: "storeのgetAllにデータがあればそれを返す",
			getAll: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}},
			want:   []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &registerSymbol{registerSymbolStore: &testRegisterSymbolStore{getAll: test.getAll}}
			got := service.Get()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_Set(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  []*kabuspb.RegisterSymbol
		want []*kabuspb.RegisterSymbol
	}{
		{name: "データの入っている引数をstoreに渡せる",
			arg:  []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}},
			want: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
		{name: "空の引数をstoreに渡せる",
			arg:  []*kabuspb.RegisterSymbol{},
			want: []*kabuspb.RegisterSymbol{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &testRegisterSymbolStore{}
			service := &registerSymbol{registerSymbolStore: store}
			service.Set(test.arg)
			got := store.lastSet
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

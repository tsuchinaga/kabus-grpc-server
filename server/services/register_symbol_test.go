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

func Test_registerSymbol_CountAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		countAll int
		want     int
	}{
		{name: "storeが0を返したら0を返す", countAll: 0, want: 0},
		{name: "storeが10を返したら10を返す", countAll: 10, want: 10},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &registerSymbol{registerSymbolStore: &testRegisterSymbolStore{countAll: test.countAll}}
			got := service.CountAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []*kabuspb.RegisterSymbol
		want   []*kabuspb.RegisterSymbol
	}{
		{name: "storeのGetAllが空なら空を返す",
			getAll: []*kabuspb.RegisterSymbol{},
			want:   []*kabuspb.RegisterSymbol{}},
		{name: "storeのGetAllにデータがあればそれを返す",
			getAll: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want:   []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &registerSymbol{registerSymbolStore: &testRegisterSymbolStore{getAll: test.getAll}}
			got := service.GetAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		getByRequester []*kabuspb.RegisterSymbol
		want           []*kabuspb.RegisterSymbol
	}{
		{name: "storeのGetByRequesterが空なら空を返す",
			getByRequester: []*kabuspb.RegisterSymbol{},
			want:           []*kabuspb.RegisterSymbol{}},
		{name: "storeのGetByRequesterにデータがあればそれを返す",
			getByRequester: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want:           []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &registerSymbol{registerSymbolStore: &testRegisterSymbolStore{getByRequester: test.getByRequester}}
			got := service.Get("requester")
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		arg1            string
		arg2            []*kabuspb.RegisterSymbol
		callAddAllCount int
	}{
		{name: "storeのAddAllが叩かれる",
			arg1:            "requester",
			arg2:            []*kabuspb.RegisterSymbol{{SymbolCode: "1234"}},
			callAddAllCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolStore := &testRegisterSymbolStore{}
			service := &registerSymbol{registerSymbolStore: registerSymbolStore}
			service.Add(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.callAddAllCount, registerSymbolStore.callAddAllCount) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.callAddAllCount, registerSymbolStore.callAddAllCount)
			}
		})
	}
}

func Test_registerSymbol_Remove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		arg1               string
		arg2               []*kabuspb.RegisterSymbol
		callRemoveAllCount int
	}{
		{name: "storeのRemoveAllが叩かれる",
			arg1:               "requester",
			arg2:               []*kabuspb.RegisterSymbol{{SymbolCode: "1234"}},
			callRemoveAllCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolStore := &testRegisterSymbolStore{}
			service := &registerSymbol{registerSymbolStore: registerSymbolStore}
			service.Remove(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.callRemoveAllCount, registerSymbolStore.callRemoveAllCount) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.callRemoveAllCount, registerSymbolStore.callRemoveAllCount)
			}
		})
	}
}

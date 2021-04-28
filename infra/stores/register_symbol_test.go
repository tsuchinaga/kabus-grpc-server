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
	want := &registerSymbol{store: []*RegisterSymbol{}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_registerSymbol_Count(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store repositories.RegisterSymbolStore
		want  int
	}{
		{name: "storeが空配列なら0",
			store: &registerSymbol{store: []*RegisterSymbol{}},
			want:  0},
		{name: "storeにデータがあればその要素数を返す",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "baz"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar", "baz"}}}},
			want: 3},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.store.CountAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store repositories.RegisterSymbolStore
		want  []*kabuspb.RegisterSymbol
	}{
		{name: "storeが空配列なら空配列を返す",
			store: &registerSymbol{store: []*RegisterSymbol{}},
			want:  []*kabuspb.RegisterSymbol{}},
		{name: "storeにデータがあればそれを返す",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "baz"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar", "baz"}}}},
			want: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
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

func Test_registerSymbol_GetByRequester(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store repositories.RegisterSymbolStore
		arg   string
		want  []*kabuspb.RegisterSymbol
	}{
		{name: "storeが空配列なら空配列を返す",
			store: &registerSymbol{store: []*RegisterSymbol{}},
			arg:   "foo",
			want:  []*kabuspb.RegisterSymbol{}},
		{name: "storeにデータがあればそれを返す",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "baz"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar", "baz"}}}},
			arg:  "foo",
			want: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, {SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
		{name: "storeにデータがあってもrequesterが一致しなければ何も返されない",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "baz"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar", "baz"}}}},
			arg:  "hoge",
			want: []*kabuspb.RegisterSymbol{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.store.GetByRequester(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_AddAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *registerSymbol
		arg1  string
		arg2  []*kabuspb.RegisterSymbol
		want  []*RegisterSymbol
	}{
		{name: "登録する銘柄がstoreになければ追加される",
			store: &registerSymbol{store: []*RegisterSymbol{}},
			arg1:  "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo"}},
			}},
		{name: "登録する銘柄がstoreにあってもrequesterに登録されていなければ登録する",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz"}}}},
			arg1: "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar", "foo"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo"}},
			}},
		{name: "登録する銘柄がstoreにあってrequesterに登録されていれば登録しない",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo", "bar"}}}},
			arg1: "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo", "bar"}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.AddAll(test.arg1, test.arg2)
			got := test.store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerSymbol_RemoveAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *registerSymbol
		arg1  string
		arg2  []*kabuspb.RegisterSymbol
		want  []*RegisterSymbol
	}{
		{name: "storeが空なら何も起こらない",
			store: &registerSymbol{store: []*RegisterSymbol{}},
			arg1:  "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{}},
		{name: "削除する銘柄がstoreになければなにもされない",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo"}}}},
			arg1: "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "3456", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_MEISHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo"}},
			}},
		{name: "削除する銘柄がstoreにあればrequesterから削る",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo", "bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo"}}}},
			arg1: "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"bar"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz"}},
			}},
		{name: "削除する銘柄がstoreにあればrequesterから削り、requesterが空になったらstoreからも消す",
			store: &registerSymbol{store: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"foo"}},
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz", "foo"}}}},
			arg1: "foo",
			arg2: []*kabuspb.RegisterSymbol{
				{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
				{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: []*RegisterSymbol{
				{Symbol: &kabuspb.RegisterSymbol{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}, Requesters: []string{"baz"}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.RemoveAll(test.arg1, test.arg2)
			got := test.store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_RegisterSymbol_AddRequester(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		requesters []string
		arg        string
		want       []string
	}{
		{name: "requestersに存在しないrequesterなら追加される",
			requesters: []string{"foo", "bar", "baz"},
			arg:        "hoge",
			want:       []string{"foo", "bar", "baz", "hoge"}},
		{name: "requestersに存在するrequesterなら追加されない",
			requesters: []string{"foo", "bar", "baz"},
			arg:        "bar",
			want:       []string{"foo", "bar", "baz"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbol := &RegisterSymbol{Requesters: test.requesters}
			symbol.AddRequester(test.arg)
			got := symbol.Requesters
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_RegisterSymbol_HasRequester(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		requesters []string
		arg        string
		want       bool
	}{
		{name: "requesterが含まれていればtrue",
			requesters: []string{"foo", "bar", "baz"},
			arg:        "bar",
			want:       true},
		{name: "requesterが含まれていなければfalse",
			requesters: []string{"foo", "bar", "baz"},
			arg:        "hoge",
			want:       false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &RegisterSymbol{Requesters: test.requesters}
			got := store.HasRequester(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_RegisterSymbol_IsEmptyRequesters(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		requesters []string
		want       bool
	}{
		{name: "requestersが空ならtrue",
			requesters: []string{},
			want:       true},
		{name: "requestersが空じゃないならfalse",
			requesters: []string{"foo", "bar", "baz"},
			want:       false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbol := &RegisterSymbol{Requesters: test.requesters}
			got := symbol.IsRequestersEmpty()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_SymbolRequester_RemoveRequester(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		requesters []string
		arg        string
		want       []string
	}{
		{name: "指定したrequesterが含まれていたら削除する",
			requesters: []string{"foo", "bar", "baz", "bar"},
			arg:        "bar",
			want:       []string{"foo", "baz"}},
		{name: "指定したrequesterが複数含まれていたら該当するrequesterをすべて削除する",
			requesters: []string{"foo", "bar", "baz", "bar"},
			arg:        "bar",
			want:       []string{"foo", "baz"}},
		{name: "指定したrequesterが含まれていなければ何もしない",
			requesters: []string{"foo", "bar", "baz"},
			arg:        "hoge",
			want:       []string{"foo", "bar", "baz"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbol := &RegisterSymbol{Requesters: test.requesters}
			symbol.RemoveRequester(test.arg)
			got := symbol.Requesters
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

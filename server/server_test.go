package server

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

type testSecurity struct {
	repositories.Security
	register1         *kabuspb.RegisteredSymbols
	register2         error
	unregister1       *kabuspb.RegisteredSymbols
	unregister2       error
	unregisterAll1    *kabuspb.RegisteredSymbols
	unregisterAll2    error
	symbolNameFuture1 *kabuspb.SymbolCodeInfo
	symbolNameFuture2 error
	symbolNameOption1 *kabuspb.SymbolCodeInfo
	symbolNameOption2 error
}

func (t *testSecurity) RegisterSymbols(context.Context, string, *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return t.register1, t.register2
}

func (t *testSecurity) UnregisterSymbols(context.Context, string, *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return t.unregister1, t.unregister2
}

func (t *testSecurity) UnregisterAll(context.Context, string, *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return t.unregisterAll1, t.unregisterAll2
}

func (t *testSecurity) SymbolNameFuture(context.Context, string, *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	return t.symbolNameFuture1, t.symbolNameFuture2
}

func (t *testSecurity) SymbolNameOption(context.Context, string, *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	return t.symbolNameOption1, t.symbolNameOption2
}

type testTokenService struct {
	services.TokenService
	getToken1 string
	getToken2 error
}

func (t *testTokenService) GetToken(context.Context) (string, error) {
	return t.getToken1, t.getToken2
}

type testRegisterSymbolService struct {
	services.RegisterSymbolService
	get     []*kabuspb.RegisterSymbol
	lastSet []*kabuspb.RegisterSymbol
}

func (t *testRegisterSymbolService) Get() []*kabuspb.RegisterSymbol {
	return t.get
}
func (t *testRegisterSymbolService) Set(registeredList []*kabuspb.RegisterSymbol) {
	t.lastSet = registeredList
}

func Test_NewServer(t *testing.T) {
	security := &testSecurity{}
	tokenService := &testTokenService{}
	registerSymbolService := &testRegisterSymbolService{}
	got := NewServer(security, tokenService, registerSymbolService)
	want := &server{security: security, tokenService: tokenService, registerSymbolService: registerSymbolService}
	t.Parallel()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_server_GetRegisteredSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		get      []*kabuspb.RegisterSymbol
		want     *kabuspb.RegisteredSymbols
		hasError bool
	}{
		{name: "registerSymbolServiceの結果が返される",
			get:  []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{registerSymbolService: &testRegisterSymbolService{get: test.get}}
			got1, got2 := server.GetRegisteredSymbols(context.Background(), &kabuspb.GetRegisteredSymbolsRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_RegisterSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		register1 *kabuspb.RegisteredSymbols
		register2 error
		want      *kabuspb.RegisteredSymbols
		hasError  bool
		wantSet   []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Registerでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			register2: errors.New("register error message"),
			hasError:  true},
		{name: "Registerの結果をStoreに保存してから結果を返す",
			getToken1: "TOKEN_STRING",
			register1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:      &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			wantSet:   []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{}
			server := &server{
				security:              &testSecurity{register1: test.register1, register2: test.register2},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				registerSymbolService: registerSymbolService}
			got1, got2 := server.RegisterSymbols(context.Background(), &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{}})
			got3 := registerSymbolService.lastSet
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantSet, got3) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantSet, got1, got2, got3)
			}
		})
	}
}

func Test_server_UnregisterSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		getToken1   string
		getToken2   error
		unregister1 *kabuspb.RegisteredSymbols
		unregister2 error
		want        *kabuspb.RegisteredSymbols
		hasError    bool
		wantSet     []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Unregisterでエラーがあればエラーを返す",
			getToken1:   "TOKEN_STRING",
			unregister2: errors.New("register error message"),
			hasError:    true},
		{name: "Unregisterの結果をStoreに保存してから結果を返す",
			getToken1:   "TOKEN_STRING",
			unregister1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:        &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			wantSet:     []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{}
			server := &server{
				security:              &testSecurity{unregister1: test.unregister1, unregister2: test.unregister2},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				registerSymbolService: registerSymbolService}
			got1, got2 := server.UnregisterSymbols(context.Background(), &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{}})
			got3 := registerSymbolService.lastSet
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantSet, got3) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantSet, got1, got2, got3)
			}
		})
	}
}

func Test_server_UnregisterAllSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		getToken1      string
		getToken2      error
		unregisterAll1 *kabuspb.RegisteredSymbols
		unregisterAll2 error
		want           *kabuspb.RegisteredSymbols
		hasError       bool
		wantSet        []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "UnregisterAllでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			unregisterAll2: errors.New("register error message"),
			hasError:       true},
		{name: "UnregisterAllの結果をStoreに保存してから結果を返す",
			getToken1:      "TOKEN_STRING",
			unregisterAll1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:           &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			wantSet:        []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{}
			server := &server{
				security:              &testSecurity{unregisterAll1: test.unregisterAll1, unregisterAll2: test.unregisterAll2},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				registerSymbolService: registerSymbolService}
			got1, got2 := server.UnregisterAllSymbols(context.Background(), &kabuspb.UnregisterAllSymbolsRequest{})
			got3 := registerSymbolService.lastSet
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantSet, got3) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantSet, got1, got2, got3)
			}
		})
	}
}

func Test_server_GetFutureSymbolCodeInfo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		getToken1         string
		getToken2         error
		symbolNameFuture1 *kabuspb.SymbolCodeInfo
		symbolNameFuture2 error
		want              *kabuspb.SymbolCodeInfo
		hasError          bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "SymbolNameFutureでエラーがあればエラーを返す",
			getToken1:         "TOKEN_STRING",
			symbolNameFuture2: errors.New("register error message"),
			hasError:          true},
		{name: "SymbolNameFutureの結果をStoreに保存してから結果を返す",
			getToken1:         "TOKEN_STRING",
			symbolNameFuture1: &kabuspb.SymbolCodeInfo{Code: "166060018", Name: "日経平均先物 21/06"},
			want:              &kabuspb.SymbolCodeInfo{Code: "166060018", Name: "日経平均先物 21/06"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{symbolNameFuture1: test.symbolNameFuture1, symbolNameFuture2: test.symbolNameFuture2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetFutureSymbolCodeInfo(context.Background(), &kabuspb.GetFutureSymbolCodeInfoRequest{
				FutureCode:      kabuspb.FutureCode_FUTURE_CODE_NK225,
				DerivativeMonth: timestamppb.Now()})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetOptionSymbolCodeInfo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		getToken1         string
		getToken2         error
		symbolNameOption1 *kabuspb.SymbolCodeInfo
		symbolNameOption2 error
		want              *kabuspb.SymbolCodeInfo
		hasError          bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "SymbolNameOptionでエラーがあればエラーを返す",
			getToken1:         "TOKEN_STRING",
			symbolNameOption2: errors.New("register error message"),
			hasError:          true},
		{name: "SymbolNameOptionの結果をStoreに保存してから結果を返す",
			getToken1:         "TOKEN_STRING",
			symbolNameOption1: &kabuspb.SymbolCodeInfo{Code: "166060018", Name: "日経平均先物 21/06"},
			want:              &kabuspb.SymbolCodeInfo{Code: "166060018", Name: "日経平均先物 21/06"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{symbolNameOption1: test.symbolNameOption1, symbolNameOption2: test.symbolNameOption2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetOptionSymbolCodeInfo(context.Background(), &kabuspb.GetOptionSymbolCodeInfoRequest{
				DerivativeMonth: timestamppb.Now(),
				CallOrPut:       kabuspb.CallPut_CALL_PUT_CALL,
				StrikePrice:     0})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}
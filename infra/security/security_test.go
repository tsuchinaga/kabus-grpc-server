package security

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

type testRESTClient struct {
	kabus.RESTClient
	tokenWithContext1         *kabus.TokenResponse
	tokenWithContext2         error
	registerWithContext1      *kabus.RegisterResponse
	registerWithContext2      error
	lastRegisterWithContext   kabus.RegisterRequest
	unregisterWithContext1    *kabus.UnregisterResponse
	unregisterWithContext2    error
	lastUnregisterWithContext kabus.UnregisterRequest
	unregisterAllWithContext1 *kabus.UnregisterAllResponse
	unregisterAllWithContext2 error
}

func (t *testRESTClient) TokenWithContext(context.Context, kabus.TokenRequest) (*kabus.TokenResponse, error) {
	return t.tokenWithContext1, t.tokenWithContext2
}

func (t *testRESTClient) RegisterWithContext(_ context.Context, _ string, request kabus.RegisterRequest) (*kabus.RegisterResponse, error) {
	t.lastRegisterWithContext = request
	return t.registerWithContext1, t.registerWithContext2
}

func (t *testRESTClient) UnregisterWithContext(_ context.Context, _ string, request kabus.UnregisterRequest) (*kabus.UnregisterResponse, error) {
	t.lastUnregisterWithContext = request
	return t.unregisterWithContext1, t.unregisterWithContext2
}

func (t *testRESTClient) UnregisterAllWithContext(context.Context, string) (*kabus.UnregisterAllResponse, error) {
	return t.unregisterAllWithContext1, t.unregisterAllWithContext2
}

func Test_NewSecurity(t *testing.T) {
	t.Parallel()

	got := NewSecurity(&testRESTClient{})
	want := &security{restClient: &testRESTClient{}}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_security_Token(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		tokenWithContext1 *kabus.TokenResponse
		tokenWithContext2 error
		want              string
		hasError          bool
	}{
		{name: "errが返されればerrを返す", tokenWithContext2: errors.New("error message"), hasError: true},
		{name: "tokenが返されればtokenの中身を取り出して返す", tokenWithContext1: &kabus.TokenResponse{Token: "TOKEN_STRING"}, want: "TOKEN_STRING"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			security := &security{restClient: &testRESTClient{tokenWithContext1: test.tokenWithContext1, tokenWithContext2: test.tokenWithContext2}}
			got1, got2 := security.Token(context.Background(), "")
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_Register(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		registerWithContext1 *kabus.RegisterResponse
		registerWithContext2 error
		argReq               *kabuspb.RegisterSymbolsRequest
		want                 *kabuspb.RegisteredSymbols
		hasError             bool
		wantReq              kabus.RegisterRequest
	}{
		{name: "errorを返されたらerrorを返す",
			registerWithContext2: errors.New("error message"),
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			hasError:             true,
			wantReq:              kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			registerWithContext1: &kabus.RegisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			want:                 &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			wantReq:              kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{registerWithContext1: test.registerWithContext1, registerWithContext2: test.registerWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.RegisterSymbols(context.Background(), "", test.argReq)
			got3 := restClient.lastRegisterWithContext
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantReq, got3) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantReq, got1, got2, got3)
			}
		})
	}
}

func Test_security_Unregister(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                   string
		unregisterWithContext1 *kabus.UnregisterResponse
		unregisterWithContext2 error
		argReq                 *kabuspb.UnregisterSymbolsRequest
		want                   *kabuspb.RegisteredSymbols
		hasError               bool
		wantReq                kabus.UnregisterRequest
	}{
		{name: "errorを返されたらerrorを返す",
			unregisterWithContext2: errors.New("error message"),
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			hasError:               true,
			wantReq:                kabus.UnregisterRequest{Symbols: []kabus.UnregisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			unregisterWithContext1: &kabus.UnregisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			want:                   &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_TOUSHOU}}},
			wantReq:                kabus.UnregisterRequest{Symbols: []kabus.UnregisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{unregisterWithContext1: test.unregisterWithContext1, unregisterWithContext2: test.unregisterWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.UnregisterSymbols(context.Background(), "", test.argReq)
			got3 := restClient.lastUnregisterWithContext
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantReq, got3) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantReq, got1, got2, got3)
			}
		})
	}
}

func Test_security_UnregisterAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                      string
		unregisterAllWithContext1 *kabus.UnregisterAllResponse
		unregisterAllWithContext2 error
		want                      *kabuspb.RegisteredSymbols
		hasError                  bool
	}{
		{name: "errorを返されたらerrorを返す",
			unregisterAllWithContext2: errors.New("error message"),
			hasError:                  true},
		{name: "responseが返されたらresponseを変換して返す",
			unregisterAllWithContext1: &kabus.UnregisterAllResponse{RegistList: []kabus.RegisteredSymbol{}},
			want:                      &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{unregisterAllWithContext1: test.unregisterAllWithContext1, unregisterAllWithContext2: test.unregisterAllWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.UnregisterAll(context.Background(), "", &kabuspb.UnregisterAllSymbolsRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

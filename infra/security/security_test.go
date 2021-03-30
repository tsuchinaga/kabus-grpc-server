package security

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

type testRESTClient struct {
	kabus.RESTClient
	tokenWithContext1            *kabus.TokenResponse
	tokenWithContext2            error
	symbolNameFutureWithContext1 *kabus.SymbolNameFutureResponse
	symbolNameFutureWithContext2 error
	symbolNameOptionWithContext1 *kabus.SymbolNameOptionResponse
	symbolNameOptionWithContext2 error
	registerWithContext1         *kabus.RegisterResponse
	registerWithContext2         error
	lastRegisterWithContext      kabus.RegisterRequest
	unregisterWithContext1       *kabus.UnregisterResponse
	unregisterWithContext2       error
	lastUnregisterWithContext    kabus.UnregisterRequest
	unregisterAllWithContext1    *kabus.UnregisterAllResponse
	unregisterAllWithContext2    error
	boardWithContext1            *kabus.BoardResponse
	boardWithContext2            error
	symbolWithContext1           *kabus.SymbolResponse
	symbolWithContext2           error
}

func (t *testRESTClient) TokenWithContext(context.Context, kabus.TokenRequest) (*kabus.TokenResponse, error) {
	return t.tokenWithContext1, t.tokenWithContext2
}

func (t *testRESTClient) SymbolNameFutureWithContext(context.Context, string, kabus.SymbolNameFutureRequest) (*kabus.SymbolNameFutureResponse, error) {
	return t.symbolNameFutureWithContext1, t.symbolNameFutureWithContext2
}

func (t *testRESTClient) SymbolNameOptionWithContext(context.Context, string, kabus.SymbolNameOptionRequest) (*kabus.SymbolNameOptionResponse, error) {
	return t.symbolNameOptionWithContext1, t.symbolNameOptionWithContext2
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

func (t *testRESTClient) BoardWithContext(context.Context, string, kabus.BoardRequest) (*kabus.BoardResponse, error) {
	return t.boardWithContext1, t.boardWithContext2
}

func (t *testRESTClient) SymbolWithContext(context.Context, string, kabus.SymbolRequest) (*kabus.SymbolResponse, error) {
	return t.symbolWithContext1, t.symbolWithContext2
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
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			hasError:             true,
			wantReq:              kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			registerWithContext1: &kabus.RegisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:                 &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
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
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			hasError:               true,
			wantReq:                kabus.UnregisterRequest{Symbols: []kabus.UnregisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			unregisterWithContext1: &kabus.UnregisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:                   &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
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

func Test_security_SymbolNameFuture(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                         string
		symbolNameFutureWithContext1 *kabus.SymbolNameFutureResponse
		symbolNameFutureWithContext2 error
		arg                          *kabuspb.GetFutureSymbolCodeInfoRequest
		want                         *kabuspb.SymbolCodeInfo
		hasError                     bool
	}{
		{name: "errorを返されたらerrorを返す",
			symbolNameFutureWithContext2: errors.New("error message"),
			arg:                          &kabuspb.GetFutureSymbolCodeInfoRequest{FutureCode: kabuspb.FutureCode_FUTURE_CODE_NK225, DerivativeMonth: timestamppb.Now()},
			hasError:                     true},
		{name: "responseが返されたらresponseを変換して返す",
			symbolNameFutureWithContext1: &kabus.SymbolNameFutureResponse{Symbol: "166060018", SymbolName: "日経平均先物 21/06"},
			arg:                          &kabuspb.GetFutureSymbolCodeInfoRequest{FutureCode: kabuspb.FutureCode_FUTURE_CODE_NK225, DerivativeMonth: timestamppb.Now()},
			want:                         &kabuspb.SymbolCodeInfo{Code: "166060018", Name: "日経平均先物 21/06"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{symbolNameFutureWithContext1: test.symbolNameFutureWithContext1, symbolNameFutureWithContext2: test.symbolNameFutureWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SymbolNameFuture(context.Background(), "", test.arg)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_SymbolNameOption(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                         string
		symbolNameOptionWithContext1 *kabus.SymbolNameOptionResponse
		symbolNameOptionWithContext2 error
		arg                          *kabuspb.GetOptionSymbolCodeInfoRequest
		want                         *kabuspb.SymbolCodeInfo
		hasError                     bool
	}{
		{name: "errorを返されたらerrorを返す",
			symbolNameOptionWithContext2: errors.New("error message"),
			arg: &kabuspb.GetOptionSymbolCodeInfoRequest{
				DerivativeMonth: timestamppb.Now(),
				CallOrPut:       kabuspb.CallPut_CALL_PUT_PUT,
				StrikePrice:     0,
			},
			hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			symbolNameOptionWithContext1: &kabus.SymbolNameOptionResponse{Symbol: "136049118", SymbolName: "日経平均オプション 21/04 プット 29125"},
			arg: &kabuspb.GetOptionSymbolCodeInfoRequest{
				DerivativeMonth: timestamppb.Now(),
				CallOrPut:       kabuspb.CallPut_CALL_PUT_PUT,
				StrikePrice:     0,
			},
			want: &kabuspb.SymbolCodeInfo{Code: "136049118", Name: "日経平均オプション 21/04 プット 29125"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{symbolNameOptionWithContext1: test.symbolNameOptionWithContext1, symbolNameOptionWithContext2: test.symbolNameOptionWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SymbolNameOption(context.Background(), "", test.arg)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_Board(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		boardWithContext1 *kabus.BoardResponse
		boardWithContext2 error
		want              *kabuspb.Board
		hasError          bool
	}{
		{name: "errorを返されたらerrorを返す", boardWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			boardWithContext1: &kabus.BoardResponse{
				Symbol:                   "5401",
				SymbolName:               "新日鐵住金",
				Exchange:                 kabus.ExchangeToushou,
				ExchangeName:             "東証１部",
				CurrentPrice:             2408,
				CurrentPriceTime:         time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local),
				CurrentPriceChangeStatus: kabus.CurrentPriceChangeStatusDown,
				CurrentPriceStatus:       kabus.CurrentPriceStatusCurrentPrice,
				CalcPrice:                343.7,
				PreviousClose:            1048,
				PreviousCloseTime:        time.Date(2020, 7, 21, 0, 0, 0, 0, time.Local),
				ChangePreviousClose:      1360,
				ChangePreviousClosePer:   129.77,
				OpeningPrice:             2380,
				OpeningPriceTime:         time.Date(2020, 7, 22, 9, 0, 0, 0, time.Local),
				HighPrice:                2418,
				HighPriceTime:            time.Date(2020, 7, 22, 13, 25, 47, 0, time.Local),
				LowPrice:                 2370,
				LowPriceTime:             time.Date(2020, 7, 22, 10, 0, 4, 0, time.Local),
				TradingVolume:            4571500,
				TradingVolumeTime:        time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local),
				VWAP:                     2394.4262,
				TradingValue:             10946119350,
				BidQty:                   100,
				BidPrice:                 2408.5,
				BidTime:                  time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
				BidSign:                  kabus.BidAskSignGeneral,
				MarketOrderSellQty:       0,
				Sell1:                    kabus.FirstBoardSign{Time: time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local), Sign: kabus.BidAskSignGeneral, Price: 2408.5, Qty: 100},
				Sell2:                    kabus.BoardSign{Price: 2409, Qty: 800},
				Sell3:                    kabus.BoardSign{Price: 2409.5, Qty: 2100},
				Sell4:                    kabus.BoardSign{Price: 2410, Qty: 800},
				Sell5:                    kabus.BoardSign{Price: 2410.5, Qty: 500},
				Sell6:                    kabus.BoardSign{Price: 2411, Qty: 8400},
				Sell7:                    kabus.BoardSign{Price: 2411.5, Qty: 1200},
				Sell8:                    kabus.BoardSign{Price: 2412, Qty: 27200},
				Sell9:                    kabus.BoardSign{Price: 2412.5, Qty: 400},
				Sell10:                   kabus.BoardSign{Price: 2413, Qty: 16400},
				AskQty:                   200,
				AskPrice:                 2407.5,
				AskTime:                  time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
				AskSign:                  kabus.BidAskSignGeneral,
				MarketOrderBuyQty:        0,
				Buy1:                     kabus.FirstBoardSign{Time: time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local), Sign: kabus.BidAskSignGeneral, Price: 2407.5, Qty: 200},
				Buy2:                     kabus.BoardSign{Price: 2407, Qty: 400},
				Buy3:                     kabus.BoardSign{Price: 2406.5, Qty: 1000},
				Buy4:                     kabus.BoardSign{Price: 2406, Qty: 5800},
				Buy5:                     kabus.BoardSign{Price: 2405.5, Qty: 7500},
				Buy6:                     kabus.BoardSign{Price: 2405, Qty: 2200},
				Buy7:                     kabus.BoardSign{Price: 2404.5, Qty: 16700},
				Buy8:                     kabus.BoardSign{Price: 2404, Qty: 30100},
				Buy9:                     kabus.BoardSign{Price: 2403.5, Qty: 1300},
				Buy10:                    kabus.BoardSign{Price: 2403, Qty: 3000},
				OverSellQty:              974900,
				UnderBuyQty:              756000,
				TotalMarketValue:         3266254659361.4,
				ClearingPrice:            23000,
				IV:                       22.11,
				Gamma:                    0.000183,
				Theta:                    -6.5073,
				Vega:                     39.3109,
				Delta:                    0.4794,
			},
			want: &kabuspb.Board{
				SymbolCode:               "5401",
				SymbolName:               "新日鐵住金",
				Exchange:                 kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:             "東証１部",
				CurrentPrice:             2408,
				CurrentPriceTime:         timestamppb.New(time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local)),
				CurrentPriceChangeStatus: "0058",
				CurrentPriceStatus:       1,
				CalculationPrice:         343.7,
				PreviousClose:            1048,
				PreviousCloseTime:        timestamppb.New(time.Date(2020, 7, 21, 0, 0, 0, 0, time.Local)),
				ChangePreviousClose:      1360,
				ChangePreviousClosePer:   129.77,
				OpeningPrice:             2380,
				OpeningPriceTime:         timestamppb.New(time.Date(2020, 7, 22, 9, 0, 0, 0, time.Local)),
				HighPrice:                2418,
				HighPriceTime:            timestamppb.New(time.Date(2020, 7, 22, 13, 25, 47, 0, time.Local)),
				LowPrice:                 2370,
				LowPriceTime:             timestamppb.New(time.Date(2020, 7, 22, 10, 0, 4, 0, time.Local)),
				TradingVolume:            4571500,
				TradingVolumeTime:        timestamppb.New(time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local)),
				Vwap:                     2394.4262,
				TradingValue:             10946119350,
				BidQuantity:              100,
				BidPrice:                 2408.5,
				BidTime:                  timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)),
				BidSign:                  "0101",
				MarketOrderSellQuantity:  0,
				Sell1:                    &kabuspb.FirstQuote{Time: timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)), Sign: "0101", Price: 2408.5, Quantity: 100},
				Sell2:                    &kabuspb.Quote{Price: 2409, Quantity: 800},
				Sell3:                    &kabuspb.Quote{Price: 2409.5, Quantity: 2100},
				Sell4:                    &kabuspb.Quote{Price: 2410, Quantity: 800},
				Sell5:                    &kabuspb.Quote{Price: 2410.5, Quantity: 500},
				Sell6:                    &kabuspb.Quote{Price: 2411, Quantity: 8400},
				Sell7:                    &kabuspb.Quote{Price: 2411.5, Quantity: 1200},
				Sell8:                    &kabuspb.Quote{Price: 2412, Quantity: 27200},
				Sell9:                    &kabuspb.Quote{Price: 2412.5, Quantity: 400},
				Sell10:                   &kabuspb.Quote{Price: 2413, Quantity: 16400},
				AskQuantity:              200,
				AskPrice:                 2407.5,
				AskTime:                  timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)),
				AskSign:                  "0101",
				MarketOrderBuyQuantity:   0,
				Buy1:                     &kabuspb.FirstQuote{Time: timestamppb.New(time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local)), Sign: "0101", Price: 2407.5, Quantity: 200},
				Buy2:                     &kabuspb.Quote{Price: 2407, Quantity: 400},
				Buy3:                     &kabuspb.Quote{Price: 2406.5, Quantity: 1000},
				Buy4:                     &kabuspb.Quote{Price: 2406, Quantity: 5800},
				Buy5:                     &kabuspb.Quote{Price: 2405.5, Quantity: 7500},
				Buy6:                     &kabuspb.Quote{Price: 2405, Quantity: 2200},
				Buy7:                     &kabuspb.Quote{Price: 2404.5, Quantity: 16700},
				Buy8:                     &kabuspb.Quote{Price: 2404, Quantity: 30100},
				Buy9:                     &kabuspb.Quote{Price: 2403.5, Quantity: 1300},
				Buy10:                    &kabuspb.Quote{Price: 2403, Quantity: 3000},
				OverSellQuantity:         974900,
				UnderBuyQuantity:         756000,
				TotalMarketValue:         3266254659361.4,
				ClearingPrice:            23000,
				ImpliedVolatility:        22.11,
				Gamma:                    0.000183,
				Theta:                    -6.5073,
				Vega:                     39.3109,
				Delta:                    0.4794,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{boardWithContext1: test.boardWithContext1, boardWithContext2: test.boardWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.Board(context.Background(), "", &kabuspb.GetBoardRequest{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_Symbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		symbolWithContext1 *kabus.SymbolResponse
		symbolWithContext2 error
		want               *kabuspb.Symbol
		hasError           bool
	}{
		{name: "errorを返されたらerrorを返す", symbolWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			symbolWithContext1: &kabus.SymbolResponse{
				Symbol:             "9433",
				SymbolName:         "ＫＤＤＩ",
				DisplayName:        "ＫＤＤＩ",
				Exchange:           kabus.ExchangeToushou,
				ExchangeName:       "東証１部",
				BisCategory:        "5250",
				TotalMarketValue:   7654484465100,
				TotalStocks:        4484,
				TradingUnit:        100,
				FiscalYearEndBasic: kabus.YmdNUM{Time: time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)},
				PriceRangeGroup:    kabus.PriceRangeGroup10003,
				KCMarginBuy:        true,
				KCMarginSell:       true,
				MarginBuy:          true,
				MarginSell:         true,
				UpperLimit:         4041,
				LowerLimit:         2641,
				Underlyer:          kabus.UnderlyerNK225,
				DerivMonth:         kabus.YmString{Time: time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)},
				TradeStart:         kabus.YmdNUM{Time: time.Date(2015, 12, 11, 0, 0, 0, 0, time.Local)},
				TradeEnd:           kabus.YmdNUM{Time: time.Date(2020, 12, 10, 0, 0, 0, 0, time.Local)},
				StrikePrice:        23250,
				PutOrCall:          kabus.PutOrCallNumCall,
				ClearingPrice:      23000,
			},
			want: &kabuspb.Symbol{
				Code:               "9433",
				Name:               "ＫＤＤＩ",
				DisplayName:        "ＫＤＤＩ",
				Exchange:           kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:       "東証１部",
				IndustryCategory:   "5250",
				TotalMarketValue:   7654484465100,
				TotalStocks:        4484,
				TradingUnit:        100,
				FiscalYearEndBasic: timestamppb.New(time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)),
				PriceRangeGroup:    "10003",
				KabucomMarginBuy:   true,
				KabucomMarginSell:  true,
				MarginBuy:          true,
				MarginSell:         true,
				UpperLimit:         4041,
				LowerLimit:         2641,
				Underlyer:          "NK225",
				DerivativeMonth:    timestamppb.New(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)),
				TradeStart:         timestamppb.New(time.Date(2015, 12, 11, 0, 0, 0, 0, time.Local)),
				TradeEnd:           timestamppb.New(time.Date(2020, 12, 10, 0, 0, 0, 0, time.Local)),
				StrikePrice:        23250,
				CallOrPut:          kabuspb.CallPut_CALL_PUT_CALL,
				ClearingPrice:      23000,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{symbolWithContext1: test.symbolWithContext1, symbolWithContext2: test.symbolWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.Symbol(context.Background(), "", &kabuspb.GetSymbolRequest{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

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
	tokenWithContext1              *kabus.TokenResponse
	tokenWithContext2              error
	symbolNameFutureWithContext1   *kabus.SymbolNameFutureResponse
	symbolNameFutureWithContext2   error
	symbolNameOptionWithContext1   *kabus.SymbolNameOptionResponse
	symbolNameOptionWithContext2   error
	registerWithContext1           *kabus.RegisterResponse
	registerWithContext2           error
	lastRegisterWithContext        kabus.RegisterRequest
	unregisterWithContext1         *kabus.UnregisterResponse
	unregisterWithContext2         error
	lastUnregisterWithContext      kabus.UnregisterRequest
	unregisterAllWithContext1      *kabus.UnregisterAllResponse
	unregisterAllWithContext2      error
	boardWithContext1              *kabus.BoardResponse
	boardWithContext2              error
	symbolWithContext1             *kabus.SymbolResponse
	symbolWithContext2             error
	ordersWithContext1             *kabus.OrdersResponse
	ordersWithContext2             error
	positionsWithContext1          *kabus.PositionsResponse
	positionsWithContext2          error
	rankingWithContext1            *kabus.RankingResponse
	rankingWithContext2            error
	sendOrderStockWithContext1     *kabus.SendOrderStockResponse
	sendOrderStockWithContext2     error
	sendOrderFutureWithContext1    *kabus.SendOrderFutureResponse
	sendOrderFutureWithContext2    error
	sendOrderOptionWithContext1    *kabus.SendOrderOptionResponse
	sendOrderOptionWithContext2    error
	cancelOrderWithContext1        *kabus.CancelOrderResponse
	cancelOrderWithContext2        error
	walletCashWithContext1         *kabus.WalletCashResponse
	walletCashWithContext2         error
	walletCashSymbolWithContext1   *kabus.WalletCashResponse
	walletCashSymbolWithContext2   error
	walletMarginWithContext1       *kabus.WalletMarginResponse
	walletMarginWithContext2       error
	walletMarginSymbolWithContext1 *kabus.WalletMarginResponse
	walletMarginSymbolWithContext2 error
	walletFutureWithContext1       *kabus.WalletFutureResponse
	walletFutureWithContext2       error
	walletFutureSymbolWithContext1 *kabus.WalletFutureResponse
	walletFutureSymbolWithContext2 error
	walletOptionWithContext1       *kabus.WalletOptionResponse
	walletOptionWithContext2       error
	walletOptionSymbolWithContext1 *kabus.WalletOptionResponse
	walletOptionSymbolWithContext2 error
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

func (t *testRESTClient) OrdersWithContext(context.Context, string, kabus.OrdersRequest) (*kabus.OrdersResponse, error) {
	return t.ordersWithContext1, t.ordersWithContext2
}

func (t *testRESTClient) PositionsWithContext(context.Context, string, kabus.PositionsRequest) (*kabus.PositionsResponse, error) {
	return t.positionsWithContext1, t.positionsWithContext2
}

func (t *testRESTClient) RankingWithContext(context.Context, string, kabus.RankingRequest) (*kabus.RankingResponse, error) {
	return t.rankingWithContext1, t.rankingWithContext2
}

func (t *testRESTClient) SendOrderStockWithContext(context.Context, string, kabus.SendOrderStockRequest) (*kabus.SendOrderStockResponse, error) {
	return t.sendOrderStockWithContext1, t.sendOrderStockWithContext2
}

func (t *testRESTClient) SendOrderFutureWithContext(context.Context, string, kabus.SendOrderFutureRequest) (*kabus.SendOrderFutureResponse, error) {
	return t.sendOrderFutureWithContext1, t.sendOrderFutureWithContext2
}

func (t *testRESTClient) SendOrderOptionWithContext(context.Context, string, kabus.SendOrderOptionRequest) (*kabus.SendOrderOptionResponse, error) {
	return t.sendOrderOptionWithContext1, t.sendOrderOptionWithContext2
}

func (t *testRESTClient) CancelOrderWithContext(context.Context, string, kabus.CancelOrderRequest) (*kabus.CancelOrderResponse, error) {
	return t.cancelOrderWithContext1, t.cancelOrderWithContext2
}

func (t *testRESTClient) WalletCashWithContext(context.Context, string) (*kabus.WalletCashResponse, error) {
	return t.walletCashWithContext1, t.walletCashWithContext2
}

func (t *testRESTClient) WalletCashSymbolWithContext(context.Context, string, kabus.WalletCashSymbolRequest) (*kabus.WalletCashResponse, error) {
	return t.walletCashSymbolWithContext1, t.walletCashSymbolWithContext2
}

func (t *testRESTClient) WalletMarginWithContext(context.Context, string) (*kabus.WalletMarginResponse, error) {
	return t.walletMarginWithContext1, t.walletMarginWithContext2
}

func (t *testRESTClient) WalletMarginSymbolWithContext(context.Context, string, kabus.WalletMarginSymbolRequest) (*kabus.WalletMarginResponse, error) {
	return t.walletMarginSymbolWithContext1, t.walletMarginSymbolWithContext2
}

func (t *testRESTClient) WalletFutureWithContext(context.Context, string) (*kabus.WalletFutureResponse, error) {
	return t.walletFutureWithContext1, t.walletFutureWithContext2
}

func (t *testRESTClient) WalletFutureSymbolWithContext(context.Context, string, kabus.WalletFutureSymbolRequest) (*kabus.WalletFutureResponse, error) {
	return t.walletFutureSymbolWithContext1, t.walletFutureSymbolWithContext2
}

func (t *testRESTClient) WalletOptionWithContext(context.Context, string) (*kabus.WalletOptionResponse, error) {
	return t.walletOptionWithContext1, t.walletOptionWithContext2
}

func (t *testRESTClient) WalletOptionSymbolWithContext(context.Context, string, kabus.WalletOptionSymbolRequest) (*kabus.WalletOptionResponse, error) {
	return t.walletOptionSymbolWithContext1, t.walletOptionSymbolWithContext2
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
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			hasError:             true,
			wantReq:              kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			registerWithContext1: &kabus.RegisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:               &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:                 &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
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
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			hasError:               true,
			wantReq:                kabus.UnregisterRequest{Symbols: []kabus.UnregisterSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}}},
		{name: "responseが返されたらresponseを変換して返す",
			unregisterWithContext1: &kabus.UnregisterResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			argReq:                 &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:                   &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
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
			unregisterAllWithContext1: &kabus.UnregisterAllResponse{RegisterList: []kabus.RegisteredSymbol{{Symbol: "1234", Exchange: kabus.ExchangeToushou}}},
			want:                      &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}}},
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

func Test_security_Orders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		ordersWithContext1 *kabus.OrdersResponse
		ordersWithContext2 error
		want               *kabuspb.Orders
		hasError           bool
	}{
		{name: "errorを返されたらerrorを返す", ordersWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			ordersWithContext1: &kabus.OrdersResponse{{
				ID:              "20210331A02N36008375",
				State:           kabus.StateDone,
				OrderState:      kabus.OrderStateDone,
				OrdType:         kabus.OrdTypeInTrading,
				RecvTime:        time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local),
				Symbol:          "1475",
				SymbolName:      "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:        kabus.OrderExchangeToushou,
				ExchangeName:    "東証ETF/ETN",
				TimeInForce:     kabus.TimeInForceUnspecified,
				Price:           0,
				OrderQty:        1,
				CumQty:          1,
				Side:            kabus.SideBuy,
				CashMargin:      kabus.CashMarginUnspecified,
				AccountType:     kabus.AccountTypeSpecific,
				DelivType:       kabus.DelivTypeCash,
				ExpireDay:       kabus.YmdNUM{Time: time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)},
				MarginTradeType: kabus.MarginTradeTypeUnspecified,
				Details:         []kabus.OrderDetail{},
			}},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{{
				Id:                 "20210331A02N36008375",
				State:              kabuspb.State_STATE_DONE,
				OrderState:         kabuspb.OrderState_ORDER_STATE_DONE,
				OrderType:          kabuspb.OrderType_ORDER_TYPE_ZARABA,
				ReceiveTime:        timestamppb.New(time.Date(2021, 3, 31, 11, 28, 19, 398248000, time.Local)),
				SymbolCode:         "1475",
				SymbolName:         "ｉシェアーズ・コア　ＴＯＰＩＸ　ＥＴＦ",
				Exchange:           kabuspb.OrderExchange_ORDER_EXCHANGE_TOUSHOU,
				ExchangeName:       "東証ETF/ETN",
				TimeInForce:        kabuspb.TimeInForce_TIME_IN_FORCE_UNSPECIFIED,
				Price:              0,
				OrderQuantity:      1,
				CumulativeQuantity: 1,
				Side:               kabuspb.Side_SIDE_BUY,
				TradeType:          kabuspb.TradeType_TRADE_TYPE_UNSPECIFIED,
				AccountType:        kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				DeliveryType:       kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
				ExpireDay:          timestamppb.New(time.Date(2021, 3, 31, 0, 0, 0, 0, time.Local)),
				MarginTradeType:    kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_UNSPECIFIED,
				Details:            []*kabuspb.OrderDetail{},
			}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{ordersWithContext1: test.ordersWithContext1, ordersWithContext2: test.ordersWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.Orders(context.Background(), "", &kabuspb.GetOrdersRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_Positions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		positionsWithContext1 *kabus.PositionsResponse
		positionsWithContext2 error
		want                  *kabuspb.Positions
		hasError              bool
	}{
		{name: "errorを返されたらerrorを返す", positionsWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			positionsWithContext1: &kabus.PositionsResponse{{
				ExecutionID:     "20200715E02N04738464",
				AccountType:     kabus.AccountTypeSpecific,
				Symbol:          "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabus.ExchangeToushou,
				ExchangeName:    "東証１部",
				SecurityType:    kabus.SecurityTypeNK225,
				ExecutionDay:    kabus.NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQty:       500,
				HoldQty:         0,
				Side:            kabus.SideSell,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       kabus.NewYmdNUM(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabus.MarginTradeTypeSystem,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}},
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{{
				ExecutionId:     "20200715E02N04738464",
				AccountType:     kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SymbolCode:      "8306",
				SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
				Exchange:        kabuspb.Exchange_EXCHANGE_TOUSHOU,
				ExchangeName:    "東証１部",
				SecurityType:    kabuspb.SecurityType_SECURITY_TYPE_NK225,
				ExecutionDay:    timestamppb.New(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
				Price:           704,
				LeavesQuantity:  500,
				HoldQuantity:    0,
				Side:            kabuspb.Side_SIDE_SELL,
				Expenses:        0,
				Commission:      1620,
				CommissionTax:   162,
				ExpireDay:       timestamppb.New(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
				MarginTradeType: kabuspb.MarginTradeType_MARGIN_TRADE_TYPE_SYSTEM,
				CurrentPrice:    414.5,
				Valuation:       207250,
				ProfitLoss:      144750,
				ProfitLossRate:  41.12215909090909,
			}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{positionsWithContext1: test.positionsWithContext1, positionsWithContext2: test.positionsWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.Positions(context.Background(), "", &kabuspb.GetPositionsRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_PriceRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.PriceRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRate,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				PriceRanking: []kabus.PriceRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1689", SymbolName: "ガスETF/ETF(C)", CurrentPrice: 2, ChangeRatio: 1, ChangePercentage: 100, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 5722.4, Turnover: 10.4136, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, ChangePercentage: 54.65, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 3117.5, Turnover: 3194.7121, ExchangeName: "東証JQS", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.PriceRanking{
				Type:             kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.PriceRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1689", SymbolName: "ガスETF/ETF(C)", CurrentPrice: 2, ChangeRatio: 1, ChangePercentage: 100, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), TradingVolume: 5722.4, Turnover: 10.4136, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, ChangePercentage: 54.65, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), TradingVolume: 3117.5, Turnover: 3194.7121, ExchangeName: "東証JQS", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.PriceRanking(context.Background(), "", &kabuspb.GetPriceRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_TickRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.TickRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypeTickCount,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				TickRanking: []kabus.TickRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 22, Symbol: "2929", SymbolName: "ﾌｧｰﾏﾌｰｽﾞ", CurrentPrice: 2748, ChangeRatio: 99, TickCount: 40579, UpCount: 12722, DownCount: 12798, ChangePercentage: 3.73, TradingVolume: 16086.8, Turnover: 43810.0498, ExchangeName: "東証２部", CategoryName: "食料品"},
					{No: 2, Trend: kabus.RankingTrendUnchanged, AverageRanking: 2, Symbol: "9984", SymbolName: "ｿﾌﾄﾊﾞﾝｸG", CurrentPrice: 8285, ChangeRatio: -309, TickCount: 32219, UpCount: 8655, DownCount: 8562, ChangePercentage: -3.59, TradingVolume: 16688.8, Turnover: 138143.1773, ExchangeName: "東証１部", CategoryName: "情報・通信業"},
				},
			},
			want: &kabuspb.TickRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.TickRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 22, SymbolCode: "2929", SymbolName: "ﾌｧｰﾏﾌｰｽﾞ", CurrentPrice: 2748, ChangeRatio: 99, TickCount: 40579, UpCount: 12722, DownCount: 12798, ChangePercentage: 3.73, TradingVolume: 16086.8, Turnover: 43810.0498, ExchangeName: "東証２部", IndustryName: "食料品"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_NO_CHANGE, AverageRanking: 2, SymbolCode: "9984", SymbolName: "ｿﾌﾄﾊﾞﾝｸG", CurrentPrice: 8285, ChangeRatio: -309, TickCount: 32219, UpCount: 8655, DownCount: 8562, ChangePercentage: -3.59, TradingVolume: 16688.8, Turnover: 138143.1773, ExchangeName: "東証１部", IndustryName: "情報・通信業"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.TickRanking(context.Background(), "", &kabuspb.GetTickRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_VolumeRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.VolumeRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypeVolumeRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				VolumeRapidRanking: []kabus.VolumeRapidRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidTradePercentage: 49900, TradingVolume: 1, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidTradePercentage: 28189.47, TradingVolume: 3117.5, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.VolumeRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.VolumeRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidTradePercentage: 49900, TradingVolume: 1, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)), ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidTradePercentage: 28189.47, TradingVolume: 3117.5, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 54.65, ExchangeName: "東証JQS", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.VolumeRanking(context.Background(), "", &kabuspb.GetVolumeRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_ValueRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.ValueRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypeValueRapidIncrease,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				ValueRapidRanking: []kabus.ValueRapidRanking{
					{No: 1, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidPaymentPercentage: 55381.47, Turnover: 3194.7121, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
					{No: 2, Trend: kabus.RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidPaymentPercentage: 50159.4, Turnover: 7.75, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
				},
			},
			want: &kabuspb.ValueRanking{
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.ValueRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidPaymentPercentage: 55381.47, Turnover: 3194.7121, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 54.65, ExchangeName: "東証JQS", IndustryName: "電気機器"},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE_OVER_20, AverageRanking: 999, SymbolCode: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidPaymentPercentage: 50159.4, Turnover: 7.75, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)), ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", IndustryName: "その他"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.ValueRanking(context.Background(), "", &kabuspb.GetValueRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_MarginRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.MarginRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypeMarginHighMagnification,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				MarginRanking: []kabus.MarginRanking{
					{No: 1, Symbol: "3150", SymbolName: "グリムス", Ratio: 14467, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.5, BuyRapidPaymentPercentage: 1446.7, BuyLastWeekRatio: 139.7, ExchangeName: "東証１部", CategoryName: "卸売業"},
					{No: 2, Symbol: "6955", SymbolName: "ＦＤＫ", Ratio: 10536.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: -0.8, BuyRapidPaymentPercentage: 2107.3, BuyLastWeekRatio: 121.6, ExchangeName: "東証２部", CategoryName: "電気機器"},
				},
			},
			want: &kabuspb.MarginRanking{
				Type:             kabuspb.MarginRankingType_MARGIN_RANKING_TYPE_HIGH_MAGNIFICATION,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.MarginRankingInfo{
					{No: 1, SymbolCode: "3150", SymbolName: "グリムス", Ratio: 14467, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.5, BuyRapidPaymentPercentage: 1446.7, BuyLastWeekRatio: 139.7, ExchangeName: "東証１部", IndustryName: "卸売業"},
					{No: 2, SymbolCode: "6955", SymbolName: "ＦＤＫ", Ratio: 10536.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: -0.8, BuyRapidPaymentPercentage: 2107.3, BuyLastWeekRatio: 121.6, ExchangeName: "東証２部", IndustryName: "電気機器"},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.MarginRanking(context.Background(), "", &kabuspb.GetMarginRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_IndustryRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		rankingWithContext1 *kabus.RankingResponse
		rankingWithContext2 error
		want                *kabuspb.IndustryRanking
		hasError            bool
	}{
		{name: "errorを返されたらerrorを返す", rankingWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			rankingWithContext1: &kabus.RankingResponse{
				Type:             kabus.RankingTypePriceIncreaseRateByCategory,
				ExchangeDivision: kabus.ExchangeDivisionALL,
				CategoryPriceRanking: []kabus.CategoryPriceRanking{
					{No: 1, Trend: kabus.RankingTrendRise, AverageRanking: 18, Category: "343", CategoryName: "IS 空運", CurrentPrice: 170.97, ChangeRatio: 6.72, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 4.09},
					{No: 2, Trend: kabus.RankingTrendRise, AverageRanking: 16, Category: "341", CategoryName: "IS 陸運", CurrentPrice: 1895.49, ChangeRatio: 15.41, CurrentPriceTime: kabus.HmString{Time: time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.82},
				},
			},
			want: &kabuspb.IndustryRanking{
				Type:             kabuspb.IndustryRankingType_INDUSTRY_RANKING_TYPE_INCREASE_RATE,
				ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL,
				Ranking: []*kabuspb.IndustryRankingInfo{
					{No: 1, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE, AverageRanking: 18, IndustryCode: "343", IndustryName: "IS 空運", CurrentPrice: 170.97, ChangeRatio: 6.72, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 4.09},
					{No: 2, Trend: kabuspb.RankingTrend_RANKING_TREND_RISE, AverageRanking: 16, IndustryCode: "341", IndustryName: "IS 陸運", CurrentPrice: 1895.49, ChangeRatio: 15.41, CurrentPriceTime: timestamppb.New(time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)), ChangePercentage: 0.82},
				}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{rankingWithContext1: test.rankingWithContext1, rankingWithContext2: test.rankingWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.IndustryRanking(context.Background(), "", &kabuspb.GetIndustryRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_SendOrderStock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                       string
		sendOrderStockWithContext1 *kabus.SendOrderStockResponse
		sendOrderStockWithContext2 error
		want                       *kabuspb.OrderResponse
		hasError                   bool
	}{
		{name: "errorを返されたらerrorを返す", sendOrderStockWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			sendOrderStockWithContext1: &kabus.SendOrderStockResponse{Result: 0, OrderID: "ORDER-ID"},
			want:                       &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{sendOrderStockWithContext1: test.sendOrderStockWithContext1, sendOrderStockWithContext2: test.sendOrderStockWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SendOrderStock(context.Background(), "", &kabuspb.SendStockOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_SendOrderMargin(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                       string
		sendOrderStockWithContext1 *kabus.SendOrderStockResponse
		sendOrderStockWithContext2 error
		want                       *kabuspb.OrderResponse
		hasError                   bool
	}{
		{name: "errorを返されたらerrorを返す", sendOrderStockWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			sendOrderStockWithContext1: &kabus.SendOrderStockResponse{Result: 0, OrderID: "ORDER-ID"},
			want:                       &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{sendOrderStockWithContext1: test.sendOrderStockWithContext1, sendOrderStockWithContext2: test.sendOrderStockWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SendOrderMargin(context.Background(), "", &kabuspb.SendMarginOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_SendOrderFuture(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                        string
		sendOrderFutureWithContext1 *kabus.SendOrderFutureResponse
		sendOrderFutureWithContext2 error
		want                        *kabuspb.OrderResponse
		hasError                    bool
	}{
		{name: "errorを返されたらerrorを返す", sendOrderFutureWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			sendOrderFutureWithContext1: &kabus.SendOrderFutureResponse{Result: 0, OrderID: "ORDER-ID"},
			want:                        &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{sendOrderFutureWithContext1: test.sendOrderFutureWithContext1, sendOrderFutureWithContext2: test.sendOrderFutureWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SendOrderFuture(context.Background(), "", &kabuspb.SendFutureOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_SendOrderOption(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                        string
		sendOrderOptionWithContext1 *kabus.SendOrderOptionResponse
		sendOrderOptionWithContext2 error
		want                        *kabuspb.OrderResponse
		hasError                    bool
	}{
		{name: "errorを返されたらerrorを返す", sendOrderOptionWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			sendOrderOptionWithContext1: &kabus.SendOrderOptionResponse{Result: 0, OrderID: "ORDER-ID"},
			want:                        &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{sendOrderOptionWithContext1: test.sendOrderOptionWithContext1, sendOrderOptionWithContext2: test.sendOrderOptionWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.SendOrderOption(context.Background(), "", &kabuspb.SendOptionOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_CancelOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		cancelOrderWithContext1 *kabus.CancelOrderResponse
		cancelOrderWithContext2 error
		want                    *kabuspb.OrderResponse
		hasError                bool
	}{
		{name: "errorを返されたらerrorを返す", cancelOrderWithContext2: errors.New("error message"), hasError: true},
		{name: "responseが返されたらresponseを変換して返す",
			cancelOrderWithContext1: &kabus.CancelOrderResponse{Result: 0, OrderID: "ORDER-ID"},
			want:                    &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{cancelOrderWithContext1: test.cancelOrderWithContext1, cancelOrderWithContext2: test.cancelOrderWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.CancelOrder(context.Background(), "", &kabuspb.CancelOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_GetStockWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                         string
		req                          *kabuspb.GetStockWalletRequest
		walletCashWithContext1       *kabus.WalletCashResponse
		walletCashWithContext2       error
		walletCashSymbolWithContext1 *kabus.WalletCashResponse
		walletCashSymbolWithContext2 error
		want                         *kabuspb.StockWallet
		hasError                     bool
	}{
		{name: "symbolが空でerrorを返されたらerrorを返す", req: &kabuspb.GetStockWalletRequest{}, walletCashWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空じゃなくてerrorを返されたらerrorを返す", req: &kabuspb.GetStockWalletRequest{SymbolCode: "1320"}, walletCashSymbolWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空でresponseが返されたらresponseを変換して返す",
			req:                    &kabuspb.GetStockWalletRequest{},
			walletCashWithContext1: &kabus.WalletCashResponse{StockAccountWallet: 30000},
			want:                   &kabuspb.StockWallet{StockAccountWallet: 30000}},
		{name: "symbolが空じゃなくてresponseが返されたらresponseを変換して返す",
			req:                          &kabuspb.GetStockWalletRequest{SymbolCode: "1320"},
			walletCashSymbolWithContext1: &kabus.WalletCashResponse{StockAccountWallet: 30000},
			want:                         &kabuspb.StockWallet{StockAccountWallet: 30000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{
				walletCashWithContext1: test.walletCashWithContext1, walletCashWithContext2: test.walletCashWithContext2,
				walletCashSymbolWithContext1: test.walletCashSymbolWithContext1, walletCashSymbolWithContext2: test.walletCashSymbolWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.GetStockWallet(context.Background(), "", test.req)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_GetMarginWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                           string
		req                            *kabuspb.GetMarginWalletRequest
		walletMarginWithContext1       *kabus.WalletMarginResponse
		walletMarginWithContext2       error
		walletMarginSymbolWithContext1 *kabus.WalletMarginResponse
		walletMarginSymbolWithContext2 error
		want                           *kabuspb.MarginWallet
		hasError                       bool
	}{
		{name: "symbolが空でerrorを返されたらerrorを返す", req: &kabuspb.GetMarginWalletRequest{}, walletMarginWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空じゃなくてerrorを返されたらerrorを返す", req: &kabuspb.GetMarginWalletRequest{SymbolCode: "1320"}, walletMarginSymbolWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空でresponseが返されたらresponseを変換して返す",
			req:                      &kabuspb.GetMarginWalletRequest{},
			walletMarginWithContext1: &kabus.WalletMarginResponse{MarginAccountWallet: 30000},
			want:                     &kabuspb.MarginWallet{MarginAccountWallet: 30000}},
		{name: "symbolが空じゃなくてresponseが返されたらresponseを変換して返す",
			req:                            &kabuspb.GetMarginWalletRequest{SymbolCode: "1320"},
			walletMarginSymbolWithContext1: &kabus.WalletMarginResponse{MarginAccountWallet: 30000},
			want:                           &kabuspb.MarginWallet{MarginAccountWallet: 30000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{
				walletMarginWithContext1: test.walletMarginWithContext1, walletMarginWithContext2: test.walletMarginWithContext2,
				walletMarginSymbolWithContext1: test.walletMarginSymbolWithContext1, walletMarginSymbolWithContext2: test.walletMarginSymbolWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.GetMarginWallet(context.Background(), "", test.req)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_GetFutureWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                           string
		req                            *kabuspb.GetFutureWalletRequest
		walletFutureWithContext1       *kabus.WalletFutureResponse
		walletFutureWithContext2       error
		walletFutureSymbolWithContext1 *kabus.WalletFutureResponse
		walletFutureSymbolWithContext2 error
		want                           *kabuspb.FutureWallet
		hasError                       bool
	}{
		{name: "symbolが空でerrorを返されたらerrorを返す", req: &kabuspb.GetFutureWalletRequest{}, walletFutureWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空じゃなくてerrorを返されたらerrorを返す", req: &kabuspb.GetFutureWalletRequest{SymbolCode: "1320"}, walletFutureSymbolWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空でresponseが返されたらresponseを変換して返す",
			req:                      &kabuspb.GetFutureWalletRequest{},
			walletFutureWithContext1: &kabus.WalletFutureResponse{FutureTradeLimit: 300000, MarginRequirement: 0},
			want:                     &kabuspb.FutureWallet{FutureTradeLimit: 300000, MarginRequirement: 0}},
		{name: "symbolが空じゃなくてresponseが返されたらresponseを変換して返す",
			req:                            &kabuspb.GetFutureWalletRequest{SymbolCode: "1320"},
			walletFutureSymbolWithContext1: &kabus.WalletFutureResponse{FutureTradeLimit: 900000, MarginRequirement: 300000},
			want:                           &kabuspb.FutureWallet{FutureTradeLimit: 900000, MarginRequirement: 300000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{
				walletFutureWithContext1: test.walletFutureWithContext1, walletFutureWithContext2: test.walletFutureWithContext2,
				walletFutureSymbolWithContext1: test.walletFutureSymbolWithContext1, walletFutureSymbolWithContext2: test.walletFutureSymbolWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.GetFutureWallet(context.Background(), "", test.req)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_security_GetOptionWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                           string
		req                            *kabuspb.GetOptionWalletRequest
		walletOptionWithContext1       *kabus.WalletOptionResponse
		walletOptionWithContext2       error
		walletOptionSymbolWithContext1 *kabus.WalletOptionResponse
		walletOptionSymbolWithContext2 error
		want                           *kabuspb.OptionWallet
		hasError                       bool
	}{
		{name: "symbolが空でerrorを返されたらerrorを返す", req: &kabuspb.GetOptionWalletRequest{}, walletOptionWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空じゃなくてerrorを返されたらerrorを返す", req: &kabuspb.GetOptionWalletRequest{SymbolCode: "1320"}, walletOptionSymbolWithContext2: errors.New("error message"), hasError: true},
		{name: "symbolが空でresponseが返されたらresponseを変換して返す",
			req:                      &kabuspb.GetOptionWalletRequest{},
			walletOptionWithContext1: &kabus.WalletOptionResponse{OptionBuyTradeLimit: 300000, OptionSellTradeLimit: 300000, MarginRequirement: 0},
			want:                     &kabuspb.OptionWallet{OptionBuyTradeLimit: 300000, OptionSellTradeLimit: 300000, MarginRequirement: 0}},
		{name: "symbolが空じゃなくてresponseが返されたらresponseを変換して返す",
			req:                            &kabuspb.GetOptionWalletRequest{SymbolCode: "1320"},
			walletOptionSymbolWithContext1: &kabus.WalletOptionResponse{OptionBuyTradeLimit: 900000, OptionSellTradeLimit: 900000, MarginRequirement: 300000},
			want:                           &kabuspb.OptionWallet{OptionBuyTradeLimit: 900000, OptionSellTradeLimit: 900000, MarginRequirement: 300000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			restClient := &testRESTClient{
				walletOptionWithContext1: test.walletOptionWithContext1, walletOptionWithContext2: test.walletOptionWithContext2,
				walletOptionSymbolWithContext1: test.walletOptionSymbolWithContext1, walletOptionSymbolWithContext2: test.walletOptionSymbolWithContext2}
			security := &security{restClient: restClient}
			got1, got2 := security.GetOptionWallet(context.Background(), "", test.req)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

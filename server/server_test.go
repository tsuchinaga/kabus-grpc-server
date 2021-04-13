package server

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

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
	board1            *kabuspb.Board
	board2            error
	symbol1           *kabuspb.Symbol
	symbol2           error
	orders1           *kabuspb.Orders
	orders2           error
	positions1        *kabuspb.Positions
	positions2        error
	priceRanking1     *kabuspb.PriceRanking
	priceRanking2     error
	tickRanking1      *kabuspb.TickRanking
	tickRanking2      error
	volumeRanking1    *kabuspb.VolumeRanking
	volumeRanking2    error
	valueRanking1     *kabuspb.ValueRanking
	valueRanking2     error
	marginRanking1    *kabuspb.MarginRanking
	marginRanking2    error
	industryRanking1  *kabuspb.IndustryRanking
	industryRanking2  error
	sendOrderStock1   *kabuspb.OrderResponse
	sendOrderStock2   error
	sendOrderMargin1  *kabuspb.OrderResponse
	sendOrderMargin2  error
	sendOrderFuture1  *kabuspb.OrderResponse
	sendOrderFuture2  error
	sendOrderOption1  *kabuspb.OrderResponse
	sendOrderOption2  error
	cancelOrder1      *kabuspb.OrderResponse
	cancelOrder2      error
	getStockWallet1   *kabuspb.StockWallet
	getStockWallet2   error
	getMarginWallet1  *kabuspb.MarginWallet
	getMarginWallet2  error
	getFutureWallet1  *kabuspb.FutureWallet
	getFutureWallet2  error
	getOptionWallet1  *kabuspb.OptionWallet
	getOptionWallet2  error
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

func (t *testSecurity) Board(context.Context, string, *kabuspb.GetBoardRequest) (*kabuspb.Board, error) {
	return t.board1, t.board2
}

func (t *testSecurity) Symbol(context.Context, string, *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error) {
	return t.symbol1, t.symbol2
}

func (t *testSecurity) Orders(context.Context, string, *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	return t.orders1, t.orders2
}

func (t *testSecurity) Positions(context.Context, string, *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	return t.positions1, t.positions2
}

func (t *testSecurity) PriceRanking(context.Context, string, *kabuspb.GetPriceRankingRequest) (*kabuspb.PriceRanking, error) {
	return t.priceRanking1, t.priceRanking2
}

func (t *testSecurity) TickRanking(context.Context, string, *kabuspb.GetTickRankingRequest) (*kabuspb.TickRanking, error) {
	return t.tickRanking1, t.tickRanking2
}

func (t *testSecurity) VolumeRanking(context.Context, string, *kabuspb.GetVolumeRankingRequest) (*kabuspb.VolumeRanking, error) {
	return t.volumeRanking1, t.volumeRanking2
}

func (t *testSecurity) ValueRanking(context.Context, string, *kabuspb.GetValueRankingRequest) (*kabuspb.ValueRanking, error) {
	return t.valueRanking1, t.valueRanking2
}

func (t *testSecurity) MarginRanking(context.Context, string, *kabuspb.GetMarginRankingRequest) (*kabuspb.MarginRanking, error) {
	return t.marginRanking1, t.marginRanking2
}

func (t *testSecurity) IndustryRanking(context.Context, string, *kabuspb.GetIndustryRankingRequest) (*kabuspb.IndustryRanking, error) {
	return t.industryRanking1, t.industryRanking2
}

func (t *testSecurity) SendOrderStock(context.Context, string, *kabuspb.SendStockOrderRequest, string) (*kabuspb.OrderResponse, error) {
	return t.sendOrderStock1, t.sendOrderStock2
}

func (t *testSecurity) SendOrderMargin(context.Context, string, *kabuspb.SendMarginOrderRequest, string) (*kabuspb.OrderResponse, error) {
	return t.sendOrderMargin1, t.sendOrderMargin2
}

func (t *testSecurity) SendOrderFuture(context.Context, string, *kabuspb.SendFutureOrderRequest, string) (*kabuspb.OrderResponse, error) {
	return t.sendOrderFuture1, t.sendOrderFuture2
}

func (t *testSecurity) SendOrderOption(context.Context, string, *kabuspb.SendOptionOrderRequest, string) (*kabuspb.OrderResponse, error) {
	return t.sendOrderOption1, t.sendOrderOption2
}

func (t *testSecurity) CancelOrder(context.Context, string, *kabuspb.CancelOrderRequest, string) (*kabuspb.OrderResponse, error) {
	return t.cancelOrder1, t.cancelOrder2
}

func (t *testSecurity) GetStockWallet(context.Context, string, *kabuspb.GetStockWalletRequest) (*kabuspb.StockWallet, error) {
	return t.getStockWallet1, t.getStockWallet2
}

func (t *testSecurity) GetMarginWallet(context.Context, string, *kabuspb.GetMarginWalletRequest) (*kabuspb.MarginWallet, error) {
	return t.getMarginWallet1, t.getMarginWallet2
}

func (t *testSecurity) GetFutureWallet(context.Context, string, *kabuspb.GetFutureWalletRequest) (*kabuspb.FutureWallet, error) {
	return t.getFutureWallet1, t.getFutureWallet2
}

func (t *testSecurity) GetOptionWallet(context.Context, string, *kabuspb.GetOptionWalletRequest) (*kabuspb.OptionWallet, error) {
	return t.getOptionWallet1, t.getOptionWallet2
}

type testTokenService struct {
	services.TokenService
	getToken1    string
	getToken2    error
	refresh1     string
	refresh2     error
	getExpiredAt time.Time
}

func (t *testTokenService) GetToken(context.Context) (string, error) { return t.getToken1, t.getToken2 }
func (t *testTokenService) GetExpiredAt() time.Time                  { return t.getExpiredAt }
func (t *testTokenService) Refresh(context.Context) (string, error)  { return t.refresh1, t.refresh2 }

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

type testSetting struct {
	repositories.Setting
	password string
}

func (t *testSetting) Password() string { return t.password }

type testBoardStreamService struct {
	services.BoardStreamService
	connect error
}

func (t *testBoardStreamService) Connect(kabuspb.KabusService_GetBoardsStreamingServer) error {
	return t.connect
}

func Test_NewServer(t *testing.T) {
	security := &testSecurity{}
	tokenService := &testTokenService{}
	registerSymbolService := &testRegisterSymbolService{}
	setting := &testSetting{}
	boardStreamService := &testBoardStreamService{}
	got := NewServer(security, tokenService, registerSymbolService, setting, boardStreamService)
	want := &server{security: security, tokenService: tokenService, registerSymbolService: registerSymbolService, setting: setting, boardStreamService: boardStreamService}
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
		{name: "SymbolNameFutureの結果を結果を返す",
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
		{name: "SymbolNameOptionの結果を結果を返す",
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

func Test_server_GetBoard(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		board1    *kabuspb.Board
		board2    error
		want      *kabuspb.Board
		hasError  bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Boardでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			board2:    errors.New("register error message"),
			hasError:  true},
		{name: "Boardの結果を結果を返す",
			getToken1: "TOKEN_STRING",
			board1:    &kabuspb.Board{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
			want:      &kabuspb.Board{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{board1: test.board1, board2: test.board2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetBoard(context.Background(), &kabuspb.GetBoardRequest{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetSymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		symbol1   *kabuspb.Symbol
		symbol2   error
		want      *kabuspb.Symbol
		hasError  bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Symbolでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			symbol2:   errors.New("register error message"),
			hasError:  true},
		{name: "Symbolの結果を結果を返す",
			getToken1: "TOKEN_STRING",
			symbol1:   &kabuspb.Symbol{Code: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU},
			want:      &kabuspb.Symbol{Code: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{symbol1: test.symbol1, symbol2: test.symbol2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetSymbol(context.Background(), &kabuspb.GetSymbolRequest{SymbolCode: "5401", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetOrders(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		orders1   *kabuspb.Orders
		orders2   error
		want      *kabuspb.Orders
		hasError  bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Ordersでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			orders2:   errors.New("register error message"),
			hasError:  true},
		{name: "Ordersの結果を結果を返す",
			getToken1: "TOKEN_STRING",
			orders1:   &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}},
			want:      &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{orders1: test.orders1, orders2: test.orders2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetOrders(context.Background(), &kabuspb.GetOrdersRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetPositions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		getToken1  string
		getToken2  error
		positions1 *kabuspb.Positions
		positions2 error
		want       *kabuspb.Positions
		hasError   bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Positionsでエラーがあればエラーを返す",
			getToken1:  "TOKEN_STRING",
			positions2: errors.New("register error message"),
			hasError:   true},
		{name: "Positionsの結果を結果を返す",
			getToken1:  "TOKEN_STRING",
			positions1: &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}},
			want:       &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{positions1: test.positions1, positions2: test.positions2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetPositions(context.Background(), &kabuspb.GetPositionsRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetPriceRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		getToken1     string
		getToken2     error
		priceRanking1 *kabuspb.PriceRanking
		priceRanking2 error
		want          *kabuspb.PriceRanking
		hasError      bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "PriceRankingでエラーがあればエラーを返す",
			getToken1:     "TOKEN_STRING",
			priceRanking2: errors.New("register error message"),
			hasError:      true},
		{name: "PriceRankingの結果を結果を返す",
			getToken1:     "TOKEN_STRING",
			priceRanking1: &kabuspb.PriceRanking{Type: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE},
			want:          &kabuspb.PriceRanking{Type: kabuspb.PriceRankingType_PRICE_RANKING_TYPE_INCREASE_RATE}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{priceRanking1: test.priceRanking1, priceRanking2: test.priceRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetPriceRanking(context.Background(), &kabuspb.GetPriceRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetTickRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		getToken1    string
		getToken2    error
		tickRanking1 *kabuspb.TickRanking
		tickRanking2 error
		want         *kabuspb.TickRanking
		hasError     bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "TickRankingでエラーがあればエラーを返す",
			getToken1:    "TOKEN_STRING",
			tickRanking2: errors.New("register error message"),
			hasError:     true},
		{name: "TickRankingの結果を結果を返す",
			getToken1:    "TOKEN_STRING",
			tickRanking1: &kabuspb.TickRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
			want:         &kabuspb.TickRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{tickRanking1: test.tickRanking1, tickRanking2: test.tickRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetTickRanking(context.Background(), &kabuspb.GetTickRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetVolumeRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		getToken1      string
		getToken2      error
		volumeRanking1 *kabuspb.VolumeRanking
		volumeRanking2 error
		want           *kabuspb.VolumeRanking
		hasError       bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "VolumeRankingでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			volumeRanking2: errors.New("register error message"),
			hasError:       true},
		{name: "VolumeRankingの結果を結果を返す",
			getToken1:      "TOKEN_STRING",
			volumeRanking1: &kabuspb.VolumeRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
			want:           &kabuspb.VolumeRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{volumeRanking1: test.volumeRanking1, volumeRanking2: test.volumeRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetVolumeRanking(context.Background(), &kabuspb.GetVolumeRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetValueRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		getToken1     string
		getToken2     error
		valueRanking1 *kabuspb.ValueRanking
		valueRanking2 error
		want          *kabuspb.ValueRanking
		hasError      bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "ValueRankingでエラーがあればエラーを返す",
			getToken1:     "TOKEN_STRING",
			valueRanking2: errors.New("register error message"),
			hasError:      true},
		{name: "ValueRankingの結果を結果を返す",
			getToken1:     "TOKEN_STRING",
			valueRanking1: &kabuspb.ValueRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
			want:          &kabuspb.ValueRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{valueRanking1: test.valueRanking1, valueRanking2: test.valueRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetValueRanking(context.Background(), &kabuspb.GetValueRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetMarginRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		getToken1      string
		getToken2      error
		marginRanking1 *kabuspb.MarginRanking
		marginRanking2 error
		want           *kabuspb.MarginRanking
		hasError       bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "MarginRankingでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			marginRanking2: errors.New("register error message"),
			hasError:       true},
		{name: "MarginRankingの結果を結果を返す",
			getToken1:      "TOKEN_STRING",
			marginRanking1: &kabuspb.MarginRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
			want:           &kabuspb.MarginRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{marginRanking1: test.marginRanking1, marginRanking2: test.marginRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetMarginRanking(context.Background(), &kabuspb.GetMarginRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetIndustryRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		industryRanking1 *kabuspb.IndustryRanking
		industryRanking2 error
		want             *kabuspb.IndustryRanking
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "IndustryRankingでエラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			industryRanking2: errors.New("register error message"),
			hasError:         true},
		{name: "IndustryRankingの結果を結果を返す",
			getToken1:        "TOKEN_STRING",
			industryRanking1: &kabuspb.IndustryRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL},
			want:             &kabuspb.IndustryRanking{ExchangeDivision: kabuspb.ExchangeDivision_EXCHANGE_DIVISION_ALL}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{industryRanking1: test.industryRanking1, industryRanking2: test.industryRanking2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := server.GetIndustryRanking(context.Background(), &kabuspb.GetIndustryRankingRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendStockOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		getToken1       string
		getToken2       error
		sendOrderStock1 *kabuspb.OrderResponse
		sendOrderStock2 error
		want            *kabuspb.OrderResponse
		hasError        bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:       "TOKEN_STRING",
			sendOrderStock2: errors.New("register error message"),
			hasError:        true},
		{name: "エラーがなければ結果を返す",
			getToken1:       "TOKEN_STRING",
			sendOrderStock1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:            &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{sendOrderStock1: test.sendOrderStock1, sendOrderStock2: test.sendOrderStock2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.SendStockOrder(context.Background(), &kabuspb.SendStockOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendMarginOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		sendOrderMargin1 *kabuspb.OrderResponse
		sendOrderMargin2 error
		want             *kabuspb.OrderResponse
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderMargin2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			sendOrderMargin1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:             &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{sendOrderMargin1: test.sendOrderMargin1, sendOrderMargin2: test.sendOrderMargin2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.SendMarginOrder(context.Background(), &kabuspb.SendMarginOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendFutureOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		sendOrderFuture1 *kabuspb.OrderResponse
		sendOrderFuture2 error
		want             *kabuspb.OrderResponse
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderFuture2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			sendOrderFuture1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:             &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{sendOrderFuture1: test.sendOrderFuture1, sendOrderFuture2: test.sendOrderFuture2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.SendFutureOrder(context.Background(), &kabuspb.SendFutureOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendOptionOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		sendOrderOption1 *kabuspb.OrderResponse
		sendOrderOption2 error
		want             *kabuspb.OrderResponse
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderOption2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			sendOrderOption1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:             &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{sendOrderOption1: test.sendOrderOption1, sendOrderOption2: test.sendOrderOption2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.SendOptionOrder(context.Background(), &kabuspb.SendOptionOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_CancelOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		getToken1    string
		getToken2    error
		cancelOrder1 *kabuspb.OrderResponse
		cancelOrder2 error
		want         *kabuspb.OrderResponse
		hasError     bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:    "TOKEN_STRING",
			cancelOrder2: errors.New("register error message"),
			hasError:     true},
		{name: "エラーがなければ結果を返す",
			getToken1:    "TOKEN_STRING",
			cancelOrder1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:         &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{cancelOrder1: test.cancelOrder1, cancelOrder2: test.cancelOrder2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.CancelOrder(context.Background(), &kabuspb.CancelOrderRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetStockWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		getToken1       string
		getToken2       error
		getStockWallet1 *kabuspb.StockWallet
		getStockWallet2 error
		want            *kabuspb.StockWallet
		hasError        bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:       "TOKEN_STRING",
			getStockWallet2: errors.New("register error message"),
			hasError:        true},
		{name: "エラーがなければ結果を返す",
			getToken1:       "TOKEN_STRING",
			getStockWallet1: &kabuspb.StockWallet{StockAccountWallet: 300000},
			want:            &kabuspb.StockWallet{StockAccountWallet: 300000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{getStockWallet1: test.getStockWallet1, getStockWallet2: test.getStockWallet2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.GetStockWallet(context.Background(), &kabuspb.GetStockWalletRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetMarginWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		getMarginWallet1 *kabuspb.MarginWallet
		getMarginWallet2 error
		want             *kabuspb.MarginWallet
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getMarginWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			getMarginWallet1: &kabuspb.MarginWallet{MarginAccountWallet: 300000},
			want:             &kabuspb.MarginWallet{MarginAccountWallet: 300000}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{getMarginWallet1: test.getMarginWallet1, getMarginWallet2: test.getMarginWallet2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.GetMarginWallet(context.Background(), &kabuspb.GetMarginWalletRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetFutureWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		getFutureWallet1 *kabuspb.FutureWallet
		getFutureWallet2 error
		want             *kabuspb.FutureWallet
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getFutureWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			getFutureWallet1: &kabuspb.FutureWallet{FutureTradeLimit: 300000, MarginRequirement: 0},
			want:             &kabuspb.FutureWallet{FutureTradeLimit: 300000, MarginRequirement: 0}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{getFutureWallet1: test.getFutureWallet1, getFutureWallet2: test.getFutureWallet2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.GetFutureWallet(context.Background(), &kabuspb.GetFutureWalletRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetOptionWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		getToken1        string
		getToken2        error
		getOptionWallet1 *kabuspb.OptionWallet
		getOptionWallet2 error
		want             *kabuspb.OptionWallet
		hasError         bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getOptionWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			getOptionWallet1: &kabuspb.OptionWallet{OptionBuyTradeLimit: 300000, OptionSellTradeLimit: 300000, MarginRequirement: 0},
			want:             &kabuspb.OptionWallet{OptionBuyTradeLimit: 300000, OptionSellTradeLimit: 300000, MarginRequirement: 0}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{getOptionWallet1: test.getOptionWallet1, getOptionWallet2: test.getOptionWallet2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2},
				setting:      &testSetting{password: "PASSWORD"}}
			got1, got2 := server.GetOptionWallet(context.Background(), &kabuspb.GetOptionWalletRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetBoardsStreaming(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		connect  error
		hasError bool
	}{
		{name: "connectがerrorならerrorを返す", connect: errors.New("error message"), hasError: true},
		{name: "connectがnilを返したらnilを返す", connect: nil, hasError: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{boardStreamService: &testBoardStreamService{connect: test.connect}}
			got := server.GetBoardsStreaming(nil, nil)
			if (got != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.hasError, got)
			}
		})
	}
}

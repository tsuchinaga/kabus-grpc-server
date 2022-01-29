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
	register1               *kabuspb.RegisteredSymbols
	register2               error
	unregister1             *kabuspb.RegisteredSymbols
	unregister2             error
	unregisterAll1          *kabuspb.RegisteredSymbols
	unregisterAll2          error
	symbolNameFuture1       *kabuspb.SymbolCodeInfo
	symbolNameFuture2       error
	symbolNameOption1       *kabuspb.SymbolCodeInfo
	symbolNameOption2       error
	board1                  *kabuspb.Board
	board2                  error
	symbol1                 *kabuspb.Symbol
	symbol2                 error
	orders1                 *kabuspb.Orders
	orders2                 error
	positions1              *kabuspb.Positions
	positions2              error
	priceRanking1           *kabuspb.PriceRanking
	priceRanking2           error
	tickRanking1            *kabuspb.TickRanking
	tickRanking2            error
	volumeRanking1          *kabuspb.VolumeRanking
	volumeRanking2          error
	valueRanking1           *kabuspb.ValueRanking
	valueRanking2           error
	marginRanking1          *kabuspb.MarginRanking
	marginRanking2          error
	industryRanking1        *kabuspb.IndustryRanking
	industryRanking2        error
	sendOrderStock1         *kabuspb.OrderResponse
	sendOrderStock2         error
	sendOrderMargin1        *kabuspb.OrderResponse
	sendOrderMargin2        error
	sendOrderFuture1        *kabuspb.OrderResponse
	sendOrderFuture2        error
	sendOrderOption1        *kabuspb.OrderResponse
	sendOrderOption2        error
	cancelOrder1            *kabuspb.OrderResponse
	cancelOrder2            error
	getStockWallet1         *kabuspb.StockWallet
	getStockWallet2         error
	getMarginWallet1        *kabuspb.MarginWallet
	getMarginWallet2        error
	getFutureWallet1        *kabuspb.FutureWallet
	getFutureWallet2        error
	getOptionWallet1        *kabuspb.OptionWallet
	getOptionWallet2        error
	exchange1               *kabuspb.ExchangeInfo
	exchange2               error
	regulation1             *kabuspb.Regulation
	regulation2             error
	primaryExchange1        *kabuspb.PrimaryExchange
	primaryExchange2        error
	softLimit1              *kabuspb.SoftLimit
	softLimit2              error
	marginPremium1          *kabuspb.MarginPremium
	marginPremium2          error
	isMissMatchApiKeyError1 bool
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

func (t *testSecurity) SendOrderStock(context.Context, string, *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderStock1, t.sendOrderStock2
}

func (t *testSecurity) SendOrderMargin(context.Context, string, *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderMargin1, t.sendOrderMargin2
}

func (t *testSecurity) SendOrderFuture(context.Context, string, *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderFuture1, t.sendOrderFuture2
}

func (t *testSecurity) SendOrderOption(context.Context, string, *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderOption1, t.sendOrderOption2
}

func (t *testSecurity) CancelOrder(context.Context, string, *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
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

func (t *testSecurity) Exchange(context.Context, string, *kabuspb.GetExchangeRequest) (*kabuspb.ExchangeInfo, error) {
	return t.exchange1, t.exchange2
}

func (t *testSecurity) Regulation(context.Context, string, *kabuspb.GetRegulationRequest) (*kabuspb.Regulation, error) {
	return t.regulation1, t.regulation2
}

func (t *testSecurity) PrimaryExchange(context.Context, string, *kabuspb.GetPrimaryExchangeRequest) (*kabuspb.PrimaryExchange, error) {
	return t.primaryExchange1, t.primaryExchange2
}

func (t *testSecurity) SoftLimit(context.Context, string, *kabuspb.GetSoftLimitRequest) (*kabuspb.SoftLimit, error) {
	return t.softLimit1, t.softLimit2
}

func (t *testSecurity) MarginPremium(context.Context, string, *kabuspb.GetMarginPremiumRequest) (*kabuspb.MarginPremium, error) {
	return t.marginPremium1, t.marginPremium2
}

func (t *testSecurity) IsMissMatchApiKeyError(error) bool {
	return t.isMissMatchApiKeyError1
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
	countAll            int
	getAll              []*kabuspb.RegisterSymbol
	get                 []*kabuspb.RegisterSymbol
	lastGetRequester    string
	lastAddRequester    string
	lastAddSymbols      []*kabuspb.RegisterSymbol
	lastRemoveRequester string
	lastRemoveSymbols   []*kabuspb.RegisterSymbol
}

func (t *testRegisterSymbolService) CountAll() int                     { return t.countAll }
func (t *testRegisterSymbolService) GetAll() []*kabuspb.RegisterSymbol { return t.getAll }
func (t *testRegisterSymbolService) Get(requester string) []*kabuspb.RegisterSymbol {
	t.lastGetRequester = requester
	return t.get
}
func (t *testRegisterSymbolService) Add(requester string, symbols []*kabuspb.RegisterSymbol) {
	t.lastAddRequester = requester
	t.lastAddSymbols = symbols
}
func (t *testRegisterSymbolService) Remove(requester string, symbols []*kabuspb.RegisterSymbol) {
	t.lastRemoveRequester = requester
	t.lastRemoveSymbols = symbols
}

type testBoardStreamService struct {
	services.BoardStreamService
	connect error
}

func (t *testBoardStreamService) Start() {}
func (t *testBoardStreamService) Connect(kabuspb.KabusService_GetBoardsStreamingServer) error {
	return t.connect
}

type testVirtualSecurity struct {
	repositories.VirtualSecurity
	orders1          *kabuspb.Orders
	orders2          error
	positions1       *kabuspb.Positions
	positions2       error
	sendOrderStock1  *kabuspb.OrderResponse
	sendOrderStock2  error
	sendOrderMargin1 *kabuspb.OrderResponse
	sendOrderMargin2 error
	cancelOrder1     *kabuspb.OrderResponse
	cancelOrder2     error
}

func (t *testVirtualSecurity) Orders(context.Context, string, *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	return t.orders1, t.orders2
}
func (t *testVirtualSecurity) Positions(context.Context, string, *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	return t.positions1, t.positions2
}
func (t *testVirtualSecurity) SendOrderStock(context.Context, string, *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderStock1, t.sendOrderStock2
}
func (t *testVirtualSecurity) SendOrderMargin(context.Context, string, *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.sendOrderMargin1, t.sendOrderMargin2
}
func (t *testVirtualSecurity) CancelOrder(context.Context, string, *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
	return t.cancelOrder1, t.cancelOrder2
}

func Test_NewServer(t *testing.T) {
	security := &testSecurity{}
	virtual := &testVirtualSecurity{}
	tokenService := &testTokenService{}
	registerSymbolService := &testRegisterSymbolService{}
	boardStreamService := &testBoardStreamService{}
	got := NewServer(security, virtual, tokenService, registerSymbolService, boardStreamService)
	want := &server{security: security, virtual: virtual, tokenService: tokenService, registerSymbolService: registerSymbolService, boardStreamService: boardStreamService}
	t.Parallel()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_server_GetRegisteredSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		countAll int
		get      []*kabuspb.RegisterSymbol
		want     *kabuspb.RegisteredSymbols
		hasError bool
	}{
		{name: "registerSymbolServiceの結果が返される",
			countAll: 10,
			get:      []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			want: &kabuspb.RegisteredSymbols{
				Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
				Count:   10,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{registerSymbolService: &testRegisterSymbolService{countAll: test.countAll, get: test.get}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		register1               *kabuspb.RegisteredSymbols
		register2               error
		isMissMatchApiKeyError1 bool
		countAll                int
		get                     []*kabuspb.RegisterSymbol
		arg                     *kabuspb.RegisterSymbolsRequest
		want                    *kabuspb.RegisteredSymbols
		hasError                bool
		wantAdd1                string
		wantAdd2                []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Registerでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			register2: errors.New("register error message"),
			hasError:  true},
		{name: "RegisterのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			register2:               errors.New("miss match api key error message"),
			refresh2:                errors.New("refresh error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "RegisterのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			register2:               errors.New("miss match api key error message"),
			refresh1:                "REFRESHED_TOKEN_STRING",
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "リクエストをStoreに保存してから結果を返す",
			getToken1: "TOKEN_STRING",
			register1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			arg:       &kabuspb.RegisterSymbolsRequest{RequesterName: "requester", Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			get:       []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			countAll:  2,
			want:      &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}, Count: 2},
			wantAdd1:  "requester",
			wantAdd2:  []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{countAll: test.countAll, get: test.get}
			server := &server{
				security:              &testSecurity{register1: test.register1, register2: test.register2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2},
				registerSymbolService: registerSymbolService,
				boardStreamService:    &testBoardStreamService{}}
			got1, got2 := server.RegisterSymbols(context.Background(), test.arg)
			got3 := registerSymbolService.lastAddRequester
			got4 := registerSymbolService.lastAddSymbols
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantAdd1, got3) || !reflect.DeepEqual(test.wantAdd2, got4) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					test.want, test.hasError, test.wantAdd1, test.wantAdd2,
					got1, got2, got3, got4)
			}
		})
	}
}

func Test_server_UnregisterSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		unregister1             *kabuspb.RegisteredSymbols
		unregister2             error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.RegisteredSymbols
		hasError                bool
		countAll                int
		get                     []*kabuspb.RegisterSymbol
		arg                     *kabuspb.UnregisterSymbolsRequest
		wantRemove1             string
		wantRemove2             []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Unregisterでエラーがあればエラーを返す",
			getToken1:   "TOKEN_STRING",
			unregister2: errors.New("register error message"),
			hasError:    true},
		{name: "UnregisterのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			unregister2:             errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "UnregisterのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			unregister2:             errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "リクエストの結果をStoreに保存してから結果を返す",
			getToken1:   "TOKEN_STRING",
			unregister1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:        &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}, Count: 2},
			countAll:    2,
			get:         []*kabuspb.RegisterSymbol{{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			arg:         &kabuspb.UnregisterSymbolsRequest{RequesterName: "requester", Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			wantRemove1: "requester",
			wantRemove2: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{countAll: test.countAll, get: test.get}
			server := &server{
				security:              &testSecurity{unregister1: test.unregister1, unregister2: test.unregister2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2},
				registerSymbolService: registerSymbolService,
				boardStreamService:    &testBoardStreamService{}}
			got1, got2 := server.UnregisterSymbols(context.Background(), test.arg)
			got3 := registerSymbolService.lastRemoveRequester
			got4 := registerSymbolService.lastRemoveSymbols
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantRemove1, got3) || !reflect.DeepEqual(test.wantRemove2, got4) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantRemove1, test.wantRemove2, got1, got2, got3, got4)
			}
		})
	}
}

func Test_server_UnregisterAllSymbols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		unregisterAll1          *kabuspb.RegisteredSymbols
		unregisterAll2          error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.RegisteredSymbols
		hasError                bool
		countAll                int
		get                     []*kabuspb.RegisterSymbol
		arg                     *kabuspb.UnregisterAllSymbolsRequest
		wantRemove1             string
		wantRemove2             []*kabuspb.RegisterSymbol
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "UnregisterAllでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			unregisterAll2: errors.New("register error message"),
			hasError:       true},
		{name: "UnregisterAllのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			unregisterAll2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "UnregisterAllのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			unregisterAll2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "UnregisterAllの結果をStoreに保存してから結果を返す",
			getToken1:      "TOKEN_STRING",
			unregisterAll1: &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1234", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
			want:           &kabuspb.RegisteredSymbols{Symbols: []*kabuspb.RegisterSymbol{}, Count: 2},
			countAll:       2,
			get:            []*kabuspb.RegisterSymbol{{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}},
			arg:            &kabuspb.UnregisterAllSymbolsRequest{RequesterName: "requester"},
			wantRemove1:    "requester",
			wantRemove2:    []*kabuspb.RegisterSymbol{{SymbolCode: "2345", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerSymbolService := &testRegisterSymbolService{countAll: test.countAll, get: test.get}
			server := &server{
				security:              &testSecurity{unregisterAll1: test.unregisterAll1, unregisterAll2: test.unregisterAll2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService:          &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2},
				registerSymbolService: registerSymbolService}
			got1, got2 := server.UnregisterAllSymbols(context.Background(), test.arg)
			got3 := registerSymbolService.lastRemoveRequester
			got4 := registerSymbolService.lastRemoveSymbols
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError || !reflect.DeepEqual(test.wantRemove1, got3) || !reflect.DeepEqual(test.wantRemove2, got4) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(), test.want, test.hasError, test.wantRemove1, test.wantRemove2, got1, got2, got3, got4)
			}
		})
	}
}

func Test_server_GetFutureSymbolCodeInfo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		symbolNameFuture1       *kabuspb.SymbolCodeInfo
		symbolNameFuture2       error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.SymbolCodeInfo
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "SymbolNameFutureでエラーがあればエラーを返す",
			getToken1:         "TOKEN_STRING",
			symbolNameFuture2: errors.New("register error message"),
			hasError:          true},
		{name: "SymbolNameFutureのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			symbolNameFuture2:       errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "SymbolNameFutureのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			symbolNameFuture2:       errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{symbolNameFuture1: test.symbolNameFuture1, symbolNameFuture2: test.symbolNameFuture2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		symbolNameOption1       *kabuspb.SymbolCodeInfo
		symbolNameOption2       error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.SymbolCodeInfo
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "SymbolNameOptionでエラーがあればエラーを返す",
			getToken1:         "TOKEN_STRING",
			symbolNameOption2: errors.New("register error message"),
			hasError:          true},
		{name: "SymbolNameOptionのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			symbolNameOption2:       errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "SymbolNameOptionのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			symbolNameOption2:       errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{symbolNameOption1: test.symbolNameOption1, symbolNameOption2: test.symbolNameOption2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		board1                  *kabuspb.Board
		board2                  error
		refresh1                string
		refresh2                error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.Board
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Boardでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			board2:    errors.New("register error message"),
			hasError:  true},
		{name: "BoardのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			board2:                  errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "BoardのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			board2:                  errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{board1: test.board1, board2: test.board2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		symbol1                 *kabuspb.Symbol
		symbol2                 error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.Symbol
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "Symbolでエラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			symbol2:   errors.New("register error message"),
			hasError:  true},
		{name: "SymbolのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			symbol2:                 errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "SymbolのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			symbol2:                 errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{symbol1: test.symbol1, symbol2: test.symbol2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name     string
		server   kabuspb.KabusServiceServer
		arg      *kabuspb.GetOrdersRequest
		want     *kabuspb.Orders
		hasError bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			server: &server{
				tokenService: &testTokenService{getToken2: errors.New("get token error message")}},
			arg:      &kabuspb.GetOrdersRequest{IsVirtual: false},
			hasError: true},
		{name: "Ordersでエラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{orders2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.GetOrdersRequest{IsVirtual: false},
			hasError: true},
		{name: "OrdersのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{orders2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh2: errors.New("refresh error message")}},
			arg:      &kabuspb.GetOrdersRequest{IsVirtual: false},
			hasError: true},
		{name: "OrdersのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			server: &server{
				security:     &testSecurity{orders2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh1: "REFRESHED_TOKEN_STRING"}},
			arg:      &kabuspb.GetOrdersRequest{IsVirtual: false},
			hasError: true},
		{name: "Ordersの結果を結果を返す",
			server: &server{
				security:     &testSecurity{orders1: &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.GetOrdersRequest{IsVirtual: false},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}}},
		{name: "仮想証券会社を指定して、Ordersでエラーがあればエラーを返す",
			server: &server{
				virtual:      &testVirtualSecurity{orders2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.GetOrdersRequest{IsVirtual: true},
			hasError: true},
		{name: "仮想証券会社を指定して、Ordersの結果を結果を返す",
			server: &server{
				virtual:      &testVirtualSecurity{orders1: &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.GetOrdersRequest{IsVirtual: true},
			want: &kabuspb.Orders{Orders: []*kabuspb.Order{{Id: "20210331A02N36008399"}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1, got2 := test.server.GetOrders(context.Background(), test.arg)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetPositions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		server   kabuspb.KabusServiceServer
		arg      *kabuspb.GetPositionsRequest
		want     *kabuspb.Positions
		hasError bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			server:   &server{tokenService: &testTokenService{getToken2: errors.New("get token error message")}},
			arg:      &kabuspb.GetPositionsRequest{IsVirtual: false},
			hasError: true},
		{name: "Positionsでエラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{positions2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.GetPositionsRequest{IsVirtual: false},
			hasError: true},
		{name: "PositionsのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{positions2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh2: errors.New("refresh error message")}},
			arg:      &kabuspb.GetPositionsRequest{IsVirtual: false},
			hasError: true},
		{name: "PositionsのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			server: &server{
				security:     &testSecurity{positions2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh1: "REFRESHED_TOKEN_STRING"}},
			arg:      &kabuspb.GetPositionsRequest{IsVirtual: false},
			hasError: true},
		{name: "Positionsの結果を結果を返す",
			server: &server{
				security:     &testSecurity{positions1: &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.GetPositionsRequest{IsVirtual: false},
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}}},
		{name: "仮想証券会社を指定して、Positionsでエラーがあればエラーを返す",
			server: &server{
				virtual:      &testVirtualSecurity{positions2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.GetPositionsRequest{IsVirtual: true},
			hasError: true},
		{name: "仮想証券会社を指定して、Positionsの結果を結果を返す",
			server: &server{
				virtual:      &testVirtualSecurity{positions1: &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.GetPositionsRequest{IsVirtual: true},
			want: &kabuspb.Positions{Positions: []*kabuspb.Position{{ExecutionId: "20210331A02N36008399"}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1, got2 := test.server.GetPositions(context.Background(), test.arg)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetPriceRanking(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		priceRanking1           *kabuspb.PriceRanking
		priceRanking2           error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.PriceRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "PriceRankingでエラーがあればエラーを返す",
			getToken1:     "TOKEN_STRING",
			priceRanking2: errors.New("register error message"),
			hasError:      true},
		{name: "PriceRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			priceRanking2:           errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "PriceRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			priceRanking2:           errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{priceRanking1: test.priceRanking1, priceRanking2: test.priceRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		tickRanking1            *kabuspb.TickRanking
		tickRanking2            error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.TickRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "TickRankingでエラーがあればエラーを返す",
			getToken1:    "TOKEN_STRING",
			tickRanking2: errors.New("register error message"),
			hasError:     true},
		{name: "TickRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			tickRanking2:            errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "TickRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			tickRanking2:            errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{tickRanking1: test.tickRanking1, tickRanking2: test.tickRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		volumeRanking1          *kabuspb.VolumeRanking
		volumeRanking2          error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.VolumeRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "VolumeRankingでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			volumeRanking2: errors.New("register error message"),
			hasError:       true},
		{name: "VolumeRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			volumeRanking2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "VolumeRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			volumeRanking2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{volumeRanking1: test.volumeRanking1, volumeRanking2: test.volumeRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		valueRanking1           *kabuspb.ValueRanking
		valueRanking2           error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.ValueRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "ValueRankingでエラーがあればエラーを返す",
			getToken1:     "TOKEN_STRING",
			valueRanking2: errors.New("register error message"),
			hasError:      true},
		{name: "ValueRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			valueRanking2:           errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "ValueRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			valueRanking2:           errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{valueRanking1: test.valueRanking1, valueRanking2: test.valueRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		marginRanking1          *kabuspb.MarginRanking
		marginRanking2          error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.MarginRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "MarginRankingでエラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			marginRanking2: errors.New("register error message"),
			hasError:       true},
		{name: "MarginRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			marginRanking2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "MarginRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			marginRanking2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{marginRanking1: test.marginRanking1, marginRanking2: test.marginRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		industryRanking1        *kabuspb.IndustryRanking
		industryRanking2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.IndustryRanking
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "IndustryRankingでエラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			industryRanking2: errors.New("register error message"),
			hasError:         true},
		{name: "IndustryRankingのエラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			industryRanking2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "IndustryRankingのエラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			industryRanking2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			refresh1:                "REFRESHED_TOKEN_STRING",
			hasError:                true},
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
				security:     &testSecurity{industryRanking1: test.industryRanking1, industryRanking2: test.industryRanking2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name     string
		server   kabuspb.KabusServiceServer
		want     *kabuspb.OrderResponse
		arg      *kabuspb.SendStockOrderRequest
		hasError bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			server:   &server{tokenService: &testTokenService{getToken2: errors.New("get token error message")}},
			arg:      &kabuspb.SendStockOrderRequest{IsVirtual: false},
			hasError: true},
		{name: "エラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{sendOrderStock2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.SendStockOrderRequest{IsVirtual: false},
			hasError: true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			server: &server{
				security:     &testSecurity{sendOrderStock2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh2: errors.New("refresh error message")}},
			arg:      &kabuspb.SendStockOrderRequest{IsVirtual: false},
			hasError: true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			server: &server{
				security:     &testSecurity{sendOrderStock2: errors.New("miss match api key error message"), isMissMatchApiKeyError1: true},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING", refresh1: "REFRESHED_TOKEN_STRING"}},
			arg:      &kabuspb.SendStockOrderRequest{IsVirtual: false},
			hasError: true},
		{name: "エラーがなければ結果を返す",
			server: &server{
				security:     &testSecurity{sendOrderStock1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.SendStockOrderRequest{IsVirtual: false},
			want: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
		{name: "仮想証券会社を指定していて、エラーがあればエラーを返す",
			server: &server{
				virtual:      &testVirtualSecurity{sendOrderStock2: errors.New("register error message")},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:      &kabuspb.SendStockOrderRequest{IsVirtual: true},
			hasError: true},
		{name: "仮想証券会社を指定していて、エラーがなければ結果を返す",
			server: &server{
				virtual:      &testVirtualSecurity{sendOrderStock1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
				tokenService: &testTokenService{getToken1: "TOKEN_STRING"}},
			arg:  &kabuspb.SendStockOrderRequest{IsVirtual: true},
			want: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1, got2 := test.server.SendStockOrder(context.Background(), test.arg)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendMarginOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		sendOrderMargin1        *kabuspb.OrderResponse
		sendOrderMargin2        error
		virtualSendOrderMargin1 *kabuspb.OrderResponse
		virtualSendOrderMargin2 error
		isMissMatchApiKeyError1 bool
		arg2                    *kabuspb.SendMarginOrderRequest
		want                    *kabuspb.OrderResponse
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			arg2:      &kabuspb.SendMarginOrderRequest{},
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderMargin2: errors.New("register error message"),
			arg2:             &kabuspb.SendMarginOrderRequest{},
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			sendOrderMargin2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			arg2:                    &kabuspb.SendMarginOrderRequest{},
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			sendOrderMargin2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			arg2:                    &kabuspb.SendMarginOrderRequest{},
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1:        "TOKEN_STRING",
			sendOrderMargin1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			arg2:             &kabuspb.SendMarginOrderRequest{},
			want:             &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
		{name: "仮想証券会社が指定されていれば仮想証券会社の結果を返す",
			virtualSendOrderMargin1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			arg2:                    &kabuspb.SendMarginOrderRequest{IsVirtual: true},
			want:                    &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{sendOrderMargin1: test.sendOrderMargin1, sendOrderMargin2: test.sendOrderMargin2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				virtual:      &testVirtualSecurity{sendOrderMargin1: test.virtualSendOrderMargin1, sendOrderMargin2: test.virtualSendOrderMargin2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.SendMarginOrder(context.Background(), test.arg2)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_SendFutureOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		sendOrderFuture1        *kabuspb.OrderResponse
		sendOrderFuture2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.OrderResponse
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderFuture2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			sendOrderFuture2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			sendOrderFuture2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{sendOrderFuture1: test.sendOrderFuture1, sendOrderFuture2: test.sendOrderFuture2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		sendOrderOption1        *kabuspb.OrderResponse
		sendOrderOption2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.OrderResponse
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			sendOrderOption2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			sendOrderOption2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			sendOrderOption2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{sendOrderOption1: test.sendOrderOption1, sendOrderOption2: test.sendOrderOption2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		cancelOrder1            *kabuspb.OrderResponse
		cancelOrder2            error
		virtualCancelOrder1     *kabuspb.OrderResponse
		virtualCancelOrder2     error
		isMissMatchApiKeyError1 bool
		arg2                    *kabuspb.CancelOrderRequest
		want                    *kabuspb.OrderResponse
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			arg2:      &kabuspb.CancelOrderRequest{},
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			arg2:         &kabuspb.CancelOrderRequest{},
			getToken1:    "TOKEN_STRING",
			cancelOrder2: errors.New("register error message"),
			hasError:     true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			cancelOrder2:            errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			arg2:                    &kabuspb.CancelOrderRequest{},
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			cancelOrder2:            errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			arg2:                    &kabuspb.CancelOrderRequest{},
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			arg2:         &kabuspb.CancelOrderRequest{},
			getToken1:    "TOKEN_STRING",
			cancelOrder1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:         &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
		{name: "仮想証券会社が指定されていれば、仮想証券会社を叩く",
			arg2:                &kabuspb.CancelOrderRequest{IsVirtual: true},
			virtualCancelOrder1: &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"},
			want:                &kabuspb.OrderResponse{ResultCode: 0, OrderId: "ORDER-ID"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{cancelOrder1: test.cancelOrder1, cancelOrder2: test.cancelOrder2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				virtual:      &testVirtualSecurity{cancelOrder1: test.virtualCancelOrder1, cancelOrder2: test.virtualCancelOrder2},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.CancelOrder(context.Background(), test.arg2)
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetStockWallet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		getStockWallet1         *kabuspb.StockWallet
		getStockWallet2         error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.StockWallet
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:       "TOKEN_STRING",
			getStockWallet2: errors.New("register error message"),
			hasError:        true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			getStockWallet2:         errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			getStockWallet2:         errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{getStockWallet1: test.getStockWallet1, getStockWallet2: test.getStockWallet2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		getMarginWallet1        *kabuspb.MarginWallet
		getMarginWallet2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.MarginWallet
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getMarginWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			getMarginWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			getMarginWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			refresh1:                "REFRESHED_TOKEN_STRING",
			hasError:                true},
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
				security:     &testSecurity{getMarginWallet1: test.getMarginWallet1, getMarginWallet2: test.getMarginWallet2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		getFutureWallet1        *kabuspb.FutureWallet
		getFutureWallet2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.FutureWallet
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getFutureWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			getFutureWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			getFutureWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{getFutureWallet1: test.getFutureWallet1, getFutureWallet2: test.getFutureWallet2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
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
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		getOptionWallet1        *kabuspb.OptionWallet
		getOptionWallet2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.OptionWallet
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			getOptionWallet2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			getOptionWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			getOptionWallet2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
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
				security:     &testSecurity{getOptionWallet1: test.getOptionWallet1, getOptionWallet2: test.getOptionWallet2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetOptionWallet(context.Background(), &kabuspb.GetOptionWalletRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		exchange1               *kabuspb.ExchangeInfo
		exchange2               error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.ExchangeInfo
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1: "TOKEN_STRING",
			exchange2: errors.New("register error message"),
			hasError:  true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			exchange2:               errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			exchange2:               errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1: "TOKEN_STRING",
			exchange1: &kabuspb.ExchangeInfo{
				Currency: kabuspb.Currency_CURRENCY_USD_JPY,
				BidPrice: 105.502,
				Spread:   0.2,
				AskPrice: 105.504,
				Change:   -0.055,
				Time:     timestamppb.New(time.Date(0, 1, 1, 16, 10, 45, 0, time.Local)),
			},
			want: &kabuspb.ExchangeInfo{
				Currency: kabuspb.Currency_CURRENCY_USD_JPY,
				BidPrice: 105.502,
				Spread:   0.2,
				AskPrice: 105.504,
				Change:   -0.055,
				Time:     timestamppb.New(time.Date(0, 1, 1, 16, 10, 45, 0, time.Local)),
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{exchange1: test.exchange1, exchange2: test.exchange2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetExchange(context.Background(), &kabuspb.GetExchangeRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetRegulation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		regulation1             *kabuspb.Regulation
		regulation2             error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.Regulation
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:   "TOKEN_STRING",
			regulation2: errors.New("register error message"),
			hasError:    true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			regulation2:             errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			regulation2:             errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1: "TOKEN_STRING",
			regulation1: &kabuspb.Regulation{
				SymbolCode: "5614",
				RegulationInfoList: []*kabuspb.RegulationInfo{
					{
						Exchange:      kabuspb.RegulationExchange_REGULATION_EXCHANGE_TOUSHOU,
						Product:       kabuspb.RegulationProduct_REGULATION_PRODUCT_RECEIPT,
						Side:          kabuspb.RegulationSide_REGULATION_SIDE_BUY,
						Reason:        "品受停止（貸借申込停止銘柄（日証金規制））",
						LimitStartDay: timestamppb.New(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)),
						LimitEndDay:   timestamppb.New(time.Date(2999, 12, 31, 0, 0, 0, 0, time.Local)),
						Level:         kabuspb.RegulationLevel_REGULATION_LEVEL_ERROR,
					}, {
						Exchange:      kabuspb.RegulationExchange_REGULATION_EXCHANGE_UNSPECIFIED,
						Product:       kabuspb.RegulationProduct_REGULATION_PRODUCT_STOCK,
						Side:          kabuspb.RegulationSide_REGULATION_SIDE_BUY,
						Reason:        "その他（代用不適格銘柄）",
						LimitStartDay: timestamppb.New(time.Date(2021, 1, 27, 0, 0, 0, 0, time.Local)),
						LimitEndDay:   timestamppb.New(time.Date(2021, 2, 17, 0, 0, 0, 0, time.Local)),
						Level:         kabuspb.RegulationLevel_REGULATION_LEVEL_ERROR,
					},
				},
			},
			want: &kabuspb.Regulation{
				SymbolCode: "5614",
				RegulationInfoList: []*kabuspb.RegulationInfo{
					{
						Exchange:      kabuspb.RegulationExchange_REGULATION_EXCHANGE_TOUSHOU,
						Product:       kabuspb.RegulationProduct_REGULATION_PRODUCT_RECEIPT,
						Side:          kabuspb.RegulationSide_REGULATION_SIDE_BUY,
						Reason:        "品受停止（貸借申込停止銘柄（日証金規制））",
						LimitStartDay: timestamppb.New(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)),
						LimitEndDay:   timestamppb.New(time.Date(2999, 12, 31, 0, 0, 0, 0, time.Local)),
						Level:         kabuspb.RegulationLevel_REGULATION_LEVEL_ERROR,
					}, {
						Exchange:      kabuspb.RegulationExchange_REGULATION_EXCHANGE_UNSPECIFIED,
						Product:       kabuspb.RegulationProduct_REGULATION_PRODUCT_STOCK,
						Side:          kabuspb.RegulationSide_REGULATION_SIDE_BUY,
						Reason:        "その他（代用不適格銘柄）",
						LimitStartDay: timestamppb.New(time.Date(2021, 1, 27, 0, 0, 0, 0, time.Local)),
						LimitEndDay:   timestamppb.New(time.Date(2021, 2, 17, 0, 0, 0, 0, time.Local)),
						Level:         kabuspb.RegulationLevel_REGULATION_LEVEL_ERROR,
					},
				},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{regulation1: test.regulation1, regulation2: test.regulation2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetRegulation(context.Background(), &kabuspb.GetRegulationRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetPrimaryExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		primaryExchange1        *kabuspb.PrimaryExchange
		primaryExchange2        error
		isMissMatchApiKeyError1 bool
		want                    *kabuspb.PrimaryExchange
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:        "TOKEN_STRING",
			primaryExchange2: errors.New("register error message"),
			hasError:         true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			primaryExchange2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			primaryExchange2:        errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1: "TOKEN_STRING",
			primaryExchange1: &kabuspb.PrimaryExchange{
				SymbolCode:      "2928",
				PrimaryExchange: kabuspb.StockExchange_STOCK_EXCHANGE_SATSUSHOU,
			},
			want: &kabuspb.PrimaryExchange{
				SymbolCode:      "2928",
				PrimaryExchange: kabuspb.StockExchange_STOCK_EXCHANGE_SATSUSHOU,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{primaryExchange1: test.primaryExchange1, primaryExchange2: test.primaryExchange2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetPrimaryExchange(context.Background(), &kabuspb.GetPrimaryExchangeRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_GetSoftLimit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		isMissMatchApiKeyError1 bool
		softLimit1              *kabuspb.SoftLimit
		softLimit2              error
		want                    *kabuspb.SoftLimit
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:  "TOKEN_STRING",
			softLimit2: errors.New("register error message"),
			hasError:   true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			softLimit2:              errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			softLimit2:              errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1: "TOKEN_STRING",
			softLimit1: &kabuspb.SoftLimit{
				Stock:        200,
				Margin:       200,
				Future:       10,
				FutureMini:   100,
				Option:       20,
				KabusVersion: "5.13.1.0",
			},
			want: &kabuspb.SoftLimit{
				Stock:        200,
				Margin:       200,
				Future:       10,
				FutureMini:   100,
				Option:       20,
				KabusVersion: "5.13.1.0",
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{softLimit1: test.softLimit1, softLimit2: test.softLimit2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetSoftLimit(context.Background(), &kabuspb.GetSoftLimitRequest{})
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

func Test_server_MarginPremium(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		getToken1               string
		getToken2               error
		refresh1                string
		refresh2                error
		isMissMatchApiKeyError1 bool
		marginPremium1          *kabuspb.MarginPremium
		marginPremium2          error
		want                    *kabuspb.MarginPremium
		hasError                bool
	}{
		{name: "token取得でエラーがあればエラーを返す",
			getToken2: errors.New("get token error message"),
			hasError:  true},
		{name: "エラーがあればエラーを返す",
			getToken1:      "TOKEN_STRING",
			marginPremium2: errors.New("register error message"),
			hasError:       true},
		{name: "エラーがAPIキー不一致なら再発行をたたき、再発行でエラーがあればエラーを返す",
			getToken1:               "TOKEN_STRING",
			refresh2:                errors.New("refresh error message"),
			marginPremium2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがAPIキー不一致なら再発行をたたき再発行に成功すれば再度リクエストを送る",
			getToken1:               "TOKEN_STRING",
			refresh1:                "REFRESHED_TOKEN_STRING",
			marginPremium2:          errors.New("miss match api key error message"),
			isMissMatchApiKeyError1: true,
			hasError:                true},
		{name: "エラーがなければ結果を返す",
			getToken1: "TOKEN_STRING",
			marginPremium1: &kabuspb.MarginPremium{
				SymbolCode: "9433",
				GeneralMargin: &kabuspb.MarginPremiumDetail{
					MarginPremiumType:  kabuspb.MarginPremiumType_MARGIN_PREMIUM_TYPE_UNSPECIFIED,
					MarginPremium:      0,
					UpperMarginPremium: 0,
					LowerMarginPremium: 0,
					TickMarginPremium:  0,
				},
				DayTrade: &kabuspb.MarginPremiumDetail{
					MarginPremiumType:  kabuspb.MarginPremiumType_MARGIN_PREMIUM_TYPE_AUCTION,
					MarginPremium:      0.55,
					UpperMarginPremium: 1,
					LowerMarginPremium: 0.3,
					TickMarginPremium:  0.01,
				},
			},
			want: &kabuspb.MarginPremium{
				SymbolCode: "9433",
				GeneralMargin: &kabuspb.MarginPremiumDetail{
					MarginPremiumType:  kabuspb.MarginPremiumType_MARGIN_PREMIUM_TYPE_UNSPECIFIED,
					MarginPremium:      0,
					UpperMarginPremium: 0,
					LowerMarginPremium: 0,
					TickMarginPremium:  0,
				},
				DayTrade: &kabuspb.MarginPremiumDetail{
					MarginPremiumType:  kabuspb.MarginPremiumType_MARGIN_PREMIUM_TYPE_AUCTION,
					MarginPremium:      0.55,
					UpperMarginPremium: 1,
					LowerMarginPremium: 0.3,
					TickMarginPremium:  0.01,
				},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{
				security:     &testSecurity{marginPremium1: test.marginPremium1, marginPremium2: test.marginPremium2, isMissMatchApiKeyError1: test.isMissMatchApiKeyError1},
				tokenService: &testTokenService{getToken1: test.getToken1, getToken2: test.getToken2, refresh1: test.refresh1, refresh2: test.refresh2}}
			got1, got2 := server.GetMarginPremium(context.Background(), &kabuspb.GetMarginPremiumRequest{})
			if !reflect.DeepEqual(test.want, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want, test.hasError, got1, got2)
			}
		})
	}
}

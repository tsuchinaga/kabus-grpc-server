package server

import (
	"context"
	"sync"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/services"
)

func NewServer(
	security repositories.Security,
	virtual repositories.VirtualSecurity,
	tokenService services.TokenService,
	registerSymbolService services.RegisterSymbolService,
	boardStreamService services.BoardStreamService) kabuspb.KabusServiceServer {
	return &server{
		security:              security,
		virtual:               virtual,
		tokenService:          tokenService,
		registerSymbolService: registerSymbolService,
		boardStreamService:    boardStreamService,
	}
}

type server struct {
	kabuspb.UnimplementedKabusServiceServer
	security              repositories.Security
	virtual               repositories.VirtualSecurity
	tokenService          services.TokenService
	registerSymbolService services.RegisterSymbolService
	boardStreamService    services.BoardStreamService
	orderMtx              sync.Mutex
	walletMtx             sync.Mutex
	infoMtx               sync.Mutex
}

func (s *server) SendStockOrder(ctx context.Context, req *kabuspb.SendStockOrderRequest) (*kabuspb.OrderResponse, error) {
	s.orderMtx.Lock()
	defer func() {
		<-time.After(200 * time.Millisecond) // 0.2s
		s.orderMtx.Unlock()
	}()

	// 仮想証券会社の利用
	if req.IsVirtual {
		return s.virtual.SendOrderStock(ctx, "", req)
	}

	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := s.security.SendOrderStock(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SendOrderStock(ctx, token, req)
	}
	return res, err
}

func (s *server) SendMarginOrder(ctx context.Context, req *kabuspb.SendMarginOrderRequest) (*kabuspb.OrderResponse, error) {
	s.orderMtx.Lock()
	defer func() {
		<-time.After(200 * time.Millisecond) // 0.2s
		s.orderMtx.Unlock()
	}()

	// 仮想証券会社の利用
	if req.IsVirtual {
		return s.virtual.SendOrderMargin(ctx, "", req)
	}

	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := s.security.SendOrderMargin(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SendOrderMargin(ctx, token, req)
	}
	return res, err
}

func (s *server) SendFutureOrder(ctx context.Context, req *kabuspb.SendFutureOrderRequest) (*kabuspb.OrderResponse, error) {
	s.orderMtx.Lock()
	defer func() {
		<-time.After(200 * time.Millisecond) // 0.2s
		s.orderMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := s.security.SendOrderFuture(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SendOrderFuture(ctx, token, req)
	}
	return res, err
}

func (s *server) SendOptionOrder(ctx context.Context, req *kabuspb.SendOptionOrderRequest) (*kabuspb.OrderResponse, error) {
	s.orderMtx.Lock()
	defer func() {
		<-time.After(200 * time.Millisecond) // 0.2s
		s.orderMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := s.security.SendOrderOption(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SendOrderOption(ctx, token, req)
	}
	return res, err
}

func (s *server) CancelOrder(ctx context.Context, req *kabuspb.CancelOrderRequest) (*kabuspb.OrderResponse, error) {
	s.orderMtx.Lock()
	defer func() {
		<-time.After(200 * time.Millisecond) // 0.2s
		s.orderMtx.Unlock()
	}()

	// 仮想証券会社の利用
	if req.IsVirtual {
		return s.virtual.CancelOrder(ctx, "", req)
	}

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.CancelOrder(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.CancelOrder(ctx, token, req)
	}
	return res, err
}

func (s *server) GetStockWallet(ctx context.Context, req *kabuspb.GetStockWalletRequest) (*kabuspb.StockWallet, error) {
	s.walletMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.walletMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.GetStockWallet(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.GetStockWallet(ctx, token, req)
	}
	return res, err
}

func (s *server) GetMarginWallet(ctx context.Context, req *kabuspb.GetMarginWalletRequest) (*kabuspb.MarginWallet, error) {
	s.walletMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.walletMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.GetMarginWallet(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.GetMarginWallet(ctx, token, req)
	}
	return res, err
}

func (s *server) GetFutureWallet(ctx context.Context, req *kabuspb.GetFutureWalletRequest) (*kabuspb.FutureWallet, error) {
	s.walletMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.walletMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.GetFutureWallet(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.GetFutureWallet(ctx, token, req)
	}
	return res, err
}

func (s *server) GetOptionWallet(ctx context.Context, req *kabuspb.GetOptionWalletRequest) (*kabuspb.OptionWallet, error) {
	s.walletMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.walletMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.GetOptionWallet(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.GetOptionWallet(ctx, token, req)
	}
	return res, err
}

func (s *server) GetBoard(ctx context.Context, req *kabuspb.GetBoardRequest) (*kabuspb.Board, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Board(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Board(ctx, token, req)
	}
	return res, err
}

func (s *server) GetOrders(ctx context.Context, req *kabuspb.GetOrdersRequest) (*kabuspb.Orders, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	// 仮想証券会社の利用
	if req.IsVirtual {
		return s.virtual.Orders(ctx, "", req)
	}

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Orders(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Orders(ctx, token, req)
	}
	return res, err
}

func (s *server) GetPositions(ctx context.Context, req *kabuspb.GetPositionsRequest) (*kabuspb.Positions, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	// 仮想証券会社の利用
	if req.IsVirtual {
		return s.virtual.Positions(ctx, "", req)
	}

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Positions(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Positions(ctx, token, req)
	}
	return res, err
}

func (s *server) GetFutureSymbolCodeInfo(ctx context.Context, req *kabuspb.GetFutureSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.SymbolNameFuture(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SymbolNameFuture(ctx, token, req)
	}
	return res, err
}

func (s *server) GetOptionSymbolCodeInfo(ctx context.Context, req *kabuspb.GetOptionSymbolCodeInfoRequest) (*kabuspb.SymbolCodeInfo, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.SymbolNameOption(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SymbolNameOption(ctx, token, req)
	}
	return res, err
}

func (s *server) GetRegisteredSymbols(_ context.Context, req *kabuspb.GetRegisteredSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) RegisterSymbols(ctx context.Context, req *kabuspb.RegisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.security.RegisterSymbols(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		_, err = s.security.RegisterSymbols(ctx, token, req)
	}
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Add(req.RequesterName, req.Symbols)
	s.boardStreamService.Start() // 銘柄を登録された段階で仮想証券会社への通知を始める

	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) UnregisterSymbols(ctx context.Context, req *kabuspb.UnregisterSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.security.UnregisterSymbols(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		_, err = s.security.UnregisterSymbols(ctx, token, req)
	}
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Remove(req.RequesterName, req.Symbols)
	s.boardStreamService.Start() // 銘柄を登録された段階で仮想証券会社への通知を始める

	return &kabuspb.RegisteredSymbols{
		Symbols: s.registerSymbolService.Get(req.RequesterName),
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) UnregisterAllSymbols(ctx context.Context, req *kabuspb.UnregisterAllSymbolsRequest) (*kabuspb.RegisteredSymbols, error) {
	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.security.UnregisterAll(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		_, err = s.security.UnregisterAll(ctx, token, req)
	}
	if err != nil {
		return nil, err
	}

	s.registerSymbolService.Remove(req.RequesterName, s.registerSymbolService.Get(req.RequesterName))
	return &kabuspb.RegisteredSymbols{
		Symbols: []*kabuspb.RegisterSymbol{},
		Count:   int32(s.registerSymbolService.CountAll()),
	}, nil
}

func (s *server) GetSymbol(ctx context.Context, req *kabuspb.GetSymbolRequest) (*kabuspb.Symbol, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Symbol(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Symbol(ctx, token, req)
	}
	return res, err
}

func (s *server) GetPriceRanking(ctx context.Context, req *kabuspb.GetPriceRankingRequest) (*kabuspb.PriceRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.PriceRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.PriceRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetTickRanking(ctx context.Context, req *kabuspb.GetTickRankingRequest) (*kabuspb.TickRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.TickRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.TickRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetVolumeRanking(ctx context.Context, req *kabuspb.GetVolumeRankingRequest) (*kabuspb.VolumeRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.VolumeRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.VolumeRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetValueRanking(ctx context.Context, req *kabuspb.GetValueRankingRequest) (*kabuspb.ValueRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.ValueRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.ValueRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetMarginRanking(ctx context.Context, req *kabuspb.GetMarginRankingRequest) (*kabuspb.MarginRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.MarginRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.MarginRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetIndustryRanking(ctx context.Context, req *kabuspb.GetIndustryRankingRequest) (*kabuspb.IndustryRanking, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.IndustryRanking(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.IndustryRanking(ctx, token, req)
	}
	return res, err
}

func (s *server) GetExchange(ctx context.Context, req *kabuspb.GetExchangeRequest) (*kabuspb.ExchangeInfo, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Exchange(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Exchange(ctx, token, req)
	}
	return res, err
}

func (s *server) GetRegulation(ctx context.Context, req *kabuspb.GetRegulationRequest) (*kabuspb.Regulation, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.Regulation(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.Regulation(ctx, token, req)
	}
	return res, err
}

func (s *server) GetPrimaryExchange(ctx context.Context, req *kabuspb.GetPrimaryExchangeRequest) (*kabuspb.PrimaryExchange, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.PrimaryExchange(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.PrimaryExchange(ctx, token, req)
	}
	return res, err
}

func (s *server) GetSoftLimit(ctx context.Context, req *kabuspb.GetSoftLimitRequest) (*kabuspb.SoftLimit, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.SoftLimit(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.SoftLimit(ctx, token, req)
	}
	return res, err
}

func (s *server) GetMarginPremium(ctx context.Context, req *kabuspb.GetMarginPremiumRequest) (*kabuspb.MarginPremium, error) {
	s.infoMtx.Lock()
	defer func() {
		<-time.After(100 * time.Millisecond) // 0.1s
		s.infoMtx.Unlock()
	}()

	token, err := s.tokenService.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.security.MarginPremium(ctx, token, req)
	if s.security.IsMissMatchApiKeyError(err) { // APIキー不一致なら再発行して再実行
		token, err = s.tokenService.Refresh(ctx)
		if err != nil {
			return nil, err
		}

		res, err = s.security.MarginPremium(ctx, token, req)
	}
	return res, err
}

func (s *server) GetBoardsStreaming(_ *kabuspb.GetBoardsStreamingRequest, stream kabuspb.KabusService_GetBoardsStreamingServer) error {
	return s.boardStreamService.Connect(stream)
}

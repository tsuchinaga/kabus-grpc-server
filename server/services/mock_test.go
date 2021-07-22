package services

import (
	"context"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

type testTokenStore struct {
	repositories.TokenStore
	getToken      string
	getExpiredAt  time.Time
	isExpired     bool
	lastSetToken1 string
	lastSetToken2 time.Time
	resetCount    int
}

func (t *testTokenStore) GetToken() string         { return t.getToken }
func (t *testTokenStore) GetExpiredAt() time.Time  { return t.getExpiredAt }
func (t *testTokenStore) IsExpired(time.Time) bool { return t.isExpired }
func (t *testTokenStore) SetToken(token string, expire time.Time) {
	t.lastSetToken1 = token
	t.lastSetToken2 = expire
}
func (t *testTokenStore) Reset() {
	t.resetCount++
}

type testSecurity struct {
	repositories.Security
	token1 string
	token2 error
}

func (t *testSecurity) Token(context.Context, string) (string, error) { return t.token1, t.token2 }

type testClock struct {
	repositories.Clock
	now time.Time
}

func (t *testClock) Now() time.Time {
	if t.now.IsZero() {
		return time.Now()
	}
	return t.now
}

type testSetting struct {
	repositories.Setting
}

func (t *testSetting) Password() string { return "" }

type testRegisterSymbolStore struct {
	repositories.RegisterSymbolStore
	countAll           int
	getAll             []*kabuspb.RegisterSymbol
	getByRequester     []*kabuspb.RegisterSymbol
	callAddAllCount    int
	callRemoveAllCount int
}

func (t *testRegisterSymbolStore) CountAll() int                     { return t.countAll }
func (t *testRegisterSymbolStore) GetAll() []*kabuspb.RegisterSymbol { return t.getAll }
func (t *testRegisterSymbolStore) GetByRequester(string) []*kabuspb.RegisterSymbol {
	return t.getByRequester
}
func (t *testRegisterSymbolStore) AddAll(string, []*kabuspb.RegisterSymbol) { t.callAddAllCount++ }
func (t *testRegisterSymbolStore) RemoveAll(string, []*kabuspb.RegisterSymbol) {
	t.callRemoveAllCount++
}

type testBoardStreamStore struct {
	repositories.BoardStreamStore
	hasStream   bool
	all         map[int]kabuspb.KabusService_GetBoardsStreamingServer
	addCount    int
	removeCount int
	chErr       error
}

func (t *testBoardStreamStore) HasStream() bool { return t.hasStream }
func (t *testBoardStreamStore) All() map[int]kabuspb.KabusService_GetBoardsStreamingServer {
	return t.all
}
func (t *testBoardStreamStore) Add(_ kabuspb.KabusService_GetBoardsStreamingServer, ch chan error) {
	go func() {
		<-time.After(time.Second)
		close(ch)
	}()
	t.addCount++
}
func (t *testBoardStreamStore) Remove(int, error) { t.removeCount++ }

type testBoardWS struct {
	repositories.BoardWS
	isConnected     bool
	connect         error
	disconnect      error
	disconnectCount int
}

func (t *testBoardWS) IsConnected() bool                              { return t.isConnected }
func (t *testBoardWS) Connect(func(board *kabuspb.Board) error) error { return t.connect }
func (t *testBoardWS) Disconnect() error {
	t.disconnectCount++
	return t.disconnect
}

type testGetBoardsStreamingServer struct {
	kabuspb.KabusService_GetBoardsStreamingServer
	send      error
	sendCount int
}

func (t *testGetBoardsStreamingServer) Send(*kabuspb.Board) error {
	t.sendCount++
	return t.send
}

type testVirtualSecurity struct {
	repositories.VirtualSecurity
	sendPrice      error
	sendPriceCount int
}

func (t *testVirtualSecurity) SendPrice(_ context.Context, _ *kabuspb.Board) error {
	t.sendPriceCount++
	return t.sendPrice
}

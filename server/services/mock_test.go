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
	getAll  []*kabuspb.RegisterSymbol
	lastSet []*kabuspb.RegisterSymbol
}

func (t *testRegisterSymbolStore) GetAll() []*kabuspb.RegisterSymbol {
	return t.getAll
}
func (t *testRegisterSymbolStore) SetAll(registeredList []*kabuspb.RegisterSymbol) {
	t.lastSet = registeredList
}

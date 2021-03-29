package stores

import (
	"sync"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

var (
	tokenSingleton      repositories.TokenStore
	tokenSingletonMutex sync.Mutex
)

func GetTokenStore() repositories.TokenStore {
	tokenSingletonMutex.Lock()
	defer tokenSingletonMutex.Unlock()

	if tokenSingleton == nil {
		tokenSingleton = &token{}
	}

	return tokenSingleton
}

type token struct {
	token  string
	expire time.Time
	mtx    sync.Mutex
}

func (s *token) GetToken() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.token
}

func (s *token) GetExpiredAt() time.Time {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.expire
}

func (s *token) SetToken(token string, expire time.Time) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.token = token
	s.expire = expire
}

func (s *token) IsExpired(now time.Time) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.expire.Before(now)
}

func (s *token) Reset() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.token = ""
	s.expire = time.Time{}
}

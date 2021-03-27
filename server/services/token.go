package services

import (
	"context"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewTokenService(
	tokenStore repositories.TokenStore,
	security repositories.Security,
	clock repositories.Clock,
	setting repositories.Setting) TokenService {
	return &token{tokenStore: tokenStore, security: security, clock: clock, setting: setting}
}

type TokenService interface {
	GetToken(ctx context.Context) (string, error)
}

type token struct {
	tokenStore repositories.TokenStore
	security   repositories.Security
	clock      repositories.Clock
	setting    repositories.Setting
}

func (s *token) GetToken(ctx context.Context) (string, error) {
	now := s.clock.Now()
	if s.tokenStore.IsExpired(now) {
		token, err := s.security.Token(ctx, s.setting.Password())
		if err != nil {
			return "", err
		}
		expire := time.Date(now.Year(), now.Month(), now.Day(), 6, 30, 0, 0, now.Location())
		if expire.Before(now) {
			expire = expire.AddDate(0, 0, 1)
		}
		s.tokenStore.SetToken(token, expire)
	}

	return s.tokenStore.GetToken(), nil
}

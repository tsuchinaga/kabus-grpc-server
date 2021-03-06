package repositories

import "time"

type TokenStore interface {
	GetToken() string
	GetExpiredAt() time.Time
	SetToken(token string, expire time.Time)
	IsExpired(now time.Time) bool
	Reset()
}

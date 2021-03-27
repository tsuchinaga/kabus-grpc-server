package infra

import (
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewClock() repositories.Clock {
	return new(clock)
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
}

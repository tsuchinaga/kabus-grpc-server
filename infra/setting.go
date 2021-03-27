package infra

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

var (
	settingSingleton      repositories.Setting
	settingSingletonMutex sync.Mutex
)

func InitSetting(isProd bool, password string) {
	settingSingletonMutex.Lock()
	defer settingSingletonMutex.Unlock()

	settingSingleton = &setting{isProd: isProd, password: password}
}

func GetSetting() repositories.Setting {
	settingSingletonMutex.Lock()
	defer settingSingletonMutex.Unlock()

	return settingSingleton
}

type setting struct {
	isProd   bool
	password string
}

func (s *setting) IsProduction() bool {
	return s.isProd
}

func (s *setting) Password() string {
	return s.password
}

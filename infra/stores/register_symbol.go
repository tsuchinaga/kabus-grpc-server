package stores

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

var (
	registerSymbolSingleton      repositories.RegisterSymbolStore
	registerSymbolSingletonMutex sync.Mutex
)

func GetRegisterSymbolStore() repositories.RegisterSymbolStore {
	registerSymbolSingletonMutex.Lock()
	defer registerSymbolSingletonMutex.Unlock()

	if registerSymbolSingleton == nil {
		registerSymbolSingleton = &registerSymbol{store: []*kabuspb.RegisterSymbol{}}
	}

	return registerSymbolSingleton
}

type registerSymbol struct {
	store []*kabuspb.RegisterSymbol
	mtx   sync.Mutex
}

func (s *registerSymbol) GetAll() []*kabuspb.RegisterSymbol {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.store
}

func (s *registerSymbol) SetAll(registeredList []*kabuspb.RegisterSymbol) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.store = registeredList
}

package services

import (
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewRegisterSymbolService(registerSymbolStore repositories.RegisterSymbolStore) RegisterSymbolService {
	return &registerSymbol{registerSymbolStore: registerSymbolStore}
}

type RegisterSymbolService interface {
	Get() []*kabuspb.RegisterSymbol
	Set(registeredSymbols []*kabuspb.RegisterSymbol)
}

type registerSymbol struct {
	registerSymbolStore repositories.RegisterSymbolStore
}

func (s *registerSymbol) Get() []*kabuspb.RegisterSymbol {
	return s.registerSymbolStore.GetAll()
}

func (s *registerSymbol) Set(registeredSymbols []*kabuspb.RegisterSymbol) {
	s.registerSymbolStore.SetAll(registeredSymbols)
}

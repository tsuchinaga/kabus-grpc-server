package services

import (
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewRegisterSymbolService(registerSymbolStore repositories.RegisterSymbolStore) RegisterSymbolService {
	return &registerSymbol{registerSymbolStore: registerSymbolStore}
}

type RegisterSymbolService interface {
	CountAll() int
	GetAll() []*kabuspb.RegisterSymbol
	Get(requester string) []*kabuspb.RegisterSymbol
	Add(requester string, symbols []*kabuspb.RegisterSymbol)
	Remove(requester string, symbols []*kabuspb.RegisterSymbol)
}

type registerSymbol struct {
	registerSymbolStore repositories.RegisterSymbolStore
}

func (s *registerSymbol) CountAll() int {
	return s.registerSymbolStore.CountAll()
}

func (s *registerSymbol) GetAll() []*kabuspb.RegisterSymbol {
	return s.registerSymbolStore.GetAll()
}

func (s *registerSymbol) Get(requester string) []*kabuspb.RegisterSymbol {
	return s.registerSymbolStore.GetByRequester(requester)
}

func (s *registerSymbol) Add(requester string, symbols []*kabuspb.RegisterSymbol) {
	s.registerSymbolStore.AddAll(requester, symbols)
}

func (s *registerSymbol) Remove(requester string, symbols []*kabuspb.RegisterSymbol) {
	s.registerSymbolStore.RemoveAll(requester, symbols)
}

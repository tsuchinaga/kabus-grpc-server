package repositories

import "gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

type RegisterSymbolStore interface {
	CountAll() int
	GetAll() []*kabuspb.RegisterSymbol
	GetByRequester(requester string) []*kabuspb.RegisterSymbol
	AddAll(requester string, symbols []*kabuspb.RegisterSymbol)
	RemoveAll(requester string, symbols []*kabuspb.RegisterSymbol)
}

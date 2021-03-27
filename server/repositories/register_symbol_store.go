package repositories

import "gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

type RegisterSymbolStore interface {
	GetAll() []*kabuspb.RegisterSymbol
	SetAll(registeredList []*kabuspb.RegisterSymbol)
}

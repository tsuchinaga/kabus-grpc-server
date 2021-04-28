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
		registerSymbolSingleton = &registerSymbol{store: []*RegisterSymbol{}}
	}

	return registerSymbolSingleton
}

type registerSymbol struct {
	store []*RegisterSymbol
	mtx   sync.Mutex
}

func (s *registerSymbol) CountAll() int {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return len(s.store)
}

func (s *registerSymbol) GetAll() []*kabuspb.RegisterSymbol {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make([]*kabuspb.RegisterSymbol, len(s.store))
	for i, symbol := range s.store {
		res[i] = symbol.Symbol
	}
	return res
}

func (s *registerSymbol) GetByRequester(requester string) []*kabuspb.RegisterSymbol {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make([]*kabuspb.RegisterSymbol, 0)
	for _, symbol := range s.store {
		if symbol.HasRequester(requester) {
			res = append(res, symbol.Symbol)
		}
	}

	return res
}

func (s *registerSymbol) AddAll(requester string, symbols []*kabuspb.RegisterSymbol) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, symbol := range symbols {
		index := -1
		for i, storedSymbol := range s.store {
			if storedSymbol.Symbol.SymbolCode == symbol.SymbolCode && storedSymbol.Symbol.Exchange == symbol.Exchange {
				index = i
				break
			}
		}

		if index == -1 {
			s.store = append(s.store, &RegisterSymbol{
				Symbol:     symbol,
				Requesters: []string{requester},
			})
		} else {
			s.store[index].AddRequester(requester)
		}
	}
}

func (s *registerSymbol) RemoveAll(requester string, symbols []*kabuspb.RegisterSymbol) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, symbol := range symbols {
		index := -1
		for i, storedSymbol := range s.store {
			if storedSymbol.Symbol.SymbolCode == symbol.SymbolCode && storedSymbol.Symbol.Exchange == symbol.Exchange {
				index = i
				break
			}
		}

		if index > -1 {
			s.store[index].RemoveRequester(requester)
		}
	}

	ns := make([]*RegisterSymbol, len(s.store))
	nsIndex := 0
	for _, symbol := range s.store {
		if !symbol.IsRequestersEmpty() {
			ns[nsIndex] = symbol
			nsIndex++
		}
	}

	s.store = ns[:nsIndex]
}

type RegisterSymbol struct {
	Symbol     *kabuspb.RegisterSymbol
	Requesters []string
	mtx        sync.Mutex
}

func (r *RegisterSymbol) AddRequester(requester string) {
	if !r.HasRequester(requester) {
		r.mtx.Lock()
		defer r.mtx.Unlock()
		r.Requesters = append(r.Requesters, requester)
	}
}

func (r *RegisterSymbol) RemoveRequester(requester string) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	nr := make([]string, len(r.Requesters))
	i := 0
	for _, req := range r.Requesters {
		if req != requester {
			nr[i] = req
			i++
		}
	}
	r.Requesters = nr[:i]
}

func (r *RegisterSymbol) HasRequester(requester string) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, req := range r.Requesters {
		if req == requester {
			return true
		}
	}
	return false
}

func (r *RegisterSymbol) IsRequestersEmpty() bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return len(r.Requesters) == 0
}

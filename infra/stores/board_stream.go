package stores

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

var (
	boardStreamSingleton      repositories.BoardStreamStore
	boardStreamSingletonMutex sync.Mutex
)

func GetBoardStreamStore() repositories.BoardStreamStore {
	boardStreamSingletonMutex.Lock()
	defer boardStreamSingletonMutex.Unlock()

	if boardStreamSingleton == nil {
		boardStreamSingleton = &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{}}
	}

	return boardStreamSingleton
}

type boardStream struct {
	store []kabuspb.KabusService_GetBoardsStreamingServer
	mtx   sync.Mutex
}

func (s *boardStream) All() []kabuspb.KabusService_GetBoardsStreamingServer {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.store
}

func (s *boardStream) Add(stream kabuspb.KabusService_GetBoardsStreamingServer) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.store = append(s.store, stream)
}

func (s *boardStream) Remove(index int) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if index < 0 || len(s.store)-1 < index {
		return
	}

	s.store = append(s.store[:index], s.store[index+1:]...)
}

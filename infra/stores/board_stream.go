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
		boardStreamSingleton = &boardStream{store: []streamChan{}}
	}

	return boardStreamSingleton
}

type boardStream struct {
	store []streamChan
	mtx   sync.Mutex
}

type streamChan struct {
	stream kabuspb.KabusService_GetBoardsStreamingServer
	ch     chan error
}

func (s *boardStream) HasStream() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return len(s.store) > 0
}

func (s *boardStream) All() []kabuspb.KabusService_GetBoardsStreamingServer {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make([]kabuspb.KabusService_GetBoardsStreamingServer, len(s.store))
	for i, stream := range s.store {
		res[i] = stream.stream
	}
	return res
}

func (s *boardStream) Add(stream kabuspb.KabusService_GetBoardsStreamingServer, ch chan error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.store = append(s.store, streamChan{stream: stream, ch: ch})
}

func (s *boardStream) Remove(index int, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if index < 0 || len(s.store)-1 < index {
		return
	}

	stream := s.store[index]
	stream.ch <- err
	s.store = append(s.store[:index], s.store[index+1:]...)
}

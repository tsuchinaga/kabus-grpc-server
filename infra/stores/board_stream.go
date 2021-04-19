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
		boardStreamSingleton = &boardStream{store: map[int]streamChan{}}
	}

	return boardStreamSingleton
}

type boardStream struct {
	store map[int]streamChan
	seq   int
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

func (s *boardStream) All() map[int]kabuspb.KabusService_GetBoardsStreamingServer {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make(map[int]kabuspb.KabusService_GetBoardsStreamingServer)
	for i, stream := range s.store {
		res[i] = stream.stream
	}
	return res
}

func (s *boardStream) Add(stream kabuspb.KabusService_GetBoardsStreamingServer, ch chan error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.store[s.seq] = streamChan{stream: stream, ch: ch}
	s.seq++
}

func (s *boardStream) Remove(seq int, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.store[seq]; !ok {
		return
	}

	stream := s.store[seq]
	stream.ch <- err
	delete(s.store, seq)
}

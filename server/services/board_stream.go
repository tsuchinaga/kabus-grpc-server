package services

import (
	"log"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewBoardStreamService(streamStore repositories.BoardStreamStore, boardWS repositories.BoardWS) BoardStreamService {
	return &boardStream{
		streamStore: streamStore,
		boardWS:     boardWS,
	}
}

type BoardStreamService interface {
	Connect(stream kabuspb.KabusService_GetBoardsStreamingServer) error
}

type boardStream struct {
	streamStore repositories.BoardStreamStore
	boardWS     repositories.BoardWS
}

func (s *boardStream) Connect(stream kabuspb.KabusService_GetBoardsStreamingServer) error {
	if !s.boardWS.IsConnected() {
		var err error
		go func() {
			err = s.boardWS.Connect(s.onNext)
		}()

		// wsの接続に失敗するかもなので1s程度待ってあげる
		<-time.After(1 * time.Second)
		if err != nil {
			return err
		}
	}

	ch := make(chan error)
	s.streamStore.Add(stream, ch)
	return <-ch
}

func (s *boardStream) onNext(board *kabuspb.Board) error {
	for i, stream := range s.streamStore.All() {
		if err := stream.Send(board); err != nil {
			log.Println(err) // デバッグのためにおいとく
			s.streamStore.Remove(i, err)
		}
	}
	if !s.streamStore.HasStream() {
		if err := s.boardWS.Disconnect(); err != nil {
			log.Println(err) // デバッグのためにおいとく
		}
	}
	return nil
}

package services

import (
	"context"
	"log"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func NewBoardStreamService(streamStore repositories.BoardStreamStore, boardWS repositories.BoardWS, virtual repositories.VirtualSecurity) BoardStreamService {
	return &boardStream{
		streamStore: streamStore,
		boardWS:     boardWS,
		virtual:     virtual,
	}
}

type BoardStreamService interface {
	Start()
	Connect(stream kabuspb.KabusService_GetBoardsStreamingServer) error
}

type boardStream struct {
	streamStore repositories.BoardStreamStore
	boardWS     repositories.BoardWS
	virtual     repositories.VirtualSecurity
}

func (s *boardStream) Start() {
	if !s.boardWS.IsConnected() {
		go func() {
			err := s.boardWS.Connect(s.onNext)
			for i := range s.streamStore.All() {
				s.streamStore.Remove(i, err)
			}
		}()
	}
}

func (s *boardStream) Connect(stream kabuspb.KabusService_GetBoardsStreamingServer) error {
	ch := make(chan error)
	s.streamStore.Add(stream, ch)
	s.Start()
	return <-ch
}

func (s *boardStream) onNext(board *kabuspb.Board) error {
	for i, stream := range s.streamStore.All() {
		if err := stream.Send(board); err != nil {
			log.Println(err) // デバッグのためにおいとく
			s.streamStore.Remove(i, err)
		}
	}

	// 仮想証券会社への価格情報送信
	go func() {
		_ = s.virtual.SendPrice(context.Background(), board)
	}()

	// 仮想証券会社に価格情報を送る必要があるので切断をやめる
	// 登録銘柄がなくなれば切断してもいいけど、その場合は価格情報も来ないから影響がほとんどない
	//if !s.streamStore.HasStream() {
	//	if err := s.boardWS.Disconnect(); err != nil {
	//		log.Println(err) // デバッグのためにおいとく
	//	}
	//}
	return nil
}

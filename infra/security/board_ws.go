package security

import (
	"sync"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

var (
	boardWSSingleton      repositories.BoardWS
	boardWSSingletonMutex sync.Mutex
)

func GetBoardWS(isProduction bool) repositories.BoardWS {
	boardWSSingletonMutex.Lock()
	defer boardWSSingletonMutex.Unlock()

	if boardWSSingleton == nil {
		boardWSSingleton = &boardWS{ws: kabus.NewWSRequester(isProduction)}
	}

	return boardWSSingleton
}

type boardWS struct {
	ws     kabus.WSRequester
	onNext func(board *kabuspb.Board) error
}

func (s *boardWS) IsConnected() bool {
	return s.ws.IsOpened()
}

func (s *boardWS) Connect(onNext func(board *kabuspb.Board) error) error {
	if s.ws.IsOpened() {
		return nil
	}
	s.onNext = onNext
	s.ws.SetOnNext(s.callOnNext)

	return s.ws.Open()
}

func (s *boardWS) callOnNext(message kabus.PriceMessage) error {
	return s.onNext(fromPriceMessage(message))
}

func (s *boardWS) Disconnect() error {
	return s.ws.Close()
}

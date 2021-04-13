package repositories

import "gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

type BoardWS interface {
	IsConnected() bool
	Connect(onNext func(board *kabuspb.Board) error) error
	Disconnect() error
}

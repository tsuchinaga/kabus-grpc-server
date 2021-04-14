package repositories

import "gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

type BoardStreamStore interface {
	All() []kabuspb.KabusService_GetBoardsStreamingServer
	Add(stream kabuspb.KabusService_GetBoardsStreamingServer, ch chan error)
	Remove(index int, err error)
}

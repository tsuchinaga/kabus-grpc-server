package repositories

import "gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

type BoardStreamStore interface {
	HasStream() bool
	All() map[int]kabuspb.KabusService_GetBoardsStreamingServer
	Add(stream kabuspb.KabusService_GetBoardsStreamingServer, ch chan error)
	Remove(seq int, err error)
}

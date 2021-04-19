package services

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_NewBoardStreamService(t *testing.T) {
	streamStore := &testBoardStreamStore{}
	boardWS := &testBoardWS{}
	got := NewBoardStreamService(streamStore, boardWS)
	want := &boardStream{
		streamStore: streamStore,
		boardWS:     boardWS,
	}

	t.Parallel()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_boardStream_Connect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		isConnected bool
		connect     error
		all         map[int]kabuspb.KabusService_GetBoardsStreamingServer
		addCount    int
		removeCount int
	}{
		{name: "ws接続済みならAddするだけ", isConnected: true, addCount: 1},
		{name: "ws未接続ならws接続し、connectがreturnすればstoreのstream分だけremoveを叩く",
			isConnected: false,
			all:         map[int]kabuspb.KabusService_GetBoardsStreamingServer{0: &testGetBoardsStreamingServer{}, 1: &testGetBoardsStreamingServer{}, 2: &testGetBoardsStreamingServer{}},
			connect:     errors.New("error message"),
			addCount:    1,
			removeCount: 3},
		{name: "ws未接続ならws接続し、connectがreturnしたときにstoreのstreamがなければremoveは叩かれない",
			isConnected: false,
			all:         map[int]kabuspb.KabusService_GetBoardsStreamingServer{},
			connect:     nil,
			addCount:    1,
			removeCount: 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all}
			boardWS := &testBoardWS{isConnected: test.isConnected, connect: test.connect}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS}
			_ = service.Connect(nil)
			if test.addCount != streamStore.addCount || test.removeCount != streamStore.removeCount {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.addCount, test.removeCount, streamStore.addCount, streamStore.removeCount)
			}
		})
	}
}

func Test_boardStream_onNext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		all             map[int]kabuspb.KabusService_GetBoardsStreamingServer
		hasStream       bool
		removeCount     int
		disconnect      error
		disconnectCount int
		hasError        bool
	}{
		{name: "allで何もなければsendしない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: true, removeCount: 0},
		{name: "sendしたときにerrorがなければremoveされない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{
			0: &testGetBoardsStreamingServer{send: nil}, 1: &testGetBoardsStreamingServer{send: nil}, 2: &testGetBoardsStreamingServer{send: nil},
		}, hasStream: true, removeCount: 0},
		{name: "sendでerrorがあればremoveする", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{
			0: &testGetBoardsStreamingServer{send: errors.New("error message1")}, 1: &testGetBoardsStreamingServer{send: nil}, 2: &testGetBoardsStreamingServer{send: errors.New("error message1")},
		}, hasStream: true, removeCount: 2},
		{name: "hasStreamがfalseならDisconnectが叩かれる", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: false, removeCount: 0, disconnectCount: 1},
		{name: "hasStreamがfalseならDisconnectでエラーが出てもエラーは返されない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: false, disconnect: errors.New("error message"), removeCount: 0, disconnectCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all, hasStream: test.hasStream}
			boardWS := &testBoardWS{disconnect: test.disconnect}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS}
			got := service.onNext(nil)
			if (got != nil) != test.hasError || test.removeCount != streamStore.removeCount || test.disconnectCount != boardWS.disconnectCount {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.hasError, test.removeCount, test.disconnectCount, got, streamStore.removeCount, boardWS.disconnectCount)
			}
		})
	}
}

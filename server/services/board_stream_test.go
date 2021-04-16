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
		addCount    int
		hasError    bool
	}{
		{name: "ws接続済みならAddするだけ", isConnected: true, addCount: 1},
		{name: "ws未接続ならws接続し、errがあればerrを返す", isConnected: false, connect: errors.New("error message"), hasError: true},
		{name: "ws未接続ならws接続し、errがなければAddする", isConnected: false, connect: nil, addCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{}
			boardWS := &testBoardWS{isConnected: test.isConnected, connect: test.connect}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS}
			got := service.Connect(nil)
			if (got != nil) != test.hasError || test.addCount != streamStore.addCount {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.hasError, test.addCount, got, streamStore.addCount)
			}
		})
	}
}

func Test_boardStream_onNext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		all             []kabuspb.KabusService_GetBoardsStreamingServer
		hasStream       bool
		removeCount     int
		disconnectCount int
		hasError        bool
	}{
		{name: "allで何もなければsendしない", all: []kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: true, removeCount: 0},
		{name: "sendしたときにerrorがなければremoveされない", all: []kabuspb.KabusService_GetBoardsStreamingServer{
			&testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: nil},
		}, hasStream: true, removeCount: 0},
		{name: "sendでerrorがあればremoveする", all: []kabuspb.KabusService_GetBoardsStreamingServer{
			&testGetBoardsStreamingServer{send: errors.New("error message1")}, &testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: errors.New("error message1")},
		}, hasStream: true, removeCount: 2},
		{name: "hasStreamがfalseならDisconnectが叩かれる", all: []kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: false, removeCount: 0, disconnectCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all, hasStream: test.hasStream}
			boardWS := &testBoardWS{disconnect: nil}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS}
			got := service.onNext(nil)
			if (got != nil) != test.hasError || test.removeCount != streamStore.removeCount || test.disconnectCount != boardWS.disconnectCount {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(), test.hasError, test.removeCount, test.disconnectCount, got, streamStore.removeCount, boardWS.disconnectCount)
			}
		})
	}
}

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
		name        string
		all         []kabuspb.KabusService_GetBoardsStreamingServer
		removeCount int
		hasError    bool
	}{
		{name: "allで何もなければsendしない", all: []kabuspb.KabusService_GetBoardsStreamingServer{}, removeCount: 0},
		{name: "sendしたときにerrorがなければremoveされない", all: []kabuspb.KabusService_GetBoardsStreamingServer{
			&testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: nil},
		}, removeCount: 0},
		{name: "sendでerrorがあればremoveする", all: []kabuspb.KabusService_GetBoardsStreamingServer{
			&testGetBoardsStreamingServer{send: errors.New("error message1")}, &testGetBoardsStreamingServer{send: nil}, &testGetBoardsStreamingServer{send: errors.New("error message1")},
		}, removeCount: 2},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all}
			service := &boardStream{streamStore: streamStore}
			got := service.onNext(nil)
			if (got != nil) != test.hasError || test.removeCount != streamStore.removeCount {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.hasError, test.removeCount, got, streamStore.removeCount)
			}
		})
	}
}

package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_NewBoardStreamService(t *testing.T) {
	streamStore := &testBoardStreamStore{}
	boardWS := &testBoardWS{}
	virtual := &testVirtualSecurity{}
	got := NewBoardStreamService(streamStore, boardWS, virtual)
	want := &boardStream{
		streamStore: streamStore,
		boardWS:     boardWS,
		virtual:     virtual,
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
			time.Sleep(time.Second) // 非同期処理があるの少し待つ
			if test.addCount != streamStore.addCount || test.removeCount != streamStore.removeCount {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.addCount, test.removeCount, streamStore.addCount, streamStore.removeCount)
			}
		})
	}
}

func Test_boardStream_onNext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		all            map[int]kabuspb.KabusService_GetBoardsStreamingServer
		hasStream      bool
		removeCount    int
		hasError       bool
		sendPriceCount int
	}{
		{name: "allで何もなければsendしない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: true, removeCount: 0, sendPriceCount: 1},
		{name: "sendしたときにerrorがなければremoveされない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{
			0: &testGetBoardsStreamingServer{send: nil}, 1: &testGetBoardsStreamingServer{send: nil}, 2: &testGetBoardsStreamingServer{send: nil},
		}, hasStream: true, removeCount: 0, sendPriceCount: 1},
		{name: "sendでerrorがあればremoveする", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{
			0: &testGetBoardsStreamingServer{send: errors.New("error message1")}, 1: &testGetBoardsStreamingServer{send: nil}, 2: &testGetBoardsStreamingServer{send: errors.New("error message1")},
		}, hasStream: true, removeCount: 2, sendPriceCount: 1},
		// 仮想証券会社連携で切断する処理を停止したので、テストもコメントアウト
		//{name: "hasStreamがfalseならDisconnectが叩かれる", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: false, removeCount: 0, disconnectCount: 1},
		//{name: "hasStreamがfalseならDisconnectでエラーが出てもエラーは返されない", all: map[int]kabuspb.KabusService_GetBoardsStreamingServer{}, hasStream: false, disconnect: errors.New("error message"), removeCount: 0, disconnectCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all, hasStream: test.hasStream}
			boardWS := &testBoardWS{}
			virtual := &testVirtualSecurity{}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS, virtual: virtual}
			got := service.onNext(nil)
			time.Sleep(time.Second) // 非同期処理があるの少し待つ
			if (got != nil) != test.hasError || test.removeCount != streamStore.removeCount || test.sendPriceCount != virtual.sendPriceCount {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(),
					test.hasError, test.removeCount, test.sendPriceCount,
					got, streamStore.removeCount, virtual.sendPriceCount)
			}
		})
	}
}

func Test_boardStream_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		isConnected bool
		connect     error
		all         map[int]kabuspb.KabusService_GetBoardsStreamingServer
		removeCount int
	}{
		{name: "ws接続済みならなにもしない", isConnected: true},
		{name: "ws未接続ならws接続し、connectがreturnすればstoreのstream分だけremoveを叩く",
			isConnected: false,
			all:         map[int]kabuspb.KabusService_GetBoardsStreamingServer{0: &testGetBoardsStreamingServer{}, 1: &testGetBoardsStreamingServer{}, 2: &testGetBoardsStreamingServer{}},
			connect:     errors.New("error message"),
			removeCount: 3},
		{name: "ws未接続ならws接続し、connectがreturnしたときにstoreのstreamがなければremoveは叩かれない",
			isConnected: false,
			all:         map[int]kabuspb.KabusService_GetBoardsStreamingServer{},
			connect:     nil,
			removeCount: 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamStore := &testBoardStreamStore{all: test.all}
			boardWS := &testBoardWS{isConnected: test.isConnected, connect: test.connect}
			service := &boardStream{streamStore: streamStore, boardWS: boardWS}
			service.Start()
			time.Sleep(time.Second) // 非同期処理があるの少し待つ
			if test.removeCount != streamStore.removeCount {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.removeCount, streamStore.removeCount)
			}
		})
	}
}

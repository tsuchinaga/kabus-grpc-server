package stores

import (
	"reflect"
	"testing"

	"google.golang.org/grpc"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
)

func Test_GetBoardStreamStore(t *testing.T) {
	t.Parallel()
	got := GetBoardStreamStore()
	want := &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

type testBoardStreamServer struct {
	Foo string
	grpc.ServerStream
}

func (t *testBoardStreamServer) Send(*kabuspb.Board) error { return nil }

func Test_boardStream_All(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store repositories.BoardStreamStore
		want  []kabuspb.KabusService_GetBoardsStreamingServer
	}{
		{name: "storeが空配列なら空配列を返す", store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{}}, want: []kabuspb.KabusService_GetBoardsStreamingServer{}},
		{name: "storeに要素があればそれらすべてを返す",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}}},
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.store.All()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_boardStream_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *boardStream
		arg   kabuspb.KabusService_GetBoardsStreamingServer
		want  []kabuspb.KabusService_GetBoardsStreamingServer
	}{
		{name: "store空配列でも要素が追加される",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{}},
			arg:   &testBoardStreamServer{Foo: "foo"},
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}}},
		{name: "storeに要素があっても要素が追加される",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}}},
			arg:   &testBoardStreamServer{Foo: "foo"},
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "foo"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.Add(test.arg)
			got := test.store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_boardStream_Remove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *boardStream
		arg   int
		want  []kabuspb.KabusService_GetBoardsStreamingServer
	}{
		{name: "storeの先頭の要素を削除できる",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
			arg:   0,
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
		{name: "storeの中間の要素を削除できる",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
			arg:   1,
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "baz"}}},
		{name: "storeの末尾の要素を削除できる",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
			arg:   2,
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}}},
		{name: "storeに存在しないindexを指定されたら何も起こらない(-1)",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
			arg:   -1,
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
		{name: "storeに存在しないindexを指定されたら何も起こらない(1000)",
			store: &boardStream{store: []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
			arg:   1000,
			want:  []kabuspb.KabusService_GetBoardsStreamingServer{&testBoardStreamServer{Foo: "foo"}, &testBoardStreamServer{Foo: "bar"}, &testBoardStreamServer{Foo: "baz"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.Remove(test.arg)
			got := test.store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

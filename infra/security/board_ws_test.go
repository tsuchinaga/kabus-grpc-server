package security

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

type testWSRequester struct {
	kabus.WSRequester
	Foo           string
	isOpened      bool
	open          error
	close         error
	callSetOnNext int
}

func (t *testWSRequester) SetOnNext(func(kabus.PriceMessage) error) { t.callSetOnNext++ }
func (t *testWSRequester) IsOpened() bool                           { return t.isOpened }
func (t *testWSRequester) Open() error                              { return t.open }
func (t *testWSRequester) Close() error                             { return t.close }

func Test_GetBoardWS(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		boardWSSingleton repositories.BoardWS
		arg              bool
		want             repositories.BoardWS
	}{
		{name: "boardWSSingletonがnilなら引数の値をもったboardWSが返される",
			boardWSSingleton: nil,
			arg:              true,
			want:             &boardWS{ws: kabus.NewWSRequester(true)}},
		{name: "boardWSSingletonがnilでないならboardWSSingletonが返される",
			boardWSSingleton: &boardWS{ws: kabus.NewWSRequester(true)},
			arg:              false,
			want:             &boardWS{ws: kabus.NewWSRequester(true)}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			boardWSSingleton = test.boardWSSingleton
			got := GetBoardWS(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_boardWS_IsConnected(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		ws   *boardWS
		want bool
	}{
		{name: "IsOpenedがtrueならtrue", ws: &boardWS{ws: &testWSRequester{isOpened: true}}, want: true},
		{name: "IsOpenedがfalseならfalse", ws: &boardWS{ws: &testWSRequester{isOpened: false}}, want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.ws.IsConnected()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_boardWS_Disconnect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		ws       *boardWS
		hasError bool
	}{
		{name: "closeがerrorを返さなければnil", ws: &boardWS{ws: &testWSRequester{close: nil}}, hasError: false},
		{name: "closeがerrorを返せばそのerror", ws: &boardWS{ws: &testWSRequester{close: errors.New("error message")}}, hasError: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.ws.Disconnect()
			if (got != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.hasError, got)
			}
		})
	}
}

func Test_boardWS_Connect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		ws              *testWSRequester
		isCallSetOnNext bool
		hasError        bool
	}{
		{name: "isOpenedがtrueなら何もせずnilを返す", ws: &testWSRequester{isOpened: true}, hasError: false},
		{name: "isOpenがfalseならSetOnNextしてからopenし、openがnilならnil", ws: &testWSRequester{}, isCallSetOnNext: true, hasError: false},
		{name: "openがerrorならerrorを返す", ws: &testWSRequester{open: errors.New("error message")}, isCallSetOnNext: true, hasError: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ws := &boardWS{ws: test.ws}
			got := ws.Connect(nil)
			if (got != nil) != test.hasError || (test.ws.callSetOnNext > 0) != test.isCallSetOnNext {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.hasError, test.isCallSetOnNext, got, test.ws.callSetOnNext)
			}
		})
	}
}

func Test_boardWS_callOnNext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		onNext   func(board *kabuspb.Board) error
		hasError bool
	}{
		{name: "onNextがerrorを返したらerrorあり", onNext: func(*kabuspb.Board) error { return errors.New("error message") }, hasError: true},
		{name: "onNextがnilを返したらerrorなし", onNext: func(*kabuspb.Board) error { return nil }, hasError: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ws := &boardWS{onNext: test.onNext}
			got := ws.callOnNext(kabus.PriceMessage{})
			if (got != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.hasError, got)
			}
		})
	}
}

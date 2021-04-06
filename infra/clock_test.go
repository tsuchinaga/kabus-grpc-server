package infra

import (
	"reflect"
	"testing"
	"time"
)

func Test_NewClock(t *testing.T) {
	t.Parallel()

	got := NewClock()
	want := &clock{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_clock_Now(t *testing.T) {
	t.Parallel()
	got := NewClock().Now()
	want := time.Now().Add(-time.Millisecond)
	if !got.After(want) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

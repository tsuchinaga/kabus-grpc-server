package infra

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func Test_setting_Password(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		setting repositories.Setting
		want    string
	}{
		{name: "空文字なら空文字を返す", setting: &setting{password: ""}, want: ""},
		{name: "文字列が設定されていれば文字列を返す", setting: &setting{password: "Password1234"}, want: "Password1234"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.setting.Password()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_IsProduction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		setting repositories.Setting
		want    bool
	}{
		{name: "isProdがtrueならtrue", setting: &setting{isProd: true}, want: true},
		{name: "isProdがfalseならfalse", setting: &setting{isProd: false}, want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.setting.IsProduction()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

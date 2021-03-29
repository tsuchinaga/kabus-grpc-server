package stores

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/server/repositories"
)

func Test_GetTokenStore(t *testing.T) {
	t.Parallel()
	got := GetTokenStore()
	want := &token{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_token_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token repositories.TokenStore
		want  string
	}{
		{name: "tokenが空文字なら空文字を返す", token: &token{token: ""}, want: ""},
		{name: "tokenに文字列があればその文字列を返す", token: &token{token: "TOKEN_STRING"}, want: "TOKEN_STRING"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.token.GetToken()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_token_SetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token repositories.TokenStore
		arg1  string
		arg2  time.Time
		want  repositories.TokenStore
	}{
		{name: "引数のトークンと有効期限がセットされる",
			token: &token{token: "TOKEN_STRING", expire: time.Date(2021, 3, 23, 8, 0, 0, 0, time.Local)},
			arg1:  "NEW_TOKEN_STRING",
			arg2:  time.Date(2021, 3, 23, 16, 0, 0, 0, time.Local),
			want:  &token{token: "NEW_TOKEN_STRING", expire: time.Date(2021, 3, 23, 16, 0, 0, 0, time.Local)}},
		{name: "元がゼロ値でも引数のトークンと有効期限がセットされる",
			token: &token{},
			arg1:  "NEW_TOKEN_STRING",
			arg2:  time.Date(2021, 3, 23, 16, 0, 0, 0, time.Local),
			want:  &token{token: "NEW_TOKEN_STRING", expire: time.Date(2021, 3, 23, 16, 0, 0, 0, time.Local)}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.token.SetToken(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, test.token) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, test.token)
			}
		})
	}
}

func Test_token_IsExpired(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token repositories.TokenStore
		arg   time.Time
		want  bool
	}{
		{name: "有効期限がゼロ値ならtrue",
			token: &token{},
			arg:   time.Date(2021, 3, 23, 9, 0, 0, 0, time.Local),
			want:  true},
		{name: "有効期限が引数より過去ならtrue",
			token: &token{expire: time.Date(2021, 3, 23, 9, 0, 0, 0, time.Local)},
			arg:   time.Date(2021, 3, 23, 9, 1, 0, 0, time.Local),
			want:  true},
		{name: "有効期限が引数と同じならfalse",
			token: &token{expire: time.Date(2021, 3, 23, 9, 0, 0, 0, time.Local)},
			arg:   time.Date(2021, 3, 23, 9, 0, 0, 0, time.Local),
			want:  false},
		{name: "有効期限が引数より未来ならfalse",
			token: &token{expire: time.Date(2021, 3, 23, 9, 0, 0, 0, time.Local)},
			arg:   time.Date(2021, 3, 23, 8, 59, 0, 0, time.Local),
			want:  false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.token.IsExpired(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_token_Reset(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token repositories.TokenStore
		want  repositories.TokenStore
	}{
		{name: "元々ゼロ値ならゼロ値のままにする", token: &token{}, want: &token{}},
		{name: "ゼロ値で更新する",
			token: &token{token: "TOKEN_STRING", expire: time.Date(2021, 3, 24, 2, 12, 0, 0, time.Local)},
			want:  &token{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.token.Reset()
			if !reflect.DeepEqual(test.want, test.token) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, test.token)
			}
		})
	}
}

func Test_token_GetExpiredAt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token repositories.TokenStore
		want  time.Time
	}{
		{name: "ゼロ値でも返される",
			token: &token{expire: time.Time{}},
			want:  time.Time{}},
		{name: "有効な日付を返される",
			token: &token{expire: time.Date(2021, 3, 30, 2, 49, 0, 0, time.Local)},
			want:  time.Date(2021, 3, 30, 2, 49, 0, 0, time.Local)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.token.GetExpiredAt()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

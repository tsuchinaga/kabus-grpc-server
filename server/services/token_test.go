package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func Test_NewTokenService(t *testing.T) {
	t.Parallel()
	got := NewTokenService(&testTokenStore{}, &testSecurity{}, &testClock{}, &testSetting{})
	want := &token{tokenStore: &testTokenStore{}, security: &testSecurity{}, clock: &testClock{}, setting: &testSetting{}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_token_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		now           time.Time
		isExpired     bool
		getToken      string
		token1        string
		token2        error
		want          string
		hasError      bool
		wantSetToken1 string
		wantSetToken2 time.Time
	}{
		{name: "有効期限切れでなければ現在のトークンを返す",
			now:       time.Date(2021, 3, 25, 12, 0, 0, 0, time.Local),
			isExpired: false,
			getToken:  "TOKEN_STRING",
			want:      "TOKEN_STRING",
			hasError:  false},
		{name: "有効期限切れでトークン取得に失敗したらエラーを返す",
			now:       time.Date(2021, 3, 25, 12, 0, 0, 0, time.Local),
			isExpired: true,
			getToken:  "TOKEN_STRING",
			token1:    "",
			token2:    errors.New("error message"),
			want:      "",
			hasError:  true},
		{name: "有効期限切れでトークン取得に成功したらトークンと次の有効期限(翌日)をセットする",
			now:           time.Date(2021, 3, 25, 12, 0, 0, 0, time.Local),
			isExpired:     true,
			getToken:      "TOKEN_STRING",
			token1:        "NEW_TOKEN_STRING",
			token2:        nil,
			want:          "TOKEN_STRING",
			hasError:      false,
			wantSetToken1: "NEW_TOKEN_STRING",
			wantSetToken2: time.Date(2021, 3, 26, 6, 30, 0, 0, time.Local)},
		{name: "有効期限切れでトークン取得に成功したらトークンと次の有効期限(当日)をセットする",
			now:           time.Date(2021, 3, 25, 3, 0, 0, 0, time.Local),
			isExpired:     true,
			getToken:      "TOKEN_STRING",
			token1:        "NEW_TOKEN_STRING",
			token2:        nil,
			want:          "TOKEN_STRING",
			hasError:      false,
			wantSetToken1: "NEW_TOKEN_STRING",
			wantSetToken2: time.Date(2021, 3, 25, 6, 30, 0, 0, time.Local)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			tokenStore := &testTokenStore{getToken: test.getToken, isExpired: test.isExpired}
			security := &testSecurity{token1: test.token1, token2: test.token2}
			clock := &testClock{now: test.now}
			setting := &testSetting{}
			token := &token{tokenStore: tokenStore, security: security, clock: clock, setting: setting}
			got, err := token.GetToken(context.Background())
			if !reflect.DeepEqual(test.want, got) ||
				(err != nil) != test.hasError ||
				!reflect.DeepEqual(tokenStore.lastSetToken1, test.wantSetToken1) ||
				!reflect.DeepEqual(tokenStore.lastSetToken2, test.wantSetToken2) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					test.want, test.hasError, test.wantSetToken1, test.wantSetToken2,
					got, err, tokenStore.lastSetToken1, tokenStore.lastSetToken2)
			}
		})
	}
}

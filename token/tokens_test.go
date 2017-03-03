package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestNewTokens(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		input token.Tokens
		want  token.Tokens
	}{
		{
			token.Tokens{},
			token.Tokens{},
		},
		{
			token.Tokens{nil},
			token.Tokens{},
		},
		{
			token.Tokens{
				token.Word("foo"),
			},
			token.Tokens{
				token.Word("foo"),
			},
		},
		{
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
				token.Word("baz"),
				token.Word("qux"),
			},
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
				token.Word("baz"),
				token.Word("qux"),
			},
		},
		{
			token.Tokens{
				nil,
				token.Word("foo"),
				nil,
				token.Word("bar"),
				nil,
				token.Word("baz"),
				nil,
				token.Word("qux"),
				nil,
			},
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
				token.Word("baz"),
				token.Word("qux"),
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := token.NewTokens(c.input...)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestTokensAppend(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		tokens1 token.Tokens
		tokens2 token.Tokens
		want    token.Tokens
	}{
		{
			token.Tokens{},
			token.Tokens{},
			token.Tokens{},
		},
		{
			token.Tokens{
				token.Word("foo"),
			},
			token.Tokens{},
			token.Tokens{
				token.Word("foo"),
			},
		},
		{
			token.Tokens{},
			token.Tokens{
				token.Word("foo"),
			},
			token.Tokens{
				token.Word("foo"),
			},
		},
		{
			token.Tokens{
				token.Word("foo"),
			},
			token.Tokens{
				token.Word("bar"),
			},
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
			},
		},
		{
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
			},
			token.Tokens{
				token.Word("baz"),
				token.Word("qux"),
			},
			token.Tokens{
				token.Word("foo"),
				token.Word("bar"),
				token.Word("baz"),
				token.Word("qux"),
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := c.tokens1.Append(c.tokens2...)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestTokensSprint(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		tokens   token.Tokens
		standard string
		mySQL    string
	}{
		{
			token.Tokens{
				token.QuoteStart,
				token.Word("foo"),
				token.QuoteEnd,
			},
			`"foo"`,
			"`foo`",
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Standard", i), func(t *testing.T) {
			t.Parallel()
			sql := c.tokens.Sprint(token.StandardFormat)
			if sql != c.standard {
				t.Error(diff.SQL(sql, c.standard))
			}
		})
		t.Run(fmt.Sprintf("%d MySQL", i), func(t *testing.T) {
			t.Parallel()
			sql := c.tokens.Sprint(token.MySQLFormat)
			if sql != c.mySQL {
				t.Error(diff.SQL(sql, c.mySQL))
			}
		})
	}
}

package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/token"
)

func TestRepeat(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		token token.Token
		n     int
		want  token.Tokens
	}{
		{
			token.Space,
			0,
			token.Tokens{},
		},
		{
			token.Space,
			1,
			token.Tokens{
				token.Space,
			},
		},
		{
			token.Space,
			2,
			token.Tokens{
				token.Space,
				token.Space,
			},
		},
		{
			token.Space,
			5,
			token.Tokens{
				token.Space,
				token.Space,
				token.Space,
				token.Space,
				token.Space,
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := token.Repeat(c.token, c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		tokensList []token.Tokens
		want       token.Tokens
	}{
		{
			[]token.Tokens{},
			token.Tokens{},
		},
		{
			[]token.Tokens{
				{
					token.Space,
					token.Comma,
				},
			},
			token.Tokens{
				token.Space,
				token.Comma,
			},
		},
		{
			[]token.Tokens{
				{
					token.Space,
					token.Comma,
				},
				{
					token.Comma,
					token.LParen,
					token.RParen,
				},
			},
			token.Tokens{
				token.Space,
				token.Comma,
				token.Comma,
				token.LParen,
				token.RParen,
			},
		},
		{
			[]token.Tokens{
				{
					token.Space,
				},
				{
					token.Comma,
				},
				{
					token.Comma,
				},
				{},
				{
					token.RParen,
				},
			},
			token.Tokens{
				token.Space,
				token.Comma,
				token.Comma,
				token.RParen,
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := token.Flatten(c.tokensList...)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestPlaceholders(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		n    int
		want token.Tokens
	}{
		{
			0,
			token.Tokens{},
		},
		{
			1,
			token.Tokens{
				token.Placeholder,
			},
		},
		{
			2,
			token.Tokens{
				token.Placeholder,
				token.Comma,
				token.Placeholder,
			},
		},
		{
			5,
			token.Tokens{
				token.Placeholder,
				token.Comma,
				token.Placeholder,
				token.Comma,
				token.Placeholder,
				token.Comma,
				token.Placeholder,
				token.Comma,
				token.Placeholder,
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := token.Placeholders(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

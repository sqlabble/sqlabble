package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestRepeat(t *testing.T) {
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := token.Repeat(c.token, c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestFlatten(t *testing.T) {
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
					token.ParenStart,
					token.ParenEnd,
				},
			},
			token.Tokens{
				token.Space,
				token.Comma,
				token.Comma,
				token.ParenStart,
				token.ParenEnd,
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
					token.ParenEnd,
				},
			},
			token.Tokens{
				token.Space,
				token.Comma,
				token.Comma,
				token.ParenEnd,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := token.Flatten(c.tokensList...)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestPlaceholders(t *testing.T) {
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := token.Placeholders(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestWord(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		token  token.Token
		sprint string
		string string
	}{
		{
			token.Space,
			" ",
			"Word( )",
		},
		{
			token.Placeholder,
			"?",
			"Word(?)",
		},
		{
			token.Comma,
			",",
			"Word(,)",
		},
		{
			token.ParenStart,
			"(",
			"Word(()",
		},
		{
			token.ParenEnd,
			")",
			"Word())",
		},
		{
			token.FuncParenStart,
			"(",
			"Word(()",
		},
		{
			token.FuncParenEnd,
			")",
			"Word())",
		},
		{
			token.NewWord("foo"),
			"foo",
			"Word(foo)",
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Sprint(StandardFormat)", i), func(t *testing.T) {
			t.Parallel()
			got := c.token.Sprint(token.StandardFormat)
			if got != c.sprint {
				t.Error(diff.Values(got, c.sprint))
			}
		})
		t.Run(fmt.Sprintf("%d String()", i), func(t *testing.T) {
			t.Parallel()
			got := c.token.String()
			if got != c.string {
				t.Error(diff.Values(got, c.string))
			}
		})
	}
}

func TestSameWord(t *testing.T) {
	t.Parallel()
	if token.ParenStart == token.FuncParenStart {
		t.Error("ParenStart and FuncParenStart should be different")
	}
	if token.ParenEnd == token.FuncParenEnd {
		t.Error("ParenEnd and FuncParenEnd should be different")
	}
}

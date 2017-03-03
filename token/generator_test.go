package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestGenerate(t *testing.T) {
	for i, c := range []struct {
		tokens          token.Tokens
		withBreaking    string
		withoutBreaking string
	}{
		{
			token.Tokens{},
			``,
			``,
		},
		{
			token.Tokens{
				token.LineStart, token.Word("foo"), token.LineEnd,
				token.LineStart, token.Indent, token.Word("bar"), token.ParenStart, token.Word("a"), token.Comma, token.Word("b"), token.ParenEnd, token.LineEnd,
				token.LineStart, token.Indent, token.Word("baz"), token.FuncParenStart, token.Word("c"), token.Comma, token.Word("d"), token.FuncParenEnd, token.LineEnd,
			},
			`foo
  bar (a, b)
  baz(c, d)
`,
			`foo bar (a, b) baz(c, d)`,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d WithBreaking", i), func(t *testing.T) {
			got := token.Generate(c.tokens, token.StandardIndentedFormat)
			if got != c.withBreaking {
				t.Error(diff.SQL(got, c.withBreaking))
			}
		})
		t.Run(fmt.Sprintf("%d WithoutBreaking", i), func(t *testing.T) {
			got := token.Generate(c.tokens, token.StandardFormat)
			if got != c.withoutBreaking {
				t.Error(diff.SQL(got, c.withBreaking))
			}
		})
	}
}

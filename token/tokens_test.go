package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestToken(t *testing.T) {
	for i, c := range []struct {
		tokens   token.Tokens
		standard string
		mySQL    string
	}{
		{
			token.Tokens{
				token.Quote,
				token.Word("foo"),
				token.Quote,
			},
			`"foo"`,
			"`foo`",
		},
	} {
		t.Run(fmt.Sprintf("%d Standard", i), func(t *testing.T) {
			sql := c.tokens.String(token.StandardFormat)
			if sql != c.standard {
				t.Error(diff.SQL(sql, c.standard))
			}
		})
		t.Run(fmt.Sprintf("%d MySQL", i), func(t *testing.T) {
			sql := c.tokens.String(token.MySQLFormat)
			if sql != c.mySQL {
				t.Error(diff.SQL(sql, c.mySQL))
			}
		})
	}
}

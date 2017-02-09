package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/token"
)

func TestPlaceholders(t *testing.T) {
	for _, c := range []struct {
		input int
		want  string
	}{
		{0, ""},
		{1, "?"},
		{2, "?, ?"},
		{3, "?, ?, ?"},
	} {
		t.Run(fmt.Sprintf("Placeholder(%d) should be '%s'", c.input, c.want), func(t *testing.T) {
			got := token.Placeholders(c.input)
			if got != c.want {
				t.Error(diff.SQL(got, c.want))
			}
		})
	}
}

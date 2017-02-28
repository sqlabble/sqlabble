package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestNewFormat(t *testing.T) {
	for i, c := range []struct {
		prefix  string
		indent  string
		quote   string
		lineEnd string
		want    token.Format
	}{
		{
			"",
			"",
			"",
			"",
			token.Format{
				Prefix:     "",
				Indent:     "",
				IsBreaking: false,
				Quote:      `"`,
				LineEnd:    "\n",
			},
		},
		{
			">>",
			"--",
			"'",
			"\r",
			token.Format{
				Prefix:     ">>",
				Indent:     "--",
				IsBreaking: true,
				Quote:      `'`,
				LineEnd:    "\r",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := token.NewFormat(c.prefix, c.indent, c.quote, c.lineEnd)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

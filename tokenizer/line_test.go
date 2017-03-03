package tokenizer_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

func TestLine(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		line  tokenizer.Line
		depth int
		want  token.Tokens
	}{
		{
			tokenizer.NewLine(),
			0,
			token.NewTokens(),
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got := c.line.Tokenize(c.depth)
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

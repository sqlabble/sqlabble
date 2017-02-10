package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/token"
)

func TestSelectClauseIsCaluse(t *testing.T) {
	if _, ok := interface{}(token.Select{}).(grammar.Clause); !ok {
		t.Errorf("token.SelectClause doesn't implement grammar.Clause")
	}
}

func TestSelectClause(t *testing.T) {
	for i, c := range []struct {
		statement token.Select
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewSelect(
				token.NewColumn("*"),
			),
			"SELECT *",
			`> SELECT
>   *
`,
			[]interface{}{},
		},
		{
			token.NewSelect(
				token.NewColumn("foo"),
			),
			"SELECT foo",
			`> SELECT
>   foo
`,
			[]interface{}{},
		},
		{
			token.NewSelect(
				token.NewColumn("foo").As("a"),
			),
			"SELECT foo AS a",
			`> SELECT
>   foo AS a
`,
			[]interface{}{},
		},
		{
			token.NewSelect(
				token.NewColumn("foo").As("a"),
				token.NewColumn("bar").As("b"),
				token.NewColumn("baz").As("c"),
			),
			"SELECT foo AS a, bar AS b, baz AS c",
			`> SELECT
>   foo AS a
>   , bar AS b
>   , baz AS c
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "> ", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

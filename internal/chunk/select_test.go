package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestSelectClauseIsCaluse(t *testing.T) {
	if _, ok := interface{}(chunk.Select{}).(grammar.Clause); !ok {
		t.Errorf("chunk.SelectClause doesn't implement grammar.Clause")
	}
}

func TestSelectClause(t *testing.T) {
	for i, c := range []struct {
		statement chunk.Select
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewSelect(
				chunk.NewColumn("*"),
			),
			"SELECT *",
			`> SELECT
>   *
`,
			[]interface{}{},
		},
		{
			chunk.NewSelect(
				chunk.NewColumn("foo"),
			),
			"SELECT foo",
			`> SELECT
>   foo
`,
			[]interface{}{},
		},
		{
			chunk.NewSelect(
				chunk.NewColumn("foo").As("a"),
			),
			"SELECT foo AS a",
			`> SELECT
>   foo AS a
`,
			[]interface{}{},
		},
		{
			chunk.NewSelect(
				chunk.NewColumn("foo").As("a"),
				chunk.NewColumn("bar").As("b"),
				chunk.NewColumn("baz").As("c"),
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

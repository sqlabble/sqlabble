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

func TestHavingType(t *testing.T) {
	if _, ok := interface{}(chunk.Having{}).(grammar.Clause); !ok {
		t.Errorf("chunk.Having doesn't implement grammar.Clause")
	}
}

func TestHavingSQL(t *testing.T) {
	for i, c := range []struct {
		statement chunk.Having
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewHaving(
				chunk.NewColumn("foo").Eq(100),
			),
			"HAVING foo = ?",
			`> HAVING
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			chunk.NewHaving(
				chunk.NewAnd(
					chunk.NewColumn("foo").Eq(100),
					chunk.NewColumn("bar").Eq("abc"),
				),
			),
			"HAVING foo = ? AND bar = ?",
			`> HAVING
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{100, "abc"},
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

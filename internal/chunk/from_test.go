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

func TestFromType(t *testing.T) {
	if _, ok := interface{}(chunk.From{}).(grammar.Clause); !ok {
		t.Errorf("chunk.FromClause doesn't implement grammar.Clause")
	}
}

func TestFromSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewFrom(
				chunk.NewTable("foo"),
			),
			"FROM foo",
			`> FROM
>   foo
`,
			[]interface{}{},
		},
		{
			chunk.NewFrom(
				chunk.NewTable("foo").As("a"),
			),
			"FROM foo AS a",
			`> FROM
>   foo AS a
`,
			[]interface{}{},
		},
		{
			chunk.NewFrom(
				chunk.NewTable("foo").As("a").Join(
					chunk.NewTable("bar"),
				),
			),
			"FROM foo AS a JOIN bar",
			`> FROM
>   foo AS a
>   JOIN bar
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

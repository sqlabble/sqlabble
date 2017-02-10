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

func TestGroupByType(t *testing.T) {
	if _, ok := interface{}(chunk.GroupBy{}).(grammar.Clause); !ok {
		t.Errorf("chunk.GroupBy doesn't implement grammar.Clause")
	}
}

func TestGroupBySQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewGroupBy(
				chunk.NewColumn("foo"),
			),
			"GROUP BY foo",
			`> GROUP BY
>   foo
`,
			[]interface{}{},
		},
		{
			chunk.NewGroupBy(
				chunk.NewColumn("foo"),
				chunk.NewColumn("bar"),
				chunk.NewColumn("baz"),
			),
			"GROUP BY foo, bar, baz",
			`> GROUP BY
>   foo
>   , bar
>   , baz
`,
			[]interface{}{},
		},
		{
			chunk.NewGroupBy(
				chunk.NewColumn("foo"),
			).Having(
				chunk.NewColumn("COUNT(*)").Gt(200),
			),
			"GROUP BY foo HAVING COUNT(*) > ?",
			`> GROUP BY
>   foo
> HAVING
>   COUNT(*) > ?
`,
			[]interface{}{
				200,
			},
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

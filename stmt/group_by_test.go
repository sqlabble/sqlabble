package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestGroupBySQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewGroupBy(
				stmt.NewColumn("foo"),
			),
			"GROUP BY foo",
			`> GROUP BY
>   foo
`,
			nil,
		},
		{
			stmt.NewGroupBy(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
				stmt.NewColumn("baz"),
			),
			"GROUP BY foo, bar, baz",
			`> GROUP BY
>   foo
>   , bar
>   , baz
`,
			nil,
		},
		{
			stmt.NewGroupBy(
				stmt.NewColumn("foo"),
			).Having(
				stmt.NewColumn("COUNT(*)").Gt(stmt.NewParam(200)),
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
		{
			stmt.NewGroupBy(
				stmt.NewColumn("foo"),
			).OrderBy(
				stmt.NewColumn("age").Asc(),
			),
			"GROUP BY foo ORDER BY age ASC",
			`> GROUP BY
>   foo
> ORDER BY
>   age ASC
`,
			nil,
		},
		{
			stmt.NewGroupBy(
				stmt.NewColumn("foo"),
			).Limit(20),
			"GROUP BY foo LIMIT ?",
			`> GROUP BY
>   foo
> LIMIT
>   ?
`,
			[]interface{}{
				20,
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

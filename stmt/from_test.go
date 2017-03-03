package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestFromSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewFrom(
				stmt.NewTable("foo"),
			),
			`FROM foo`,
			`> FROM
>   foo
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo").As("a"),
			),
			`FROM foo AS "a"`,
			`> FROM
>   foo AS "a"
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo").As("a").
					Join(stmt.NewTable("bar")).Using(stmt.NewColumn("id")),
			),
			`FROM foo AS "a" JOIN bar USING id`,
			`> FROM
>   foo AS "a"
>   JOIN bar USING id
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
					),
				).As("f"),
			),
			`FROM (SELECT foo) AS f`,
			`> FROM
>   (
>     SELECT
>       foo
>   ) AS f
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo"),
			).Where(
				stmt.NewColumn("age").Gt(stmt.NewVal(20)),
			),
			`FROM foo WHERE age > ?`,
			`> FROM
>   foo
> WHERE
>   age > ?
`,
			[]interface{}{
				20,
			},
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo"),
			).GroupBy(
				stmt.NewColumn("id"),
			),
			`FROM foo GROUP BY id`,
			`> FROM
>   foo
> GROUP BY
>   id
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo"),
			).OrderBy(
				stmt.NewColumn("age").Asc(),
			),
			`FROM foo ORDER BY age ASC`,
			`> FROM
>   foo
> ORDER BY
>   age ASC
`,
			nil,
		},
		{
			stmt.NewFrom(
				stmt.NewTable("foo"),
			).Limit(20),
			`FROM foo LIMIT ?`,
			`> FROM
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

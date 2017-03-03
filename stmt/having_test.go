package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestHavingSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewHaving(
				stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
			),
			"HAVING foo = ?",
			`> HAVING
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			stmt.NewHaving(
				stmt.NewAnd(
					stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
					stmt.NewColumn("bar").Eq(stmt.NewParam("abc")),
				),
			),
			"HAVING foo = ? AND bar = ?",
			`> HAVING
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{100, "abc"},
		},
		{
			stmt.NewHaving(
				stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
			).OrderBy(
				stmt.NewColumn("age").Asc(),
			),
			"HAVING foo = ? ORDER BY age ASC",
			`> HAVING
>   foo = ?
> ORDER BY
>   age ASC
`,
			[]interface{}{100},
		},
		{
			stmt.NewHaving(
				stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
			).Limit(20),
			"HAVING foo = ? LIMIT ?",
			`> HAVING
>   foo = ?
> LIMIT
>   ?
`,
			[]interface{}{100, 20},
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

package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestHavingSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewHaving(
				statement.NewColumn("foo").Eq(statement.NewParam(100)),
			),
			"HAVING foo = ?",
			`> HAVING
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			statement.NewHaving(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam("abc")),
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
			statement.NewHaving(
				statement.NewColumn("foo").Eq(statement.NewParam(100)),
			).OrderBy(
				statement.NewColumn("age").Asc(),
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
			statement.NewHaving(
				statement.NewColumn("foo").Eq(statement.NewParam(100)),
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
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

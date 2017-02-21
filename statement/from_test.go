package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestFromType(t *testing.T) {
	f := statement.From{}
	if _, ok := interface{}(f).(statement.Clause); !ok {
		t.Errorf("%T should implement statement.Clause", f)
	}
}

func TestFromSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement.NewFrom(
				statement.NewTable("foo"),
			),
			`FROM foo`,
			`> FROM
>   foo
`,
			[]interface{}{},
		},
		// 1
		{
			statement.NewFrom(
				statement.NewTable("foo").As("a"),
			),
			`FROM foo AS "a"`,
			`> FROM
>   foo AS "a"
`,
			[]interface{}{},
		},
		// 2
		{
			statement.NewFrom(
				statement.NewTable("foo").As("a").Join(
					statement.NewTable("bar"),
				),
			),
			`FROM foo AS "a" JOIN bar`,
			`> FROM
>   foo AS "a"
>   JOIN bar
`,
			[]interface{}{},
		},
		// 3
		{
			statement.NewFrom(
				statement.NewTable("foo"),
			).Where(
				statement.NewColumn("age").Gt(statement.NewParam(20)),
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
		// 4
		{
			statement.NewFrom(
				statement.NewTable("foo"),
			).GroupBy(
				statement.NewColumn("id"),
			),
			`FROM foo GROUP BY id`,
			`> FROM
>   foo
> GROUP BY
>   id
`,
			[]interface{}{},
		},
		// 5
		{
			statement.NewFrom(
				statement.NewTable("foo"),
			).OrderBy(
				statement.NewColumn("age").Asc(),
			),
			`FROM foo ORDER BY age ASC`,
			`> FROM
>   foo
> ORDER BY
>   age ASC
`,
			[]interface{}{},
		},
		// 6
		{
			statement.NewFrom(
				statement.NewTable("foo"),
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
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
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

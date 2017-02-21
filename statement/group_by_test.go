package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestGroupByType(t *testing.T) {
	g := statement.GroupBy{}
	if _, ok := interface{}(g).(statement.Clause); !ok {
		t.Errorf("%T doesn't implement statement.Clause", g)
	}
}

func TestGroupBySQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewGroupBy(
				statement.NewColumn("foo"),
			),
			"GROUP BY foo",
			`> GROUP BY
>   foo
`,
			[]interface{}{},
		},
		{
			statement.NewGroupBy(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
				statement.NewColumn("baz"),
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
			statement.NewGroupBy(
				statement.NewColumn("foo"),
			).Having(
				statement.NewColumn("COUNT(*)").Gt(statement.NewParam(200)),
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
			statement.NewGroupBy(
				statement.NewColumn("foo"),
			).OrderBy(
				statement.NewColumn("age").Asc(),
			),
			"GROUP BY foo ORDER BY age ASC",
			`> GROUP BY
>   foo
> ORDER BY
>   age ASC
`,
			[]interface{}{},
		},
		{
			statement.NewGroupBy(
				statement.NewColumn("foo"),
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

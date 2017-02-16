package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestGroupByType(t *testing.T) {
// 	if _, ok := interface{}(statement.GroupBy{}).(statement.ClauseNode); !ok {
// 		t.Errorf("statement.GroupBy doesn't implement statement.Clause")
// 	}
// }

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
				statement.NewColumn("COUNT(*)").Gt(200),
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

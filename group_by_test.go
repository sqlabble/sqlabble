package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestGroupByType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.GroupBy{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.GroupBy doesn't implement sqlabble.Clause")
// 	}
// }

func TestGroupBySQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewGroupBy(
				sqlabble.NewColumn("foo"),
			),
			"GROUP BY foo",
			`> GROUP BY
>   foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewGroupBy(
				sqlabble.NewColumn("foo"),
				sqlabble.NewColumn("bar"),
				sqlabble.NewColumn("baz"),
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
			sqlabble.NewGroupBy(
				sqlabble.NewColumn("foo"),
			).Having(
				sqlabble.NewColumn("COUNT(*)").Gt(200),
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
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

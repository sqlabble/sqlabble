package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestFromType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.From{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.FromClause doesn't implement sqlabble.Clause")
// 	}
// }

func TestFromSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			sqlabble.NewFrom(
				sqlabble.NewTable("foo"),
			),
			"FROM foo",
			`> FROM
>   foo
`,
			[]interface{}{},
		},
		// 1
		{
			sqlabble.NewFrom(
				sqlabble.NewTable("foo").As("a"),
			),
			"FROM foo AS a",
			`> FROM
>   foo AS a
`,
			[]interface{}{},
		},
		// 2
		{
			sqlabble.NewFrom(
				sqlabble.NewTable("foo").As("a").Join(
					sqlabble.NewTable("bar"),
				),
			),
			"FROM foo AS a JOIN bar",
			`> FROM
>   foo AS a
>   JOIN bar
`,
			[]interface{}{},
		},
		// 3
		{
			sqlabble.NewFrom(
				sqlabble.NewTable("foo"),
			).Where(
				sqlabble.NewColumn("age").Gt(20),
			),
			"FROM foo WHERE age > ?",
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
			sqlabble.NewFrom(
				sqlabble.NewTable("foo"),
			).OrderBy(
				sqlabble.NewColumn("age").Asc(),
			),
			"FROM foo ORDER BY age ASC",
			`> FROM
>   foo
> ORDER BY
>   age ASC
`,
			[]interface{}{},
		},
		// 5
		{
			sqlabble.NewFrom(
				sqlabble.NewTable("foo"),
			).GroupBy(
				sqlabble.NewColumn("id"),
			),
			"FROM foo GROUP BY id",
			`> FROM
>   foo
> GROUP BY
>   id
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

package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestSelectType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.Select{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.Select doesn't implement sqlabble.Clause")
// 	}
// }

func TestSelectSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewSelect(),
			"SELECT",
			`> SELECT
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSelect(
				sqlabble.NewColumn("*"),
			),
			"SELECT *",
			`> SELECT
>   *
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSelect(
				sqlabble.NewColumn("foo"),
			),
			"SELECT foo",
			`> SELECT
>   foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSelect(
				sqlabble.NewColumn("foo").As("a"),
			),
			"SELECT foo AS a",
			`> SELECT
>   foo AS a
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSelect(
				sqlabble.NewColumn("foo").As("a"),
				sqlabble.NewColumn("bar").As("b"),
				sqlabble.NewColumn("baz").As("c"),
			),
			"SELECT foo AS a, bar AS b, baz AS c",
			`> SELECT
>   foo AS a
>   , bar AS b
>   , baz AS c
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSelectDistinct(
				sqlabble.NewColumn("foo").As("a"),
				sqlabble.NewColumn("bar").As("b"),
				sqlabble.NewColumn("baz").As("c"),
			),
			"SELECT DISTINCT foo AS a, bar AS b, baz AS c",
			`> SELECT DISTINCT
>   foo AS a
>   , bar AS b
>   , baz AS c
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

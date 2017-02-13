package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestSetType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Set{},
// 	} {
// 		t.Run(fmt.Sprintf("Type %T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.ClauseNode); !ok {
// 				t.Errorf("%T should implement sqlabble.Clause", c)
// 			}
// 		})
// 	}
// }

func TestSet(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewSet(
				sqlabble.NewColumn("foo").Assign(100),
			),
			"SET foo = ?",
			`> SET
>   foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewSet(
				sqlabble.NewColumn("foo").Assign(100),
				sqlabble.NewColumn("bar").Assign(200),
			),
			"SET foo = ?, bar = ?",
			`> SET
>   foo = ?
>   , bar = ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			sqlabble.NewSet(
				sqlabble.NewColumn("foo").Assign(100),
				sqlabble.NewColumn("bar").Assign(200),
				sqlabble.NewColumn("baz").Assign(300),
			),
			"SET foo = ?, bar = ?, baz = ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			sqlabble.NewSet(
				sqlabble.NewColumn("foo").Assign(100),
				sqlabble.NewColumn("bar").Assign(200),
				sqlabble.NewColumn("baz").Assign(300),
			).Where(
				sqlabble.NewColumn("qux").Lte(400),
			),
			"SET foo = ?, bar = ?, baz = ? WHERE qux <= ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
> WHERE
>   qux <= ?
`,
			[]interface{}{
				100,
				200,
				300,
				400,
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

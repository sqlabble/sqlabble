package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestUpdateType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Update{},
// 	} {
// 		t.Run(fmt.Sprintf("Type %T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.ClauseNode); !ok {
// 				t.Errorf("%T should implement sqlabble.Clause", c)
// 			}
// 		})
// 	}
// }

func TestUpdate(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewUpdate(
				sqlabble.NewTable("foo"),
			),
			"UPDATE foo",
			`> UPDATE
>   foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewUpdate(
				sqlabble.NewTable("foo"),
			).Set(
				sqlabble.NewColumn("bar").Assign(100),
			),
			"UPDATE foo SET bar = ?",
			`> UPDATE
>   foo
> SET
>   bar = ?
`,
			[]interface{}{
				100,
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

package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestUpdateType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		statement.Update{},
// 	} {
// 		t.Run(fmt.Sprintf("Type %T", c), func(t *testing.T) {
// 			if _, ok := c.(statement.ClauseNode); !ok {
// 				t.Errorf("%T should implement statement.Clause", c)
// 			}
// 		})
// 	}
// }

func TestUpdate(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewUpdate(
				statement.NewTable("foo"),
			),
			"UPDATE foo",
			`> UPDATE
>   foo
`,
			[]interface{}{},
		},
		{
			statement.NewUpdate(
				statement.NewTable("foo"),
			).Set(
				statement.NewColumn("bar").Assign(100),
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
			sql, values := builder.Build(c.statement)
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

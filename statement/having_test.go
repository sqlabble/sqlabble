package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestHavingType(t *testing.T) {
// 	if _, ok := interface{}(statement.Having{}).(statement.ClauseNode); !ok {
// 		t.Errorf("statement.Having doesn't implement statement.Clause")
// 	}
// }

func TestHavingSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewHaving(
				statement.NewColumn("foo").Eq(100),
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
					statement.NewColumn("foo").Eq(100),
					statement.NewColumn("bar").Eq("abc"),
				),
			),
			"HAVING foo = ? AND bar = ?",
			`> HAVING
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{100, "abc"},
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

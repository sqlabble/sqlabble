package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestWhereType(t *testing.T) {
// 	if _, ok := interface{}(statement.Where{}).(statement.ClauseNode); !ok {
// 		t.Errorf("statement.Where doesn't implement statement.Clause")
// 	}
// }

func TestWhereSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewWhere(
				statement.NewColumn("foo").Eq(100),
			),
			"WHERE foo = ?",
			`> WHERE
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			statement.NewWhere(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(100),
					statement.NewColumn("bar").Eq("abc"),
				),
			),
			"WHERE foo = ? AND bar = ?",
			`> WHERE
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{100, "abc"},
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

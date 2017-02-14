package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestHavingType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.Having{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.Having doesn't implement sqlabble.Clause")
// 	}
// }

func TestHavingSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewHaving(
				sqlabble.NewColumn("foo").Eq(100),
			),
			"HAVING foo = ?",
			`> HAVING
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			sqlabble.NewHaving(
				sqlabble.NewAnd(
					sqlabble.NewColumn("foo").Eq(100),
					sqlabble.NewColumn("bar").Eq("abc"),
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

package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestWhereType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.Where{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.Where doesn't implement sqlabble.Clause")
// 	}
// }

func TestWhereSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewWhere(
				sqlabble.NewColumn("foo").Eq(100),
			),
			"WHERE foo = ?",
			`> WHERE
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			sqlabble.NewWhere(
				sqlabble.NewAnd(
					sqlabble.NewColumn("foo").Eq(100),
					sqlabble.NewColumn("bar").Eq("abc"),
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

package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestLimitType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.Limit{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.Limit should implement sqlabble.Clause")
// 	}
// }

func TestLimitSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewLimit(0, 3),
			"LIMIT ?",
			`> LIMIT
>   ?
`,
			[]interface{}{
				3,
			},
		},
		{
			sqlabble.NewLimit(10, 3),
			"LIMIT ?, ?",
			`> LIMIT
>   ?, ?
`,
			[]interface{}{
				10,
				3,
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

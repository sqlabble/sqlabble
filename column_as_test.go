package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestColumnAsType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.ColumnAs{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.ColumnNode); !ok {
// 				t.Errorf("%T should implement Column", c)
// 			}
// 		})
// 	}
// }

func TestColumnAsSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewColumnAs("foo"),
			"AS foo",
			`> AS foo
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

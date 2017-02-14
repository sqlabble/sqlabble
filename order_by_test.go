package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestOrderByType(t *testing.T) {
// 	if _, ok := interface{}(sqlabble.OrderBy{}).(sqlabble.ClauseNode); !ok {
// 		t.Errorf("sqlabble.OrderBy doesn't implement sqlabble.Clause")
// 	}
// }

func TestOrderBySQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewOrderBy(
				sqlabble.NewColumn("foo").Asc(),
			),
			"ORDER BY foo ASC",
			`> ORDER BY
>   foo ASC
`,
			[]interface{}{},
		},
		{
			sqlabble.NewOrderBy(
				sqlabble.NewColumn("foo").Desc(),
				sqlabble.NewColumn("bar").Asc(),
				sqlabble.NewColumn("baz").Desc(),
			),
			"ORDER BY foo DESC, bar ASC, baz DESC",
			`> ORDER BY
>   foo DESC
>   , bar ASC
>   , baz DESC
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

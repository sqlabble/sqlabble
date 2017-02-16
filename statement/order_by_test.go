package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestOrderByType(t *testing.T) {
// 	if _, ok := interface{}(statement.OrderBy{}).(statement.ClauseNode); !ok {
// 		t.Errorf("statement.OrderBy doesn't implement statement.Clause")
// 	}
// }

func TestOrderBySQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewOrderBy(
				statement.NewColumn("foo").Asc(),
			),
			"ORDER BY foo ASC",
			`> ORDER BY
>   foo ASC
`,
			[]interface{}{},
		},
		{
			statement.NewOrderBy(
				statement.NewColumn("foo").Desc(),
				statement.NewColumn("bar").Asc(),
				statement.NewColumn("baz").Desc(),
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

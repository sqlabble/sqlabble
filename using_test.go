package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestUsing(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: sqlabble.NewUsing(
				sqlabble.NewColumn("id"),
			),
			sql: "USING id",
			sqlIndent: `> USING id
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewUsing(
				sqlabble.NewColumn("id"),
			).Join(
				sqlabble.NewTable("bar"),
			),
			sql: "USING id JOIN bar",
			sqlIndent: `> USING id
> JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewUsing(
				sqlabble.NewColumn("id"),
			).InnerJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "USING id INNER JOIN bar",
			sqlIndent: `> USING id
> INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewUsing(
				sqlabble.NewColumn("id"),
			).LeftJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "USING id LEFT JOIN bar",
			sqlIndent: `> USING id
> LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewUsing(
				sqlabble.NewColumn("id"),
			).RightJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "USING id RIGHT JOIN bar",
			sqlIndent: `> USING id
> RIGHT JOIN bar
`,
			values: []interface{}{},
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

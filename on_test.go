package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestOn(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: sqlabble.NewOn(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			),
			sql: "ON f.id = b.id",
			sqlIndent: `> ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewOn(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			).Join(
				sqlabble.NewTable("bar"),
			),
			sql: "ON f.id = b.id JOIN bar",
			sqlIndent: `> ON f.id = b.id
> JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewOn(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			).InnerJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "ON f.id = b.id INNER JOIN bar",
			sqlIndent: `> ON f.id = b.id
> INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewOn(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			).LeftJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `> ON f.id = b.id
> LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: sqlabble.NewOn(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			).RightJoin(
				sqlabble.NewTable("bar"),
			),
			sql: "ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `> ON f.id = b.id
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

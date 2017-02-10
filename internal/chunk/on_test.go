package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestOn(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: chunk.NewOn(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			),
			sql: "ON f.id = b.id",
			sqlIndent: `> ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewOn(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			).Join(
				chunk.NewTable("bar"),
			),
			sql: "ON f.id = b.id JOIN bar",
			sqlIndent: `> ON f.id = b.id
> JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewOn(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			).InnerJoin(
				chunk.NewTable("bar"),
			),
			sql: "ON f.id = b.id INNER JOIN bar",
			sqlIndent: `> ON f.id = b.id
> INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewOn(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			).LeftJoin(
				chunk.NewTable("bar"),
			),
			sql: "ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `> ON f.id = b.id
> LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewOn(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			).RightJoin(
				chunk.NewTable("bar"),
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

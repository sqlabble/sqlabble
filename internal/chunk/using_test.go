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

func TestUsing(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: chunk.NewUsing(
				chunk.NewColumn("id"),
			),
			sql: "USING id",
			sqlIndent: `> USING id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewUsing(
				chunk.NewColumn("id"),
			).Join(
				chunk.NewTable("bar"),
			),
			sql: "USING id JOIN bar",
			sqlIndent: `> USING id
> JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewUsing(
				chunk.NewColumn("id"),
			).InnerJoin(
				chunk.NewTable("bar"),
			),
			sql: "USING id INNER JOIN bar",
			sqlIndent: `> USING id
> INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewUsing(
				chunk.NewColumn("id"),
			).LeftJoin(
				chunk.NewTable("bar"),
			),
			sql: "USING id LEFT JOIN bar",
			sqlIndent: `> USING id
> LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewUsing(
				chunk.NewColumn("id"),
			).RightJoin(
				chunk.NewTable("bar"),
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

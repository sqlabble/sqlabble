package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestUsing(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewUsing(
				stmt.NewColumn("id"),
			),
			"USING id",
			`> USING id
`,
			nil,
		},
		{
			stmt.NewUsing(
				stmt.NewColumn("id"),
			).Join(
				stmt.NewTable("bar"),
			),
			"USING id JOIN bar",
			`> USING id
> JOIN bar
`,
			nil,
		},
		{
			stmt.NewUsing(
				stmt.NewColumn("id"),
			).InnerJoin(
				stmt.NewTable("bar"),
			),
			"USING id INNER JOIN bar",
			`> USING id
> INNER JOIN bar
`,
			nil,
		},
		{
			stmt.NewUsing(
				stmt.NewColumn("id"),
			).LeftJoin(
				stmt.NewTable("bar"),
			),
			"USING id LEFT JOIN bar",
			`> USING id
> LEFT JOIN bar
`,
			nil,
		},
		{
			stmt.NewUsing(
				stmt.NewColumn("id"),
			).RightJoin(
				stmt.NewTable("bar"),
			),
			"USING id RIGHT JOIN bar",
			`> USING id
> RIGHT JOIN bar
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("a").
				Join(stmt.NewTable("bar")).Using(stmt.NewColumn("id")),
			`foo AS "a" JOIN bar USING id`,
			`> foo AS "a"
> JOIN bar USING id
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestOn(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewOn(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			),
			"ON f.id = b.id",
			`> ON f.id = b.id
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			).Join(
				stmt.NewTable("bar"),
			),
			"ON f.id = b.id JOIN bar",
			`> ON f.id = b.id
> JOIN bar
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			).InnerJoin(
				stmt.NewTable("bar"),
			),
			"ON f.id = b.id INNER JOIN bar",
			`> ON f.id = b.id
> INNER JOIN bar
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			).LeftJoin(
				stmt.NewTable("bar"),
			),
			"ON f.id = b.id LEFT JOIN bar",
			`> ON f.id = b.id
> LEFT JOIN bar
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			).RightJoin(
				stmt.NewTable("bar"),
			),
			"ON f.id = b.id RIGHT JOIN bar",
			`> ON f.id = b.id
> RIGHT JOIN bar
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("a").
				Join(stmt.NewTable("bar")).On(
				stmt.NewColumn("a.id"),
				stmt.NewColumn("bar.id"),
			),
			`foo AS "a" JOIN bar ON a.id = bar.id`,
			`> foo AS "a"
> JOIN bar ON a.id = bar.id
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

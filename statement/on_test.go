package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestOn(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			"ON f.id = b.id",
			`> ON f.id = b.id
`,
			nil,
		},
		{
			statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).Join(
				statement.NewTable("bar"),
			),
			"ON f.id = b.id JOIN bar",
			`> ON f.id = b.id
> JOIN bar
`,
			nil,
		},
		{
			statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).InnerJoin(
				statement.NewTable("bar"),
			),
			"ON f.id = b.id INNER JOIN bar",
			`> ON f.id = b.id
> INNER JOIN bar
`,
			nil,
		},
		{
			statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).LeftJoin(
				statement.NewTable("bar"),
			),
			"ON f.id = b.id LEFT JOIN bar",
			`> ON f.id = b.id
> LEFT JOIN bar
`,
			nil,
		},
		{
			statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).RightJoin(
				statement.NewTable("bar"),
			),
			"ON f.id = b.id RIGHT JOIN bar",
			`> ON f.id = b.id
> RIGHT JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").As("a").
				Join(statement.NewTable("bar")).On(
				statement.NewColumn("a.id"),
				statement.NewColumn("bar.id"),
			),
			`foo AS "a" JOIN bar ON a.id = bar.id`,
			`> foo AS "a"
> JOIN bar
> ON a.id = bar.id
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

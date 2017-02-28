package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestUsing(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewUsing(
				statement.NewColumn("id"),
			),
			"USING id",
			`> USING id
`,
			nil,
		},
		{
			statement.NewUsing(
				statement.NewColumn("id"),
			).Join(
				statement.NewTable("bar"),
			),
			"USING id JOIN bar",
			`> USING id
> JOIN bar
`,
			nil,
		},
		{
			statement.NewUsing(
				statement.NewColumn("id"),
			).InnerJoin(
				statement.NewTable("bar"),
			),
			"USING id INNER JOIN bar",
			`> USING id
> INNER JOIN bar
`,
			nil,
		},
		{
			statement.NewUsing(
				statement.NewColumn("id"),
			).LeftJoin(
				statement.NewTable("bar"),
			),
			"USING id LEFT JOIN bar",
			`> USING id
> LEFT JOIN bar
`,
			nil,
		},
		{
			statement.NewUsing(
				statement.NewColumn("id"),
			).RightJoin(
				statement.NewTable("bar"),
			),
			"USING id RIGHT JOIN bar",
			`> USING id
> RIGHT JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").As("a").
				Join(statement.NewTable("bar")).Using(statement.NewColumn("id")),
			`foo AS "a" JOIN bar USING id`,
			`> foo AS "a"
> JOIN bar
> USING id
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

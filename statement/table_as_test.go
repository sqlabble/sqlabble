package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestTableAsType(t *testing.T) {
	for _, c := range []interface{}{
		statement.TableAs{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(statement.Joiner); !ok {
				t.Errorf("%T should implement statement.Joiner", c)
			}
		})
	}
}

func TestTableAs(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewTable("foo").As("f"),
			`foo AS "f"`,
			`> foo AS "f"
`,
			nil,
		},
		{
			statement.NewTable("foo").As("f").
				Join(statement.NewTable("bar")),
			`foo AS "f" JOIN bar`,
			`> foo AS "f"
> JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").As("f").
				InnerJoin(statement.NewTable("bar")),
			`foo AS "f" INNER JOIN bar`,
			`> foo AS "f"
> INNER JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").As("f").
				LeftJoin(statement.NewTable("bar")),
			`foo AS "f" LEFT JOIN bar`,
			`> foo AS "f"
> LEFT JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").As("f").
				RightJoin(statement.NewTable("bar")),
			`foo AS "f" RIGHT JOIN bar`,
			`> foo AS "f"
> RIGHT JOIN bar
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

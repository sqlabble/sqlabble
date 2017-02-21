package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestTableType(t *testing.T) {
	for _, c := range []interface{}{
		statement.Table{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(statement.Joiner); !ok {
				t.Errorf("%T should implement statement.Joiner", c)
			}
		})
	}
}

func TestTable(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewTable("foo"),
			`foo`,
			`> foo
`,
			nil,
		},
		{
			statement.NewTable("foo").
				Join(statement.NewTable("bar")),
			`foo JOIN bar`,
			`> foo
> JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").
				InnerJoin(statement.NewTable("bar")),
			`foo INNER JOIN bar`,
			`> foo
> INNER JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").
				LeftJoin(statement.NewTable("bar")),
			`foo LEFT JOIN bar`,
			`> foo
> LEFT JOIN bar
`,
			nil,
		},
		{
			statement.NewTable("foo").
				RightJoin(statement.NewTable("bar")),
			`foo RIGHT JOIN bar`,
			`> foo
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

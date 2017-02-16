package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestTableType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		statement.Table{},
// 		statement.TableAs{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(statement.TableNode); !ok {
// 				t.Errorf("%T doesn't implement statement.Table", c)
// 			}
// 		})
// 	}
// }

func TestTable(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewTable("foo"),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").
				Join(statement.NewTable("bar")),
			"foo JOIN bar",
			`> foo
> JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").
				InnerJoin(statement.NewTable("bar")),
			"foo INNER JOIN bar",
			`> foo
> INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").
				LeftJoin(statement.NewTable("bar")),
			"foo LEFT JOIN bar",
			`> foo
> LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").
				RightJoin(statement.NewTable("bar")),
			"foo RIGHT JOIN bar",
			`> foo
> RIGHT JOIN bar
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestAs(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewTable("foo").As("f"),
			"foo AS f",
			`> foo AS f
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").As("f").
				Join(statement.NewTable("bar")),
			"foo AS f JOIN bar",
			`> foo AS f
> JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").As("f").
				InnerJoin(statement.NewTable("bar")),
			"foo AS f INNER JOIN bar",
			`> foo AS f
> INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").As("f").
				LeftJoin(statement.NewTable("bar")),
			"foo AS f LEFT JOIN bar",
			`> foo AS f
> LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			statement.NewTable("foo").As("f").
				RightJoin(statement.NewTable("bar")),
			"foo AS f RIGHT JOIN bar",
			`> foo AS f
> RIGHT JOIN bar
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

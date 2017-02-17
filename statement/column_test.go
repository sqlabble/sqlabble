package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestColumnType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		statement.Column{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(statement.ColumnNode); !ok {
// 				t.Errorf("%T should implement LogicalOperation", c)
// 			}
// 		})
// 	}
// }

func TestColumnSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo"),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").As("f"),
			"foo AS `f`",
			`> foo AS ` + "`f`" + `
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Define("VARCHAR(20)"),
			"foo VARCHAR(20)",
			`> foo VARCHAR(20)
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Assign(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Eq(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").NotEq(100),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gt(100),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gte(100),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lt(100),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lte(100),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Between(100, 200),
			"foo BETWEEN ? AND ?",
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewColumn("foo").In(100, 200, 300),
			"foo IN (?, ?, ?)",
			`> foo IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewColumn("foo").NotIn(100, 200, 300),
			"foo NOT IN (?, ?, ?)",
			`> foo NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewColumn("foo").Like(`%bar%`),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				`%bar%`,
			},
		},
		{
			statement.NewColumn("foo").RegExp("^(bar|baz)"),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"^(bar|baz)",
			},
		},
		{
			statement.NewColumn("foo").IsNull(),
			"foo IS NULL",
			`> foo IS NULL
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Asc(),
			"foo ASC",
			`> foo ASC
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Desc(),
			"foo DESC",
			`> foo DESC
`,
			[]interface{}{},
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

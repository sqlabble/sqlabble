package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestColumnType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Column{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.ColumnNode); !ok {
// 				t.Errorf("%T should implement LogicalOperation", c)
// 			}
// 		})
// 	}
// }

func TestColumnSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewColumn("foo"),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").As("f"),
			"foo AS f",
			`> foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").Define("VARCHAR(20)"),
			"foo VARCHAR(20)",
			`> foo VARCHAR(20)
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").Assign(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Eq(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").NotEq(100),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Gt(100),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Gte(100),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Lt(100),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Lte(100),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").Between(100, 200),
			"foo BETWEEN ? AND ?",
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			sqlabble.NewColumn("foo").In(100, 200, 300),
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
			sqlabble.NewColumn("foo").NotIn(100, 200, 300),
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
			sqlabble.NewColumn("foo").Like(`%bar%`),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				`%bar%`,
			},
		},
		{
			sqlabble.NewColumn("foo").RegExp("^(bar|baz)"),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"^(bar|baz)",
			},
		},
		{
			sqlabble.NewColumn("foo").IsNull(),
			"foo IS NULL",
			`> foo IS NULL
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").Asc(),
			"foo ASC",
			`> foo ASC
`,
			[]interface{}{},
		},
		{
			sqlabble.NewColumn("foo").Desc(),
			"foo DESC",
			`> foo DESC
`,
			[]interface{}{},
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

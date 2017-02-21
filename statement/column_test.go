package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestColumnType(t *testing.T) {
	for _, c := range []interface{}{
		statement.Column{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(statement.ColumnOrSubquery); !ok {
				t.Errorf("%T should implement statement.ColumnOrSubquery", c)
			}
			if _, ok := c.(statement.ColumnOrColumnAsOrSubquery); !ok {
				t.Errorf("%T should implement statement.ColumnOrColumnAsOrSubquery", c)
			}
		})
	}
}

func TestColumnSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo"),
			`foo`,
			`> foo
`,
			nil,
		},
		{
			statement.NewColumn("foo").As("f"),
			`foo AS "f"`,
			`> foo AS "f"
`,
			nil,
		},
		{
			statement.NewColumn("foo").Define("VARCHAR(20)"),
			`foo VARCHAR(20)`,
			`> foo VARCHAR(20)
`,
			nil,
		},
		{
			statement.NewColumn("foo").Assign(statement.NewParam(100)),
			`foo = ?`,
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Eq(statement.NewParam(100)),
			`foo = ?`,
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").NotEq(statement.NewParam(100)),
			`foo != ?`,
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gt(statement.NewParam(100)),
			`foo > ?`,
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gte(statement.NewParam(100)),
			`foo >= ?`,
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lt(statement.NewParam(100)),
			`foo < ?`,
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lte(statement.NewParam(100)),
			`foo <= ?`,
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Between(statement.NewParam(100), statement.NewParam(200)),
			`foo BETWEEN ? AND ?`,
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewColumn("foo").
				In(statement.NewParams(100, 200, 300)),
			`foo IN (?, ?, ?)`,
			`> foo IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewColumn("foo").
				NotIn(statement.NewParams(100, 200, 300)),
			`foo NOT IN (?, ?, ?)`,
			`> foo NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewColumn("foo").Like(statement.NewParam(`%bar%`)),
			`foo LIKE ?`,
			`> foo LIKE ?
`,
			[]interface{}{
				`%bar%`,
			},
		},
		{
			statement.NewColumn("foo").RegExp(statement.NewParam("^(bar|baz)")),
			`foo REGEXP ?`,
			`> foo REGEXP ?
`,
			[]interface{}{
				`^(bar|baz)`,
			},
		},
		{
			statement.NewColumn("foo").IsNull(),
			`foo IS NULL`,
			`> foo IS NULL
`,
			nil,
		},
		{
			statement.NewColumn("foo").IsNotNull(),
			`foo IS NOT NULL`,
			`> foo IS NOT NULL
`,
			nil,
		},
		{
			statement.NewColumn("foo").Asc(),
			`foo ASC`,
			`> foo ASC
`,
			nil,
		},
		{
			statement.NewColumn("foo").Desc(),
			`foo DESC`,
			`> foo DESC
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

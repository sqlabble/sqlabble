package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestColumnType(t *testing.T) {
	for _, c := range []interface{}{
		chunk.Column{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Column); !ok {
				t.Errorf("%T should implement LogicalOperation", c)
			}
		})
	}
}

func TestColumnSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewColumn("foo"),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").As("f"),
			"foo AS f",
			`> foo AS f
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").Define("VARCHAR(20)"),
			"foo VARCHAR(20)",
			`> foo VARCHAR(20)
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").Assign(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Eq(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").NotEq(100),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Gt(100),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Gte(100),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Lt(100),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Lte(100),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewColumn("foo").Between(100, 200),
			"foo BETWEEN ? AND ?",
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			chunk.NewColumn("foo").In(100, 200, 300),
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
			chunk.NewColumn("foo").NotIn(100, 200, 300),
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
			chunk.NewColumn("foo").Like(`%bar%`),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				`%bar%`,
			},
		},
		{
			chunk.NewColumn("foo").RegExp("^(bar|baz)"),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"^(bar|baz)",
			},
		},
		{
			chunk.NewColumn("foo").IsNull(),
			"foo IS NULL",
			`> foo IS NULL
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").Asc(),
			"foo ASC",
			`> foo ASC
`,
			[]interface{}{},
		},
		{
			chunk.NewColumn("foo").Desc(),
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

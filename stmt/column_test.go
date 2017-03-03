package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestColumnType(t *testing.T) {
	t.Parallel()
	for _, c := range []interface{}{
		stmt.Column{},
	} {
		c := c
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			t.Parallel()
			if _, ok := c.(stmt.ColOrSub); !ok {
				t.Errorf("%T should implement stmt.ColumnOrSubquery", c)
			}
			if _, ok := c.(stmt.ColOrAliasOrSub); !ok {
				t.Errorf("%T should implement stmt.ColumnOrColumnAsOrSubquery", c)
			}
		})
	}
}

func TestColumn(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo"),
			`foo`,
			`> foo
`,
			nil,
		},
		{
			stmt.NewColumn("foo").As("f"),
			`foo AS "f"`,
			`> foo AS "f"
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Define("VARCHAR(20)"),
			`foo VARCHAR(20)`,
			`> foo VARCHAR(20)
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Assign(stmt.NewParam(100)),
			`foo = ?`,
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
			`foo = ?`,
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").NotEq(stmt.NewParam(100)),
			`foo != ?`,
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Gt(stmt.NewParam(100)),
			`foo > ?`,
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Gte(stmt.NewParam(100)),
			`foo >= ?`,
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Lt(stmt.NewParam(100)),
			`foo < ?`,
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Lte(stmt.NewParam(100)),
			`foo <= ?`,
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Between(stmt.NewParam(100), stmt.NewParam(200)),
			`foo BETWEEN ? AND ?`,
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			stmt.NewColumn("foo").
				In(stmt.NewParams(100, 200, 300)),
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
			stmt.NewColumn("foo").
				NotIn(stmt.NewParams(100, 200, 300)),
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
			stmt.NewColumn("foo").Like(stmt.NewParam(`%bar%`)),
			`foo LIKE ?`,
			`> foo LIKE ?
`,
			[]interface{}{
				`%bar%`,
			},
		},
		{
			stmt.NewColumn("foo").RegExp(stmt.NewParam("^(bar|baz)")),
			`foo REGEXP ?`,
			`> foo REGEXP ?
`,
			[]interface{}{
				`^(bar|baz)`,
			},
		},
		{
			stmt.NewColumn("foo").IsNull(),
			`foo IS NULL`,
			`> foo IS NULL
`,
			nil,
		},
		{
			stmt.NewColumn("foo").IsNotNull(),
			`foo IS NOT NULL`,
			`> foo IS NOT NULL
`,
			nil,
		},
		{
			stmt.NewColumn("foo").EqAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo = ALL (SELECT)`,
			`> foo = ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").NotEqAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo != ALL (SELECT)`,
			`> foo != ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").GtAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo > ALL (SELECT)`,
			`> foo > ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").GteAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo >= ALL (SELECT)`,
			`> foo >= ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").LtAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo < ALL (SELECT)`,
			`> foo < ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").LteAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo <= ALL (SELECT)`,
			`> foo <= ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").EqAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo = ANY (SELECT)`,
			`> foo = ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").NotEqAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo != ANY (SELECT)`,
			`> foo != ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").GtAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo > ANY (SELECT)`,
			`> foo > ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").GteAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo >= ANY (SELECT)`,
			`> foo >= ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").LtAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo < ANY (SELECT)`,
			`> foo < ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").LteAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`foo <= ANY (SELECT)`,
			`> foo <= ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Asc(),
			`foo ASC`,
			`> foo ASC
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Desc(),
			`foo DESC`,
			`> foo DESC
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

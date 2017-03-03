package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestFuncOperationLeftSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewCurDate(),
			"CURDATE()",
			`> CURDATE()
`,
			nil,
		},
		{
			stmt.NewCurDate().
				Eq(stmt.NewParam(100)),
			"CURDATE() = ?",
			`> CURDATE() = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				NotEq(stmt.NewParam(100)),
			"CURDATE() != ?",
			`> CURDATE() != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				Gt(stmt.NewParam(100)),
			"CURDATE() > ?",
			`> CURDATE() > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				Gte(stmt.NewParam(100)),
			"CURDATE() >= ?",
			`> CURDATE() >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				Lt(stmt.NewParam(100)),
			"CURDATE() < ?",
			`> CURDATE() < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				Lte(stmt.NewParam(100)),
			"CURDATE() <= ?",
			`> CURDATE() <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewCurDate().
				Like(stmt.NewParam("bar")),
			"CURDATE() LIKE ?",
			`> CURDATE() LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			stmt.NewCurDate().
				RegExp(stmt.NewParam("bar")),
			"CURDATE() REGEXP ?",
			`> CURDATE() REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			stmt.NewCurDate().
				Between(
					stmt.NewParam(100),
					stmt.NewParam(200),
				),
			"CURDATE() BETWEEN ? AND ?",
			`> CURDATE() BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			stmt.NewCurDate().
				In(stmt.NewParams(
					100, 200, 300,
				)),
			"CURDATE() IN (?, ?, ?)",
			`> CURDATE() IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			stmt.NewCurDate().
				NotIn(stmt.NewParams(
					100, 200, 300,
				)),
			"CURDATE() NOT IN (?, ?, ?)",
			`> CURDATE() NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},

		{
			stmt.NewCurDate().EqAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() = ALL (SELECT)`,
			`> CURDATE() = ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().NotEqAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() != ALL (SELECT)`,
			`> CURDATE() != ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().GtAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() > ALL (SELECT)`,
			`> CURDATE() > ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().GteAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() >= ALL (SELECT)`,
			`> CURDATE() >= ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().LtAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() < ALL (SELECT)`,
			`> CURDATE() < ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().LteAll(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() <= ALL (SELECT)`,
			`> CURDATE() <= ALL (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().EqAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() = ANY (SELECT)`,
			`> CURDATE() = ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().NotEqAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() != ANY (SELECT)`,
			`> CURDATE() != ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().GtAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() > ANY (SELECT)`,
			`> CURDATE() > ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().GteAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() >= ANY (SELECT)`,
			`> CURDATE() >= ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().LtAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() < ANY (SELECT)`,
			`> CURDATE() < ANY (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewCurDate().LteAny(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
			),
			`CURDATE() <= ANY (SELECT)`,
			`> CURDATE() <= ANY (
>   SELECT
> )
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

func TestFuncOperationRightSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo").Eq(
				stmt.NewCurDate(),
			),
			"foo = CURDATE()",
			`> foo = CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").NotEq(
				stmt.NewCurDate(),
			),
			"foo != CURDATE()",
			`> foo != CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Gt(
				stmt.NewCurDate(),
			),
			"foo > CURDATE()",
			`> foo > CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Gte(
				stmt.NewCurDate(),
			),
			"foo >= CURDATE()",
			`> foo >= CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Lt(
				stmt.NewCurDate(),
			),
			"foo < CURDATE()",
			`> foo < CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Lte(
				stmt.NewCurDate(),
			),
			"foo <= CURDATE()",
			`> foo <= CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Like(
				stmt.NewCurDate(),
			),
			"foo LIKE CURDATE()",
			`> foo LIKE CURDATE()
`,
			nil,
		},
		{
			stmt.NewColumn("foo").RegExp(
				stmt.NewCurDate(),
			),
			"foo REGEXP CURDATE()",
			`> foo REGEXP CURDATE()
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

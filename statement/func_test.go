package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestFuncOperationLeftSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewCurDate(),
			"CURDATE()",
			`> CURDATE()
`,
			nil,
		},
		{
			statement.NewCurDate().
				Eq(statement.NewParam(100)),
			"CURDATE() = ?",
			`> CURDATE() = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				NotEq(statement.NewParam(100)),
			"CURDATE() != ?",
			`> CURDATE() != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				Gt(statement.NewParam(100)),
			"CURDATE() > ?",
			`> CURDATE() > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				Gte(statement.NewParam(100)),
			"CURDATE() >= ?",
			`> CURDATE() >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				Lt(statement.NewParam(100)),
			"CURDATE() < ?",
			`> CURDATE() < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				Lte(statement.NewParam(100)),
			"CURDATE() <= ?",
			`> CURDATE() <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewCurDate().
				Like(statement.NewParam("bar")),
			"CURDATE() LIKE ?",
			`> CURDATE() LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewCurDate().
				RegExp(statement.NewParam("bar")),
			"CURDATE() REGEXP ?",
			`> CURDATE() REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewCurDate().
				Between(
					statement.NewParam(100),
					statement.NewParam(200),
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
			statement.NewCurDate().
				In(statement.NewParams(
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
			statement.NewCurDate().
				NotIn(statement.NewParams(
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
			statement.NewCurDate().EqAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().NotEqAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().GtAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().GteAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().LtAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().LteAll(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().EqAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().NotEqAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().GtAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().GteAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().LtAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
			statement.NewCurDate().LteAny(
				statement.NewSubquery(
					statement.NewSelect(),
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
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
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

func TestFuncOperationRightSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Eq(
				statement.NewCurDate(),
			),
			"foo = CURDATE()",
			`> foo = CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").NotEq(
				statement.NewCurDate(),
			),
			"foo != CURDATE()",
			`> foo != CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").Gt(
				statement.NewCurDate(),
			),
			"foo > CURDATE()",
			`> foo > CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").Gte(
				statement.NewCurDate(),
			),
			"foo >= CURDATE()",
			`> foo >= CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").Lt(
				statement.NewCurDate(),
			),
			"foo < CURDATE()",
			`> foo < CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").Lte(
				statement.NewCurDate(),
			),
			"foo <= CURDATE()",
			`> foo <= CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").Like(
				statement.NewCurDate(),
			),
			"foo LIKE CURDATE()",
			`> foo LIKE CURDATE()
`,
			nil,
		},
		{
			statement.NewColumn("foo").RegExp(
				statement.NewCurDate(),
			),
			"foo REGEXP CURDATE()",
			`> foo REGEXP CURDATE()
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
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

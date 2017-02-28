package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestSubOperationLeftSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSubquery(statement.NewSelect()),
			"(SELECT)",
			`> (
>   SELECT
> )
`,
			nil,
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Eq(statement.NewParam(100)),
			"(SELECT) = ?",
			`> (
>   SELECT
> ) = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				NotEq(statement.NewParam(100)),
			"(SELECT) != ?",
			`> (
>   SELECT
> ) != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Gt(statement.NewParam(100)),
			"(SELECT) > ?",
			`> (
>   SELECT
> ) > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Gte(statement.NewParam(100)),
			"(SELECT) >= ?",
			`> (
>   SELECT
> ) >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Lt(statement.NewParam(100)),
			"(SELECT) < ?",
			`> (
>   SELECT
> ) < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Lte(statement.NewParam(100)),
			"(SELECT) <= ?",
			`> (
>   SELECT
> ) <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Like(statement.NewParam("bar")),
			"(SELECT) LIKE ?",
			`> (
>   SELECT
> ) LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				RegExp(statement.NewParam("bar")),
			"(SELECT) REGEXP ?",
			`> (
>   SELECT
> ) REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				Between(
					statement.NewParam(100),
					statement.NewParam(200),
				),
			"(SELECT) BETWEEN ? AND ?",
			`> (
>   SELECT
> ) BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				IsNull(),
			"(SELECT) IS NULL",
			`> (
>   SELECT
> ) IS NULL
`,
			nil,
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				IsNotNull(),
			"(SELECT) IS NOT NULL",
			`> (
>   SELECT
> ) IS NOT NULL
`,
			nil,
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				In(statement.NewParams(
					100, 200, 300,
				)),
			"(SELECT) IN (?, ?, ?)",
			`> (
>   SELECT
> ) IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewSubquery(statement.NewSelect()).
				NotIn(statement.NewParams(
					100, 200, 300,
				)),
			"(SELECT) NOT IN (?, ?, ?)",
			`> (
>   SELECT
> ) NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
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

func TestSubOperationRightSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement.NewColumn("foo").Eq(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo = (SELECT bar)",
			`> foo = (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 1
		{
			statement.NewColumn("foo").NotEq(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo != (SELECT bar)",
			`> foo != (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 2
		{
			statement.NewColumn("foo").Gt(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo > (SELECT bar)",
			`> foo > (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 3
		{
			statement.NewColumn("foo").Gte(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo >= (SELECT bar)",
			`> foo >= (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 4
		{
			statement.NewColumn("foo").Lt(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo < (SELECT bar)",
			`> foo < (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 5
		{
			statement.NewColumn("foo").Lte(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo <= (SELECT bar)",
			`> foo <= (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 6
		{
			statement.NewColumn("foo").Like(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo LIKE (SELECT bar)",
			`> foo LIKE (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 7
		{
			statement.NewColumn("foo").RegExp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"foo REGEXP (SELECT bar)",
			`> foo REGEXP (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 8
		{
			statement.NewColumn("foo").
				In(statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				)),
			"foo IN (SELECT bar)",
			`> foo IN (
>   SELECT
>     bar
> )
`,
			nil,
		},
		// 9
		{
			statement.NewColumn("foo").
				NotIn(statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				)),
			"foo NOT IN (SELECT bar)",
			`> foo NOT IN (
>   SELECT
>     bar
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

func TestSubOperationBothSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSubquery(
				statement.NewSelect(statement.NewColumn("foo")),
			).Eq(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"(SELECT foo) = (SELECT bar)",
			`> (
>   SELECT
>     foo
> ) = (
>   SELECT
>     bar
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

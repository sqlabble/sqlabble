package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSubOperationLeftSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSubquery(stmt.NewSelect()),
			"(SELECT)",
			`> (
>   SELECT
> )
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				Eq(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				NotEq(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Gt(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Gte(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Lt(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Lte(stmt.NewVal(100)),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Like(stmt.NewVal("bar")),
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
			stmt.NewSubquery(stmt.NewSelect()).
				RegExp(stmt.NewVal("bar")),
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
			stmt.NewSubquery(stmt.NewSelect()).
				Between(
					stmt.NewVal(100),
					stmt.NewVal(200),
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
			stmt.NewSubquery(stmt.NewSelect()).
				IsNull(),
			"(SELECT) IS NULL",
			`> (
>   SELECT
> ) IS NULL
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				IsNotNull(),
			"(SELECT) IS NOT NULL",
			`> (
>   SELECT
> ) IS NOT NULL
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				In(stmt.NewVals(
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
			stmt.NewSubquery(stmt.NewSelect()).
				NotIn(stmt.NewVals(
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

func TestSubOperationRightSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			stmt.NewColumn("foo").Eq(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").NotEq(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").Gt(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").Gte(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").Lt(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").Lte(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").Like(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").RegExp(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").
				In(stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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
			stmt.NewColumn("foo").
				NotIn(stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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

func TestSubOperationBothSide(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSubquery(
				stmt.NewSelect(stmt.NewColumn("foo")),
			).Eq(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("bar")),
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

package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestSubOperationLeftSide(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSub(statement.NewSelect()),
			"(SELECT)",
			`> (
>   SELECT
> )
`,
			[]interface{}{},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Eq(100),
			"(SELECT) = ?",
			`> (
>   SELECT
> )
> = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				NotEq(100),
			"(SELECT) != ?",
			`> (
>   SELECT
> )
> != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Gt(100),
			"(SELECT) > ?",
			`> (
>   SELECT
> )
> > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Gte(100),
			"(SELECT) >= ?",
			`> (
>   SELECT
> )
> >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Lt(100),
			"(SELECT) < ?",
			`> (
>   SELECT
> )
> < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Lte(100),
			"(SELECT) <= ?",
			`> (
>   SELECT
> )
> <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Like("bar"),
			"(SELECT) LIKE ?",
			`> (
>   SELECT
> )
> LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				RegExp("bar"),
			"(SELECT) REGEXP ?",
			`> (
>   SELECT
> )
> REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				Between(100, 200),
			"(SELECT) BETWEEN ? AND ?",
			`> (
>   SELECT
> )
> BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				IsNull(),
			"(SELECT) IS NULL",
			`> (
>   SELECT
> )
> IS NULL
`,
			[]interface{}{},
		},
		{
			statement.NewSub(statement.NewSelect()).
				IsNotNull(),
			"(SELECT) IS NOT NULL",
			`> (
>   SELECT
> )
> IS NOT NULL
`,
			[]interface{}{},
		},
		{
			statement.NewSub(statement.NewSelect()).
				In(100, 200, 300),
			"(SELECT) IN (?, ?, ?)",
			`> (
>   SELECT
> )
> IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewSub(statement.NewSelect()).
				NotIn(100, 200, 300),
			"(SELECT) NOT IN (?, ?, ?)",
			`> (
>   SELECT
> )
> NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
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

func TestSubOperationRightSide(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Eq(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo = (SELECT bar)",
			`> foo =
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").NotEq(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo != (SELECT bar)",
			`> foo !=
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Gt(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo > (SELECT bar)",
			`> foo >
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Gte(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo >= (SELECT bar)",
			`> foo >=
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Lt(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo < (SELECT bar)",
			`> foo <
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Lte(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo <= (SELECT bar)",
			`> foo <=
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").Like(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo LIKE (SELECT bar)",
			`> foo LIKE
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").RegExp(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo REGEXP (SELECT bar)",
			`> foo REGEXP
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").In(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo IN (SELECT bar)",
			`> foo IN
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").NotIn(
				statement.NewSelect(statement.NewColumn("bar")),
			),
			"foo NOT IN (SELECT bar)",
			`> foo NOT IN
> (
>   SELECT
>     bar
> )
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

func TestSubOperationBothSide(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSub(
				statement.NewSelect(statement.NewColumn("foo")),
			).Eq(
				statement.NewSub(
					statement.NewSelect(statement.NewColumn("bar")),
				),
			),
			"(SELECT foo) = (SELECT bar)",
			`> (
>   SELECT
>     foo
> )
> =
> (
>   SELECT
>     bar
> )
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

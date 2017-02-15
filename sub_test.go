package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestSubOperationLeftSide(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewSub(sqlabble.NewSelect()),
			"(SELECT)",
			`> (
>   SELECT
> )
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sqlabble.NewSub(sqlabble.NewSelect()).
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewColumn("foo").Eq(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").NotEq(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").Gt(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").Gte(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").Lt(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").Lte(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").Like(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").RegExp(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").In(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sqlabble.NewColumn("foo").NotIn(
				sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewSub(
				sqlabble.NewSelect(sqlabble.NewColumn("foo")),
			).Eq(
				sqlabble.NewSub(
					sqlabble.NewSelect(sqlabble.NewColumn("bar")),
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
			sql, values := sqlabble.Build(c.statement)
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

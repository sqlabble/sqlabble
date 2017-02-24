package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestWhereOperation(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewWhere(
				statement.NewColumn("foo").
					Eq(statement.NewParam(100)),
			),
			"WHERE foo = ?",
			`> WHERE
>   foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewWhere(
				statement.NewColumn("foo").
					EqAll(
						statement.NewSubquery(
							statement.NewSelect(),
						),
					),
			),
			"WHERE foo = ALL (SELECT)",
			`> WHERE
>   foo = ALL (
>     SELECT
>   )
`,
			nil,
		},
		{
			statement.NewWhere(
				statement.NewExists(
					statement.NewSubquery(
						statement.NewSelect(),
					),
				),
			),
			"WHERE EXISTS (SELECT)",
			`> WHERE
>   EXISTS (
>     SELECT
>   )
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

func TestWhereSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewWhere(
				statement.NewColumn("foo").Eq(statement.NewParam(100)),
			),
			"WHERE foo = ?",
			`> WHERE
>   foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewWhere(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam("abc")),
				),
			),
			"WHERE foo = ? AND bar = ?",
			`> WHERE
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{
				100,
				"abc",
			},
		},
		{
			statement.NewWhere(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam("abc")),
				),
			).GroupBy(
				statement.NewColumn("baz"),
			),
			"WHERE foo = ? AND bar = ? GROUP BY baz",
			`> WHERE
>   foo = ?
>   AND bar = ?
> GROUP BY
>   baz
`,
			[]interface{}{
				100,
				"abc",
			},
		},
		{
			statement.NewWhere(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam("abc")),
				),
			).OrderBy(
				statement.NewColumn("baz").Asc(),
			),
			"WHERE foo = ? AND bar = ? ORDER BY baz ASC",
			`> WHERE
>   foo = ?
>   AND bar = ?
> ORDER BY
>   baz ASC
`,
			[]interface{}{
				100,
				"abc",
			},
		},
		{
			statement.NewWhere(
				statement.NewAnd(
					statement.NewColumn("foo").Eq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam("abc")),
				),
			).Limit(20),
			"WHERE foo = ? AND bar = ? LIMIT ?",
			`> WHERE
>   foo = ?
>   AND bar = ?
> LIMIT
>   ?
`,
			[]interface{}{
				100,
				"abc",
				20,
			},
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

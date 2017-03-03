package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestWhereOperation(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewWhere(
				stmt.NewColumn("foo").
					Eq(stmt.NewParam(100)),
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
			stmt.NewWhere(
				stmt.NewColumn("foo").
					EqAll(
						stmt.NewSubquery(
							stmt.NewSelect(),
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
			stmt.NewWhere(
				stmt.NewExists(
					stmt.NewSubquery(
						stmt.NewSelect(),
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

func TestWhereSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewWhere(
				stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
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
			stmt.NewWhere(
				stmt.NewAnd(
					stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
					stmt.NewColumn("bar").Eq(stmt.NewParam("abc")),
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
			stmt.NewWhere(
				stmt.NewAnd(
					stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
					stmt.NewColumn("bar").Eq(stmt.NewParam("abc")),
				),
			).GroupBy(
				stmt.NewColumn("baz"),
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
			stmt.NewWhere(
				stmt.NewAnd(
					stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
					stmt.NewColumn("bar").Eq(stmt.NewParam("abc")),
				),
			).OrderBy(
				stmt.NewColumn("baz").Asc(),
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
			stmt.NewWhere(
				stmt.NewAnd(
					stmt.NewColumn("foo").Eq(stmt.NewParam(100)),
					stmt.NewColumn("bar").Eq(stmt.NewParam("abc")),
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

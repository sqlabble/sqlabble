package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSelectSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			stmt.NewSelect(),
			`SELECT`,
			`> SELECT
`,
			nil,
		},
		// 1
		{
			stmt.NewSelect(
				stmt.NewWildcard(),
			),
			`SELECT *`,
			`> SELECT
>   *
`,
			nil,
		},
		// 2
		{
			stmt.NewSelect(
				stmt.NewColumn("foo"),
			),
			`SELECT "foo"`,
			`> SELECT
>   "foo"
`,
			nil,
		},
		// 3
		{
			stmt.NewSelect(
				stmt.NewColumn("foo").As("a"),
			),
			`SELECT "foo" AS "a"`,
			`> SELECT
>   "foo" AS "a"
`,
			nil,
		},
		// 4
		{
			stmt.NewSelect(
				stmt.NewColumn("foo").As("a"),
				stmt.NewColumn("bar").As("b"),
				stmt.NewColumn("baz").As("c"),
			),
			`SELECT "foo" AS "a", "bar" AS "b", "baz" AS "c"`,
			`> SELECT
>   "foo" AS "a"
>   , "bar" AS "b"
>   , "baz" AS "c"
`,
			nil,
		},
		// 5
		{
			stmt.NewSelectDistinct(
				stmt.NewColumn("foo").As("a"),
				stmt.NewColumn("bar").As("b"),
				stmt.NewColumn("baz").As("c"),
			),
			`SELECT DISTINCT "foo" AS "a", "bar" AS "b", "baz" AS "c"`,
			`> SELECT DISTINCT
>   "foo" AS "a"
>   , "bar" AS "b"
>   , "baz" AS "c"
`,
			nil,
		},
		// 6
		{
			stmt.NewSelect(
				stmt.NewColumn("foo"),
			).From(
				stmt.NewTable("bar"),
			),
			`SELECT "foo" FROM "bar"`,
			`> SELECT
>   "foo"
> FROM
>   "bar"
`,
			nil,
		},
		// 7
		{
			stmt.NewSelect(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
					),
				),
			).From(
				stmt.NewTable("bar"),
			),
			`SELECT (SELECT "foo") FROM "bar"`,
			`> SELECT
>   (
>     SELECT
>       "foo"
>   )
> FROM
>   "bar"
`,
			nil,
		},
		{
			stmt.NewSelect(
				stmt.NewConcat(
					stmt.NewColumn("fname"),
					stmt.NewVal(" "),
					stmt.NewColumn("lname"),
				),
			),
			`SELECT CONCAT("fname", ?, "lname")`,
			`> SELECT
>   CONCAT("fname", ?, "lname")
`,
			[]interface{}{
				" ",
			},
		},
		{
			stmt.NewSelect(
				stmt.NewVal(1),
			),
			`SELECT ?`,
			`> SELECT
>   ?
`,
			[]interface{}{
				1,
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

package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSet(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(stmt.NewParam(100)),
			),
			"SET foo = ?",
			`> SET
>   foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(stmt.NewCurDate()),
			),
			"SET foo = CURDATE()",
			`> SET
>   foo = CURDATE()
`,
			nil,
		},
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(
					stmt.NewSubquery(
						stmt.NewSelect(
							stmt.NewColumn("count(*)"),
						),
					),
				),
			),
			"SET foo = (SELECT count(*))",
			`> SET
>   foo = (
>     SELECT
>       count(*)
>   )
`,
			nil,
		},
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(stmt.NewParam(100)),
				stmt.NewColumn("bar").Assign(stmt.NewParam(200)),
			),
			"SET foo = ?, bar = ?",
			`> SET
>   foo = ?
>   , bar = ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(stmt.NewParam(100)),
				stmt.NewColumn("bar").Assign(stmt.NewParam(200)),
				stmt.NewColumn("baz").Assign(stmt.NewParam(300)),
			),
			"SET foo = ?, bar = ?, baz = ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			stmt.NewSet(
				stmt.NewColumn("foo").Assign(stmt.NewParam(100)),
				stmt.NewColumn("bar").Assign(stmt.NewParam(200)),
				stmt.NewColumn("baz").Assign(stmt.NewParam(300)),
			).Where(
				stmt.NewColumn("qux").Lte(stmt.NewParam(400)),
			),
			"SET foo = ?, bar = ?, baz = ? WHERE qux <= ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
> WHERE
>   qux <= ?
`,
			[]interface{}{
				100,
				200,
				300,
				400,
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

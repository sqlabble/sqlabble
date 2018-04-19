package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestDeleteSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewDelete(),
			`DELETE`,
			`> DELETE
`,
			nil,
		},
		{
			stmt.NewDelete().From(
				stmt.NewTable("foo"),
			),
			`DELETE FROM "foo"`,
			`> DELETE
> FROM
>   "foo"
`,
			nil,
		},
		{
			stmt.NewDelete(
				stmt.NewTable("hoge"),
			).From(
				stmt.NewTable("foo"),
			),
			`DELETE "hoge" FROM "foo"`,
			`> DELETE "hoge"
> FROM
>   "foo"
`,
			nil,
		},
		{
			stmt.NewDelete(
				stmt.NewTable("foo"),
			).From(
				stmt.NewTable("foo").
					InnerJoin(stmt.NewTable("hoge")).
					On(stmt.NewColumn("foo_id"), stmt.NewColumn("hoge_id")),
			),
			`DELETE "foo" FROM "foo" INNER JOIN "hoge" ON "foo_id" = "hoge_id"`,
			`> DELETE "foo"
> FROM
>   "foo"
>   INNER JOIN "hoge" ON "foo_id" = "hoge_id"
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

package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestOn(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewOn(
				stmt.NewTable("f").Column("id"),
				stmt.NewTable("b").Column("id"),
			),
			`ON "f"."id" = "b"."id"`,
			`> ON "f"."id" = "b"."id"
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewTable("f").Column("id"),
				stmt.NewTable("b").Column("id"),
			).Join(
				stmt.NewTable("bar"),
			),
			`ON "f"."id" = "b"."id" JOIN "bar"`,
			`> ON "f"."id" = "b"."id"
> JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewTable("f").Column("id"),
				stmt.NewTable("b").Column("id"),
			).InnerJoin(
				stmt.NewTable("bar"),
			),
			`ON "f"."id" = "b"."id" INNER JOIN "bar"`,
			`> ON "f"."id" = "b"."id"
> INNER JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewTable("f").Column("id"),
				stmt.NewTable("b").Column("id"),
			).LeftJoin(
				stmt.NewTable("bar"),
			),
			`ON "f"."id" = "b"."id" LEFT JOIN "bar"`,
			`> ON "f"."id" = "b"."id"
> LEFT JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewOn(
				stmt.NewTable("f").Column("id"),
				stmt.NewTable("b").Column("id"),
			).RightJoin(
				stmt.NewTable("bar"),
			),
			`ON "f"."id" = "b"."id" RIGHT JOIN "bar"`,
			`> ON "f"."id" = "b"."id"
> RIGHT JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("a").
				Join(stmt.NewTable("bar")).
				On(
					stmt.NewTable("a").Column("id"),
					stmt.NewTable("bar").Column("id"),
				),
			`"foo" AS "a" JOIN "bar" ON "a"."id" = "bar"."id"`,
			`> "foo" AS "a"
> JOIN "bar" ON "a"."id" = "bar"."id"
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

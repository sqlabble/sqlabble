package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestTableType(t *testing.T) {
	t.Parallel()
	for _, c := range []interface{}{
		stmt.Table{},
	} {
		c := c
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			t.Parallel()
			if _, ok := c.(stmt.Joiner); !ok {
				t.Errorf("%T should implement stmt.Joiner", c)
			}
		})
	}
}

func TestTable(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewTable(""),
			``,
			``,
			nil,
		},
		{
			stmt.NewTable("foo"),
			`"foo"`,
			`> "foo"
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("f"),
			`"foo" AS "f"`,
			`> "foo" AS "f"
`,
			nil,
		},
		{
			stmt.NewTable("foo").Column("foo_id"),
			`"foo"."foo_id"`,
			`> "foo"."foo_id"
`,
			nil,
		},
		{
			stmt.NewTable("foo").
				Join(stmt.NewTable("bar")),
			`"foo" JOIN "bar"`,
			`> "foo"
> JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewTable("foo").
				InnerJoin(stmt.NewTable("bar")),
			`"foo" INNER JOIN "bar"`,
			`> "foo"
> INNER JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewTable("foo").
				LeftJoin(stmt.NewTable("bar")),
			`"foo" LEFT JOIN "bar"`,
			`> "foo"
> LEFT JOIN "bar"
`,
			nil,
		},
		{
			stmt.NewTable("foo").
				RightJoin(stmt.NewTable("bar")),
			`"foo" RIGHT JOIN "bar"`,
			`> "foo"
> RIGHT JOIN "bar"
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

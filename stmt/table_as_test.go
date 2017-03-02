package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestTableAsType(t *testing.T) {
	t.Parallel()
	for _, c := range []interface{}{
		stmt.TableAs{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			t.Parallel()
			if _, ok := c.(stmt.Joiner); !ok {
				t.Errorf("%T should implement stmt.Joiner", c)
			}
		})
	}
}

func TestTableAs(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewTable("foo").As("f"),
			`foo AS "f"`,
			`> foo AS "f"
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("f").
				Join(stmt.NewTable("bar")),
			`foo AS "f" JOIN bar`,
			`> foo AS "f"
> JOIN bar
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("f").
				InnerJoin(stmt.NewTable("bar")),
			`foo AS "f" INNER JOIN bar`,
			`> foo AS "f"
> INNER JOIN bar
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("f").
				LeftJoin(stmt.NewTable("bar")),
			`foo AS "f" LEFT JOIN bar`,
			`> foo AS "f"
> LEFT JOIN bar
`,
			nil,
		},
		{
			stmt.NewTable("foo").As("f").
				RightJoin(stmt.NewTable("bar")),
			`foo AS "f" RIGHT JOIN bar`,
			`> foo AS "f"
> RIGHT JOIN bar
`,
			nil,
		},
	} {
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

package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSubqueryAlias(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSubquery(stmt.NewSelect()).
				As("foo").
				Join(stmt.NewTable("bar")),
			`(SELECT) AS foo JOIN bar`,
			`> (
>   SELECT
> ) AS foo
> JOIN bar
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				As("foo").
				InnerJoin(stmt.NewTable("bar")),
			`(SELECT) AS foo INNER JOIN bar`,
			`> (
>   SELECT
> ) AS foo
> INNER JOIN bar
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				As("foo").
				LeftJoin(stmt.NewTable("bar")),
			`(SELECT) AS foo LEFT JOIN bar`,
			`> (
>   SELECT
> ) AS foo
> LEFT JOIN bar
`,
			nil,
		},
		{
			stmt.NewSubquery(stmt.NewSelect()).
				As("foo").
				RightJoin(stmt.NewTable("bar")),
			`(SELECT) AS foo RIGHT JOIN bar`,
			`> (
>   SELECT
> ) AS foo
> RIGHT JOIN bar
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

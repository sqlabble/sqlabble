package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestOrderBySQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewOrderBy(
				stmt.NewColumn("foo").Asc(),
			),
			`ORDER BY "foo" ASC`,
			`> ORDER BY
>   "foo" ASC
`,
			nil,
		},
		{
			stmt.NewOrderBy(
				stmt.NewColumn("foo").Desc(),
				stmt.NewColumn("bar").Asc(),
				stmt.NewColumn("baz").Desc(),
			),
			`ORDER BY "foo" DESC, "bar" ASC, "baz" DESC`,
			`> ORDER BY
>   "foo" DESC
>   , "bar" ASC
>   , "baz" DESC
`,
			nil,
		},
		{
			stmt.NewOrderBy(
				stmt.NewColumn("foo").Asc(),
			).Limit(10),
			`ORDER BY "foo" ASC LIMIT ?`,
			`> ORDER BY
>   "foo" ASC
> LIMIT
>   ?
`,
			[]interface{}{
				10,
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

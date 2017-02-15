package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestSub(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewSub(sqlabble.NewSelect()),
			"(SELECT)",
			`> (
>   SELECT
> )
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSub(sqlabble.NewSelect()).
				Eq(100),
			"(SELECT) = ?",
			`> (
>   SELECT
> )
> = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewColumn("foo").
				Eq(
					sqlabble.NewSub(
						sqlabble.NewSelect(
							sqlabble.NewColumn("bar"),
						),
					),
				),
			"foo = (SELECT bar)",
			`> foo =
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
		{
			sqlabble.NewSub(
				sqlabble.NewSelect(sqlabble.NewColumn("foo")),
			).Eq(
				sqlabble.NewSub(
					sqlabble.NewSelect(sqlabble.NewColumn("bar")),
				),
			),
			"(SELECT foo) = (SELECT bar)",
			`> (
>   SELECT
>     foo
> )
> =
> (
>   SELECT
>     bar
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestSelectType(t *testing.T) {
	s := statement.Select{}
	if _, ok := interface{}(s).(statement.Clause); !ok {
		t.Errorf("%T should implement statement.Clause", s)
	}
}

func TestSelectSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSelect(),
			`SELECT`,
			`> SELECT
`,
			nil,
		},
		{
			statement.NewSelect(
				statement.NewColumn("*"),
			),
			`SELECT *`,
			`> SELECT
>   *
`,
			nil,
		},
		{
			statement.NewSelect(
				statement.NewColumn("foo"),
			),
			`SELECT foo`,
			`> SELECT
>   foo
`,
			nil,
		},
		{
			statement.NewSelect(
				statement.NewColumn("foo").As("a"),
			),
			`SELECT foo AS "a"`,
			`> SELECT
>   foo AS "a"
`,
			nil,
		},
		{
			statement.NewSelect(
				statement.NewColumn("foo").As("a"),
				statement.NewColumn("bar").As("b"),
				statement.NewColumn("baz").As("c"),
			),
			`SELECT foo AS "a", bar AS "b", baz AS "c"`,
			`> SELECT
>   foo AS "a"
>   , bar AS "b"
>   , baz AS "c"
`,
			nil,
		},
		{
			statement.NewSelectDistinct(
				statement.NewColumn("foo").As("a"),
				statement.NewColumn("bar").As("b"),
				statement.NewColumn("baz").As("c"),
			),
			`SELECT DISTINCT foo AS "a", bar AS "b", baz AS "c"`,
			`> SELECT DISTINCT
>   foo AS "a"
>   , bar AS "b"
>   , baz AS "c"
`,
			nil,
		},
		{
			statement.NewSelect(
				statement.NewColumn("foo"),
			).From(
				statement.NewTable("bar"),
			),
			`SELECT foo FROM bar`,
			`> SELECT
>   foo
> FROM
>   bar
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestNonScalarAll(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewEqAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`= ALL (SELECT foo)`,
			`> = ALL (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewNotEqAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`!= ALL (SELECT foo)`,
			`> != ALL (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewGtAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`> ALL (SELECT foo)`,
			`> > ALL (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewGteAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`>= ALL (SELECT foo)`,
			`> >= ALL (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewLtAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`< ALL (SELECT foo)`,
			`> < ALL (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewLteAll(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`<= ALL (SELECT foo)`,
			`> <= ALL (
>   SELECT
>     foo
> )
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

func TestNonScalarAny(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewEqAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`= ANY (SELECT foo)`,
			`> = ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewNotEqAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`!= ANY (SELECT foo)`,
			`> != ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewGtAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`> ANY (SELECT foo)`,
			`> > ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewGteAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`>= ANY (SELECT foo)`,
			`> >= ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewLtAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`< ANY (SELECT foo)`,
			`> < ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewLteAny(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("foo"),
					),
				),
			),
			`<= ANY (SELECT foo)`,
			`> <= ANY (
>   SELECT
>     foo
> )
`,
			nil,
		},
		{
			statement.NewExists(
				statement.NewSubquery(
					statement.NewSelect(),
				),
			),
			`EXISTS (SELECT)`,
			`> EXISTS (
>   SELECT
> )
`,
			nil,
		},
		{
			statement.NewNotExists(
				statement.NewSubquery(
					statement.NewSelect(),
				),
			),
			`NOT EXISTS (SELECT)`,
			`> NOT EXISTS (
>   SELECT
> )
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

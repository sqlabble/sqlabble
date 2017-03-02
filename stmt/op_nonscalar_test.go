package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestNonScalarAll(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewEqAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewNotEqAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewGtAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewGteAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewLtAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewLteAll(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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

func TestNonScalarAny(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewEqAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewNotEqAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewGtAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewGteAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewLtAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewLteAny(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("foo"),
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
			stmt.NewExists(
				stmt.NewSubquery(
					stmt.NewSelect(),
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
			stmt.NewNotExists(
				stmt.NewSubquery(
					stmt.NewSelect(),
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

package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestFormulaSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewAdd(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? + ?`,
			`> ?
> + ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewSub(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? - ?`,
			`> ?
> - ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewMul(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? * ?`,
			`> ?
> * ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewDiv(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? / ?`,
			`> ?
> / ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewIntegerDiv(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? DIV ?`,
			`> ?
> DIV ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewMod(
				stmt.NewVal(1),
				stmt.NewVal(1),
			),
			`? % ?`,
			`> ?
> % ?
`,
			[]interface{}{
				1,
				1,
			},
		},
		{
			stmt.NewAdd(
				stmt.NewColumn(`foo`),
				stmt.NewVal(1),
			),
			`"foo" + ?`,
			`> "foo"
> + ?
`,
			[]interface{}{
				1,
			},
		},
		{
			stmt.NewAdd(
				stmt.NewSubquery(
					stmt.NewSelect(),
				),
				stmt.NewVal(1),
			),
			`(SELECT) + ?`,
			`> (
>   SELECT
> )
> + ?
`,
			[]interface{}{
				1,
			},
		},
		{
			stmt.NewAdd(
				stmt.NewMod(
					stmt.NewVal(1),
					stmt.NewVal(1),
				),
				stmt.NewVal(1),
			),
			`(? % ?) + ?`,
			`> (
>   ?
>   % ?
> )
> + ?
`,
			[]interface{}{
				1,
				1,
				1,
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

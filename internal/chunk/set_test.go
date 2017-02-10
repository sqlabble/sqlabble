package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestSetType(t *testing.T) {
	for _, c := range []interface{}{
		chunk.Set{},
	} {
		t.Run(fmt.Sprintf("Type %T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Clause); !ok {
				t.Errorf("%T should implement grammar.Clause", c)
			}
		})
	}
}

func TestSet(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewSet(
				chunk.NewColumn("foo").Assign(100),
			),
			"SET foo = ?",
			`> SET
>   foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			chunk.NewSet(
				chunk.NewColumn("foo").Assign(100),
				chunk.NewColumn("bar").Assign(200),
			),
			"SET foo = ?, bar = ?",
			`> SET
>   foo = ?
>   , bar = ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			chunk.NewSet(
				chunk.NewColumn("foo").Assign(100),
				chunk.NewColumn("bar").Assign(200),
				chunk.NewColumn("baz").Assign(300),
			),
			"SET foo = ?, bar = ?, baz = ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			chunk.NewSet(
				chunk.NewColumn("foo").Assign(100),
				chunk.NewColumn("bar").Assign(200),
				chunk.NewColumn("baz").Assign(300),
			).Where(
				chunk.NewColumn("qux").Lte(400),
			),
			"SET foo = ?, bar = ?, baz = ? WHERE qux <= ?",
			`> SET
>   foo = ?
>   , bar = ?
>   , baz = ?
> WHERE
>   qux <= ?
`,
			[]interface{}{
				100,
				200,
				300,
				400,
			},
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
			sql, values := sqlabble.BuildIndent(c.statement, "> ", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

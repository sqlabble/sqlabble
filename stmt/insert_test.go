package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestInsertSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
			),
			"INSERT INTO foo ()",
			`> INSERT INTO
>   foo (
>   )
`,
			nil,
		},
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
				stmt.NewColumn("name"),
				stmt.NewColumn("age"),
			),
			"INSERT INTO foo (name, age)",
			`> INSERT INTO
>   foo (
>     name
>     , age
>   )
`,
			nil,
		},
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
				stmt.NewColumn("name"),
				stmt.NewColumn("age"),
			).Values(
				stmt.NewVals("Obi-Wan Kenobi", 63),
			),
			"INSERT INTO foo (name, age) VALUES (?, ?)",
			`> INSERT INTO
>   foo (
>     name
>     , age
>   )
> VALUES
>   (?, ?)
`,
			[]interface{}{
				"Obi-Wan Kenobi",
				63,
			},
		},
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
				stmt.NewColumn("name"),
				stmt.NewColumn("age"),
			).Values(
				stmt.NewVals("Obi-Wan Kenobi", 63),
				stmt.NewVals("Luke Skywalker", 19),
			),
			"INSERT INTO foo (name, age) VALUES (?, ?), (?, ?)",
			`> INSERT INTO
>   foo (
>     name
>     , age
>   )
> VALUES
>   (?, ?)
>   , (?, ?)
`,
			[]interface{}{
				"Obi-Wan Kenobi",
				63,
				"Luke Skywalker",
				19,
			},
		},
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
				stmt.NewColumn("name"),
				stmt.NewColumn("age"),
			).DefaultValues(),
			"INSERT INTO foo (name, age) DEFAULT VALUES",
			`> INSERT INTO
>   foo (
>     name
>     , age
>   )
> DEFAULT VALUES
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

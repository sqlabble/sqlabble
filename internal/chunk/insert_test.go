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

func TestInsertType(t *testing.T) {
	if _, ok := interface{}(chunk.InsertInto{}).(grammar.Clause); !ok {
		t.Errorf("chunk.Insert doesn't implement grammar.Clause")
	}
}

func TestInsertSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewInsertInto(
				chunk.NewTable("foo"),
				chunk.NewColumn("name"),
				chunk.NewColumn("age"),
			),
			"INSERT INTO foo (name, age)",
			`> INSERT INTO
>   foo (name, age)
`,
			[]interface{}{},
		},
		{
			chunk.NewInsertInto(
				chunk.NewTable("foo"),
				chunk.NewColumn("name"),
				chunk.NewColumn("age"),
			).Values(
				"Obi-Wan Kenobi",
				53,
			),
			"INSERT INTO foo (name, age) VALUES (?, ?)",
			`> INSERT INTO
>   foo (name, age)
> VALUES
>   (?, ?)
`,
			[]interface{}{
				"Obi-Wan Kenobi",
				53,
			},
		},
		{
			chunk.NewInsertInto(
				chunk.NewTable("foo"),
				chunk.NewColumn("name"),
				chunk.NewColumn("age"),
			).Values(
				"Obi-Wan Kenobi",
				63,
			).Values(
				"Luke Skywalker",
				19,
			),
			"INSERT INTO foo (name, age) VALUES (?, ?), (?, ?)",
			`> INSERT INTO
>   foo (name, age)
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

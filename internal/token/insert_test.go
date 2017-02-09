package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/token"
)

func TestInsertType(t *testing.T) {
	if _, ok := interface{}(token.InsertInto{}).(grammar.Clause); !ok {
		t.Errorf("token.Insert doesn't implement grammar.Clause")
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
			token.NewInsertInto(
				token.NewTable("foo"),
				token.NewColumn("name"),
				token.NewColumn("age"),
			),
			"INSERT INTO foo (name, age)",
			`> INSERT INTO
>   foo (name, age)
`,
			[]interface{}{},
		},
		{
			token.NewInsertInto(
				token.NewTable("foo"),
				token.NewColumn("name"),
				token.NewColumn("age"),
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
			token.NewInsertInto(
				token.NewTable("foo"),
				token.NewColumn("name"),
				token.NewColumn("age"),
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

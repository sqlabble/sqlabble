package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestInsertType(t *testing.T) {
// 	if _, ok := interface{}(statement.InsertInto{}).(statement.ClauseNode); !ok {
// 		t.Errorf("statement.Insert doesn't implement statement.Clause")
// 	}
// }

func TestInsertSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewInsertInto(
				statement.NewTable("foo"),
			),
			"INSERT INTO foo",
			`> INSERT INTO
>   foo
`,
			[]interface{}{},
		},
		{
			statement.NewInsertInto(
				statement.NewTable("foo"),
				statement.NewColumn("name"),
				statement.NewColumn("age"),
			),
			"INSERT INTO foo (name, age)",
			`> INSERT INTO
>   foo (name, age)
`,
			[]interface{}{},
		},
		{
			statement.NewInsertInto(
				statement.NewTable("foo"),
				statement.NewColumn("name"),
				statement.NewColumn("age"),
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
			statement.NewInsertInto(
				statement.NewTable("foo"),
				statement.NewColumn("name"),
				statement.NewColumn("age"),
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
		{
			statement.NewInsertInto(
				statement.NewTable("foo"),
				statement.NewColumn("name"),
				statement.NewColumn("age"),
			).DefaultValues(),
			"INSERT INTO foo (name, age) DEFAULT VALUES",
			`> INSERT INTO
>   foo (name, age)
> DEFAULT VALUES
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
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

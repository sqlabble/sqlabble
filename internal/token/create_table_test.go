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

func TestCreateTableType(t *testing.T) {
	// if _, ok := interface{}(token.From{}).(grammar.Clause); !ok {
	// 	t.Errorf("token.FromClause doesn't implement grammar.Clause")
	// }
}

func TestCreateTableSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewCreateTable(
				token.NewTable("foo"),
			),
			"CREATE TABLE foo",
			`> CREATE TABLE
>   foo
`,
			[]interface{}{},
		},
		{
			token.NewCreateTable(
				token.NewTable("foo"),
			).Definitions(),
			"CREATE TABLE foo ()",
			`> CREATE TABLE
>   foo
>   (
>   )
`,
			[]interface{}{},
		},
		{
			token.NewCreateTable(
				token.NewTable("foo"),
			).Definitions(
				token.NewColumn("name").Definition("VARCHAR(255)"),
			),
			"CREATE TABLE foo (name VARCHAR(255))",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>   )
`,
			[]interface{}{},
		},
		{
			token.NewCreateTable(
				token.NewTable("foo"),
			).Definitions(
				token.NewColumn("name").Definition("VARCHAR(255)"),
				token.NewColumn("gender").Definition("ENUM('M', 'F')"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'))",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>   )
`,
			[]interface{}{},
		},
		{
			token.NewCreateTable(
				token.NewTable("foo"),
			).Definitions(
				token.NewColumn("name").Definition("VARCHAR(255)"),
				token.NewColumn("gender").Definition("ENUM('M', 'F')"),
				token.NewColumn("birth_date").Definition("DATE"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'), birth_date DATE)",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>     , birth_date DATE
>   )
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

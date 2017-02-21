package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestCreateTableType(t *testing.T) {
}

func TestCreateTableSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewCreateTable(
				statement.NewTable("foo"),
			),
			"CREATE TABLE foo",
			`> CREATE TABLE
>   foo
`,
			[]interface{}{},
		},
		{
			statement.NewCreateTableIfNotExists(
				statement.NewTable("foo"),
			),
			"CREATE TABLE IF NOT EXISTS foo",
			`> CREATE TABLE IF NOT EXISTS
>   foo
`,
			[]interface{}{},
		},
		{
			statement.NewCreateTable(
				statement.NewTable("foo"),
			).Definitions(),
			"CREATE TABLE foo ()",
			`> CREATE TABLE
>   foo (
>   )
`,
			[]interface{}{},
		},
		{
			statement.NewCreateTable(
				statement.NewTable("foo"),
			).Definitions(
				statement.NewColumn("name").Define("VARCHAR(255)"),
			),
			"CREATE TABLE foo (name VARCHAR(255))",
			`> CREATE TABLE
>   foo (
>     name VARCHAR(255)
>   )
`,
			[]interface{}{},
		},
		{
			statement.NewCreateTable(
				statement.NewTable("foo"),
			).Definitions(
				statement.NewColumn("name").Define("VARCHAR(255)"),
				statement.NewColumn("gender").Define("ENUM('M', 'F')"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'))",
			`> CREATE TABLE
>   foo (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>   )
`,
			[]interface{}{},
		},
		{
			statement.NewCreateTable(
				statement.NewTable("foo"),
			).Definitions(
				statement.NewColumn("name").Define("VARCHAR(255)"),
				statement.NewColumn("gender").Define("ENUM('M', 'F')"),
				statement.NewColumn("birth_date").Define("DATE"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'), birth_date DATE)",
			`> CREATE TABLE
>   foo (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>     , birth_date DATE
>   )
`,
			[]interface{}{},
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

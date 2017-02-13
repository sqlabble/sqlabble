package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestCreateTableType(t *testing.T) {
}

func TestCreateTableSQL(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewCreateTable(
				sqlabble.NewTable("foo"),
			),
			"CREATE TABLE foo",
			`> CREATE TABLE
>   foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewCreateTable(
				sqlabble.NewTable("foo"),
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
			sqlabble.NewCreateTable(
				sqlabble.NewTable("foo"),
			).Definitions(
				sqlabble.NewColumn("name").Define("VARCHAR(255)"),
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
			sqlabble.NewCreateTable(
				sqlabble.NewTable("foo"),
			).Definitions(
				sqlabble.NewColumn("name").Define("VARCHAR(255)"),
				sqlabble.NewColumn("gender").Define("ENUM('M', 'F')"),
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
			sqlabble.NewCreateTable(
				sqlabble.NewTable("foo"),
			).Definitions(
				sqlabble.NewColumn("name").Define("VARCHAR(255)"),
				sqlabble.NewColumn("gender").Define("ENUM('M', 'F')"),
				sqlabble.NewColumn("birth_date").Define("DATE"),
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

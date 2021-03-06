package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestCreateTableType(t *testing.T) {
	t.Parallel()
}

func TestCreateTableSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewCreateTable(
				stmt.NewTable("foo"),
			),
			`CREATE TABLE "foo"`,
			`> CREATE TABLE
>   "foo"
`,
			nil,
		},
		{
			stmt.NewCreateTableIfNotExists(
				stmt.NewTable("foo"),
			),
			`CREATE TABLE IF NOT EXISTS "foo"`,
			`> CREATE TABLE IF NOT EXISTS
>   "foo"
`,
			nil,
		},
		{
			stmt.NewCreateTable(
				stmt.NewTable("foo"),
			).Definitions(),
			`CREATE TABLE "foo" ()`,
			`> CREATE TABLE
>   "foo" (
>   )
`,
			nil,
		},
		{
			stmt.NewCreateTable(
				stmt.NewTable("foo"),
			).Definitions(
				stmt.NewColumn("name").Define("VARCHAR(255)"),
			),
			`CREATE TABLE "foo" ("name" VARCHAR(255))`,
			`> CREATE TABLE
>   "foo" (
>     "name" VARCHAR(255)
>   )
`,
			nil,
		},
		{
			stmt.NewCreateTable(
				stmt.NewTable("foo"),
			).Definitions(
				stmt.NewColumn("name").Define("VARCHAR(255)"),
				stmt.NewColumn("gender").Define("ENUM('M', 'F')"),
			),
			`CREATE TABLE "foo" ("name" VARCHAR(255), "gender" ENUM('M', 'F'))`,
			`> CREATE TABLE
>   "foo" (
>     "name" VARCHAR(255)
>     , "gender" ENUM('M', 'F')
>   )
`,
			nil,
		},
		{
			stmt.NewCreateTable(
				stmt.NewTable("foo"),
			).Definitions(
				stmt.NewColumn("name").Define("VARCHAR(255)"),
				stmt.NewColumn("gender").Define("ENUM('M', 'F')"),
				stmt.NewColumn("birth_date").Define("DATE"),
			),
			`CREATE TABLE "foo" ("name" VARCHAR(255), "gender" ENUM('M', 'F'), "birth_date" DATE)`,
			`> CREATE TABLE
>   "foo" (
>     "name" VARCHAR(255)
>     , "gender" ENUM('M', 'F')
>     , "birth_date" DATE
>   )
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

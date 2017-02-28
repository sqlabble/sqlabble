package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestSimpleCase(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.
				NewSimpleCase(
					statement.NewColumn("customer.cust_type_cd"),
					statement.
						NewSimpleWhen(
							statement.NewParam("I"),
						).
						Then(
							statement.NewSubquery(
								statement.NewSelect(statement.NewConcat(
									statement.NewColumn("i.fname"),
									statement.NewParam(" "),
									statement.NewColumn("i.lname"),
								)).From(
									statement.NewTable("individual").As("i"),
								).Where(
									statement.NewColumn("i.cust_id").Eq(statement.NewColumn("customer.cust_id")),
								),
							),
						),
				),
			`CASE customer.cust_type_cd WHEN ? THEN (SELECT CONCAT(i.fname, ?, i.lname) FROM individual AS "i" WHERE i.cust_id = customer.cust_id) END`,
			`> CASE customer.cust_type_cd
>   WHEN
>     ?
>   THEN
>     (
>       SELECT
>         CONCAT(i.fname, ?, i.lname)
>       FROM
>         individual AS "i"
>       WHERE
>         i.cust_id = customer.cust_id
>     )
> END
`,
			[]interface{}{
				"I",
				" ",
			},
		},
		{
			statement.
				NewSimpleCase(
					statement.NewColumn("customer.cust_type_cd"),
					statement.
						NewSimpleWhen(
							statement.NewParam("I"),
						).
						Then(
							statement.NewSubquery(
								statement.NewSelect(statement.NewConcat(
									statement.NewColumn("i.fname"),
									statement.NewParam(" "),
									statement.NewColumn("i.lname"),
								)).From(
									statement.NewTable("individual").As("i"),
								).Where(
									statement.NewColumn("i.cust_id").Eq(statement.NewColumn("customer.cust_id")),
								),
							),
						).
						When(
							statement.NewParam("B"),
						).
						Then(
							statement.NewSubquery(
								statement.NewSelect(statement.NewColumn("b.name")).
									From(statement.NewTable("business").As("b")).
									Where(
										statement.NewColumn("b.cust_id").Eq(statement.NewColumn("customer.cust_id")),
									),
							),
						).
						Else(statement.NewParam("Unknown Customer Type")),
				),
			`CASE customer.cust_type_cd WHEN ? THEN (SELECT CONCAT(i.fname, ?, i.lname) FROM individual AS "i" WHERE i.cust_id = customer.cust_id) WHEN ? THEN (SELECT b.name FROM business AS "b" WHERE b.cust_id = customer.cust_id) ELSE ? END`,
			`> CASE customer.cust_type_cd
>   WHEN
>     ?
>   THEN
>     (
>       SELECT
>         CONCAT(i.fname, ?, i.lname)
>       FROM
>         individual AS "i"
>       WHERE
>         i.cust_id = customer.cust_id
>     )
>   WHEN
>     ?
>   THEN
>     (
>       SELECT
>         b.name
>       FROM
>         business AS "b"
>       WHERE
>         b.cust_id = customer.cust_id
>     )
>   ELSE
>     ?
> END
`,
			[]interface{}{
				"I",
				" ",
				"B",
				"Unknown Customer Type",
			},
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

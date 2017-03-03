package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSimpleCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.
				NewSimpleCase(
					stmt.NewColumn("customer.cust_type_cd"),
					stmt.
						NewSimpleWhen(
							stmt.NewParam("I"),
						).
						Then(
							stmt.NewSubquery(
								stmt.NewSelect(stmt.NewConcat(
									stmt.NewColumn("i.fname"),
									stmt.NewParam(" "),
									stmt.NewColumn("i.lname"),
								)).From(
									stmt.NewTable("individual").As("i"),
								).Where(
									stmt.NewColumn("i.cust_id").Eq(stmt.NewColumn("customer.cust_id")),
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
			stmt.
				NewSimpleCase(
					stmt.NewColumn("customer.cust_type_cd"),
					stmt.
						NewSimpleWhen(
							stmt.NewParam("I"),
						).
						Then(
							stmt.NewSubquery(
								stmt.NewSelect(stmt.NewConcat(
									stmt.NewColumn("i.fname"),
									stmt.NewParam(" "),
									stmt.NewColumn("i.lname"),
								)).From(
									stmt.NewTable("individual").As("i"),
								).Where(
									stmt.NewColumn("i.cust_id").Eq(stmt.NewColumn("customer.cust_id")),
								),
							),
						).
						When(
							stmt.NewParam("B"),
						).
						Then(
							stmt.NewSubquery(
								stmt.NewSelect(stmt.NewColumn("b.name")).
									From(stmt.NewTable("business").As("b")).
									Where(
										stmt.NewColumn("b.cust_id").Eq(stmt.NewColumn("customer.cust_id")),
									),
							),
						).
						Else(stmt.NewParam("Unknown Customer Type")),
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

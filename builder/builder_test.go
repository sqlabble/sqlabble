package builder_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/builder"
	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
	"github.com/sqlabble/sqlabble/token"
)

var (
	b  = builder.Standard
	bi = builder.NewBuilder(token.NewFormat("> ", "  ", `"`, "\n"))
)

func TestBuilder(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSelect(),
			`SELECT`,
			`> SELECT
`,
			nil,
		},
		{
			stmt.NewSelect(
				stmt.NewColumn("created_at"),
				stmt.NewColumn("name").As("n"),
				stmt.NewColumn("gender").As("g"),
				stmt.NewColumn("age"),
			).From(
				stmt.NewTable("users"),
			).Where(
				stmt.NewAnd(
					stmt.NewColumn("g").Eq(stmt.NewVal("male")),
					stmt.NewOr(
						stmt.NewColumn("age").Lt(stmt.NewVal(20)),
						stmt.NewColumn("age").Eq(stmt.NewVal(30)),
						stmt.NewColumn("age").Gte(stmt.NewVal(50)),
					),
					stmt.NewColumn("created_at").Between(stmt.NewVal("2016-01-01"), stmt.NewVal("2016-12-31")),
				),
			).OrderBy(
				stmt.NewColumn("created_at").Desc(),
				stmt.NewColumn("id").Asc(),
			).Limit(
				20,
			).Offset(
				20 * 5,
			),
			`SELECT created_at, name AS "n", gender AS "g", age FROM users WHERE g = ? AND (age < ? OR age = ? OR age >= ?) AND created_at BETWEEN ? AND ? ORDER BY created_at DESC, id ASC LIMIT ? OFFSET ?`,
			`> SELECT
>   created_at
>   , name AS "n"
>   , gender AS "g"
>   , age
> FROM
>   users
> WHERE
>   g = ?
>   AND (
>     age < ?
>     OR age = ?
>     OR age >= ?
>   )
>   AND created_at BETWEEN ? AND ?
> ORDER BY
>   created_at DESC
>   , id ASC
> LIMIT
>   ?
> OFFSET
>   ?
`,
			[]interface{}{
				`male`,
				20,
				30,
				50,
				`2016-01-01`,
				`2016-12-31`,
				20,
				100,
			},
		},
		{
			stmt.NewInsertInto(
				stmt.NewTable("foo"),
				stmt.NewColumn("name"),
				stmt.NewColumn("age"),
			).Values(
				stmt.NewVals(`Obi-Wan Kenobi`, 63),
				stmt.NewVals(`Luke Skywalker`, 19),
			),
			`INSERT INTO foo ("name", "age") VALUES (?, ?), (?, ?)`,
			`> INSERT INTO
>   foo (
>     "name"
>     , "age"
>   )
> VALUES
>   (?, ?)
>   , (?, ?)
`,
			[]interface{}{
				`Obi-Wan Kenobi`,
				63,
				`Luke Skywalker`,
				19,
			},
		},
		{
			stmt.NewDelete().From(
				stmt.NewTable("login_history"),
			).Where(
				stmt.NewColumn("login_date").Lt(stmt.NewVal("2004-07-02 09:00:00")),
			),
			`DELETE FROM login_history WHERE login_date < ?`,
			`> DELETE
> FROM
>   login_history
> WHERE
>   login_date < ?
`,
			[]interface{}{
				`2004-07-02 09:00:00`,
			},
		},
		{
			stmt.NewUnion(
				stmt.NewSelect(
					stmt.NewColumn("emp_id"),
				).From(
					stmt.NewTable("employee"),
				).Where(
					stmt.NewAnd(
						stmt.NewColumn("assigned_branch_id").Eq(stmt.NewVal(2)),
						stmt.NewOr(
							stmt.NewColumn("title").Eq(stmt.NewVal("Teller")),
							stmt.NewColumn("title").Eq(stmt.NewVal("Head Teller")),
						),
					),
				),
				stmt.NewSelectDistinct(
					stmt.NewColumn("open_emp_id"),
				).From(
					stmt.NewTable("account"),
				).Where(
					stmt.NewColumn("open_branch_id").Eq(stmt.NewVal(2)),
				),
			),
			`(SELECT "emp_id" FROM "employee" WHERE "assigned_branch_id" = ? AND ("title" = ? OR "title" = ?)) UNION (SELECT DISTINCT "open_emp_id" FROM "account" WHERE "open_branch_id" = ?)`,
			`> (
>   SELECT
>     "emp_id"
>   FROM
>     "employee"
>   WHERE
>     "assigned_branch_id" = ?
>     AND (
>       "title" = ?
>       OR "title" = ?
>     )
> )
> UNION (
>   SELECT DISTINCT
>     "open_emp_id"
>   FROM
>     "account"
>   WHERE
>     "open_branch_id" = ?
> )
`,
			[]interface{}{
				2,
				`Teller`,
				`Head Teller`,
				2,
			},
		},
	} {
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

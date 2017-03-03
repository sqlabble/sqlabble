package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestSearchedCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewSearchedCase(
				stmt.NewSearchedWhen(
					stmt.NewColumn("employee.title").Eq(stmt.NewParam("Head Teller")),
				).
					Then(stmt.NewParam("Head Teller")),
			),
			`CASE WHEN employee.title = ? THEN ? END`,
			`> CASE
>   WHEN
>     employee.title = ?
>   THEN
>     ?
> END
`,
			[]interface{}{
				"Head Teller",
				"Head Teller",
			},
		},
		{
			stmt.NewSearchedCase(
				stmt.NewSearchedWhen(
					stmt.NewColumn("employee.title").Eq(stmt.NewParam("Head Teller")),
				).
					Then(stmt.NewParam("Head Teller")).
					When(
						stmt.NewAnd(
							stmt.NewColumn("employee.title").Eq(stmt.NewParam("Teller")),
							stmt.NewYear(stmt.NewColumn("employee.start_date")).Gt(stmt.NewParam(2004)),
						),
					).
					Then(stmt.NewParam("Teller Trainee")).
					When(
						stmt.NewAnd(
							stmt.NewColumn("employee.title").Eq(stmt.NewParam("Teller")),
							stmt.NewYear(stmt.NewColumn("employee.start_date")).Lt(stmt.NewParam(2003)),
						),
					).
					Then(stmt.NewParam("Experienced Teller")).
					When(
						stmt.NewAnd(
							stmt.NewColumn("employee.title").Eq(stmt.NewParam("Teller")),
						),
					).
					Then(stmt.NewParam("Teller")).
					Else(stmt.NewParam("Non-Teller")),
			),
			`CASE WHEN employee.title = ? THEN ? WHEN employee.title = ? AND YEAR(employee.start_date) > ? THEN ? WHEN employee.title = ? AND YEAR(employee.start_date) < ? THEN ? WHEN employee.title = ? THEN ? ELSE ? END`,
			`> CASE
>   WHEN
>     employee.title = ?
>   THEN
>     ?
>   WHEN
>     employee.title = ?
>     AND YEAR(employee.start_date) > ?
>   THEN
>     ?
>   WHEN
>     employee.title = ?
>     AND YEAR(employee.start_date) < ?
>   THEN
>     ?
>   WHEN
>     employee.title = ?
>   THEN
>     ?
>   ELSE
>     ?
> END
`,
			[]interface{}{
				"Head Teller",
				"Head Teller",
				"Teller",
				2004,
				"Teller Trainee",
				"Teller",
				2003,
				"Experienced Teller",
				"Teller",
				"Teller",
				"Non-Teller",
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

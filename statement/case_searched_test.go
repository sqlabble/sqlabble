package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestSearchedCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSearchedCase(
				statement.NewSearchedWhen(
					statement.NewColumn("employee.title").Eq(statement.NewParam("Head Teller")),
				).
					Then(statement.NewParam("Head Teller")),
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
			statement.NewSearchedCase(
				statement.NewSearchedWhen(
					statement.NewColumn("employee.title").Eq(statement.NewParam("Head Teller")),
				).
					Then(statement.NewParam("Head Teller")).
					When(
						statement.NewAnd(
							statement.NewColumn("employee.title").Eq(statement.NewParam("Teller")),
							statement.NewYear(statement.NewColumn("employee.start_date")).Gt(statement.NewParam(2004)),
						),
					).
					Then(statement.NewParam("Teller Trainee")).
					When(
						statement.NewAnd(
							statement.NewColumn("employee.title").Eq(statement.NewParam("Teller")),
							statement.NewYear(statement.NewColumn("employee.start_date")).Lt(statement.NewParam(2003)),
						),
					).
					Then(statement.NewParam("Experienced Teller")).
					When(
						statement.NewAnd(
							statement.NewColumn("employee.title").Eq(statement.NewParam("Teller")),
						),
					).
					Then(statement.NewParam("Teller")).
					Else(statement.NewParam("Non-Teller")),
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
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
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

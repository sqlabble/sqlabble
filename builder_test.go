package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	q "github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/internal/diff"
)

var (
	builder       = q.Standard
	builderIndent = q.NewBuilder(
		generator.Options{
			Prefix: "> ",
			Indent: "  ",
		})
	builderMySQL4       = q.MySQL4
	builderIndentMySQL4 = q.NewBuilder(
		generator.Options{
			Prefix:   "> ",
			Indent:   "  ",
			FlatSets: true,
		})
)

func TestBuild(t *testing.T) {
	q.C("created_at").Assign(300)
	q.Select(q.Column("foo")).From(q.T("bar")).Where(q.Column("baz").Eq(100)).OrderBy(q.C("quz").Asc()).Limit(0, 20)

	for i, c := range []struct {
		statement q.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			q.Select(
				q.C("created_at"),
				q.C("name").As("n"),
				q.C("gender").As("g"),
				q.C("age"),
			).From(
				q.T("users"),
			).Where(
				q.And(
					q.C("g").Eq("male"),
					q.Or(
						q.C("age").Lt(20),
						q.C("age").Eq(30),
						q.C("age").Gte(50),
					),
					q.C("created_at").Between("2016-01-01", "2016-12-31"),
				),
			).OrderBy(
				q.C("created_at").Desc(),
				q.C("id").Asc(),
			).Limit(
				20*5, 20,
			),
			"SELECT created_at, name AS n, gender AS g, age FROM users WHERE g = ? AND (age < ? OR age = ? OR age >= ?) AND created_at BETWEEN ? AND ? ORDER BY created_at DESC, id ASC LIMIT ?, ?",
			`> SELECT
>   created_at
>   , name AS n
>   , gender AS g
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
>   ?, ?
`,
			[]interface{}{
				"male",
				20,
				30,
				50,
				"2016-01-01",
				"2016-12-31",
				100,
				20,
			},
		},
		{
			q.InsertInto(
				q.T("foo"),
				q.C("name"),
				q.C("age"),
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
			q.Delete().From(
				q.T("login_history"),
			).Where(
				q.C("login_date").Lt("2004-07-02 09:00:00"),
			),
			"DELETE FROM login_history WHERE login_date < ?",
			`> DELETE
> FROM
>   login_history
> WHERE
>   login_date < ?
`,
			[]interface{}{
				"2004-07-02 09:00:00",
			},
		},
		{
			q.Union(
				q.Select(
					q.Column("emp_id"),
				).From(
					q.Table("employee"),
				).Where(
					q.And(
						q.Column("assigned_branch_id").Eq(2),
						q.Or(
							q.Column("title").Eq("Teller"),
							q.Column("title").Eq("Head Teller"),
						),
					),
				),
				q.SelectDistinct(
					q.Column("open_emp_id"),
				).From(
					q.Table("account"),
				).Where(
					q.Column("open_branch_id").Eq(2),
				),
			),
			"(SELECT emp_id FROM employee WHERE assigned_branch_id = ? AND (title = ? OR title = ?)) UNION (SELECT DISTINCT open_emp_id FROM account WHERE open_branch_id = ?)",
			`> (
>   SELECT
>     emp_id
>   FROM
>     employee
>   WHERE
>     assigned_branch_id = ?
>     AND (
>       title = ?
>       OR title = ?
>     )
> )
> UNION
> (
>   SELECT DISTINCT
>     open_emp_id
>   FROM
>     account
>   WHERE
>     open_branch_id = ?
> )
`,
			[]interface{}{
				2,
				"Teller",
				"Head Teller",
				2,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := q.Build(c.statement)
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

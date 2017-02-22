package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestJoinType(t *testing.T) {
	for _, c := range []interface{}{
		statement.Join{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(statement.Joiner); !ok {
				t.Errorf("%T should implement statement.Joiner", c)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewJoin(
				statement.NewTable("foo"),
			),
			`JOIN foo`,
			`> JOIN foo
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			),
			`JOIN foo AS "f"`,
			`> JOIN foo AS "f"
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("emp_id"),
					).From(
						statement.NewTable("employee"),
					),
				).As("e"),
			),
			`JOIN (SELECT emp_id FROM employee) AS e`,
			`> JOIN (
>   SELECT
>     emp_id
>   FROM
>     employee
> ) AS e
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).Join(
				statement.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> JOIN bar AS "b"
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).InnerJoin(
				statement.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" INNER JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> INNER JOIN bar AS "b"
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).LeftJoin(
				statement.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" LEFT JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> LEFT JOIN bar AS "b"
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).RightJoin(
				statement.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" RIGHT JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> RIGHT JOIN bar AS "b"
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			`JOIN foo ON foo.id = bar.id`,
			`> JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			`JOIN foo AS "f" ON f.id = b.id`,
			`> JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			`JOIN foo USING id`,
			`> JOIN foo USING id
`,
			nil,
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			`JOIN foo AS "f" USING id`,
			`> JOIN foo AS "f" USING id
`,
			nil,
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

func TestInnerJoin(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			),
			`INNER JOIN foo`,
			`> INNER JOIN foo
`,
			nil,
		},
		// 1
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			),
			`INNER JOIN foo AS "f"`,
			`> INNER JOIN foo AS "f"
`,
			nil,
		},
		{
			statement.NewInnerJoin(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("emp_id"),
					).From(
						statement.NewTable("employee"),
					),
				).As("e"),
			),
			`INNER JOIN (SELECT emp_id FROM employee) AS e`,
			`> INNER JOIN (
>   SELECT
>     emp_id
>   FROM
>     employee
> ) AS e
`,
			nil,
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			`INNER JOIN foo ON foo.id = bar.id`,
			`> INNER JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		// 3
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			`INNER JOIN foo AS "f" ON f.id = b.id`,
			`> INNER JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		// 4
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			`INNER JOIN foo USING id`,
			`> INNER JOIN foo USING id
`,
			nil,
		},
		// 5
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			`INNER JOIN foo AS "f" USING id`,
			`> INNER JOIN foo AS "f" USING id
`,
			nil,
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

func TestLeftJoin(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewLeftJoin(
				statement.NewTable("foo"),
			),
			`LEFT JOIN foo`,
			`> LEFT JOIN foo
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			),
			`LEFT JOIN foo AS "f"`,
			`> LEFT JOIN foo AS "f"
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("emp_id"),
					).From(
						statement.NewTable("employee"),
					),
				).As("e"),
			),
			`LEFT JOIN (SELECT emp_id FROM employee) AS e`,
			`> LEFT JOIN (
>   SELECT
>     emp_id
>   FROM
>     employee
> ) AS e
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			`LEFT JOIN foo ON foo.id = bar.id`,
			`> LEFT JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			`LEFT JOIN foo AS "f" ON f.id = b.id`,
			`> LEFT JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			`LEFT JOIN foo USING id`,
			`> LEFT JOIN foo USING id
`,
			nil,
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			`LEFT JOIN foo AS "f" USING id`,
			`> LEFT JOIN foo AS "f" USING id
`,
			nil,
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

func TestRightJoin(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewRightJoin(
				statement.NewTable("foo"),
			),
			`RIGHT JOIN foo`,
			`> RIGHT JOIN foo
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			),
			`RIGHT JOIN foo AS "f"`,
			`> RIGHT JOIN foo AS "f"
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewSubquery(
					statement.NewSelect(
						statement.NewColumn("emp_id"),
					).From(
						statement.NewTable("employee"),
					),
				).As("e"),
			),
			`RIGHT JOIN (SELECT emp_id FROM employee) AS e`,
			`> RIGHT JOIN (
>   SELECT
>     emp_id
>   FROM
>     employee
> ) AS e
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			`RIGHT JOIN foo ON foo.id = bar.id`,
			`> RIGHT JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			`RIGHT JOIN foo AS "f" ON f.id = b.id`,
			`> RIGHT JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			`RIGHT JOIN foo USING id`,
			`> RIGHT JOIN foo USING id
`,
			nil,
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			`RIGHT JOIN foo AS "f" USING id`,
			`> RIGHT JOIN foo AS "f" USING id
`,
			nil,
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

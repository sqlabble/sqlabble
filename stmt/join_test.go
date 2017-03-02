package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestJoinType(t *testing.T) {
	t.Parallel()
	for _, c := range []interface{}{
		stmt.Join{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			t.Parallel()
			if _, ok := c.(stmt.Joiner); !ok {
				t.Errorf("%T should implement stmt.Joiner", c)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewJoin(
				stmt.NewTable("foo"),
			),
			`JOIN foo`,
			`> JOIN foo
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			),
			`JOIN foo AS "f"`,
			`> JOIN foo AS "f"
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("emp_id"),
					).From(
						stmt.NewTable("employee"),
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
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).Join(
				stmt.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> JOIN bar AS "b"
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).InnerJoin(
				stmt.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" INNER JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> INNER JOIN bar AS "b"
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).LeftJoin(
				stmt.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" LEFT JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> LEFT JOIN bar AS "b"
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).RightJoin(
				stmt.NewTable("bar").As("b"),
			),
			`JOIN foo AS "f" RIGHT JOIN bar AS "b"`,
			`> JOIN foo AS "f"
> RIGHT JOIN bar AS "b"
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo"),
			).On(
				stmt.NewColumn("foo.id"),
				stmt.NewColumn("bar.id"),
			),
			`JOIN foo ON foo.id = bar.id`,
			`> JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).On(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			),
			`JOIN foo AS "f" ON f.id = b.id`,
			`> JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo"),
			).Using(
				stmt.NewColumn("id"),
			),
			`JOIN foo USING id`,
			`> JOIN foo USING id
`,
			nil,
		},
		{
			stmt.NewJoin(
				stmt.NewTable("foo").As("f"),
			).Using(
				stmt.NewColumn("id"),
			),
			`JOIN foo AS "f" USING id`,
			`> JOIN foo AS "f" USING id
`,
			nil,
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

func TestInnerJoin(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			stmt.NewInnerJoin(
				stmt.NewTable("foo"),
			),
			`INNER JOIN foo`,
			`> INNER JOIN foo
`,
			nil,
		},
		// 1
		{
			stmt.NewInnerJoin(
				stmt.NewTable("foo").As("f"),
			),
			`INNER JOIN foo AS "f"`,
			`> INNER JOIN foo AS "f"
`,
			nil,
		},
		{
			stmt.NewInnerJoin(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("emp_id"),
					).From(
						stmt.NewTable("employee"),
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
			stmt.NewInnerJoin(
				stmt.NewTable("foo"),
			).On(
				stmt.NewColumn("foo.id"),
				stmt.NewColumn("bar.id"),
			),
			`INNER JOIN foo ON foo.id = bar.id`,
			`> INNER JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		// 3
		{
			stmt.NewInnerJoin(
				stmt.NewTable("foo").As("f"),
			).On(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			),
			`INNER JOIN foo AS "f" ON f.id = b.id`,
			`> INNER JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		// 4
		{
			stmt.NewInnerJoin(
				stmt.NewTable("foo"),
			).Using(
				stmt.NewColumn("id"),
			),
			`INNER JOIN foo USING id`,
			`> INNER JOIN foo USING id
`,
			nil,
		},
		// 5
		{
			stmt.NewInnerJoin(
				stmt.NewTable("foo").As("f"),
			).Using(
				stmt.NewColumn("id"),
			),
			`INNER JOIN foo AS "f" USING id`,
			`> INNER JOIN foo AS "f" USING id
`,
			nil,
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

func TestLeftJoin(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewLeftJoin(
				stmt.NewTable("foo"),
			),
			`LEFT JOIN foo`,
			`> LEFT JOIN foo
`,
			nil,
		},
		{
			stmt.NewLeftJoin(
				stmt.NewTable("foo").As("f"),
			),
			`LEFT JOIN foo AS "f"`,
			`> LEFT JOIN foo AS "f"
`,
			nil,
		},
		{
			stmt.NewLeftJoin(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("emp_id"),
					).From(
						stmt.NewTable("employee"),
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
			stmt.NewLeftJoin(
				stmt.NewTable("foo"),
			).On(
				stmt.NewColumn("foo.id"),
				stmt.NewColumn("bar.id"),
			),
			`LEFT JOIN foo ON foo.id = bar.id`,
			`> LEFT JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			stmt.NewLeftJoin(
				stmt.NewTable("foo").As("f"),
			).On(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			),
			`LEFT JOIN foo AS "f" ON f.id = b.id`,
			`> LEFT JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			stmt.NewLeftJoin(
				stmt.NewTable("foo"),
			).Using(
				stmt.NewColumn("id"),
			),
			`LEFT JOIN foo USING id`,
			`> LEFT JOIN foo USING id
`,
			nil,
		},
		{
			stmt.NewLeftJoin(
				stmt.NewTable("foo").As("f"),
			).Using(
				stmt.NewColumn("id"),
			),
			`LEFT JOIN foo AS "f" USING id`,
			`> LEFT JOIN foo AS "f" USING id
`,
			nil,
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

func TestRightJoin(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewRightJoin(
				stmt.NewTable("foo"),
			),
			`RIGHT JOIN foo`,
			`> RIGHT JOIN foo
`,
			nil,
		},
		{
			stmt.NewRightJoin(
				stmt.NewTable("foo").As("f"),
			),
			`RIGHT JOIN foo AS "f"`,
			`> RIGHT JOIN foo AS "f"
`,
			nil,
		},
		{
			stmt.NewRightJoin(
				stmt.NewSubquery(
					stmt.NewSelect(
						stmt.NewColumn("emp_id"),
					).From(
						stmt.NewTable("employee"),
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
			stmt.NewRightJoin(
				stmt.NewTable("foo"),
			).On(
				stmt.NewColumn("foo.id"),
				stmt.NewColumn("bar.id"),
			),
			`RIGHT JOIN foo ON foo.id = bar.id`,
			`> RIGHT JOIN foo ON foo.id = bar.id
`,
			nil,
		},
		{
			stmt.NewRightJoin(
				stmt.NewTable("foo").As("f"),
			).On(
				stmt.NewColumn("f.id"),
				stmt.NewColumn("b.id"),
			),
			`RIGHT JOIN foo AS "f" ON f.id = b.id`,
			`> RIGHT JOIN foo AS "f" ON f.id = b.id
`,
			nil,
		},
		{
			stmt.NewRightJoin(
				stmt.NewTable("foo"),
			).Using(
				stmt.NewColumn("id"),
			),
			`RIGHT JOIN foo USING id`,
			`> RIGHT JOIN foo USING id
`,
			nil,
		},
		{
			stmt.NewRightJoin(
				stmt.NewTable("foo").As("f"),
			).Using(
				stmt.NewColumn("id"),
			),
			`RIGHT JOIN foo AS "f" USING id`,
			`> RIGHT JOIN foo AS "f" USING id
`,
			nil,
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

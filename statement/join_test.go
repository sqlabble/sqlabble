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
			"JOIN foo",
			`> JOIN foo
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			),
			"JOIN foo AS f",
			`> JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).Join(
				statement.NewTable("bar").As("b"),
			),
			"JOIN foo AS f JOIN bar AS b",
			`> JOIN foo AS f
> JOIN bar AS b
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).InnerJoin(
				statement.NewTable("bar").As("b"),
			),
			"JOIN foo AS f INNER JOIN bar AS b",
			`> JOIN foo AS f
> INNER JOIN bar AS b
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).LeftJoin(
				statement.NewTable("bar").As("b"),
			),
			"JOIN foo AS f LEFT JOIN bar AS b",
			`> JOIN foo AS f
> LEFT JOIN bar AS b
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).RightJoin(
				statement.NewTable("bar").As("b"),
			),
			"JOIN foo AS f RIGHT JOIN bar AS b",
			`> JOIN foo AS f
> RIGHT JOIN bar AS b
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			"JOIN foo ON foo.id = bar.id",
			`> JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			"JOIN foo AS f ON f.id = b.id",
			`> JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			"JOIN foo USING id",
			`> JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			statement.NewJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			"JOIN foo AS f USING id",
			`> JOIN foo AS f USING id
`,
			[]interface{}{},
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
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			),
			"INNER JOIN foo",
			`> INNER JOIN foo
`,
			[]interface{}{},
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			),
			"INNER JOIN foo AS f",
			`> INNER JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			"INNER JOIN foo ON foo.id = bar.id",
			`> INNER JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			"INNER JOIN foo AS f ON f.id = b.id",
			`> INNER JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			"INNER JOIN foo USING id",
			`> INNER JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			statement.NewInnerJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			"INNER JOIN foo AS f USING id",
			`> INNER JOIN foo AS f USING id
`,
			[]interface{}{},
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
			"LEFT JOIN foo",
			`> LEFT JOIN foo
`,
			[]interface{}{},
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			),
			"LEFT JOIN foo AS f",
			`> LEFT JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			"LEFT JOIN foo ON foo.id = bar.id",
			`> LEFT JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			"LEFT JOIN foo AS f ON f.id = b.id",
			`> LEFT JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			"LEFT JOIN foo USING id",
			`> LEFT JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			statement.NewLeftJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			"LEFT JOIN foo AS f USING id",
			`> LEFT JOIN foo AS f USING id
`,
			[]interface{}{},
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
			"RIGHT JOIN foo",
			`> RIGHT JOIN foo
`,
			[]interface{}{},
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			),
			"RIGHT JOIN foo AS f",
			`> RIGHT JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo"),
			).On(
				statement.NewColumn("foo.id"),
				statement.NewColumn("bar.id"),
			),
			"RIGHT JOIN foo ON foo.id = bar.id",
			`> RIGHT JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			).On(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			"RIGHT JOIN foo AS f ON f.id = b.id",
			`> RIGHT JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo"),
			).Using(
				statement.NewColumn("id"),
			),
			"RIGHT JOIN foo USING id",
			`> RIGHT JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			statement.NewRightJoin(
				statement.NewTable("foo").As("f"),
			).Using(
				statement.NewColumn("id"),
			),
			"RIGHT JOIN foo AS f USING id",
			`> RIGHT JOIN foo AS f USING id
`,
			[]interface{}{},
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

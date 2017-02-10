package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/token"
)

func TestTableType(t *testing.T) {
	for _, c := range []interface{}{
		token.Table{},
		token.TableAs{},
	} {
		t.Run(fmt.Sprintf("Table type %T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Table); !ok {
				t.Errorf("%T shouldn't implement grammar.Table", c)
			}
		})
	}
}

func TestTable(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewTable("foo"),
			"foo",
			`foo
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").
				Join(token.NewTable("bar")),
			"foo JOIN bar",
			`foo
JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").
				InnerJoin(token.NewTable("bar")),
			"foo INNER JOIN bar",
			`foo
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").
				LeftJoin(token.NewTable("bar")),
			"foo LEFT JOIN bar",
			`foo
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").
				RightJoin(token.NewTable("bar")),
			"foo RIGHT JOIN bar",
			`foo
RIGHT JOIN bar
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestAs(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewTable("foo").As("f"),
			"foo AS f",
			`foo AS f
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").As("f").
				Join(token.NewTable("bar")),
			"foo AS f JOIN bar",
			`foo AS f
JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").As("f").
				InnerJoin(token.NewTable("bar")),
			"foo AS f INNER JOIN bar",
			`foo AS f
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").As("f").
				LeftJoin(token.NewTable("bar")),
			"foo AS f LEFT JOIN bar",
			`foo AS f
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			token.NewTable("foo").As("f").
				RightJoin(token.NewTable("bar")),
			"foo AS f RIGHT JOIN bar",
			`foo AS f
RIGHT JOIN bar
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestJoin(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: token.NewJoin(
				token.NewTable("foo"),
			),
			sql: "JOIN foo",
			sqlIndent: `JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			),
			sql: "JOIN foo AS f",
			sqlIndent: `JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo"),
			).On(
				token.NewColumn("foo.id"),
				token.NewColumn("bar.id"),
			),
			sql: "JOIN foo ON foo.id = bar.id",
			sqlIndent: `JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			),
			sql: "JOIN foo AS f ON f.id = b.id",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).Join(
				token.NewTable("bar"),
			),
			sql: "JOIN foo AS f ON f.id = b.id JOIN bar",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).InnerJoin(
				token.NewTable("bar"),
			),
			sql: "JOIN foo AS f ON f.id = b.id INNER JOIN bar",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).LeftJoin(
				token.NewTable("bar"),
			),
			sql: "JOIN foo AS f ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).RightJoin(
				token.NewTable("bar"),
			),
			sql: "JOIN foo AS f ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
RIGHT JOIN bar
`,
			values: []interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
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
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo"),
			),
			sql: "INNER JOIN foo",
			sqlIndent: `INNER JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			),
			sql: "INNER JOIN foo AS f",
			sqlIndent: `INNER JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo"),
			).On(
				token.NewColumn("foo.id"),
				token.NewColumn("bar.id"),
			),
			sql: "INNER JOIN foo ON foo.id = bar.id",
			sqlIndent: `INNER JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).Join(
				token.NewTable("bar"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id JOIN bar",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).InnerJoin(
				token.NewTable("bar"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id INNER JOIN bar",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).LeftJoin(
				token.NewTable("bar"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewInnerJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).RightJoin(
				token.NewTable("bar"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
RIGHT JOIN bar
`,
			values: []interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
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
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo"),
			),
			sql: "LEFT JOIN foo",
			sqlIndent: `LEFT JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			),
			sql: "LEFT JOIN foo AS f",
			sqlIndent: `LEFT JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo"),
			).On(
				token.NewColumn("foo.id"),
				token.NewColumn("bar.id"),
			),
			sql: "LEFT JOIN foo ON foo.id = bar.id",
			sqlIndent: `LEFT JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).Join(
				token.NewTable("bar"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id JOIN bar",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).InnerJoin(
				token.NewTable("bar"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id INNER JOIN bar",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).LeftJoin(
				token.NewTable("bar"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewLeftJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).RightJoin(
				token.NewTable("bar"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
RIGHT JOIN bar
`,
			values: []interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
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
		statement grammar.Table
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: token.NewRightJoin(
				token.NewTable("foo"),
			),
			sql: "RIGHT JOIN foo",
			sqlIndent: `RIGHT JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			),
			sql: "RIGHT JOIN foo AS f",
			sqlIndent: `RIGHT JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo"),
			).On(
				token.NewColumn("foo.id"),
				token.NewColumn("bar.id"),
			),
			sql: "RIGHT JOIN foo ON foo.id = bar.id",
			sqlIndent: `RIGHT JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).Join(
				token.NewTable("bar"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id JOIN bar",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).InnerJoin(
				token.NewTable("bar"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id INNER JOIN bar",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
INNER JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).LeftJoin(
				token.NewTable("bar"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
LEFT JOIN bar
`,
			values: []interface{}{},
		},
		{
			statement: token.NewRightJoin(
				token.NewTable("foo").As("f"),
			).On(
				token.NewColumn("f.id"),
				token.NewColumn("b.id"),
			).RightJoin(
				token.NewTable("bar"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
RIGHT JOIN bar
`,
			values: []interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

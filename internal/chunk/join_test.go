package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestJoinType(t *testing.T) {
	for _, c := range []interface{}{
		chunk.Join{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Table); !ok {
				t.Errorf("%T doesn't implement grammar.Table", c)
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
			statement: chunk.NewJoin(
				chunk.NewTable("foo"),
			),
			sql: "JOIN foo",
			sqlIndent: `JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewJoin(
				chunk.NewTable("foo").As("f"),
			),
			sql: "JOIN foo AS f",
			sqlIndent: `JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewJoin(
				chunk.NewTable("foo"),
			).On(
				chunk.NewColumn("foo.id"),
				chunk.NewColumn("bar.id"),
			),
			sql: "JOIN foo ON foo.id = bar.id",
			sqlIndent: `JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewJoin(
				chunk.NewTable("foo").As("f"),
			).On(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			),
			sql: "JOIN foo AS f ON f.id = b.id",
			sqlIndent: `JOIN foo AS f ON f.id = b.id
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
			statement: chunk.NewInnerJoin(
				chunk.NewTable("foo"),
			),
			sql: "INNER JOIN foo",
			sqlIndent: `INNER JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewInnerJoin(
				chunk.NewTable("foo").As("f"),
			),
			sql: "INNER JOIN foo AS f",
			sqlIndent: `INNER JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewInnerJoin(
				chunk.NewTable("foo"),
			).On(
				chunk.NewColumn("foo.id"),
				chunk.NewColumn("bar.id"),
			),
			sql: "INNER JOIN foo ON foo.id = bar.id",
			sqlIndent: `INNER JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewInnerJoin(
				chunk.NewTable("foo").As("f"),
			).On(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			),
			sql: "INNER JOIN foo AS f ON f.id = b.id",
			sqlIndent: `INNER JOIN foo AS f ON f.id = b.id
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
			statement: chunk.NewLeftJoin(
				chunk.NewTable("foo"),
			),
			sql: "LEFT JOIN foo",
			sqlIndent: `LEFT JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewLeftJoin(
				chunk.NewTable("foo").As("f"),
			),
			sql: "LEFT JOIN foo AS f",
			sqlIndent: `LEFT JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewLeftJoin(
				chunk.NewTable("foo"),
			).On(
				chunk.NewColumn("foo.id"),
				chunk.NewColumn("bar.id"),
			),
			sql: "LEFT JOIN foo ON foo.id = bar.id",
			sqlIndent: `LEFT JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewLeftJoin(
				chunk.NewTable("foo").As("f"),
			).On(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			),
			sql: "LEFT JOIN foo AS f ON f.id = b.id",
			sqlIndent: `LEFT JOIN foo AS f ON f.id = b.id
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
			statement: chunk.NewRightJoin(
				chunk.NewTable("foo"),
			),
			sql: "RIGHT JOIN foo",
			sqlIndent: `RIGHT JOIN foo
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewRightJoin(
				chunk.NewTable("foo").As("f"),
			),
			sql: "RIGHT JOIN foo AS f",
			sqlIndent: `RIGHT JOIN foo AS f
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewRightJoin(
				chunk.NewTable("foo"),
			).On(
				chunk.NewColumn("foo.id"),
				chunk.NewColumn("bar.id"),
			),
			sql: "RIGHT JOIN foo ON foo.id = bar.id",
			sqlIndent: `RIGHT JOIN foo ON foo.id = bar.id
`,
			values: []interface{}{},
		},
		{
			statement: chunk.NewRightJoin(
				chunk.NewTable("foo").As("f"),
			).On(
				chunk.NewColumn("f.id"),
				chunk.NewColumn("b.id"),
			),
			sql: "RIGHT JOIN foo AS f ON f.id = b.id",
			sqlIndent: `RIGHT JOIN foo AS f ON f.id = b.id
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

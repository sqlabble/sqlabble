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

func TestLogicalCondition(t *testing.T) {
	for _, c := range []interface{}{
		chunk.And{},
		chunk.Or{},
		chunk.Not{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(grammar.LogicalOperation); !ok {
				t.Errorf("%T should implement LogicalOperation", c)
			}
			if _, ok := c.(grammar.Operation); !ok {
				t.Errorf("%T should implement Operation", c)
			}
			if _, ok := c.(grammar.Column); ok {
				t.Errorf("%T shouldn't implement Column", c)
			}
		})
	}
}

func TestConditions(t *testing.T) {
	for _, c := range []interface{}{
		chunk.Eq{},
		chunk.NotEq{},
		chunk.Gt{},
		chunk.Gte{},
		chunk.Lt{},
		chunk.Lte{},
		chunk.Between{},
		chunk.In{},
		chunk.NotIn{},
		chunk.Like{},
		chunk.RegExp{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Operation); !ok {
				t.Errorf("%T should implement Operation", c)
			}
			if _, ok := c.(grammar.LogicalOperation); ok {
				t.Errorf("%T shouldn't implement LogicalOperation", c)
			}
			if _, ok := c.(grammar.Column); ok {
				t.Errorf("%T shouldn't implement Column", c)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement: chunk.NewAnd(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
			),
			sql: "foo != ?",
			sqlIndent: `foo != ?
`,
			values: []interface{}{
				100,
			},
		},
		// 1
		{
			statement: chunk.NewAnd(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
				chunk.NewEq(sqlabble.C("bar"), 200),
			),
			sql: "foo != ? AND bar = ?",
			sqlIndent: `foo != ?
AND bar = ?
`,
			values: []interface{}{
				100,
				200,
			},
		},
		// 2
		{
			statement: chunk.NewAnd(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
				chunk.NewEq(sqlabble.C("bar"), 200),
				chunk.NewLike(sqlabble.C("baz"), "abc"),
			),
			sql: "foo != ? AND bar = ? AND baz LIKE ?",
			sqlIndent: `foo != ?
AND bar = ?
AND baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		// 3
		{
			statement: chunk.NewAnd(
				chunk.NewAnd(
					chunk.NewNotEq(sqlabble.C("foo"), 100),
					chunk.NewEq(sqlabble.C("bar"), 200),
					chunk.NewLike(sqlabble.C("baz"), "abc"),
				),
			),
			sql: "foo != ? AND bar = ? AND baz LIKE ?",
			sqlIndent: `foo != ?
AND bar = ?
AND baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		// 4
		{
			statement: chunk.NewAnd(
				chunk.NewAnd(
					chunk.NewAnd(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
			),
			sql: "foo != ? AND bar = ? AND baz LIKE ?",
			sqlIndent: `foo != ?
AND bar = ?
AND baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		// 5
		{
			statement: chunk.NewAnd(
				chunk.NewAnd(
					chunk.NewAnd(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				chunk.NewGt(sqlabble.C("foo"), 300),
			),
			sql: "(foo != ? AND bar = ? AND baz LIKE ?) AND foo > ?",
			sqlIndent: `(
  foo != ?
  AND bar = ?
  AND baz LIKE ?
)
AND foo > ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
		// 6
		{
			statement: chunk.NewAnd(
				chunk.NewAnd(
					chunk.NewAnd(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
					chunk.NewGt(sqlabble.C("foo"), 300),
				),
			),
			sql: "(foo != ? AND bar = ? AND baz LIKE ?) AND foo > ?",
			sqlIndent: `(
  foo != ?
  AND bar = ?
  AND baz LIKE ?
)
AND foo > ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build %+v", i, c.statement), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("%d BuildIndent %+v", i, c.statement), func(t *testing.T) {
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

func TestOr(t *testing.T) {
	for _, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: chunk.NewOr(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
			),
			sql: "foo != ?",
			sqlIndent: `foo != ?
`,
			values: []interface{}{
				100,
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
				chunk.NewEq(sqlabble.C("bar"), 200),
			),
			sql: "foo != ? OR bar = ?",
			sqlIndent: `foo != ?
OR bar = ?
`,
			values: []interface{}{
				100,
				200,
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
				chunk.NewEq(sqlabble.C("bar"), 200),
				chunk.NewLike(sqlabble.C("baz"), "abc"),
			),
			sql: "foo != ? OR bar = ? OR baz LIKE ?",
			sqlIndent: `foo != ?
OR bar = ?
OR baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewOr(
					chunk.NewNotEq(sqlabble.C("foo"), 100),
					chunk.NewEq(sqlabble.C("bar"), 200),
					chunk.NewLike(sqlabble.C("baz"), "abc"),
				),
			),
			sql: "foo != ? OR bar = ? OR baz LIKE ?",
			sqlIndent: `foo != ?
OR bar = ?
OR baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewOr(
					chunk.NewOr(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
			),
			sql: "foo != ? OR bar = ? OR baz LIKE ?",
			sqlIndent: `foo != ?
OR bar = ?
OR baz LIKE ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewOr(
					chunk.NewOr(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				chunk.NewGt(sqlabble.C("foo"), 300),
			),
			sql: "(foo != ? OR bar = ? OR baz LIKE ?) OR foo > ?",
			sqlIndent: `(
  foo != ?
  OR bar = ?
  OR baz LIKE ?
)
OR foo > ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewOr(
					chunk.NewOr(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
					chunk.NewGt(sqlabble.C("foo"), 300),
				),
			),
			sql: "(foo != ? OR bar = ? OR baz LIKE ?) OR foo > ?",
			sqlIndent: `(
  foo != ?
  OR bar = ?
  OR baz LIKE ?
)
OR foo > ?
`,
			values: []interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
	} {
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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

func TestNot(t *testing.T) {
	for i, c := range []struct {
		statement chunk.Not
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement: chunk.NewNot(
				chunk.NewNotEq(sqlabble.C("foo"), 100),
			),
			sql: "NOT (foo != ?)",
			sqlIndent: `NOT (
  foo != ?
)
`,
			values: []interface{}{
				100,
			},
		},
		// 1
		{
			statement: chunk.NewNot(
				chunk.NewNot(
					chunk.NewNotEq(sqlabble.C("foo"), 100),
				),
			),
			sql: "NOT (NOT (foo != ?))",
			sqlIndent: `NOT (
  NOT (
    foo != ?
  )
)
`,
			values: []interface{}{
				100,
			},
		},
		// 3
		{
			statement: chunk.NewNot(
				chunk.NewNot(
					chunk.NewNot(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
					),
				),
			),
			sql: "NOT (NOT (NOT (foo != ?)))",
			sqlIndent: `NOT (
  NOT (
    NOT (
      foo != ?
    )
  )
)
`,
			values: []interface{}{
				100,
			},
		},
		// 4
		{
			statement: chunk.NewNot(
				chunk.NewOr(
					chunk.NewNotEq(sqlabble.C("foo"), 100),
				),
			),
			sql: "NOT (foo != ?)",
			sqlIndent: `NOT (
  foo != ?
)
`,
			values: []interface{}{
				100,
			},
		},
		// 5
		{
			statement: chunk.NewNot(
				chunk.NewOr(
					chunk.NewNotEq(sqlabble.C("foo"), 100),
					chunk.NewEq(sqlabble.C("bar"), 200),
					chunk.NewLike(sqlabble.C("baz"), "abc"),
				),
			),
			sql: "NOT (foo != ? OR bar = ? OR baz LIKE ?)",
			sqlIndent: `NOT (
  foo != ?
  OR bar = ?
  OR baz LIKE ?
)
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		// 6
		{
			statement: chunk.NewNot(
				chunk.NewNot(
					chunk.NewOr(
						chunk.NewNotEq(sqlabble.C("foo"), 100),
						chunk.NewEq(sqlabble.C("bar"), 200),
						chunk.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
			),
			sql: "NOT (NOT (foo != ? OR bar = ? OR baz LIKE ?))",
			sqlIndent: `NOT (
  NOT (
    foo != ?
    OR bar = ?
    OR baz LIKE ?
  )
)
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
		// 7
		{
			statement: chunk.NewNot(
				chunk.NewNot(
					chunk.NewNot(
						chunk.NewOr(
							chunk.NewNotEq(sqlabble.C("foo"), 100),
							chunk.NewEq(sqlabble.C("bar"), 200),
							chunk.NewLike(sqlabble.C("baz"), "abc"),
						),
					),
				),
			),
			sql: "NOT (NOT (NOT (foo != ? OR bar = ? OR baz LIKE ?)))",
			sqlIndent: `NOT (
  NOT (
    NOT (
      foo != ?
      OR bar = ?
      OR baz LIKE ?
    )
  )
)
`,
			values: []interface{}{
				100,
				200,
				"abc",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build %+v", i, c.statement), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("%d BuildIndent %+v", i, c.statement), func(t *testing.T) {
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

func TestEq(t *testing.T) {
	for _, c := range []struct {
		eq            chunk.Eq
		wantSQL       string
		wantIndentSQL string
		wantValues    []interface{}
	}{
		{
			eq:      chunk.NewEq(chunk.NewColumn("foo"), 100),
			wantSQL: "foo = ?",
			wantIndentSQL: `foo = ?
`,
			wantValues: []interface{}{
				100,
			},
		},
	} {
		t.Run(fmt.Sprintf("Build %+v", c.eq), func(t *testing.T) {
			sql, values := sqlabble.Build(c.eq)
			if sql != c.wantSQL {
				t.Error(diff.SQL(sql, c.wantSQL))
			}
			if !reflect.DeepEqual(values, c.wantValues) {
				t.Error(diff.Values(values, c.wantValues))
			}
		})
		t.Run(fmt.Sprintf("BuildIndent %+v", c.eq), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.eq, "", "  ")
			if sql != c.wantIndentSQL {
				t.Error(diff.SQL(sql, c.wantIndentSQL))
			}
			if !reflect.DeepEqual(values, c.wantValues) {
				t.Error(diff.Values(values, c.wantValues))
			}
		})
	}
}

func TestComplexOperation(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: chunk.NewAnd(
				chunk.NewOr(
					chunk.NewNot(
						chunk.NewAnd(
							chunk.NewOr(
								chunk.NewNot(
									sqlabble.C("foo").Eq(100),
								),
							),
						),
					),
				),
			),
			sql: "NOT (NOT (foo = ?))",
			sqlIndent: `NOT (
  NOT (
    foo = ?
  )
)
`,
			values: []interface{}{
				100,
			},
		},
		{
			statement: chunk.NewOr(
				chunk.NewNot(
					chunk.NewRegExp(sqlabble.C("baz"), "def"),
				),
				chunk.NewOr(
					chunk.NewAnd(
						chunk.NewBetween(sqlabble.C("qux"), 400, 500),
						chunk.NewNot(
							chunk.NewAnd(
								chunk.NewOr(
									chunk.NewNotEq(sqlabble.C("foo"), 100),
									chunk.NewEq(sqlabble.C("bar"), 200),
									chunk.NewLike(sqlabble.C("baz"), "abc"),
									chunk.NewOr(
										chunk.NewIn(sqlabble.C("baz"), "a", "b", "c"),
										chunk.NewNotIn(sqlabble.C("qux"), 600, 700, 800),
									),
								),
							),
						),
						chunk.NewGt(sqlabble.C("foo"), 300),
					),
				),
			),
			sql: "NOT (baz REGEXP ?) OR (qux BETWEEN ? AND ? AND NOT (foo != ? OR bar = ? OR baz LIKE ? OR (baz IN (?, ?, ?) OR qux NOT IN (?, ?, ?))) AND foo > ?)",
			sqlIndent: `NOT (
  baz REGEXP ?
)
OR (
  qux BETWEEN ? AND ?
  AND NOT (
    foo != ?
    OR bar = ?
    OR baz LIKE ?
    OR (
      baz IN (?, ?, ?)
      OR qux NOT IN (?, ?, ?)
    )
  )
  AND foo > ?
)
`,
			values: []interface{}{
				"def",
				400,
				500,
				100,
				200,
				"abc",
				"a",
				"b",
				"c",
				600,
				700,
				800,
				300,
			},
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

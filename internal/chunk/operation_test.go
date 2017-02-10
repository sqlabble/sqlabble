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

func TestLogicalCondition(t *testing.T) {
	for _, c := range []interface{}{
		token.And{},
		token.Or{},
		token.Not{},
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
		token.Eq{},
		token.NotEq{},
		token.Gt{},
		token.Gte{},
		token.Lt{},
		token.Lte{},
		token.Between{},
		token.In{},
		token.NotIn{},
		token.Like{},
		token.RegExp{},
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
			statement: token.NewAnd(
				token.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: token.NewAnd(
				token.NewNotEq(sqlabble.C("foo"), 100),
				token.NewEq(sqlabble.C("bar"), 200),
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
			statement: token.NewAnd(
				token.NewNotEq(sqlabble.C("foo"), 100),
				token.NewEq(sqlabble.C("bar"), 200),
				token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewAnd(
				token.NewAnd(
					token.NewNotEq(sqlabble.C("foo"), 100),
					token.NewEq(sqlabble.C("bar"), 200),
					token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewAnd(
				token.NewAnd(
					token.NewAnd(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewAnd(
				token.NewAnd(
					token.NewAnd(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				token.NewGt(sqlabble.C("foo"), 300),
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
			statement: token.NewAnd(
				token.NewAnd(
					token.NewAnd(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
					),
					token.NewGt(sqlabble.C("foo"), 300),
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
			statement: token.NewOr(
				token.NewNotEq(sqlabble.C("foo"), 100),
			),
			sql: "foo != ?",
			sqlIndent: `foo != ?
`,
			values: []interface{}{
				100,
			},
		},
		{
			statement: token.NewOr(
				token.NewNotEq(sqlabble.C("foo"), 100),
				token.NewEq(sqlabble.C("bar"), 200),
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
			statement: token.NewOr(
				token.NewNotEq(sqlabble.C("foo"), 100),
				token.NewEq(sqlabble.C("bar"), 200),
				token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewOr(
				token.NewOr(
					token.NewNotEq(sqlabble.C("foo"), 100),
					token.NewEq(sqlabble.C("bar"), 200),
					token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewOr(
				token.NewOr(
					token.NewOr(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewOr(
				token.NewOr(
					token.NewOr(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				token.NewGt(sqlabble.C("foo"), 300),
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
			statement: token.NewOr(
				token.NewOr(
					token.NewOr(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
					),
					token.NewGt(sqlabble.C("foo"), 300),
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
		statement token.Not
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement: token.NewNot(
				token.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: token.NewNot(
				token.NewNot(
					token.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: token.NewNot(
				token.NewNot(
					token.NewNot(
						token.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: token.NewNot(
				token.NewOr(
					token.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: token.NewNot(
				token.NewOr(
					token.NewNotEq(sqlabble.C("foo"), 100),
					token.NewEq(sqlabble.C("bar"), 200),
					token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewNot(
				token.NewNot(
					token.NewOr(
						token.NewNotEq(sqlabble.C("foo"), 100),
						token.NewEq(sqlabble.C("bar"), 200),
						token.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: token.NewNot(
				token.NewNot(
					token.NewNot(
						token.NewOr(
							token.NewNotEq(sqlabble.C("foo"), 100),
							token.NewEq(sqlabble.C("bar"), 200),
							token.NewLike(sqlabble.C("baz"), "abc"),
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
		eq            token.Eq
		wantSQL       string
		wantIndentSQL string
		wantValues    []interface{}
	}{
		{
			eq:      token.NewEq(token.NewColumn("foo"), 100),
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
			statement: token.NewAnd(
				token.NewOr(
					token.NewNot(
						token.NewAnd(
							token.NewOr(
								token.NewNot(
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
			statement: token.NewOr(
				token.NewNot(
					token.NewRegExp(sqlabble.C("baz"), "def"),
				),
				token.NewOr(
					token.NewAnd(
						token.NewBetween(sqlabble.C("qux"), 400, 500),
						token.NewNot(
							token.NewAnd(
								token.NewOr(
									token.NewNotEq(sqlabble.C("foo"), 100),
									token.NewEq(sqlabble.C("bar"), 200),
									token.NewLike(sqlabble.C("baz"), "abc"),
									token.NewOr(
										token.NewIn(sqlabble.C("baz"), "a", "b", "c"),
										token.NewNotIn(sqlabble.C("qux"), 600, 700, 800),
									),
								),
							),
						),
						token.NewGt(sqlabble.C("foo"), 300),
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

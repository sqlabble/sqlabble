package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestLogicalCondition(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.And{},
// 		sqlabble.Or{},
// 		sqlabble.Not{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.LogicalOperationNode); !ok {
// 				t.Errorf("%T should implement LogicalOperationNode", c)
// 			}
// 			if _, ok := c.(sqlabble.OperationNode); !ok {
// 				t.Errorf("%T should implement Operation", c)
// 			}
// 			if _, ok := c.(sqlabble.ColumnNode); ok {
// 				t.Errorf("%T shouldn't implement Column", c)
// 			}
// 		})
// 	}
// }
//
// func TestConditions(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Eq{},
// 		sqlabble.NotEq{},
// 		sqlabble.Gt{},
// 		sqlabble.Gte{},
// 		sqlabble.Lt{},
// 		sqlabble.Lte{},
// 		sqlabble.Between{},
// 		sqlabble.In{},
// 		sqlabble.NotIn{},
// 		sqlabble.Like{},
// 		sqlabble.RegExp{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.OperationNode); !ok {
// 				t.Errorf("%T should implement Operation", c)
// 			}
// 			if _, ok := c.(sqlabble.LogicalOperationNode); ok {
// 				t.Errorf("%T shouldn't implement LogicalOperation", c)
// 			}
// 			if _, ok := c.(sqlabble.ColumnNode); ok {
// 				t.Errorf("%T shouldn't implement Column", c)
// 			}
// 		})
// 	}
// }

func TestAnd(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement: sqlabble.NewAnd(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
				sqlabble.NewEq(sqlabble.C("bar"), 200),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
				sqlabble.NewEq(sqlabble.C("bar"), 200),
				sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewNotEq(sqlabble.C("foo"), 100),
					sqlabble.NewEq(sqlabble.C("bar"), 200),
					sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				sqlabble.NewGt(sqlabble.C("foo"), 300),
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
			statement: sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
					),
					sqlabble.NewGt(sqlabble.C("foo"), 300),
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: sqlabble.NewOr(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
			),
			sql: "foo != ?",
			sqlIndent: `foo != ?
`,
			values: []interface{}{
				100,
			},
		},
		{
			statement: sqlabble.NewOr(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
				sqlabble.NewEq(sqlabble.C("bar"), 200),
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
			statement: sqlabble.NewOr(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
				sqlabble.NewEq(sqlabble.C("bar"), 200),
				sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewNotEq(sqlabble.C("foo"), 100),
					sqlabble.NewEq(sqlabble.C("bar"), 200),
					sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
					),
				),
				sqlabble.NewGt(sqlabble.C("foo"), 300),
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
			statement: sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
					),
					sqlabble.NewGt(sqlabble.C("foo"), 300),
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement: sqlabble.NewNot(
				sqlabble.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewNot(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: sqlabble.NewNot(
				sqlabble.NewOr(
					sqlabble.NewNotEq(sqlabble.C("foo"), 100),
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
			statement: sqlabble.NewNot(
				sqlabble.NewOr(
					sqlabble.NewNotEq(sqlabble.C("foo"), 100),
					sqlabble.NewEq(sqlabble.C("bar"), 200),
					sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewOr(
						sqlabble.NewNotEq(sqlabble.C("foo"), 100),
						sqlabble.NewEq(sqlabble.C("bar"), 200),
						sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
			statement: sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewNot(
						sqlabble.NewOr(
							sqlabble.NewNotEq(sqlabble.C("foo"), 100),
							sqlabble.NewEq(sqlabble.C("bar"), 200),
							sqlabble.NewLike(sqlabble.C("baz"), "abc"),
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
		eq            sqlabble.Statement
		wantSQL       string
		wantIndentSQL string
		wantValues    []interface{}
	}{
		{
			eq:      sqlabble.NewEq(sqlabble.NewColumn("foo"), 100),
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: sqlabble.NewAnd(
				sqlabble.NewOr(
					sqlabble.NewNot(
						sqlabble.NewAnd(
							sqlabble.NewOr(
								sqlabble.NewNot(
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
			statement: sqlabble.NewOr(
				sqlabble.NewNot(
					sqlabble.NewRegExp(sqlabble.C("baz"), "def"),
				),
				sqlabble.NewOr(
					sqlabble.NewAnd(
						sqlabble.NewBetween(sqlabble.C("qux"), 400, 500),
						sqlabble.NewNot(
							sqlabble.NewAnd(
								sqlabble.NewOr(
									sqlabble.NewNotEq(sqlabble.C("foo"), 100),
									sqlabble.NewEq(sqlabble.C("bar"), 200),
									sqlabble.NewLike(sqlabble.C("baz"), "abc"),
									sqlabble.NewOr(
										sqlabble.NewIn(sqlabble.C("baz"), "a", "b", "c"),
										sqlabble.NewNotIn(sqlabble.C("qux"), 600, 700, 800),
									),
								),
							),
						),
						sqlabble.NewGt(sqlabble.C("foo"), 300),
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

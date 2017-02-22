package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

// func TestLogicalCondition(t *testing.T) {
// 	for _, c := range []interface{}{
// 		statement.And{},
// 		statement.Or{},
// 		statement.Not{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(statement.LogicalOperationNode); !ok {
// 				t.Errorf("%T should implement LogicalOperationNode", c)
// 			}
// 			if _, ok := c.(statement.OperationNode); !ok {
// 				t.Errorf("%T should implement Operation", c)
// 			}
// 			if _, ok := c.(statement.ColumnNode); ok {
// 				t.Errorf("%T shouldn't implement Column", c)
// 			}
// 		})
// 	}
// }
//
// func TestConditions(t *testing.T) {
// 	for _, c := range []interface{}{
// 		statement.Eq{},
// 		statement.NotEq{},
// 		statement.Gt{},
// 		statement.Gte{},
// 		statement.Lt{},
// 		statement.Lte{},
// 		statement.Between{},
// 		statement.In{},
// 		statement.NotIn{},
// 		statement.Like{},
// 		statement.RegExp{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(statement.OperationNode); !ok {
// 				t.Errorf("%T should implement Operation", c)
// 			}
// 			if _, ok := c.(statement.LogicalOperationNode); ok {
// 				t.Errorf("%T shouldn't implement LogicalOperation", c)
// 			}
// 			if _, ok := c.(statement.ColumnNode); ok {
// 				t.Errorf("%T shouldn't implement Column", c)
// 			}
// 		})
// 	}
// }

func TestAnd(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement.NewAnd(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
			),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		// 1
		{
			statement.NewAnd(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				statement.NewColumn("bar").Eq(statement.NewParam(200)),
			),
			"foo != ? AND bar = ?",
			`> foo != ?
> AND bar = ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		// 2
		{
			statement.NewAnd(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				statement.NewColumn("bar").Eq(statement.NewParam(200)),
				statement.NewColumn("baz").Like(statement.NewParam("abc")),
			),
			"foo != ? AND bar = ? AND baz LIKE ?",
			`> foo != ?
> AND bar = ?
> AND baz LIKE ?
`,
			[]interface{}{
				100,
				200,
				"abc",
			},
		},
		// 3
		{
			statement.NewAnd(
				statement.NewAnd(
					statement.NewColumn("foo").NotEq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam(200)),
					statement.NewColumn("baz").Like(statement.NewParam("abc")),
				),
				statement.NewColumn("foo").Between(
					statement.NewParam(300),
					statement.NewParam(400),
				),
			),
			"(foo != ? AND bar = ? AND baz LIKE ?) AND foo BETWEEN ? AND ?",
			`> (
>   foo != ?
>   AND bar = ?
>   AND baz LIKE ?
> )
> AND foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
				"abc",
				300,
				400,
			},
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

func TestOr(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewOr(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
			),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewOr(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				statement.NewColumn("bar").Eq(statement.NewParam(200)),
			),
			"foo != ? OR bar = ?",
			`> foo != ?
> OR bar = ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewOr(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				statement.NewColumn("bar").Eq(statement.NewParam(200)),
				statement.NewColumn("baz").Like(statement.NewParam("abc")),
			),
			"foo != ? OR bar = ? OR baz LIKE ?",
			`> foo != ?
> OR bar = ?
> OR baz LIKE ?
`,
			[]interface{}{
				100,
				200,
				"abc",
			},
		},
		{
			statement.NewOr(
				statement.NewOr(
					statement.NewColumn("foo").NotEq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam(200)),
					statement.NewColumn("baz").Like(statement.NewParam("abc")),
				),
				statement.NewColumn("foo").Gt(statement.NewParam(300)),
			),
			"(foo != ? OR bar = ? OR baz LIKE ?) OR foo > ?",
			`> (
>   foo != ?
>   OR bar = ?
>   OR baz LIKE ?
> )
> OR foo > ?
`,
			[]interface{}{
				100,
				200,
				"abc",
				300,
			},
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

func TestNot(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			statement.NewNot(
				statement.NewColumn("foo").NotEq(statement.NewParam(100)),
			),
			"NOT (foo != ?)",
			`> NOT (
>   foo != ?
> )
`,
			[]interface{}{
				100,
			},
		},
		// 1
		{
			statement.NewNot(
				statement.NewNot(
					statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				),
			),
			"NOT (NOT (foo != ?))",
			`> NOT (
>   NOT (
>     foo != ?
>   )
> )
`,
			[]interface{}{
				100,
			},
		},
		// 3
		{
			statement.NewNot(
				statement.NewNot(
					statement.NewNot(
						statement.NewColumn("foo").NotEq(statement.NewParam(100)),
					),
				),
			),
			"NOT (NOT (NOT (foo != ?)))",
			`> NOT (
>   NOT (
>     NOT (
>       foo != ?
>     )
>   )
> )
`,
			[]interface{}{
				100,
			},
		},
		// 4
		{
			statement.NewNot(
				statement.NewOr(
					statement.NewColumn("foo").NotEq(statement.NewParam(100)),
				),
			),
			"NOT (foo != ?)",
			`> NOT (
>   foo != ?
> )
`,
			[]interface{}{
				100,
			},
		},
		// 5
		{
			statement.NewNot(
				statement.NewOr(
					statement.NewColumn("foo").NotEq(statement.NewParam(100)),
					statement.NewColumn("bar").Eq(statement.NewParam(200)),
					statement.NewColumn("baz").Like(statement.NewParam("abc")),
				),
			),
			"NOT (foo != ? OR bar = ? OR baz LIKE ?)",
			`> NOT (
>   foo != ?
>   OR bar = ?
>   OR baz LIKE ?
> )
`,
			[]interface{}{
				100,
				200,
				"abc",
			},
		},
		// 6
		{
			statement.NewNot(
				statement.NewNot(
					statement.NewOr(
						statement.NewColumn("foo").NotEq(statement.NewParam(100)),
						statement.NewColumn("bar").Eq(statement.NewParam(200)),
						statement.NewColumn("baz").Like(statement.NewParam("abc")),
					),
				),
			),
			"NOT (NOT (foo != ? OR bar = ? OR baz LIKE ?))",
			`> NOT (
>   NOT (
>     foo != ?
>     OR bar = ?
>     OR baz LIKE ?
>   )
> )
`,
			[]interface{}{
				100,
				200,
				"abc",
			},
		},
		// 7
		{
			statement.NewNot(
				statement.NewNot(
					statement.NewNot(
						statement.NewOr(
							statement.NewColumn("foo").NotEq(statement.NewParam(100)),
							statement.NewColumn("bar").Eq(statement.NewParam(200)),
							statement.NewColumn("baz").Like(statement.NewParam("abc")),
						),
					),
				),
			),
			"NOT (NOT (NOT (foo != ? OR bar = ? OR baz LIKE ?)))",
			`> NOT (
>   NOT (
>     NOT (
>       foo != ?
>       OR bar = ?
>       OR baz LIKE ?
>     )
>   )
> )
`,
			[]interface{}{
				100,
				200,
				"abc",
			},
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

func TestComparisonOperators(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Eq(statement.NewParam(100)),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lt(statement.NewParam(100)),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lte(statement.NewParam(100)),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gt(statement.NewParam(100)),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gte(statement.NewParam(100)),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Like(statement.NewParam("bar")),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewColumn("foo").RegExp(statement.NewParam("bar")),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"bar",
			},
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

func TestBetween(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Between(
				statement.NewParam(100),
				statement.NewParam(200),
			),
			"foo BETWEEN ? AND ?",
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
		{
			statement.NewColumn("joined_users").Between(
				statement.NewParam(100),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("count(*)")).
						From(statement.NewTable("users")),
				),
			),
			"joined_users BETWEEN ? AND (SELECT count(*) FROM users)",
			`> joined_users BETWEEN ? AND (
>   SELECT
>     count(*)
>   FROM
>     users
> )
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("joined_users").Between(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("count(*)")).
						From(statement.NewTable("users")),
				),
				statement.NewParam(500),
			),
			"joined_users BETWEEN (SELECT count(*) FROM users) AND ?",
			`> joined_users BETWEEN (
>   SELECT
>     count(*)
>   FROM
>     users
> ) AND ?
`,
			[]interface{}{
				500,
			},
		},
		{
			statement.NewColumn("joined_users").Between(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("count(*)")).
						From(statement.NewTable("super_users")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("count(*)")).
						From(statement.NewTable("users")),
				),
			),
			"joined_users BETWEEN (SELECT count(*) FROM super_users) AND (SELECT count(*) FROM users)",
			`> joined_users BETWEEN (
>   SELECT
>     count(*)
>   FROM
>     super_users
> ) AND (
>   SELECT
>     count(*)
>   FROM
>     users
> )
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

func TestContainingOperators(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").
				In(statement.NewParams(
					100,
					200,
					300,
				)),
			"foo IN (?, ?, ?)",
			`> foo IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
		},
		{
			statement.NewColumn("foo").
				NotIn(statement.NewParams(
					100,
					200,
					300,
				)),
			"foo NOT IN (?, ?, ?)",
			`> foo NOT IN (?, ?, ?)
`,
			[]interface{}{
				100,
				200,
				300,
			},
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

func TestKeywordOperators(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").IsNull(),
			"foo IS NULL",
			`> foo IS NULL
`,
			nil,
		},
		{
			statement.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
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

func TestComplexOperation(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewAnd(
				statement.NewOr(
					statement.NewNot(
						statement.NewAnd(
							statement.NewOr(
								statement.NewNot(
									statement.NewColumn("foo").Eq(statement.NewParam(100)),
								),
							),
						),
					),
				),
			),
			"(NOT ((NOT (foo = ?))))",
			`> (
>   NOT (
>     (
>       NOT (
>         foo = ?
>       )
>     )
>   )
> )
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewOr(
				statement.NewNot(
					statement.NewColumn("baz").RegExp(statement.NewParam("def")),
				),
				statement.NewOr(
					statement.NewAnd(
						statement.NewColumn("qux").Between(
							statement.NewParam(400),
							statement.NewParam(500),
						),
						statement.NewNot(
							statement.NewAnd(
								statement.NewOr(
									statement.NewColumn("foo").NotEq(statement.NewParam(100)),
									statement.NewColumn("bar").Eq(statement.NewParam(200)),
									statement.NewColumn("baz").Like(statement.NewParam("abc")),
									statement.NewOr(
										statement.NewColumn("baz").
											In(statement.NewParams(
												"a", "b", "c",
											)),
										statement.NewColumn("qux").
											NotIn(statement.NewParams(
												600, 700, 800,
											)),
									),
								),
							),
						),
						statement.NewColumn("foo").Gt(statement.NewParam(300)),
					),
				),
			),
			"NOT (baz REGEXP ?) OR ((qux BETWEEN ? AND ? AND NOT ((foo != ? OR bar = ? OR baz LIKE ? OR (baz IN (?, ?, ?) OR qux NOT IN (?, ?, ?)))) AND foo > ?))",
			`> NOT (
>   baz REGEXP ?
> )
> OR (
>   (
>     qux BETWEEN ? AND ?
>     AND NOT (
>       (
>         foo != ?
>         OR bar = ?
>         OR baz LIKE ?
>         OR (
>           baz IN (?, ?, ?)
>           OR qux NOT IN (?, ?, ?)
>         )
>       )
>     )
>     AND foo > ?
>   )
> )
`,
			[]interface{}{
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

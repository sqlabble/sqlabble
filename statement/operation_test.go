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
				statement.NewColumn("foo").NotEq(100),
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
				statement.NewColumn("foo").NotEq(100),
				statement.NewColumn("bar").Eq(200),
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
				statement.NewColumn("foo").NotEq(100),
				statement.NewColumn("bar").Eq(200),
				statement.NewColumn("baz").Like("abc"),
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
					statement.NewColumn("foo").NotEq(100),
					statement.NewColumn("bar").Eq(200),
					statement.NewColumn("baz").Like("abc"),
				),
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
		// 4
		{
			statement.NewAnd(
				statement.NewAnd(
					statement.NewAnd(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
				),
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
		// 5
		{
			statement.NewAnd(
				statement.NewAnd(
					statement.NewAnd(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
				),
				statement.NewColumn("foo").Gt(300),
			),
			"(foo != ? AND bar = ? AND baz LIKE ?) AND foo > ?",
			`> (
>   foo != ?
>   AND bar = ?
>   AND baz LIKE ?
> )
> AND foo > ?
`,
			[]interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
		// 6
		{
			statement.NewAnd(
				statement.NewAnd(
					statement.NewAnd(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
					statement.NewColumn("foo").Gt(300),
				),
			),
			"(foo != ? AND bar = ? AND baz LIKE ?) AND foo > ?",
			`> (
>   foo != ?
>   AND bar = ?
>   AND baz LIKE ?
> )
> AND foo > ?
`,
			[]interface{}{
				100,
				200,
				"abc",
				300,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build %+v", i, c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("%d BuildIndent %+v", i, c.statement), func(t *testing.T) {
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

func TestOr(t *testing.T) {
	for _, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewOr(
				statement.NewColumn("foo").NotEq(100),
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
				statement.NewColumn("foo").NotEq(100),
				statement.NewColumn("bar").Eq(200),
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
				statement.NewColumn("foo").NotEq(100),
				statement.NewColumn("bar").Eq(200),
				statement.NewColumn("baz").Like("abc"),
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
					statement.NewColumn("foo").NotEq(100),
					statement.NewColumn("bar").Eq(200),
					statement.NewColumn("baz").Like("abc"),
				),
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
					statement.NewOr(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
				),
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
					statement.NewOr(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
				),
				statement.NewColumn("foo").Gt(300),
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
		{
			statement.NewOr(
				statement.NewOr(
					statement.NewOr(
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
					),
					statement.NewColumn("foo").Gt(300),
				),
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
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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
				statement.NewColumn("foo").NotEq(100),
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
					statement.NewColumn("foo").NotEq(100),
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
						statement.NewColumn("foo").NotEq(100),
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
					statement.NewColumn("foo").NotEq(100),
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
					statement.NewColumn("foo").NotEq(100),
					statement.NewColumn("bar").Eq(200),
					statement.NewColumn("baz").Like("abc"),
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
						statement.NewColumn("foo").NotEq(100),
						statement.NewColumn("bar").Eq(200),
						statement.NewColumn("baz").Like("abc"),
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
							statement.NewColumn("foo").NotEq(100),
							statement.NewColumn("bar").Eq(200),
							statement.NewColumn("baz").Like("abc"),
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
		t.Run(fmt.Sprintf("%d Build %+v", i, c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})

		t.Run(fmt.Sprintf("%d BuildIndent %+v", i, c.statement), func(t *testing.T) {
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

func TestComparisonOperators(t *testing.T) {
	for _, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Eq(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lt(100),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Lte(100),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gt(100),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Gte(100),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			statement.NewColumn("foo").Like("bar"),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			statement.NewColumn("foo").RegExp("bar"),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
	} {
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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

func TestBetween(t *testing.T) {
	for _, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").Between(
				100,
				200,
			),
			"foo BETWEEN ? AND ?",
			`> foo BETWEEN ? AND ?
`,
			[]interface{}{
				100,
				200,
			},
		},
	} {
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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

func TestContainingOperators(t *testing.T) {
	for _, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewColumn("foo").In(
				100,
				200,
				300,
			),
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
			statement.NewColumn("foo").NotIn(
				100,
				200,
				300,
			),
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
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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

func TestKeywordOperators(t *testing.T) {
	for _, c := range []struct {
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
			[]interface{}{},
		},
		{
			statement.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("Build %+v", c.statement), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("BuildIndent %+v", c.statement), func(t *testing.T) {
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
									statement.NewColumn("foo").Eq(100),
								),
							),
						),
					),
				),
			),
			"NOT (NOT (foo = ?))",
			`> NOT (
>   NOT (
>     foo = ?
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
					statement.NewColumn("baz").RegExp("def"),
				),
				statement.NewOr(
					statement.NewAnd(
						statement.NewColumn("qux").Between(400, 500),
						statement.NewNot(
							statement.NewAnd(
								statement.NewOr(
									statement.NewColumn("foo").NotEq(100),
									statement.NewColumn("bar").Eq(200),
									statement.NewColumn("baz").Like("abc"),
									statement.NewOr(
										statement.NewColumn("baz").In("a", "b", "c"),
										statement.NewColumn("qux").NotIn(600, 700, 800),
									),
								),
							),
						),
						statement.NewColumn("foo").Gt(300),
					),
				),
			),
			"NOT (baz REGEXP ?) OR (qux BETWEEN ? AND ? AND NOT (foo != ? OR bar = ? OR baz LIKE ? OR (baz IN (?, ?, ?) OR qux NOT IN (?, ?, ?))) AND foo > ?)",
			`> NOT (
>   baz REGEXP ?
> )
> OR (
>   qux BETWEEN ? AND ?
>   AND NOT (
>     foo != ?
>     OR bar = ?
>     OR baz LIKE ?
>     OR (
>       baz IN (?, ?, ?)
>       OR qux NOT IN (?, ?, ?)
>     )
>   )
>   AND foo > ?
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
			sql, values := builder.Build(c.statement)
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

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
			sqlabble.NewAnd(
				sqlabble.C("foo").NotEq(100),
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
			sqlabble.NewAnd(
				sqlabble.C("foo").NotEq(100),
				sqlabble.C("bar").Eq(200),
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
			sqlabble.NewAnd(
				sqlabble.C("foo").NotEq(100),
				sqlabble.C("bar").Eq(200),
				sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.C("foo").NotEq(100),
					sqlabble.C("bar").Eq(200),
					sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
					),
				),
				sqlabble.C("foo").Gt(300),
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
			sqlabble.NewAnd(
				sqlabble.NewAnd(
					sqlabble.NewAnd(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
					),
					sqlabble.C("foo").Gt(300),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewOr(
				sqlabble.C("foo").NotEq(100),
			),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.NewOr(
				sqlabble.C("foo").NotEq(100),
				sqlabble.C("bar").Eq(200),
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
			sqlabble.NewOr(
				sqlabble.C("foo").NotEq(100),
				sqlabble.C("bar").Eq(200),
				sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.C("foo").NotEq(100),
					sqlabble.C("bar").Eq(200),
					sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
					),
				),
				sqlabble.C("foo").Gt(300),
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
			sqlabble.NewOr(
				sqlabble.NewOr(
					sqlabble.NewOr(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
					),
					sqlabble.C("foo").Gt(300),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			sqlabble.NewNot(
				sqlabble.C("foo").NotEq(100),
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
			sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.C("foo").NotEq(100),
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
			sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewNot(
						sqlabble.C("foo").NotEq(100),
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
			sqlabble.NewNot(
				sqlabble.NewOr(
					sqlabble.C("foo").NotEq(100),
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
			sqlabble.NewNot(
				sqlabble.NewOr(
					sqlabble.C("foo").NotEq(100),
					sqlabble.C("bar").Eq(200),
					sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewOr(
						sqlabble.C("foo").NotEq(100),
						sqlabble.C("bar").Eq(200),
						sqlabble.C("baz").Like("abc"),
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
			sqlabble.NewNot(
				sqlabble.NewNot(
					sqlabble.NewNot(
						sqlabble.NewOr(
							sqlabble.C("foo").NotEq(100),
							sqlabble.C("bar").Eq(200),
							sqlabble.C("baz").Like("abc"),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.C("foo").Eq(100),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.C("foo").Lt(100),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.C("foo").Lte(100),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.C("foo").Gt(100),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.C("foo").Gte(100),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			sqlabble.C("foo").Like("bar"),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			sqlabble.C("foo").RegExp("bar"),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"bar",
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewBetween(
				sqlabble.NewColumn("foo"),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewIn(
				sqlabble.NewColumn("foo"),
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
			sqlabble.NewNotIn(
				sqlabble.NewColumn("foo"),
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
			sql, values := sqlabble.Build(c.statement)
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewIsNull(
				sqlabble.NewColumn("foo"),
			),
			"foo IS NULL",
			`> foo IS NULL
`,
			[]interface{}{},
		},
		{
			sqlabble.NewIsNotNull(
				sqlabble.NewColumn("foo"),
			),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			[]interface{}{},
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
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewAnd(
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
			sqlabble.NewOr(
				sqlabble.NewNot(
					sqlabble.C("baz").RegExp("def"),
				),
				sqlabble.NewOr(
					sqlabble.NewAnd(
						sqlabble.C("qux").Between(400, 500),
						sqlabble.NewNot(
							sqlabble.NewAnd(
								sqlabble.NewOr(
									sqlabble.C("foo").NotEq(100),
									sqlabble.C("bar").Eq(200),
									sqlabble.C("baz").Like("abc"),
									sqlabble.NewOr(
										sqlabble.C("baz").In("a", "b", "c"),
										sqlabble.C("qux").NotIn(600, 700, 800),
									),
								),
							),
						),
						sqlabble.C("foo").Gt(300),
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
			sql, values := sqlabble.Build(c.statement)
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

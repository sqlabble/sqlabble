package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestAnd(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			stmt.NewAnd(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
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
			stmt.NewAnd(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
				stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
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
			stmt.NewAnd(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
				stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
				stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
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
			stmt.NewAnd(
				stmt.NewAnd(
					stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
					stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
					stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
				),
				stmt.NewColumn("foo").Between(
					stmt.NewVal(300),
					stmt.NewVal(400),
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
		c := c
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

func TestOr(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewOr(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
			),
			"foo != ?",
			`> foo != ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewOr(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
				stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
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
			stmt.NewOr(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
				stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
				stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
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
			stmt.NewOr(
				stmt.NewOr(
					stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
					stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
					stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
				),
				stmt.NewColumn("foo").Gt(stmt.NewVal(300)),
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
		c := c
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

func TestNot(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		// 0
		{
			stmt.NewNot(
				stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
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
			stmt.NewNot(
				stmt.NewNot(
					stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
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
			stmt.NewNot(
				stmt.NewNot(
					stmt.NewNot(
						stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
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
			stmt.NewNot(
				stmt.NewOr(
					stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
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
			stmt.NewNot(
				stmt.NewOr(
					stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
					stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
					stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
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
			stmt.NewNot(
				stmt.NewNot(
					stmt.NewOr(
						stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
						stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
						stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
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
			stmt.NewNot(
				stmt.NewNot(
					stmt.NewNot(
						stmt.NewOr(
							stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
							stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
							stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
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
		c := c
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

func TestComparisonOperators(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo").Eq(stmt.NewVal(100)),
			"foo = ?",
			`> foo = ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Lt(stmt.NewVal(100)),
			"foo < ?",
			`> foo < ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Lte(stmt.NewVal(100)),
			"foo <= ?",
			`> foo <= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Gt(stmt.NewVal(100)),
			"foo > ?",
			`> foo > ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Gte(stmt.NewVal(100)),
			"foo >= ?",
			`> foo >= ?
`,
			[]interface{}{
				100,
			},
		},
		{
			stmt.NewColumn("foo").Like(stmt.NewVal("bar")),
			"foo LIKE ?",
			`> foo LIKE ?
`,
			[]interface{}{
				"bar",
			},
		},
		{
			stmt.NewColumn("foo").RegExp(stmt.NewVal("bar")),
			"foo REGEXP ?",
			`> foo REGEXP ?
`,
			[]interface{}{
				"bar",
			},
		},
	} {
		c := c
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

func TestBetween(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo").Between(
				stmt.NewVal(100),
				stmt.NewVal(200),
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
			stmt.NewColumn("joined_users").Between(
				stmt.NewVal(100),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("count(*)")).
						From(stmt.NewTable("users")),
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
			stmt.NewColumn("joined_users").Between(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("count(*)")).
						From(stmt.NewTable("users")),
				),
				stmt.NewVal(500),
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
			stmt.NewColumn("joined_users").Between(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("count(*)")).
						From(stmt.NewTable("super_users")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("count(*)")).
						From(stmt.NewTable("users")),
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
		c := c
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

func TestContainingOperators(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo").
				In(stmt.NewVals(
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
			stmt.NewColumn("foo").
				NotIn(stmt.NewVals(
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
		c := c
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

func TestKeywordOperators(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumn("foo").IsNull(),
			"foo IS NULL",
			`> foo IS NULL
`,
			nil,
		},
		{
			stmt.NewColumn("foo").IsNotNull(),
			"foo IS NOT NULL",
			`> foo IS NOT NULL
`,
			nil,
		},
	} {
		c := c
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

func TestComplexOperation(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewAnd(
				stmt.NewOr(
					stmt.NewNot(
						stmt.NewAnd(
							stmt.NewOr(
								stmt.NewNot(
									stmt.NewColumn("foo").Eq(stmt.NewVal(100)),
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
			stmt.NewOr(
				stmt.NewNot(
					stmt.NewColumn("baz").RegExp(stmt.NewVal("def")),
				),
				stmt.NewOr(
					stmt.NewAnd(
						stmt.NewColumn("qux").Between(
							stmt.NewVal(400),
							stmt.NewVal(500),
						),
						stmt.NewNot(
							stmt.NewAnd(
								stmt.NewOr(
									stmt.NewColumn("foo").NotEq(stmt.NewVal(100)),
									stmt.NewColumn("bar").Eq(stmt.NewVal(200)),
									stmt.NewColumn("baz").Like(stmt.NewVal("abc")),
									stmt.NewOr(
										stmt.NewColumn("baz").
											In(stmt.NewVals(
												"a", "b", "c",
											)),
										stmt.NewColumn("qux").
											NotIn(stmt.NewVals(
												600, 700, 800,
											)),
									),
								),
							),
						),
						stmt.NewColumn("foo").Gt(stmt.NewVal(300)),
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
		c := c
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

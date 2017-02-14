package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestJoinType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Join{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.Node); !ok {
// 				t.Errorf("%T doesn't implement sqlabble.Table", c)
// 			}
// 		})
// 	}
// }

func TestJoin(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo"),
			),
			"JOIN foo",
			`> JOIN foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo").As("f"),
			),
			"JOIN foo AS f",
			`> JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo"),
			).On(
				sqlabble.NewColumn("foo.id"),
				sqlabble.NewColumn("bar.id"),
			),
			"JOIN foo ON foo.id = bar.id",
			`> JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo").As("f"),
			).On(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			),
			"JOIN foo AS f ON f.id = b.id",
			`> JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"JOIN foo USING id",
			`> JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewJoin(
				sqlabble.NewTable("foo").As("f"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"JOIN foo AS f USING id",
			`> JOIN foo AS f USING id
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

func TestInnerJoin(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo"),
			),
			"INNER JOIN foo",
			`> INNER JOIN foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo").As("f"),
			),
			"INNER JOIN foo AS f",
			`> INNER JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo"),
			).On(
				sqlabble.NewColumn("foo.id"),
				sqlabble.NewColumn("bar.id"),
			),
			"INNER JOIN foo ON foo.id = bar.id",
			`> INNER JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo").As("f"),
			).On(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			),
			"INNER JOIN foo AS f ON f.id = b.id",
			`> INNER JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"INNER JOIN foo USING id",
			`> INNER JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewInnerJoin(
				sqlabble.NewTable("foo").As("f"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"INNER JOIN foo AS f USING id",
			`> INNER JOIN foo AS f USING id
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

func TestLeftJoin(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo"),
			),
			"LEFT JOIN foo",
			`> LEFT JOIN foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo").As("f"),
			),
			"LEFT JOIN foo AS f",
			`> LEFT JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo"),
			).On(
				sqlabble.NewColumn("foo.id"),
				sqlabble.NewColumn("bar.id"),
			),
			"LEFT JOIN foo ON foo.id = bar.id",
			`> LEFT JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo").As("f"),
			).On(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			),
			"LEFT JOIN foo AS f ON f.id = b.id",
			`> LEFT JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"LEFT JOIN foo USING id",
			`> LEFT JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewLeftJoin(
				sqlabble.NewTable("foo").As("f"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"LEFT JOIN foo AS f USING id",
			`> LEFT JOIN foo AS f USING id
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

func TestRightJoin(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo"),
			),
			"RIGHT JOIN foo",
			`> RIGHT JOIN foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo").As("f"),
			),
			"RIGHT JOIN foo AS f",
			`> RIGHT JOIN foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo"),
			).On(
				sqlabble.NewColumn("foo.id"),
				sqlabble.NewColumn("bar.id"),
			),
			"RIGHT JOIN foo ON foo.id = bar.id",
			`> RIGHT JOIN foo ON foo.id = bar.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo").As("f"),
			).On(
				sqlabble.NewColumn("f.id"),
				sqlabble.NewColumn("b.id"),
			),
			"RIGHT JOIN foo AS f ON f.id = b.id",
			`> RIGHT JOIN foo AS f ON f.id = b.id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"RIGHT JOIN foo USING id",
			`> RIGHT JOIN foo USING id
`,
			[]interface{}{},
		},
		{
			sqlabble.NewRightJoin(
				sqlabble.NewTable("foo").As("f"),
			).Using(
				sqlabble.NewColumn("id"),
			),
			"RIGHT JOIN foo AS f USING id",
			`> RIGHT JOIN foo AS f USING id
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

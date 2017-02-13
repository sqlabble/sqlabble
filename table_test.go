package sqlabble_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

// func TestTableType(t *testing.T) {
// 	for _, c := range []interface{}{
// 		sqlabble.Table{},
// 		sqlabble.TableAs{},
// 	} {
// 		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
// 			if _, ok := c.(sqlabble.TableNode); !ok {
// 				t.Errorf("%T doesn't implement sqlabble.Table", c)
// 			}
// 		})
// 	}
// }

func TestTable(t *testing.T) {
	for i, c := range []struct {
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewTable("foo"),
			"foo",
			`foo
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").
				Join(sqlabble.NewTable("bar")),
			"foo JOIN bar",
			`foo
JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").
				InnerJoin(sqlabble.NewTable("bar")),
			"foo INNER JOIN bar",
			`foo
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").
				LeftJoin(sqlabble.NewTable("bar")),
			"foo LEFT JOIN bar",
			`foo
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").
				RightJoin(sqlabble.NewTable("bar")),
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
		statement sqlabble.Node
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			sqlabble.NewTable("foo").As("f"),
			"foo AS f",
			`foo AS f
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").As("f").
				Join(sqlabble.NewTable("bar")),
			"foo AS f JOIN bar",
			`foo AS f
JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").As("f").
				InnerJoin(sqlabble.NewTable("bar")),
			"foo AS f INNER JOIN bar",
			`foo AS f
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").As("f").
				LeftJoin(sqlabble.NewTable("bar")),
			"foo AS f LEFT JOIN bar",
			`foo AS f
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			sqlabble.NewTable("foo").As("f").
				RightJoin(sqlabble.NewTable("bar")),
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

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

func TestTableType(t *testing.T) {
	for _, c := range []interface{}{
		chunk.Table{},
		chunk.TableAs{},
	} {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if _, ok := c.(grammar.Table); !ok {
				t.Errorf("%T doesn't implement grammar.Table", c)
			}
		})
	}
}

func TestTable(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewTable("foo"),
			"foo",
			`foo
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").
				Join(chunk.NewTable("bar")),
			"foo JOIN bar",
			`foo
JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").
				InnerJoin(chunk.NewTable("bar")),
			"foo INNER JOIN bar",
			`foo
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").
				LeftJoin(chunk.NewTable("bar")),
			"foo LEFT JOIN bar",
			`foo
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").
				RightJoin(chunk.NewTable("bar")),
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
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewTable("foo").As("f"),
			"foo AS f",
			`foo AS f
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").As("f").
				Join(chunk.NewTable("bar")),
			"foo AS f JOIN bar",
			`foo AS f
JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").As("f").
				InnerJoin(chunk.NewTable("bar")),
			"foo AS f INNER JOIN bar",
			`foo AS f
INNER JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").As("f").
				LeftJoin(chunk.NewTable("bar")),
			"foo AS f LEFT JOIN bar",
			`foo AS f
LEFT JOIN bar
`,
			[]interface{}{},
		},
		{
			chunk.NewTable("foo").As("f").
				RightJoin(chunk.NewTable("bar")),
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

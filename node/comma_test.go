package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/node"
)

func TestComma(t *testing.T) {
	for i, c := range []struct {
		nodes     node.Comma
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			node.NewComma(),
			"",
			``,
			[]interface{}{},
		},
		{
			node.NewComma(
				node.NewExpression("foo"),
			),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			node.NewComma(
				node.NewExpression("foo"),
				node.NewExpression("bar"),
				node.NewExpression("baz"),
			),
			"foo, bar, baz",
			`> foo
> , bar
> , baz
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := c.nodes.ToSQL(ctx)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := c.nodes.ToSQL(ctxIndent)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

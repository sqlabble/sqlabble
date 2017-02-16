package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/node"
)

func TestContainer(t *testing.T) {
	for i, c := range []struct {
		nodes     node.Container
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			node.NewContainer(
				node.NewExpression("foo"),
			),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			node.NewContainer(
				node.NewExpression("foo"),
				node.NewExpression("bar"),
			),
			"foo bar",
			`> foo
>   bar
`,
			[]interface{}{},
		},
		{
			node.NewContainer(
				node.NewExpression("foo"),
				node.NewExpression("bar"),
				node.NewExpression("baz"),
				node.NewExpression("qux"),
			),
			"foo bar baz qux",
			`> foo
>   bar
>   baz
>   qux
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

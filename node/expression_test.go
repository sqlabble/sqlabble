package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/node"
)

func TestExpression(t *testing.T) {
	for i, c := range []struct {
		nodes     node.Expression
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			node.NewExpression("foo"),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			node.NewExpression("foo", 100),
			"foo",
			`> foo
`,
			[]interface{}{
				100,
			},
		},
		{
			node.NewExpression("foo", 100, "bar", true),
			"foo",
			`> foo
`,
			[]interface{}{
				100,
				"bar",
				true,
			},
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

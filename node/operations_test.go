package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
)

func TestJoinOperation(t *testing.T) {
	for i, c := range []struct {
		nodes     node.JoinOperation
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			node.NewJoinOperation(
				operator.And,
				node.NewExpression("foo"),
			),
			"foo",
			`> foo
`,
			[]interface{}{},
		},
		{
			node.NewJoinOperation(
				operator.And,
				node.NewExpression("foo"),
				node.NewExpression("bar"),
			),
			"foo AND bar",
			`> foo
> AND bar
`,
			[]interface{}{},
		},
		{
			node.NewJoinOperation(
				operator.Or,
				node.NewExpression("foo"),
				node.NewExpression("bar"),
				node.NewExpression("baz"),
			),
			"foo OR bar OR baz",
			`> foo
> OR bar
> OR baz
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

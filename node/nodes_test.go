package node_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/node"
)

var (
	ctx       = node.NewContext("", "", false)
	ctxIndent = node.NewContext("> ", "  ", false)
)

func TestParallelNodes(t *testing.T) {
	for i, c := range []struct {
		nodes     node.Nodes
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			node.NewNodes(
				node.NewExpression("foo"),
				node.NewExpression("bar"),
			),
			"foo bar",
			`> foo
> bar
`,
			[]interface{}{},
		},
		{
			node.NewNodes(
				node.NewParentheses(
					node.NewComma(
						node.NewExpression("foo-1"),
						node.NewExpression("foo-2"),
					),
				),
				node.NewParentheses(
					node.NewComma(
						node.NewExpression("bar-1"),
						node.NewExpression("bar-2"),
					),
				),
			),
			"(foo-1, foo-2) (bar-1, bar-2)",
			`> (
>   foo-1
>   , foo-2
> )
> (
>   bar-1
>   , bar-2
> )
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

// func TestSerialNodes(t *testing.T) {
// 	for i, c := range []struct {
// 		nodes     node.SerialNodes
// 		sql       string
// 		sqlIndent string
// 		values    []interface{}
// 	}{
// 		{
// 			node.NewSerialNodes(
// 				node.NewExpression("foo"),
// 				node.NewExpression("bar"),
// 			),
// 			"foo bar",
// 			`> foo bar
// `,
// 			[]interface{}{},
// 		},
// 		{
// 			node.NewSerialNodes(
// 				node.NewParentheses(
// 					node.NewComma(
// 						node.NewExpression("foo-1"),
// 						node.NewExpression("foo-2"),
// 					),
// 				),
// 				node.NewParentheses(
// 					node.NewComma(
// 						node.NewExpression("bar-1"),
// 						node.NewExpression("bar-2"),
// 					),
// 				),
// 			),
// 			"(foo-1, foo-2) (bar-1, bar-2)",
// 			`> (
// >   foo-1
// >   , foo-2
// > ) (
// >   bar-1
// >   , bar-2
// > )
// `,
// 			[]interface{}{},
// 		},
// 	} {
// 		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
// 			sql, values := c.nodes.ToSQL(ctx)
// 			if sql != c.sql {
// 				t.Error(diff.SQL(sql, c.sql))
// 			}
// 			if !reflect.DeepEqual(values, c.values) {
// 				t.Error(diff.Values(values, c.values))
// 			}
// 		})
// 		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
// 			sql, values := c.nodes.ToSQL(ctxIndent)
// 			if sql != c.sqlIndent {
// 				t.Error(diff.SQL(sql, c.sqlIndent))
// 			}
// 			if !reflect.DeepEqual(values, c.values) {
// 				t.Error(diff.Values(values, c.values))
// 			}
// 		})
// 	}
// }

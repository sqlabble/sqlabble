package generator_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/internal/diff"
)

var (
	ctx       = generator.NewContext("", "", false)
	ctxIndent = generator.NewContext("> ", "  ", false)
)

func TestParallelNodes(t *testing.T) {
	for i, c := range []struct {
		nodes     generator.Nodes
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			generator.NewNodes(
				generator.NewExpression("foo"),
				generator.NewExpression("bar"),
			),
			"foo bar",
			`> foo
> bar
`,
			[]interface{}{},
		},
		{
			generator.NewNodes(
				generator.NewParentheses(
					generator.NewComma(
						generator.NewExpression("foo-1"),
						generator.NewExpression("foo-2"),
					),
				),
				generator.NewParentheses(
					generator.NewComma(
						generator.NewExpression("bar-1"),
						generator.NewExpression("bar-2"),
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
// 		nodes     generator.SerialNodes
// 		sql       string
// 		sqlIndent string
// 		values    []interface{}
// 	}{
// 		{
// 			generator.NewSerialNodes(
// 				generator.NewExpression("foo"),
// 				generator.NewExpression("bar"),
// 			),
// 			"foo bar",
// 			`> foo bar
// `,
// 			[]interface{}{},
// 		},
// 		{
// 			generator.NewSerialNodes(
// 				generator.NewParentheses(
// 					generator.NewComma(
// 						generator.NewExpression("foo-1"),
// 						generator.NewExpression("foo-2"),
// 					),
// 				),
// 				generator.NewParentheses(
// 					generator.NewComma(
// 						generator.NewExpression("bar-1"),
// 						generator.NewExpression("bar-2"),
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

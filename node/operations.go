package node

import (
	"fmt"

	"github.com/minodisk/sqlabble/operator"
)

// JoinOperation is a Node that joins multiple nodes with operators.
type JoinOperation struct {
	operator operator.Operator
	nodes    []Node
}

// NewJoinOperation returns a new JoinOperation.
func NewJoinOperation(operator operator.Operator, nodes ...Node) JoinOperation {
	return JoinOperation{
		operator: operator,
		nodes:    nodes,
	}
}

// ToSQL returns a query and a slice of values.
func (o JoinOperation) ToSQL(ctx Context) (string, []interface{}) {
	for {
		if len(o.nodes) != 1 {
			break
		}
		f := o.nodes[0]
		t, ok := f.(JoinOperation)
		if !ok {
			break
		}
		o = t
	}

	head := ctx.CurrentHead()
	hasParentheses := head != "" || !ctx.isTopParentheses()
	ctx = ctx.ClearHead()
	c1 := ctx.incParenthesesDepth()
	if hasParentheses {
		c1 = c1.incDepth()
	}

	sqls := make([]string, len(o.nodes))
	values := []interface{}{}
	for i, f := range o.nodes {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = f.ToSQL(c1.ClearHead())
			values = append(values, vs...)
			continue
		}
		sqls[i], vs = f.ToSQL(c1.setHead(fmt.Sprintf("%s ", o.operator)))
		values = append(values, vs...)
	}
	sql := ctx.join(sqls...)

	if !hasParentheses {
		return sql, values
	}
	if ctx.IsBreaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}
	return fmt.Sprintf("%s(%s)", head, sql), values
}

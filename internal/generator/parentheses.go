package generator

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/operator"
)

// Parentheses is a Node with parentheses around a node.
type Parentheses struct {
	node Node
}

// NewParentheses returns a new Parentheses.
func NewParentheses(node Node) Parentheses {
	return Parentheses{
		node: node,
	}
}

// ToSQL returns a query and a slice of values.
func (b Parentheses) ToSQL(ctx Context) (string, []interface{}) {
	head := ctx.currentHead()
	ctx = ctx.clearParenthesesDepth()

	sql, values := b.node.ToSQL(ctx.incDepth().clearHead())

	if ctx.isBreaking() {
		p := ctx.pre()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}

	return fmt.Sprintf("%s(%s)", head, sql), values
}

// OpParentheses is a Parentheses with an operator at the head.
type OpParentheses struct {
	op    operator.Operator
	paren Parentheses
}

// NewOpParenteses returns a new OpParentheses.
func NewOpParenteses(op operator.Operator, paren Parentheses) OpParentheses {
	return OpParentheses{
		op:    op,
		paren: paren,
	}
}

// ToSQL returns a query and a slice of values.
func (o OpParentheses) ToSQL(ctx Context) (string, []interface{}) {
	head := ctx.currentHead()

	return o.paren.ToSQL(ctx.setHead(head + string(o.op) + " "))
}

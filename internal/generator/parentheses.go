package generator

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/operator"
)

type Parentheses struct {
	node Node
}

func NewParentheses(node Node) Parentheses {
	return Parentheses{
		node: node,
	}
}

func (b Parentheses) ToSQL(ctx Context) (string, []interface{}) {
	head := ctx.Head()
	ctx = ctx.ClearParenthesesDepth()

	sql, values := b.node.ToSQL(ctx.IncDepth().ClearHead())

	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}

	return fmt.Sprintf("%s(%s)", head, sql), values
}

type OpParentheses struct {
	op    operator.Operator
	paren Parentheses
}

func NewOpParenteses(op operator.Operator, paren Parentheses) OpParentheses {
	return OpParentheses{
		op:    op,
		paren: paren,
	}
}

func (o OpParentheses) ToSQL(ctx Context) (string, []interface{}) {
	head := ctx.Head()

	return o.paren.ToSQL(ctx.SetHead(head + string(o.op) + " "))
}

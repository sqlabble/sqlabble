package statement

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
)

type Assign struct {
	column Column
	value  interface{}
}

func NewAssign(column Column, value interface{}) Assign {
	return Assign{
		column: column,
		value:  value,
	}
}

func (a Assign) node() node.Node {
	return a.expression()
}

func (a Assign) expression() node.Expression {
	return a.column.expression().
		Append(node.NewExpression(string(operator.Eq))).
		Append(node.ValuesToExpression(a.value))
}

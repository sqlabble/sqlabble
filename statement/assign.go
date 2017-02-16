package statement

import (
	"github.com/minodisk/sqlabble/generator"
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

func (a Assign) node() generator.Node {
	return a.expression()
}

func (a Assign) expression() generator.Expression {
	return a.column.expression().
		Append(generator.NewExpression(string(operator.Eq))).
		Append(generator.ValuesToExpression(a.value))
}

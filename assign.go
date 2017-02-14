package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

type assign struct {
	col column
	val interface{}
}

func newAssign(col column, val interface{}) assign {
	return assign{
		col: col,
		val: val,
	}
}

func (a assign) node() generator.Node {
	return a.expression()
}

func (a assign) expression() generator.Expression {
	return a.col.expression().
		Append(generator.NewExpression(string(operator.Eq))).
		Append(generator.ValuesToExpression(a.val))
}

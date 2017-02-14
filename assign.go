package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/operator"
)

type assign struct {
	col   column
	value interface{}
}

func newAssign(col column, value interface{}) assign {
	return assign{
		col:   col,
		value: value,
	}
}

func (a assign) node() generator.Node {
	return a.expression()
}

func (a assign) expression() generator.Expression {
	return a.col.expression().
		Append(generator.NewExpression(string(operator.Equal))).
		Append(generator.NewPlaceholders(a.value))
}

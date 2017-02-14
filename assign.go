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

func (a assign) generator() generator.Generator {
	return a.Expression()
}

func (a assign) Expression() generator.Expression {
	return a.col.expression().
		Append(generator.NewExpression(string(operator.Equal))).
		Append(generator.NewPlaceholders(a.value))
}

package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
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

func (a Assign) Expression() generator.Expression {
	return a.column.Expression().
		Append(generator.NewExpression(string(operator.Equal))).
		Append(generator.NewPlaceholders(a.value))
}

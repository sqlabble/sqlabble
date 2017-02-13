package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type columnAs struct {
	col   columnNode
	alias string
}

func newColumnAs(alias string) columnAs {
	return columnAs{
		alias: alias,
	}
}

func (c columnAs) generator() generator.Generator {
	return c.expression()
}

func (c columnAs) expression() generator.Expression {
	a := generator.NewExpression(operator.As).
		Append(generator.NewExpression(c.Alias()))
	if c.col == nil {
		return a
	}
	return c.col.expression().
		Append(a)
}

func (c columnAs) columnName() string {
	return c.col.columnName()
}

func (c columnAs) Alias() string {
	return c.alias
}

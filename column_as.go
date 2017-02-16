package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/operator"
)

type columnAs struct {
	col  columnOrColumnAs
	alia string
}

func newColumnAs(alias string) columnAs {
	return columnAs{
		alia: alias,
	}
}

func (c columnAs) node() generator.Node {
	return c.expression()
}

func (c columnAs) expression() generator.Expression {
	a := generator.NewExpression(string(operator.As)).
		Append(generator.NewExpression(c.alias()))
	if c.col == nil {
		return a
	}
	return c.col.expression().
		Append(a)
}

func (c columnAs) columnName() string {
	return c.col.columnName()
}

func (c columnAs) alias() string {
	return c.alia
}

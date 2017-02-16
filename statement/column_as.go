package statement

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
)

type ColumnAs struct {
	column ColumnOrColumnAs
	alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		alias: alias,
	}
}

func (c ColumnAs) node() node.Node {
	return c.expression()
}

func (c ColumnAs) expression() node.Expression {
	a := node.NewExpression(string(operator.As)).
		Append(node.NewExpression(c.alias))
	if c.column == nil {
		return a
	}
	return c.column.expression().
		Append(a)
}

func (c ColumnAs) columnName() string {
	return c.column.columnName()
}

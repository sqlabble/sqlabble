package statement

import (
	"fmt"

	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
)

type ColumnAs struct {
	column Column
	alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		alias: alias,
	}
}

func (c ColumnAs) node() node.Node {
	return node.JoinExpressions(
		c.column.expression(),
		node.NewExpression(string(operator.As)),
		node.NewExpression(fmt.Sprintf("`%s`", c.alias)),
	)
}

func (c ColumnAs) columnName() string {
	return c.column.columnName()
}

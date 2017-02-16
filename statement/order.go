package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/node"
)

type Order struct {
	column Column
	dir    direction.Direction
}

func (o Order) node() node.Node {
	return node.JoinExpressions(
		o.column.expression(),
		node.NewExpression(string(o.dir)),
	)
}

func (o Order) direction() direction.Direction {
	return o.dir
}

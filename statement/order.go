package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/node"
)

type Order struct {
	column Column
	dir    direction.Direction
}

func NewOrder(dir direction.Direction) Order {
	return Order{
		dir: dir,
	}
}

func (o Order) node() node.Node {
	return node.JoinExpressions(
		o.column.expression(),
		node.NewExpression(string(o.direction())),
	)
}

func (o Order) direction() direction.Direction {
	return o.dir
}

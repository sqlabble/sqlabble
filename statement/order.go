package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/generator"
)

type Order struct {
	column Column
	dir    direction.Direction
}

func (o Order) node() generator.Node {
	return generator.JoinExpressions(
		o.column.expression(),
		generator.NewExpression(string(o.dir)),
	)
}

func (o Order) direction() direction.Direction {
	return o.dir
}

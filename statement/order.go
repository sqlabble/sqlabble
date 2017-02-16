package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/generator"
)

type order struct {
	col column
	dir direction.Direction
}

func (o order) node() generator.Node {
	return generator.JoinExpressions(
		o.col.expression(),
		generator.NewExpression(string(o.dir)),
	)
}

func (o order) direction() direction.Direction {
	return o.dir
}

package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/generator"
)

type order struct {
	col column
	dir direction.Direction
}

func (o order) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", o.col.name, o.dir),
	)
}

func (o order) direction() direction.Direction {
	return o.dir
}

package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
)

type order struct {
	col column
	dir direction.Direction
}

func (o order) generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", o.col.name, o.dir),
	)
}

func (o order) direction() direction.Direction {
	return o.dir
}

package token

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
)

type Order struct {
	column    Column
	direction direction.Direction
}

func (o Order) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", o.column.name, o.direction),
	)
}

func (o Order) Direction() direction.Direction {
	return o.direction
}

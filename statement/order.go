package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/token"
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

func (o Order) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	return t1.Append(
		token.Space,
		token.Word(o.direction()),
	), v1
}

func (o Order) direction() direction.Direction {
	return o.dir
}

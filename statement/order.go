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
	return o.line()
}

func (o Order) line() (token.Line, []interface{}) {
	line, values := o.column.line()
	return line.A(token.Space, token.Word(o.direction())), values
}

func (o Order) direction() direction.Direction {
	return o.dir
}

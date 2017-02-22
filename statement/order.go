package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Order struct {
	column Column
	dir    keyword.Direction
}

func NewAsc() Order {
	return Order{
		dir: keyword.ASC,
	}
}

func NewDesc() Order {
	return Order{
		dir: keyword.DESC,
	}
}

func (o Order) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	return t1.Append(
		token.Space,
		token.Word(o.keyword()),
	), v1
}

func (o Order) keyword() keyword.Direction {
	return o.dir
}

package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Order struct {
	column *Column
	dir    keyword.Direction
}

func NewAsc() *Order {
	return &Order{
		dir: keyword.ASC,
	}
}

func NewDesc() *Order {
	return &Order{
		dir: keyword.DESC,
	}
}

func (o *Order) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if o.column == nil {
		return tokenizer.NewLine(token.Word(o.keyword())), nil
	}

	t1, v1 := o.column.nodeize()
	return t1.Append(
		token.Word(o.keyword()),
	), v1
}

func (o *Order) keyword() keyword.Direction {
	return o.dir
}

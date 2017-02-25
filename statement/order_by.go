package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type OrderBy struct {
	Prev
	Next
	orders []Order
}

func NewOrderBy(orders ...Order) *OrderBy {
	return &OrderBy{
		orders: orders,
	}
}

func (o *OrderBy) Limit(count int) *Limit {
	l := NewLimit(count)
	Link(o, l)
	return l
}

func (o *OrderBy) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(o.orders))
	values := []interface{}{}
	for i, o := range o.orders {
		var vals []interface{}
		tokenizers[i], vals = o.nodeize()
		values = append(values, vals...)
	}
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.OrderBy)),
	).SetMiddle(
		tokenizers.Prefix(token.Comma),
	), values
}

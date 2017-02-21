package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type OrderBy struct {
	prev   Clause
	orders []Order
}

func NewOrderBy(os ...Order) OrderBy {
	return OrderBy{orders: os}
}

func (o OrderBy) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(o)
}

func (o OrderBy) self() (token.Tokenizer, []interface{}) {
	tokenizers := make(token.Tokenizers, len(o.orders))
	values := []interface{}{}
	for i, o := range o.orders {
		var vals []interface{}
		tokenizers[i], vals = o.nodeize()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.OrderBy)),
	).SetMiddle(
		tokenizers.Prefix(token.Comma, token.Space),
	), values
}

func (o OrderBy) previous() Clause {
	return o.prev
}

func (o OrderBy) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = o
	return l
}

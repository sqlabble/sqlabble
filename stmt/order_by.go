package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type OrderBy struct {
	prev   Prever
	orders []Order
}

func NewOrderBy(os ...Order) OrderBy {
	return OrderBy{orders: os}
}

func (o OrderBy) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = o
	return l
}

func (o OrderBy) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(o)
}

func (o OrderBy) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
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

func (o OrderBy) previous() Prever {
	return o.prev
}

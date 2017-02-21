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
	clauses := clauseNodes(o)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (o OrderBy) container() (token.Container, []interface{}) {
	lines := make(token.Lines, len(o.orders))
	values := []interface{}{}
	for i, o := range o.orders {
		var vals []interface{}
		lines[i], vals = o.line()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.OrderBy)),
	).SetMiddle(
		lines.Prefix(token.Comma, token.Space),
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

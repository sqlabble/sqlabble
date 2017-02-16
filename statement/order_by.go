package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type OrderBy struct {
	prev   Clause
	orders []Order
}

func NewOrderBy(os ...Order) OrderBy {
	return OrderBy{orders: os}
}

func (o OrderBy) node() generator.Node {
	cs := clauseNodes(o)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewNodes(fs...)
}

func (o OrderBy) myNode() generator.Node {
	fs := make([]generator.Node, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.node()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.OrderBy)),
		generator.NewComma(fs...),
	)
}

func (o OrderBy) previous() Clause {
	return o.prev
}

func (o OrderBy) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = o
	return l
}

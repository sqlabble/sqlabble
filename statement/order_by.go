package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type OrderBy struct {
	prev   Clause
	orders []Order
}

func NewOrderBy(os ...Order) OrderBy {
	return OrderBy{orders: os}
}

func (o OrderBy) node() node.Node {
	cs := clauseNodes(o)
	fs := make([]node.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return node.NewNodes(fs...)
}

func (o OrderBy) myNode() node.Node {
	fs := make([]node.Node, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.node()
	}
	return node.NewContainer(
		node.NewExpression(string(keyword.OrderBy)),
		node.NewComma(fs...),
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

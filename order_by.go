package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type orderBy struct {
	prev   clause
	orders []order
}

func newOrderBy(os ...order) orderBy {
	return orderBy{orders: os}
}

func (o orderBy) node() generator.Node {
	cs := clauseNodes(o)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.nodeMine()
	}
	return generator.NewNodes(fs...)
}

func (o orderBy) nodeMine() generator.Node {
	fs := make([]generator.Node, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.node()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.OrderBy)),
		generator.NewComma(fs...),
	)
}

func (o orderBy) previous() clause {
	return o.prev
}

func (o orderBy) Limit(offset, lim int) limit {
	l := newLimit(offset, lim)
	l.prev = o
	return l
}

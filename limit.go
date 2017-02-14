package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type limit struct {
	prev   clause
	offset int
	lim    int
}

func newLimit(offset, lim int) limit {
	return limit{
		offset: offset,
		lim:    lim,
	}
}

func (l limit) node() generator.Node {
	cs := clauseNodes(l)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (l limit) myNode() generator.Node {
	var p generator.Expression
	if l.offset == 0 {
		p = generator.ValuesToExpression(l.lim)
	} else {
		p = generator.ValuesToExpression(l.offset, l.lim)
	}
	return generator.NewContainer(
		generator.NewExpression(keyword.Limit),
		p,
	)
}

func (l limit) previous() clause {
	return l.prev
}

package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
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
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.nodeMine()
	}
	return generator.NewNodes(fs...)
}

func (l limit) nodeMine() generator.Node {
	var p generator.Expression
	if l.offset == 0 {
		p = generator.NewPlaceholders(l.lim)
	} else {
		p = generator.NewPlaceholders(l.offset, l.lim)
	}
	return generator.NewContainer(
		generator.NewExpression(keyword.Limit),
		p,
	)
}

func (l limit) previous() clause {
	return l.prev
}

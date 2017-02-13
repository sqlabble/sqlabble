package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type limit struct {
	prev   clauseNode
	offset int
	lim    int
}

func newLimit(offset, lim int) limit {
	return limit{
		offset: offset,
		lim:    lim,
	}
}

func (l limit) generator() generator.Generator {
	cs := clauseNodes(l)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (l limit) clauseGenerator() generator.Generator {
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

func (l limit) previous() clauseNode {
	return l.prev
}

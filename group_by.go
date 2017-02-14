package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type groupBy struct {
	prev    clauseNode
	columns []column
}

func newGroupBy(col column, columns ...column) groupBy {
	return groupBy{
		columns: append([]column{col}, columns...),
	}
}

func (g groupBy) generator() generator.Generator {
	cs := clauseNodes(g)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (g groupBy) clauseGenerator() generator.Generator {
	gs := make([]generator.Generator, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.GroupBy)),
		generator.NewComma(gs...),
	)
}

func (g groupBy) previous() clauseNode {
	return g.prev
}

func (g groupBy) Having(operation operationNode) having {
	l := newHaving(operation)
	l.prev = g
	return l
}

func (g groupBy) Limit(offset, lim int) limit {
	l := newLimit(offset, lim)
	l.prev = g
	return l
}

package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type GroupBy struct {
	prev    grammar.Clause
	columns []Column
}

func NewGroupBy(column Column, columns ...Column) GroupBy {
	return GroupBy{
		columns: append([]Column{column}, columns...),
	}
}

func (g GroupBy) Generator() generator.Generator {
	cs := grammar.Clauses(g)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (g GroupBy) ClauseGenerator() generator.Generator {
	gs := make([]generator.Generator, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.Generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.GroupBy)),
		generator.NewComma(gs...),
	)
}

func (g GroupBy) Prev() grammar.Clause {
	return g.prev
}

func (g GroupBy) Having(operation grammar.Operation) Having {
	l := NewHaving(operation)
	l.prev = g
	return l
}

func (g GroupBy) Limit(offset, limit int) Limit {
	l := NewLimit(offset, limit)
	l.prev = g
	return l
}

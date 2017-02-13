package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Limit struct {
	prev   grammar.Clause
	offset int
	limit  int
}

func NewLimit(offset, limit int) Limit {
	return Limit{
		offset: offset,
		limit:  limit,
	}
}

func (l Limit) Generator() generator.Generator {
	cs := grammar.Clauses(l)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (l Limit) ClauseGenerator() generator.Generator {
	var p generator.Expression
	if l.offset == 0 {
		p = generator.NewPlaceholders(l.limit)
	} else {
		p = generator.NewPlaceholders(l.offset, l.limit)
	}
	return generator.NewContainer(
		generator.NewExpression(keyword.Limit),
		p,
	)
}

func (l Limit) Prev() grammar.Clause {
	return l.prev
}

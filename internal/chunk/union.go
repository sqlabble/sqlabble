package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Union struct {
	statements []grammar.Statement
}

func NewUnion(statements ...grammar.Statement) Union {
	return Union{
		statements: statements,
	}
}

func (u Union) Generator() generator.Generator {
	cs := grammar.Clauses(u)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (u Union) ClauseGenerator() generator.Generator {
	sep := generator.NewExpression(keyword.Union)
	gs := make([]generator.Generator, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.Generator()
	}
	return generator.NewUnions(sep, gs...)
}

func (u Union) Prev() grammar.Clause {
	return nil
}

func (u Union) OrderBy(os ...grammar.Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

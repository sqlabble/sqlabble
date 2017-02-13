package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type SetOperation struct {
	operator   string
	statements []grammar.Statement
}

func NewUnion(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Except,
		statements: statements,
	}
}

func NewExceptAll(statements ...grammar.Statement) SetOperation {
	return SetOperation{
		operator:   keyword.ExceptAll,
		statements: statements,
	}
}

func (u SetOperation) Generator() generator.Generator {
	cs := grammar.Clauses(u)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (u SetOperation) ClauseGenerator() generator.Generator {
	sep := generator.NewExpression(u.operator)
	gs := make([]generator.Generator, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.Generator()
	}
	return generator.NewUnions(sep, gs...)
}

func (u SetOperation) Prev() grammar.Clause {
	return nil
}

func (u SetOperation) OrderBy(os ...grammar.Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

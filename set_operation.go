package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type setOperation struct {
	operator   string
	statements []Node
}

func newUnion(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.Union,
		statements: statements,
	}
}

func newUnionAll(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.UnionAll,
		statements: statements,
	}
}

func newIntersect(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.Intersect,
		statements: statements,
	}
}

func newIntersectAll(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.IntersectAll,
		statements: statements,
	}
}

func newExcept(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.Except,
		statements: statements,
	}
}

func newExceptAll(statements ...Node) setOperation {
	return setOperation{
		operator:   keyword.ExceptAll,
		statements: statements,
	}
}

func (u setOperation) generator() generator.Generator {
	cs := clauseNodes(u)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (u setOperation) clauseGenerator() generator.Generator {
	sep := generator.NewExpression(u.operator)
	gs := make([]generator.Generator, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.generator()
	}
	return generator.NewUnions(sep, gs...)
}

func (u setOperation) previous() clause {
	return nil
}

func (u setOperation) OrderBy(os ...order) orderBy {
	o := newOrderBy(os...)
	o.prev = u
	return o
}

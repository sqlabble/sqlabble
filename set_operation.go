package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type setOperation struct {
	operator   string
	statements []Statement
}

func newUnion(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Union,
		statements: statements,
	}
}

func newUnionAll(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.UnionAll,
		statements: statements,
	}
}

func newIntersect(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Intersect,
		statements: statements,
	}
}

func newIntersectAll(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.IntersectAll,
		statements: statements,
	}
}

func newExcept(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Except,
		statements: statements,
	}
}

func newExceptAll(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.ExceptAll,
		statements: statements,
	}
}

func (u setOperation) node() generator.Node {
	cs := clauseNodes(u)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return generator.NewParallelNodes(gs...)
}

func (u setOperation) myNode() generator.Node {
	sep := generator.NewExpression(u.operator)
	gs := make([]generator.Node, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.node()
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

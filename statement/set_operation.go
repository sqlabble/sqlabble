package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type SetOperation struct {
	operator   string
	statements []Statement
}

func NewUnion(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.Except,
		statements: statements,
	}
}

func NewExceptAll(statements ...Statement) SetOperation {
	return SetOperation{
		operator:   keyword.ExceptAll,
		statements: statements,
	}
}

func (u SetOperation) node() generator.Node {
	cs := clauseNodes(u)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return generator.NewNodes(gs...)
}

func (u SetOperation) myNode() generator.Node {
	sep := generator.NewExpression(u.operator)
	gs := make([]generator.Node, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.node()
	}
	return generator.NewSet(sep, gs...)
}

func (u SetOperation) previous() Clause {
	return nil
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

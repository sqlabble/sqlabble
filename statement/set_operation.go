package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (u SetOperation) node() node.Node {
	cs := clauseNodes(u)
	gs := make([]node.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return node.NewNodes(gs...)
}

func (u SetOperation) myNode() node.Node {
	sep := node.NewExpression(u.operator)
	gs := make([]node.Node, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.node()
	}
	return node.NewSet(sep, gs...)
}

func (u SetOperation) previous() Clause {
	return nil
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type setOperation struct {
	operator   string
	statements []Statement
}

func NewUnion(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...Statement) setOperation {
	return setOperation{
		operator:   keyword.Except,
		statements: statements,
	}
}

func NewExceptAll(statements ...Statement) setOperation {
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
	return generator.NewNodes(gs...)
}

func (u setOperation) myNode() generator.Node {
	sep := generator.NewExpression(u.operator)
	gs := make([]generator.Node, len(u.statements))
	for i, s := range u.statements {
		gs[i] = s.node()
	}
	return generator.NewSet(sep, gs...)
}

func (u setOperation) previous() clause {
	return nil
}

func (u setOperation) OrderBy(os ...order) orderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SetOperation struct {
	Prev
	Next
	op         keyword.Operator
	statements []Statement
}

func NewUnion(statements ...Statement) *SetOperation {
	return NewSetOperation(keyword.Union, statements...)
}

func NewUnionAll(statements ...Statement) *SetOperation {
	return NewSetOperation(keyword.UnionAll, statements...)
}

func NewIntersect(statements ...Statement) *SetOperation {
	return NewSetOperation(keyword.Intersect, statements...)
}

func NewIntersectAll(statements ...Statement) *SetOperation {
	return NewSetOperation(keyword.IntersectAll, statements...)
}

func NewExcept(statements ...Statement) *SetOperation {
	return NewSetOperation(keyword.Except, statements...)
}

func NewSetOperation(op keyword.Operator, stmts ...Statement) *SetOperation {
	s := &SetOperation{
		op:         op,
		statements: stmts,
	}
	return s
}

func NewExceptAll(statements ...Statement) *SetOperation {
	return &SetOperation{
		op:         keyword.ExceptAll,
		statements: statements,
	}
}

func (u *SetOperation) OrderBy(os ...Order) *OrderBy {
	o := NewOrderBy(os...)
	Link(u, o)
	return o
}

func (u *SetOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewLine(token.Word(u.op)), nil
}

func (u *SetOperation) list() []Statement {
	ns := []Statement{}
	for i, s := range u.statements {
		if i == 0 {
			ns = append(ns, NewSubquery(s))
			continue
		}
		ns = append(ns, u, NewSubquery(s))
	}
	return ns
}

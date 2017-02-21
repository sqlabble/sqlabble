package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type SetOperation struct {
	op         operator.Operator
	statements []Statement
}

func NewUnion(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.Except,
		statements: statements,
	}
}

func NewExceptAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         operator.ExceptAll,
		statements: statements,
	}
}

func (u SetOperation) nodeize() (token.Tokenizer, []interface{}) {
	return u.self()
}

func (u SetOperation) self() (token.Tokenizer, []interface{}) {
	tokenizers := make(token.Tokenizers, len(u.statements))
	values := []interface{}{}
	for i, s := range u.statements {
		t, vals := s.nodeize()
		t = token.NewParentheses(t)
		if i != 0 {
			t = t.Prepend(
				token.Word(u.operator()),
				token.Space,
			)
		}
		tokenizers[i] = t
		values = append(values, vals...)
	}
	if len(values) == 0 {
		values = nil
	}
	return tokenizers, values
}

func (u SetOperation) previous() Clause {
	return nil
}

func (u SetOperation) operator() operator.Operator {
	return u.op
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

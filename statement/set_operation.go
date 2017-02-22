package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type SetOperation struct {
	op         keyword.Operator
	statements []Statement
}

func NewUnion(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.Except,
		statements: statements,
	}
}

func NewExceptAll(statements ...Statement) SetOperation {
	return SetOperation{
		op:         keyword.ExceptAll,
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
				token.Word(u.keyword()),
				token.Space,
			)
		}
		tokenizers[i] = t
		values = append(values, vals...)
	}
	return tokenizers, values
}

func (u SetOperation) previous() Clause {
	return nil
}

func (u SetOperation) keyword() keyword.Operator {
	return u.op
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

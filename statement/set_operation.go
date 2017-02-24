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
	return &SetOperation{
		op:         keyword.Union,
		statements: statements,
	}
}

func NewUnionAll(statements ...Statement) *SetOperation {
	return &SetOperation{
		op:         keyword.UnionAll,
		statements: statements,
	}
}

func NewIntersect(statements ...Statement) *SetOperation {
	return &SetOperation{
		op:         keyword.Intersect,
		statements: statements,
	}
}

func NewIntersectAll(statements ...Statement) *SetOperation {
	return &SetOperation{
		op:         keyword.IntersectAll,
		statements: statements,
	}
}

func NewExcept(statements ...Statement) *SetOperation {
	return &SetOperation{
		op:         keyword.Except,
		statements: statements,
	}
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
	tokenizers := make(tokenizer.Tokenizers, len(u.statements))
	values := []interface{}{}
	for i, s := range u.statements {
		t, vals := s.nodeize()
		t = tokenizer.NewParentheses(t)
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

func (u *SetOperation) keyword() keyword.Operator {
	return u.op
}

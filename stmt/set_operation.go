package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SetOperation struct {
	op    keyword.Operator
	stmts []Statement
}

func NewUnion(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.Union,
		stmts: stmts,
	}
}

func NewUnionAll(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.UnionAll,
		stmts: stmts,
	}
}

func NewIntersect(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.Intersect,
		stmts: stmts,
	}
}

func NewIntersectAll(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.IntersectAll,
		stmts: stmts,
	}
}

func NewExcept(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.Except,
		stmts: stmts,
	}
}

func NewExceptAll(stmts ...Statement) SetOperation {
	return SetOperation{
		op:    keyword.ExceptAll,
		stmts: stmts,
	}
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

func (u SetOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(u)
}

func (u SetOperation) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(u.stmts))
	values := []interface{}{}
	for i, s := range u.stmts {
		t, vals := s.nodeize()
		t = tokenizer.NewParentheses(t)
		if i != 0 {
			t = t.Prepend(
				token.Word(u.keyword()),
			)
		}
		tokenizers[i] = t
		values = append(values, vals...)
	}
	return tokenizers, values
}

func (u SetOperation) previous() Prever {
	return nil
}

func (u SetOperation) keyword() keyword.Operator {
	return u.op
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
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

func (u SetOperation) nodeize() (token.Tokenizer, []interface{}) {
	return u.tokenizers()
}

func (u SetOperation) tokenizers() (token.Tokenizers, []interface{}) {
	ts := make([]token.Tokenizer, len(u.statements))
	values := []interface{}{}
	for i, s := range u.statements {
		var vals []interface{}
		ts[i], vals = s.nodeize()
		values = append(values, vals...)
	}
	return token.NewTokenizers(ts...).Prefix(token.Word(u.operator)), values
}

func (u SetOperation) container() (token.Container, []interface{}) {
	return token.NewContainer(token.NewLine(token.Word(""))), nil
}

func (u SetOperation) previous() Clause {
	return nil
}

func (u SetOperation) OrderBy(os ...Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = u
	return o
}

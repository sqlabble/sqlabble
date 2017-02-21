package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Set struct {
	prev    Clause
	assigns []Assign
}

func NewSet(assigns ...Assign) Set {
	return Set{
		assigns: assigns,
	}
}

func (s Set) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(s)
}

func (s Set) self() (token.Tokenizer, []interface{}) {
	tokenizers := make(token.Tokenizers, len(s.assigns))
	values := []interface{}{}
	for i, a := range s.assigns {
		var vals []interface{}
		tokenizers[i], vals = a.nodeize()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(
			token.Word(keyword.Set),
		),
	).SetMiddle(
		tokenizers.Prefix(token.Comma, token.Space),
	), values
}

func (s Set) previous() Clause {
	return s.prev
}

func (s Set) Where(operation ComparisonOrLogicalOperation) Where {
	w := NewWhere(operation)
	w.prev = s
	return w
}

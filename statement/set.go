package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Set struct {
	Prev
	Next
	assigns []Assign
}

func NewSet(assigns ...Assign) *Set {
	return &Set{
		assigns: assigns,
	}
}

func (s *Set) Where(operation ComparisonOrLogicalOperation) *Where {
	w := NewWhere(operation)
	Link(s, w)
	return w
}

func (s *Set) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(s.assigns))
	values := []interface{}{}
	for i, a := range s.assigns {
		var vals []interface{}
		tokenizers[i], vals = a.nodeize()
		values = append(values, vals...)
	}
	return tokenizer.NewContainer(
		tokenizer.NewLine(
			token.Word(keyword.Set),
		),
	).SetMiddle(
		tokenizers.Prefix(token.Comma, token.Space),
	), values
}

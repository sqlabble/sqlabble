package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Set struct {
	prev    Prever
	assigns []*Assign
}

func NewSet(assigns ...*Assign) *Set {
	return &Set{
		assigns: assigns,
	}
}

func (s *Set) Where(operation ComparisonOrLogicalOperation) *Where {
	w := NewWhere(operation)
	w.prev = s
	return w
}

func (s *Set) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(s)
}

func (s *Set) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
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
		tokenizers.Prefix(token.Comma),
	), values
}

func (s *Set) previous() Prever {
	return s.prev
}

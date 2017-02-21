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
	clauses := clauseNodes(s)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (s Set) container() (token.Container, []interface{}) {
	lines := make(token.Lines, len(s.assigns))
	values := []interface{}{}
	for i, a := range s.assigns {
		var vals []interface{}
		lines[i], vals = a.line()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Set)),
	).SetMiddle(
		lines,
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

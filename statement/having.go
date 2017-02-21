package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Having struct {
	prev      Clause
	operation ComparisonOrLogicalOperation
}

func NewHaving(operation ComparisonOrLogicalOperation) Having {
	return Having{
		operation: operation,
	}
}

func (h Having) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(h)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (h Having) container() (token.Container, []interface{}) {
	middle, values := h.operation.nodeize()
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Having)),
	).SetMiddle(
		middle,
	), values
}

func (h Having) previous() Clause {
	return h.prev
}

func (h Having) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = h
	return o
}

func (h Having) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = h
	return l
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Where struct {
	prev      Clause
	operation ComparisonOrLogicalOperation
}

func NewWhere(operation ComparisonOrLogicalOperation) Where {
	return Where{
		operation: operation,
	}
}

func (w Where) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(w)
	ts := make(token.Tokenizers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		ts[i], vals = c.container()
		values = append(values, vals...)
	}
	return ts, values
}

func (w Where) container() (token.Container, []interface{}) {
	middle, values := w.operation.nodeize()
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Where)),
	).SetMiddle(
		middle,
	), values
}

func (w Where) previous() Clause {
	return w.prev
}

func (w Where) GroupBy(column Column, columns ...Column) GroupBy {
	g := NewGroupBy(column, columns...)
	g.prev = w
	return g
}

func (w Where) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = w
	return o
}

func (w Where) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = w
	return l
}

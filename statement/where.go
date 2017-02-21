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
	return nodeizeClauses(w)
}

func (w Where) self() (token.Tokenizer, []interface{}) {
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

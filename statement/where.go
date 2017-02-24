package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Where struct {
	Prev
	Next
	operation ComparisonOrLogicalOperation
}

func NewWhere(operation ComparisonOrLogicalOperation) *Where {
	return &Where{
		operation: operation,
	}
}

func (w *Where) GroupBy(column Column, columns ...Column) *GroupBy {
	g := NewGroupBy(column, columns...)
	Link(w, g)
	return g
}

func (w *Where) OrderBy(orders ...Order) *OrderBy {
	o := NewOrderBy(orders...)
	Link(w, o)
	return o
}

func (w *Where) Limit(count int) *Limit {
	l := NewLimit(count)
	Link(w, l)
	return l
}

func (w *Where) nodeize() (tokenizer.Tokenizer, []interface{}) {
	middle, values := w.operation.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Where)),
	).SetMiddle(
		middle,
	), values
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Having struct {
	Prev
	Next
	operation ComparisonOrLogicalOperation
}

func NewHaving(operation ComparisonOrLogicalOperation) *Having {
	return &Having{
		operation: operation,
	}
}

func (h *Having) OrderBy(orders ...Order) *OrderBy {
	o := NewOrderBy(orders...)
	Link(h, o)
	return o
}

func (h *Having) Limit(count int) *Limit {
	l := NewLimit(count)
	Link(h, l)
	return l
}

func (h *Having) nodeize() (tokenizer.Tokenizer, []interface{}) {
	middle, values := h.operation.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Having)),
	).SetMiddle(
		middle,
	), values
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
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

func (h Having) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizeClauses(h)
}

func (h Having) self() (tokenizer.Tokenizer, []interface{}) {
	middle, values := h.operation.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Having)),
	).SetMiddle(
		middle,
	), values
}

func (h Having) previous() Clause {
	return h.prev
}

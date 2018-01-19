package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Having struct {
	prev      Prever
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
	return nodeizePrevs(h)
}

func (h Having) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	middle, values := h.operation.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Having)),
	).SetMiddle(
		middle,
	), values
}

func (h Having) previous() Prever {
	return h.prev
}

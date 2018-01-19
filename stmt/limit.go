package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Limit struct {
	prev  Prever
	count int
}

func NewLimit(count int) Limit {
	return Limit{
		count: count,
	}
}

func (l Limit) Offset(count int) Offset {
	o := NewOffset(count)
	o.prev = l
	return o
}

func (l Limit) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(l)
}

func (l Limit) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	line, values := tokenizer.ParamsToLine(l.count)
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Limit)),
	).SetMiddle(
		line,
	), values
}

func (l Limit) previous() Prever {
	return l.prev
}

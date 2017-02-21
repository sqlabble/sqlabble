package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Limit struct {
	prev  Clause
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

func (l Limit) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(l)
}

func (l Limit) self() (token.Tokenizer, []interface{}) {
	line, values := token.ParamsToLine(l.count)
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Limit)),
	).SetMiddle(
		line,
	), values
}

func (l Limit) previous() Clause {
	return l.prev
}

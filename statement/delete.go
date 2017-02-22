package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}

func (d Delete) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(d)
}

func (d Delete) self() (token.Tokenizer, []interface{}) {
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Delete)),
	), nil
}

func (d Delete) previous() Clause {
	return nil
}

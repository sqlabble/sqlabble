package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Delete struct {
	Prev
}

func NewDelete() *Delete {
	return &Delete{}
}

func (d *Delete) From(t *Table) *From {
	f := NewFrom(t)
	Link(d, f)
	return f
}

func (d *Delete) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Delete)),
	), nil
}

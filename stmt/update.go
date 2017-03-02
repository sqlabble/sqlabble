package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Update struct {
	prev  Prever
	table Table
}

func NewUpdate(table Table) Update {
	return Update{
		table: table,
	}
}

func (u Update) Set(assigns ...Assign) Set {
	s := NewSet(assigns...)
	s.prev = u
	return s
}

func (u Update) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(u)
}

func (u Update) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	middle, values := u.table.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Update)),
	).SetMiddle(
		middle,
	), values
}

func (u Update) previous() Prever {
	return u.prev
}

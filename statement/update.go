package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Update struct {
	prev  Clause
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
	return nodeizeClauses(u)
}

func (u Update) self() (tokenizer.Tokenizer, []interface{}) {
	middle, values := u.table.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Update)),
	).SetMiddle(
		middle,
	), values
}

func (u Update) previous() Clause {
	return u.prev
}

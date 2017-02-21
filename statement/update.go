package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
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

func (u Update) nodeize() (token.Tokenizer, []interface{}) {
	return u.container()
}

func (u Update) container() (token.Container, []interface{}) {
	middle, values := u.table.nodeize()
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Update)),
	).SetMiddle(
		middle,
	), values
}

func (u Update) previous() Clause {
	return u.prev
}

func (u Update) Set(assigns ...Assign) Set {
	s := NewSet(assigns...)
	s.prev = u
	return s
}

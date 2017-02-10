package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Update struct {
	prev  grammar.Clause
	table Table
}

func NewUpdate(table Table) Update {
	return Update{
		table: table,
	}
}

func (u Update) Generator() generator.Generator {
	return u.Container()
}

func (u Update) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(keyword.Update),
		u.table.Generator(),
	)
}

func (u Update) Prev() grammar.Clause {
	return u.prev
}

func (u Update) Set(assigns ...Assign) Set {
	s := NewSet(assigns...)
	s.prev = u
	return s
}

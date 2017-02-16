package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
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

func (u Update) node() generator.Node {
	return u.myNode()
}

func (u Update) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Update),
		u.table.node(),
	)
}

func (u Update) previous() Clause {
	return u.prev
}

func (u Update) Set(assigns ...Assign) Set {
	s := NewSet(assigns...)
	s.prev = u
	return s
}

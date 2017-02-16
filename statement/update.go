package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type update struct {
	prev clause
	tbl  table
}

func NewUpdate(tbl table) update {
	return update{
		tbl: tbl,
	}
}

func (u update) node() generator.Node {
	return u.myNode()
}

func (u update) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Update),
		u.tbl.node(),
	)
}

func (u update) previous() clause {
	return u.prev
}

func (u update) Set(assigns ...assign) set {
	s := NewSet(assigns...)
	s.prev = u
	return s
}

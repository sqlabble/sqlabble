package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type update struct {
	prev clause
	tbl  table
}

func newUpdate(tbl table) update {
	return update{
		tbl: tbl,
	}
}

func (u update) node() generator.Node {
	return u.nodeMine()
}

func (u update) nodeMine() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Update),
		u.tbl.node(),
	)
}

func (u update) previous() clause {
	return u.prev
}

func (u update) Set(assigns ...assign) set {
	s := newSet(assigns...)
	s.prev = u
	return s
}

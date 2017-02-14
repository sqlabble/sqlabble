package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type update struct {
	prev clauseNode
	tbl  table
}

func newUpdate(tbl table) update {
	return update{
		tbl: tbl,
	}
}

func (u update) generator() generator.Generator {
	return u.clauseGenerator()
}

func (u update) clauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(keyword.Update),
		u.tbl.generator(),
	)
}

func (u update) previous() clauseNode {
	return u.prev
}

func (u update) Set(assigns ...assign) set {
	s := newSet(assigns...)
	s.prev = u
	return s
}

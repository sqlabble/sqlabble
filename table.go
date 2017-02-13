package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
)

type table struct {
	name string
}

func newTable(name string) table {
	return table{
		name: name,
	}
}

func (t table) generator() generator.Generator {
	ts := tableNodes(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewGenerators(fs...)
}

func (t table) expression() generator.Expression {
	return generator.NewExpression(
		t.TableName(),
	)
}

func (t table) TableName() string {
	return t.name
}

func (t table) As(alias string) tableAs {
	return tableAs{
		table: t,
		alias: alias,
	}
}

func (t table) previous() tableNode {
	return nil
}

func (t table) Join(table tableNode) tableNode {
	nj := newJoin(table)
	nj.prev = t
	return nj
}

func (t table) InnerJoin(table tableNode) tableNode {
	ij := newInnerJoin(table)
	ij.prev = t
	return ij
}

func (t table) LeftJoin(table tableNode) tableNode {
	lj := newLeftJoin(table)
	lj.prev = t
	return lj
}

func (t table) RightJoin(table tableNode) tableNode {
	rj := newRightJoin(table)
	rj.prev = t
	return rj
}

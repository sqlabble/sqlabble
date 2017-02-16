package statement

import (
	"github.com/minodisk/sqlabble/generator"
)

type table struct {
	name string
}

func NewTable(name string) table {
	return table{
		name: name,
	}
}

func (t table) node() generator.Node {
	ts := tableNodes(t)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
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

func (t table) previous() joiner {
	return nil
}

func (t table) Join(table joiner) joiner {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t table) InnerJoin(table joiner) joiner {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t table) LeftJoin(table joiner) joiner {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t table) RightJoin(table joiner) joiner {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

package statement

import (
	"github.com/minodisk/sqlabble/generator"
)

type Table struct {
	name string
}

func NewTable(name string) Table {
	return Table{
		name: name,
	}
}

func (t Table) node() generator.Node {
	ts := tableNodes(t)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
}

func (t Table) expression() generator.Expression {
	return generator.NewExpression(
		t.TableName(),
	)
}

func (t Table) TableName() string {
	return t.name
}

func (t Table) As(alias string) TableAs {
	return TableAs{
		table: t,
		alias: alias,
	}
}

func (t Table) previous() Joiner {
	return nil
}

func (t Table) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t Table) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t Table) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t Table) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

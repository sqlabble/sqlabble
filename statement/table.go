package statement

import "github.com/minodisk/sqlabble/node"

type Table struct {
	name string
}

func NewTable(name string) Table {
	return Table{
		name: name,
	}
}

func (t Table) node() node.Node {
	ts := tableNodes(t)
	ns := make([]node.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return node.NewNodes(ns...)
}

func (t Table) expression() node.Expression {
	return node.NewExpression(t.name)
}

func (t Table) previous() Joiner {
	return nil
}

func (t Table) As(alias string) TableAs {
	return TableAs{
		table: t,
		alias: alias,
	}
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

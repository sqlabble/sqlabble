package statement

import (
	"fmt"

	"github.com/minodisk/sqlabble/generator"
)

type TableAs struct {
	table Table
	alias string
}

func (t TableAs) node() generator.Node {
	ts := tableNodes(t)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
}

func (t TableAs) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("%s AS %s", t.TableName(), t.alias),
	)
}

func (t TableAs) TableName() string {
	return t.table.name
}

func (t TableAs) previous() Joiner {
	return nil
}

func (t TableAs) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t TableAs) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t TableAs) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t TableAs) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

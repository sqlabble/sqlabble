package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
)

type tableAs struct {
	table table
	alias string
}

func (t tableAs) node() generator.Node {
	ts := tableNodes(t)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
}

func (t tableAs) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("%s AS %s", t.TableName(), t.alias),
	)
}

func (t tableAs) TableName() string {
	return t.table.name
}

func (t tableAs) previous() joiner {
	return nil
}

func (t tableAs) Join(table joiner) joiner {
	nj := newJoin(table)
	nj.prev = t
	return nj
}

func (t tableAs) InnerJoin(table joiner) joiner {
	ij := newInnerJoin(table)
	ij.prev = t
	return ij
}

func (t tableAs) LeftJoin(table joiner) joiner {
	lj := newLeftJoin(table)
	lj.prev = t
	return lj
}

func (t tableAs) RightJoin(table joiner) joiner {
	rj := newRightJoin(table)
	rj.prev = t
	return rj
}

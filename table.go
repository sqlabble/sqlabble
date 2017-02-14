package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
)

type table struct {
	name string
}

func newTable(name string) table {
	return table{
		name: name,
	}
}

func (t table) node() generator.Node {
	ts := tableNodes(t)
	fs := make([]generator.Node, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewNodes(fs...)
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

func (t table) previous() tableOrTableAs {
	return nil
}

func (t table) Join(table tableOrTableAs) tableOrTableAs {
	nj := newJoin(table)
	nj.prev = t
	return nj
}

func (t table) InnerJoin(table tableOrTableAs) tableOrTableAs {
	ij := newInnerJoin(table)
	ij.prev = t
	return ij
}

func (t table) LeftJoin(table tableOrTableAs) tableOrTableAs {
	lj := newLeftJoin(table)
	lj.prev = t
	return lj
}

func (t table) RightJoin(table tableOrTableAs) tableOrTableAs {
	rj := newRightJoin(table)
	rj.prev = t
	return rj
}

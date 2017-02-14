package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/generator"
)

type tableAs struct {
	table table
	alias string
}

func (t tableAs) node() generator.Node {
	ts := tableNodes(t)
	fs := make([]generator.Node, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewNodes(fs...)
}

func (t tableAs) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("%s AS %s", t.TableName(), t.alias),
	)
}

func (t tableAs) TableName() string {
	return t.table.name
}

func (t tableAs) previous() tableOrTableAs {
	return nil
}

func (t tableAs) Join(table tableOrTableAs) tableOrTableAs {
	nj := newJoin(table)
	nj.prev = t
	return nj
}

func (t tableAs) InnerJoin(table tableOrTableAs) tableOrTableAs {
	ij := newInnerJoin(table)
	ij.prev = t
	return ij
}

func (t tableAs) LeftJoin(table tableOrTableAs) tableOrTableAs {
	lj := newLeftJoin(table)
	lj.prev = t
	return lj
}

func (t tableAs) RightJoin(table tableOrTableAs) tableOrTableAs {
	rj := newRightJoin(table)
	rj.prev = t
	return rj
}

package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type join struct {
	joinType string
	table    tableOrTableAs
	prev     tableOrTableAs
}

func newJoin(table tableOrTableAs) join {
	return join{
		joinType: keyword.Join,
		table:    table,
	}
}

func newInnerJoin(table tableOrTableAs) join {
	return join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func newLeftJoin(table tableOrTableAs) join {
	return join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func newRightJoin(table tableOrTableAs) join {
	return join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (t join) generator() generator.Generator {
	ts := tableNodes(t)
	es := make([]generator.Generator, len(ts))
	for i, t := range ts {
		es[i] = t.expression()
	}
	return generator.NewGenerators(es...)
}

func (t join) expression() generator.Expression {
	return generator.NewExpression(t.joinType).
		Append(t.table.expression())
}

func (j join) previous() tableOrTableAs {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j join) Join(table tableOrTableAs) tableOrTableAs {
	nj := newJoin(table)
	nj.prev = j
	return nj
}

func (j join) InnerJoin(table tableOrTableAs) tableOrTableAs {
	ij := newInnerJoin(table)
	ij.prev = j
	return ij
}

func (j join) LeftJoin(table tableOrTableAs) tableOrTableAs {
	lj := newLeftJoin(table)
	lj.prev = j
	return lj
}

func (j join) RightJoin(table tableOrTableAs) tableOrTableAs {
	rj := newRightJoin(table)
	rj.prev = j
	return rj
}

func (j join) On(column1, column2 column) on {
	o := newOn(column1, column2)
	o.join = j
	return o
}

func (j join) Using(col column) using {
	o := newUsing(col)
	o.join = j
	return o
}

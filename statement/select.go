package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type Select struct {
	distinct bool
	columns  []ColumnOrColumnAs
}

func NewSelect(columns ...ColumnOrColumnAs) Select {
	return Select{
		distinct: false,
		columns:  columns,
	}
}

func NewSelectDistinct(columns ...ColumnOrColumnAs) Select {
	return Select{
		distinct: true,
		columns:  columns,
	}
}

func (s Select) node() node.Node {
	cs := clauseNodes(s)
	fs := make([]node.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return node.NewNodes(fs...)
}

func (s Select) myNode() node.Node {
	fs := make([]node.Node, len(s.columns))
	for i, c := range s.columns {
		fs[i] = c.node()
	}
	k := node.NewExpression(keyword.Select)
	if s.distinct {
		k = k.Append(node.NewExpression(keyword.Distinct))
	}
	return node.NewContainer(
		k,
		node.NewComma(fs...),
	)
}

func (s Select) previous() Clause {
	return nil
}

func (s Select) From(t Joiner) From {
	f := NewFrom(t)
	f.prev = s
	return f
}

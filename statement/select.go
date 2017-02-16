package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
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

func (s Select) node() generator.Node {
	cs := clauseNodes(s)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewNodes(fs...)
}

func (s Select) myNode() generator.Node {
	fs := make([]generator.Node, len(s.columns))
	for i, c := range s.columns {
		fs[i] = c.node()
	}
	k := generator.NewExpression(keyword.Select)
	if s.distinct {
		k = k.Append(generator.NewExpression(keyword.Distinct))
	}
	return generator.NewContainer(
		k,
		generator.NewComma(fs...),
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

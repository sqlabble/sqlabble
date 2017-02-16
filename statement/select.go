package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type selec struct {
	distinct bool
	columns  []columnOrColumnAs
}

func NewSelect(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: false,
		columns:  columns,
	}
}

func NewSelectDistinct(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: true,
		columns:  columns,
	}
}

func (s selec) node() generator.Node {
	cs := clauseNodes(s)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewNodes(fs...)
}

func (s selec) myNode() generator.Node {
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

func (s selec) previous() clause {
	return nil
}

func (s selec) From(t joiner) from {
	f := NewFrom(t)
	f.prev = s
	return f
}

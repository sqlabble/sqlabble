package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type del struct{}

func NewDelete() del {
	return del{}
}

func (d del) node() generator.Node {
	cs := clauseNodes(d)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewNodes(fs...)
}

func (d del) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Delete),
	)
}

func (d del) previous() clause {
	return nil
}

func (d del) From(t table) from {
	f := NewFrom(t)
	f.prev = d
	return f
}

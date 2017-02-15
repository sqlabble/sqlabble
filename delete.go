package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type del struct{}

func newDelete() del {
	return del{}
}

func (d del) node() generator.Node {
	cs := clauseNodes(d)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewParallelNodes(fs...)
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
	f := newFrom(t)
	f.prev = d
	return f
}

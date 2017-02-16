package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) node() generator.Node {
	cs := clauseNodes(d)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewNodes(fs...)
}

func (d Delete) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Delete),
	)
}

func (d Delete) previous() Clause {
	return nil
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}

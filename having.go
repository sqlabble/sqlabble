package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type having struct {
	prev      clause
	operation comparisonOrLogicalOperation
}

func newHaving(operation comparisonOrLogicalOperation) having {
	return having{
		operation: operation,
	}
}

func (w having) node() generator.Node {
	cs := clauseNodes(w)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.nodeMine()
	}
	return generator.NewNodes(fs...)
}

func (w having) nodeMine() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		w.operation.node(),
	)
}

func (c having) previous() clause {
	return c.prev
}

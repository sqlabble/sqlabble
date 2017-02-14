package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
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
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (w having) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		w.operation.node(),
	)
}

func (c having) previous() clause {
	return c.prev
}

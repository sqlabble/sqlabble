package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type having struct {
	prev      clauseNode
	operation operationNode
}

func newHaving(operation operationNode) having {
	return having{
		operation: operation,
	}
}

func (w having) generator() generator.Generator {
	cs := clauseNodes(w)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (w having) clauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		w.operation.generator(),
	)
}

func (c having) previous() clauseNode {
	return c.prev
}

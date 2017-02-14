package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type del struct{}

func newDelete() del {
	return del{}
}

func (d del) generator() generator.Generator {
	cs := clauseNodes(d)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (d del) clauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(keyword.Delete),
	)
}

func (d del) previous() clauseNode {
	return nil
}

func (d del) From(t table) from {
	f := newFrom(t)
	f.prev = d
	return f
}

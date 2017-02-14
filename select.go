package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type selec struct {
	distinct bool
	columns  []columnOrColumnAs
}

func newSelect(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: false,
		columns:  columns,
	}
}

func newSelectDistinct(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: true,
		columns:  columns,
	}
}

func (s selec) generator() generator.Generator {
	cs := clauseNodes(s)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (s selec) clauseGenerator() generator.Generator {
	fs := make([]generator.Generator, len(s.columns))
	for i, c := range s.columns {
		fs[i] = c.generator()
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

func (c selec) previous() clause {
	return nil
}

func (c selec) From(t tableOrTableAs) from {
	f := newFrom(t)
	f.prev = c
	return f
}

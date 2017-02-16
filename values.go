package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type values struct {
	prevClause clause
	prev       vals
	vals       []interface{}
}

func newValues(vals ...interface{}) values {
	return values{
		vals: vals,
	}
}

func (v values) node() generator.Node {
	vs := valuesNodes(v)
	ns := make([]generator.Node, len(vs))
	for i, v := range vs {
		ns[i] = v.expression()
	}
	g := generator.NewContainer(
		generator.NewExpression(keyword.Values),
		generator.NewComma(ns...),
	)

	if len(vs) > 0 {
		if c := vs[0].clause(); c != nil {
			return generator.NewParallelNodes(
				c.node(),
				g,
			)
		}
	}

	return g
}

func (v values) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("(%s)", placeholders(len(v.vals))),
		v.vals...,
	)
}

func (v values) clause() clause {
	return v.prevClause
}

func (v values) previous() vals {
	return v.prev
}

func (v values) Values(vals ...interface{}) values {
	f := newValues(vals...)
	f.prev = v
	return f
}

type defaultValues struct {
	prev clause
}

func newDefaultValues() defaultValues {
	return defaultValues{}
}

func (v defaultValues) node() generator.Node {
	cs := clauseNodes(v)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewParallelNodes(ns...)
}

func (v defaultValues) myNode() generator.Node {
	return generator.NewExpression(keyword.DefaultValues)
}

func (v defaultValues) previous() clause {
	return v.prev
}

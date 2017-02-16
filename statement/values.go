package statement

import (
	"fmt"

	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Values struct {
	prevClause Clause
	prev       Vals
	values     []interface{}
}

func NewValues(values ...interface{}) Values {
	return Values{
		values: values,
	}
}

func (v Values) node() generator.Node {
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
			return generator.NewNodes(
				c.node(),
				g,
			)
		}
	}

	return g
}

func (v Values) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("(%s)", placeholders(len(v.values))),
		v.values...,
	)
}

func (v Values) clause() Clause {
	return v.prevClause
}

func (v Values) previous() Vals {
	return v.prev
}

func (v Values) Values(vals ...interface{}) Values {
	f := NewValues(vals...)
	f.prev = v
	return f
}

type DefaultValues struct {
	prev Clause
}

func NewDefaultValues() DefaultValues {
	return DefaultValues{}
}

func (v DefaultValues) node() generator.Node {
	cs := clauseNodes(v)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (v DefaultValues) myNode() generator.Node {
	return generator.NewExpression(keyword.DefaultValues)
}

func (v DefaultValues) previous() Clause {
	return v.prev
}

package statement

import (
	"fmt"

	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (v Values) node() node.Node {
	vs := valuesNodes(v)
	ns := make([]node.Node, len(vs))
	for i, v := range vs {
		ns[i] = v.expression()
	}
	g := node.NewContainer(
		node.NewExpression(keyword.Values),
		node.NewComma(ns...),
	)

	if len(vs) > 0 {
		if c := vs[0].clause(); c != nil {
			return node.NewNodes(
				c.node(),
				g,
			)
		}
	}

	return g
}

func (v Values) expression() node.Expression {
	return node.NewExpression(
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

func (v DefaultValues) node() node.Node {
	cs := clauseNodes(v)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (v DefaultValues) myNode() node.Node {
	return node.NewExpression(keyword.DefaultValues)
}

func (v DefaultValues) previous() Clause {
	return v.prev
}

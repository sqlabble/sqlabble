package token

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Values struct {
	clause grammar.Clause
	prev   grammar.Values
	values []interface{}
}

func NewValues(values ...interface{}) Values {
	return Values{
		values: values,
	}
}

func (v Values) Generator() generator.Generator {
	vs := grammar.PrevValues(v)
	es := make([]generator.Generator, len(vs))
	for i, v := range vs {
		es[i] = v.Expression()
	}
	g := generator.NewContainer(
		generator.NewExpression(keyword.Values),
		generator.NewComma(es...),
	)

	if len(vs) > 0 {
		if c := vs[0].Clause(); c != nil {
			return generator.NewGenerators(
				c.Generator(),
				g,
			)
		}
	}

	return g
}

func (v Values) Expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("(%s)", Placeholders(len(v.values))),
		v.values...,
	)
}

func (v Values) Clause() grammar.Clause {
	return v.clause
}

func (v Values) Prev() grammar.Values {
	return v.prev
}

func (v Values) Values(values ...interface{}) Values {
	f := NewValues(values...)
	f.prev = v
	return f
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
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

func (v Values) nodeize() (token.Tokenizer, []interface{}) {
	vss := valuesNodes(v)
	lines := make(token.Lines, len(vss))
	values := []interface{}{}
	for i, vs := range vss {
		var vals []interface{}
		lines[i], vals = vs.line()
		values = append(values, vals...)
	}
	c := token.NewContainer(
		token.NewLine(token.Word(keyword.Values)),
	).SetMiddle(
		lines.Prefix(token.Comma, token.Space),
	)
	return c, values
}

func (v Values) line() (token.Line, []interface{}) {
	line, values := token.ParamsToLine(v.values...)
	return line.
		P(token.ParenthesesStart).
		A(token.ParenthesesEnd), values
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

func (v DefaultValues) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(v)
	ts := make(token.Tokenizers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		ts[i], vals = c.container()
		values = append(values, vals...)
	}
	return ts, values
}

func (v DefaultValues) container() (token.Container, []interface{}) {
	return token.NewContainer(
		token.NewLine(token.Word(keyword.DefaultValues)),
	), nil
}

func (v DefaultValues) previous() Clause {
	return v.prev
}

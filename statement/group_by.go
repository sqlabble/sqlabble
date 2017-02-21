package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type GroupBy struct {
	prev    Clause
	columns []Column
}

func NewGroupBy(column Column, columns ...Column) GroupBy {
	return GroupBy{
		columns: append([]Column{column}, columns...),
	}
}

func (g GroupBy) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(g)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (g GroupBy) container() (token.Container, []interface{}) {
	lines := make(token.Lines, len(g.columns))
	values := []interface{}{}
	for i, c := range g.columns {
		var vals []interface{}
		lines[i], vals = c.line()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.GroupBy)),
	).SetMiddle(
		lines.Prefix(
			token.Comma,
			token.Space,
		),
	), values
}

func (g GroupBy) previous() Clause {
	return g.prev
}

func (g GroupBy) Having(operation ComparisonOrLogicalOperation) Having {
	l := NewHaving(operation)
	l.prev = g
	return l
}

func (g GroupBy) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = g
	return o
}

func (g GroupBy) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = g
	return l
}

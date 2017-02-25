package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type GroupBy struct {
	Prev
	Next
	columns []Column
}

func NewGroupBy(column Column, columns ...Column) *GroupBy {
	cols := append([]Column{column}, columns...)
	return &GroupBy{
		columns: cols,
	}
}

func (g *GroupBy) Having(operation ComparisonOrLogicalOperation) *Having {
	l := NewHaving(operation)
	Link(g, l)
	return l
}

func (g *GroupBy) OrderBy(orders ...Order) *OrderBy {
	o := NewOrderBy(orders...)
	Link(g, o)
	return o
}

func (g *GroupBy) Limit(count int) *Limit {
	l := NewLimit(count)
	Link(g, l)
	return l
}

func (g *GroupBy) nodeize() (tokenizer.Tokenizer, []interface{}) {
	ts := make(tokenizer.Tokenizers, len(g.columns))
	values := []interface{}{}
	for i, c := range g.columns {
		var vals []interface{}
		ts[i], vals = c.nodeize()
		values = append(values, vals...)
	}
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.GroupBy)),
	).SetMiddle(
		ts.Prefix(
			token.Comma,
		),
	), values
}

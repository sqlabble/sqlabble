package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type GroupBy struct {
	prev    Prever
	columns []Column
}

func NewGroupBy(column Column, columns ...Column) GroupBy {
	return GroupBy{
		columns: append([]Column{column}, columns...),
	}
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

func (g GroupBy) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(g)
}

func (g GroupBy) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
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

func (g GroupBy) previous() Prever {
	return g.prev
}

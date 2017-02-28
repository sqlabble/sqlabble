package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type From struct {
	Prev
	Next
	table TableOrAliasOrJoiner
}

func NewFrom(table TableOrAliasOrJoiner) *From {
	return &From{
		table: table,
	}
}

func (f *From) Where(op ComparisonOrLogicalOperation) *Where {
	w := NewWhere(op)
	Link(f, w)
	return w
}

func (f *From) GroupBy(col Column, columns ...Column) *GroupBy {
	g := NewGroupBy(col, columns...)
	Link(f, g)
	return g
}

func (f *From) OrderBy(orders ...Order) *OrderBy {
	o := NewOrderBy(orders...)
	Link(f, o)
	return o
}

func (f *From) Limit(count int) *Limit {
	l := NewLimit(count)
	Link(f, l)
	return l
}

func (f *From) nodeize() (tokenizer.Tokenizer, []interface{}) {
	middle, values := Nodize(f.table)
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.From)),
	).SetMiddle(
		middle,
	), values
}

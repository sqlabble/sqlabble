package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type From struct {
	prev  Prever
	table TableOrAliasOrJoiner
}

func NewFrom(table TableOrAliasOrJoiner) From {
	return From{
		table: table,
	}
}

func (f From) Where(op ComparisonOrLogicalOperation) Where {
	w := NewWhere(op)
	w.prev = f
	return w
}

func (f From) GroupBy(col Column, columns ...Column) GroupBy {
	g := NewGroupBy(col, columns...)
	g.prev = f
	return g
}

func (f From) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = f
	return o
}

func (f From) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = f
	return l
}

func (f From) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(f)
}

func (f From) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	middle, values := f.table.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.From)),
	).SetMiddle(
		middle,
	), values
}

func (f From) previous() Prever {
	return f.prev
}

package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type On struct {
	Prev
	Next
	column1, column2 Column
}

func NewOn(column1, column2 Column) *On {
	return &On{
		column1: column1,
		column2: column2,
	}
}

func (o *On) Join(table TableOrAlias) *Join {
	j := NewJoin(table)
	Link(o, j)
	return j
}

func (o *On) InnerJoin(table TableOrAlias) *Join {
	j := NewInnerJoin(table)
	Link(o, j)
	return j
}

func (o *On) LeftJoin(table TableOrAlias) *Join {
	j := NewLeftJoin(table)
	Link(o, j)
	return j
}

func (o *On) RightJoin(table TableOrAlias) *Join {
	j := NewRightJoin(table)
	Link(o, j)
	return j
}

func (o *On) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := o.column1.nodeize()
	t2, v2 := o.column2.nodeize()
	return tokenizer.ConcatTokenizers(
			t1,
			t2,
			tokenizer.NewLine(
				token.Word(keyword.Eq),
			),
		).Prepend(token.Word(keyword.On)),
		append(v1, v2...)
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (o *On) isTableOrAliasOrJoiner() bool {
	return true
}

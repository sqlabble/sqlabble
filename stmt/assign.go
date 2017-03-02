package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Assign struct {
	column Column
	param  ValOrFuncOrSub
}

func NewAssign(column Column, param ValOrFuncOrSub) Assign {
	return Assign{
		column: column,
		param:  param,
	}
}

func (a Assign) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := a.column.nodeize()
	t2, v2 := a.param.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(token.Word(keyword.Eq)),
	), append(v1, v2...)
}

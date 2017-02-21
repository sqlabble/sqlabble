package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type Assign struct {
	column Column
	param  interface{}
}

func NewAssign(column Column, param ParamOrSubquery) Assign {
	return Assign{
		column: column,
		param:  param,
	}
}

func (a Assign) nodeize() (token.Tokenizer, []interface{}) {
	return a.line()
}

func (a Assign) line() (token.Line, []interface{}) {
	line, values := a.column.line()
	return line.
		A(
			token.Space,
			token.Word(operator.Eq),
			token.Space,
			token.Placeholder,
		), values
}

func (a Assign) values() []interface{} {
	return []interface{}{a.param}
}

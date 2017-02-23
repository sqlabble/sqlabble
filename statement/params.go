package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Param struct {
	value interface{}
}

func NewParam(value interface{}) Param {
	return Param{
		value: value,
	}
}

func (p Param) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.ParamsToLine(p.value)
}

// isValOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (p Param) isValOrFuncOrSub() bool {
	return true
}

// isValOrColOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrColOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (p Param) isValOrColOrFuncOrSub() bool {
	return true
}

type Params []Param

func NewParams(values ...interface{}) Params {
	ps := make(Params, len(values))
	for i, value := range values {
		ps[i] = NewParam(value)
	}
	return ps
}

func (ps Params) nodeize() (tokenizer.Tokenizer, []interface{}) {
	values := make([]interface{}, len(ps))
	for i, p := range ps {
		values[i] = p.value
	}
	line, values := tokenizer.ParamsToLine(values...)
	return line.
		Prepend(token.ParenthesesStart).
		Append(token.ParenthesesEnd), values
}

// isParamsOrSubquery always returns true.
// This method exists only to implement the interface ParamsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (vs Params) isValsOrSub() bool {
	return true
}

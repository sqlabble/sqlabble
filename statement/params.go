package statement

import "github.com/minodisk/sqlabble/token"

type Param struct {
	value interface{}
}

func NewParam(value interface{}) Param {
	return Param{
		value: value,
	}
}

func (p Param) nodeize() (token.Tokenizer, []interface{}) {
	return token.ParamsToLine(p.value)
}

// isParamOrSubquery always returns true.
// This method exists only to implement the interface ParamOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (p Param) isParamOrSubquery() bool {
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

func (ps Params) nodeize() (token.Tokenizer, []interface{}) {
	values := make([]interface{}, len(ps))
	for i, p := range ps {
		values[i] = p.value
	}
	line, values := token.ParamsToLine(values...)
	return line.
		Prepend(token.ParenthesesStart).
		Append(token.ParenthesesEnd), values
}

// isParamsOrSubquery always returns true.
// This method exists only to implement the interface ParamsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (vs Params) isParamsOrSubquery() bool {
	return true
}

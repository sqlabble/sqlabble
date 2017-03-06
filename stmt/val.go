package stmt

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Val struct {
	value interface{}
}

func NewVal(value interface{}) *Val {
	return &Val{
		value: value,
	}
}

func (p *Val) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.ParamsToLine(p.value)
}

// isValOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (p *Val) isValOrFuncOrSub() bool {
	return true
}

// isValOrColOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrColOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (p *Val) isValOrColOrFuncOrSub() bool {
	return true
}

type Vals []*Val

func NewVals(values ...interface{}) Vals {
	ps := make(Vals, len(values))
	for i, value := range values {
		ps[i] = NewVal(value)
	}
	return ps
}

func (ps Vals) nodeize() (tokenizer.Tokenizer, []interface{}) {
	values := make([]interface{}, len(ps))
	for i, p := range ps {
		values[i] = p.value
	}
	line, values := tokenizer.ParamsToLine(values...)
	return line.
		Prepend(token.LParen).
		Append(token.RParen), values
}

// isParamsOrSubquery always returns true.
// This method exists only to implement the interface ParamsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (vs Vals) isValsOrSub() bool {
	return true
}

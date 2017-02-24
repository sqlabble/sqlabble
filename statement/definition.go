package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Definition struct {
	Child
	column     Column
	definition string
}

func NewDefinition(definition string) *Definition {
	return &Definition{
		definition: definition,
	}
}

func (d Definition) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, values := d.column.nodeize()
	return t.
		Append(
			token.Space,
			token.Word(d.definition),
		), values
}

type Definitions struct {
	Child
	Parent
}

func NewDefinitions(defs ...*Definition) *Definitions {
	ds := &Definitions{}
	for _, d := range defs {
		Contract(ds, d)
	}
	return ds
}

func (d Definitions) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewParentheses(nil), nil
}

func (d Definitions) separator() token.Tokens {
	return token.NewTokens(
		token.Comma,
		token.Space,
	)
}

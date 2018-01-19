package stmt

import (
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Definition struct {
	column     Column
	definition string
}

func NewDefinition(definition string) Definition {
	return Definition{
		definition: definition,
	}
}

func (d Definition) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, values := d.column.nodeize()
	return t.
		Append(
			token.Word(d.definition),
		), values
}

type Definitions struct {
	createTable CreateTable
	definitions []Definition
}

func NewDefinitions(definitions ...Definition) Definitions {
	return Definitions{
		definitions: definitions,
	}
}

func (ds Definitions) nodeize() (tokenizer.Tokenizer, []interface{}) {
	ts := make(tokenizer.Tokenizers, len(ds.definitions))
	values := []interface{}{}
	for i, d := range ds.definitions {
		var vals []interface{}
		ts[i], vals = d.nodeize()
		values = append(values, vals...)
	}
	ts = ts.Prefix(
		token.Comma,
	)

	c, values := ds.createTable.container()
	middle := c.Middle()
	def := tokenizer.NewParentheses(ts)

	return c.SetMiddle(
		tokenizer.ConcatTokenizers(
			middle,
			def,
			tokenizer.EmptyLine,
		),
	), values
}

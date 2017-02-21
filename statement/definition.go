package statement

import "github.com/minodisk/sqlabble/token"

type Definition struct {
	column     Column
	definition string
}

func NewDefinition(definition string) Definition {
	return Definition{
		definition: definition,
	}
}

func (d Definition) nodeize() (token.Tokenizer, []interface{}) {
	t, values := d.column.nodeize()
	return t.
		Append(
			token.Space,
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

func (ds Definitions) nodeize() (token.Tokenizer, []interface{}) {
	ts := make(token.Tokenizers, len(ds.definitions))
	values := []interface{}{}
	for i, d := range ds.definitions {
		var vals []interface{}
		ts[i], vals = d.nodeize()
		values = append(values, vals...)
	}
	ts = ts.Prefix(
		token.Comma,
		token.Space,
	)

	c, values := ds.createTable.container()
	middle := c.Middle()
	def := token.NewParentheses(ts)

	return c.SetMiddle(
		token.ConcatTokenizers(
			middle,
			def,
			token.NewLine(token.Space),
		),
	), values
}

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
	return d.line()
}

func (d Definition) line() (token.Line, []interface{}) {
	line, values := d.column.line()
	return line.
		A(
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
	lines := make(token.Lines, len(ds.definitions))
	values := []interface{}{}
	for i, d := range ds.definitions {
		var vals []interface{}
		lines[i], vals = d.line()
		values = append(values, vals...)
	}
	lines = lines.Prefix(
		token.Comma,
		token.Space,
	)

	c, values := ds.createTable.container()
	middle := c.Middle()
	def := token.WrapParenthesesLines(
		lines...,
	)

	return c.SetMiddle(
		token.ConcatTokenizers(
			middle,
			def,
			token.NewLine(token.Space),
		),
	), values
}

package token

import "github.com/minodisk/sqlabble/internal/generator"

type Definition struct {
	column     Column
	definition string
}

func NewDefinition(definition string) Definition {
	return Definition{
		definition: definition,
	}
}

func (d Definition) Generator() generator.Generator {
	return d.column.Expression().
		Append(generator.NewExpression(d.definition))
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

func (ds Definitions) Generator() generator.Generator {
	p := ds.createTable.Container()
	gs := make([]generator.Generator, len(ds.definitions))
	for i, d := range ds.definitions {
		gs[i] = d.Generator()
	}
	return p.AddChild(
		generator.NewBracket(
			generator.NewComma(
				gs...,
			),
		),
	)
}

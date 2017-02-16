package statement

import "github.com/minodisk/sqlabble/generator"

type Definition struct {
	column     Column
	definition string
}

func NewDefinition(definition string) Definition {
	return Definition{
		definition: definition,
	}
}

func (d Definition) node() generator.Node {
	return d.column.expression().
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

func (ds Definitions) node() generator.Node {
	p := ds.createTable.container()
	gs := make([]generator.Node, len(ds.definitions))
	for i, d := range ds.definitions {
		gs[i] = d.node()
	}
	return p.AddChild(
		generator.NewParentheses(
			generator.NewComma(
				gs...,
			),
		),
	)
}

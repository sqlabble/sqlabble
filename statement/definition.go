package statement

import "github.com/minodisk/sqlabble/node"

type Definition struct {
	column     Column
	definition string
}

func NewDefinition(definition string) Definition {
	return Definition{
		definition: definition,
	}
}

func (d Definition) node() node.Node {
	return d.column.expression().
		Append(node.NewExpression(d.definition))
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

func (ds Definitions) node() node.Node {
	p := ds.createTable.container()
	gs := make([]node.Node, len(ds.definitions))
	for i, d := range ds.definitions {
		gs[i] = d.node()
	}
	return p.AddChild(
		node.NewParentheses(
			node.NewComma(
				gs...,
			),
		),
	)
}

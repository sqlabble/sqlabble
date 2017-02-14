package sqlabble

import "github.com/minodisk/sqlabble/generator"

type definition struct {
	col column
	def string
}

func newDefinition(def string) definition {
	return definition{
		def: def,
	}
}

func (d definition) node() generator.Node {
	return d.col.expression().
		Append(generator.NewExpression(d.def))
}

type definitions struct {
	createTable createTable
	defs        []definition
}

func newDefinitions(defs ...definition) definitions {
	return definitions{
		defs: defs,
	}
}

func (ds definitions) node() generator.Node {
	p := ds.createTable.container()
	gs := make([]generator.Node, len(ds.defs))
	for i, d := range ds.defs {
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

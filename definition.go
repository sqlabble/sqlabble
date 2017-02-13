package sqlabble

import "github.com/minodisk/sqlabble/internal/generator"

type definition struct {
	col column
	def string
}

func newDefinition(def string) definition {
	return definition{
		def: def,
	}
}

func (d definition) generator() generator.Generator {
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

func (ds definitions) generator() generator.Generator {
	p := ds.createTable.Container()
	gs := make([]generator.Generator, len(ds.defs))
	for i, d := range ds.defs {
		gs[i] = d.generator()
	}
	return p.AddChild(
		generator.NewBracket(
			generator.NewComma(
				gs...,
			),
		),
	)
}

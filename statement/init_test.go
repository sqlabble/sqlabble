package statement_test

import "github.com/minodisk/sqlabble/builder"

var (
	b  = builder.NewBuilder(builder.Options{})
	bi = builder.NewBuilder(
		builder.Options{
			Prefix: "> ",
			Indent: "  ",
		})
	bm4 = builder.NewBuilder(builder.Options{
		FlatSets: true,
	})
	bim4 = builder.NewBuilder(
		builder.Options{
			Prefix:   "> ",
			Indent:   "  ",
			FlatSets: true,
		})
)

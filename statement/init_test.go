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
		Quote:    "`",
		FlatSets: true,
	})
	bim4 = builder.NewBuilder(
		builder.Options{
			Prefix:   "> ",
			Indent:   "  ",
			Quote:    "`",
			FlatSets: true,
		})
)

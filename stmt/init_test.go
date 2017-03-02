package stmt_test

import (
	"github.com/minodisk/sqlabble/builder"
	"github.com/minodisk/sqlabble/token"
)

var (
	b    = builder.Standard
	bi   = builder.NewBuilder(token.NewFormat("> ", "  ", `"`, "\n"))
	bm4  = builder.MySQL
	bim4 = builder.NewBuilder(token.NewFormat("> ", "  ", "`", "\n"))
)

package stmt_test

import (
	"github.com/sqlabble/sqlabble/builder"
	"github.com/sqlabble/sqlabble/token"
)

var (
	b    = builder.Standard
	bi   = builder.NewBuilder(token.NewFormat("> ", "  ", `"`, "\n"))
	bm4  = builder.MySQL
	bim4 = builder.NewBuilder(token.NewFormat("> ", "  ", "`", "\n"))
)

package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func Build(c grammar.Statement) (string, []interface{}) {
	return c.Generator().Generate(generator.NonBreakingContext)
}

func BuildIndent(c grammar.Statement, prefix, indent string) (string, []interface{}) {
	return c.Generator().Generate(generator.NewContext(prefix, indent))
}

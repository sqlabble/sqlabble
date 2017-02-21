package builder

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/printer"
	"github.com/minodisk/sqlabble/statement"
)

// Builder is a container for storing options
// so that you can build a query with the same options
// over and over again.
type Builder struct {
	context node.Context
}

// NewBuilder returns a Builder with a specified options.
func NewBuilder(options Options) Builder {
	return Builder{
		context: options.ToContext(),
	}
}

// Build converts a statement into a query and a slice of values.
func (b Builder) Build(stmt statement.Statement) (string, []interface{}) {
	tokenizer, values := statement.Nodeize(stmt)
	query := printer.Print(tokenizer.Tokenize(0), b.context)
	if len(values) == 0 {
		values = nil
	}
	return query, values
}

// Typical builders commonly used to build queries.
var (
	Standard = Builder{
		context: node.Standard,
	}
	IndentedStandard = Builder{
		context: node.IndentedStandard,
	}
	MySQL = Builder{
		context: node.MySQL,
	}
	MySQL4 = Builder{
		context: node.MySQL4,
	}
)

// Build builds a query.
func Build(s statement.Statement) (string, []interface{}) {
	return Standard.Build(s)
}

// BuildIndent builds an indented query.
func BuildIndent(s statement.Statement) (string, []interface{}) {
	return IndentedStandard.Build(s)
}

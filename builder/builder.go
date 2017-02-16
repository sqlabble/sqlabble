package builder

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/statement"
)

// Builder is a container for storing options
// so that you can build a query with the same options
// over and over again.
type Builder struct {
	context generator.Context
}

// NewBuilder returns a Builder with a specified options.
func NewBuilder(options Options) Builder {
	return Builder{
		context: options.ToContext(),
	}
}

// Build converts a statement into a query and a slice of values.
func (b Builder) Build(stmt statement.Statement) (string, []interface{}) {
	return statement.Node(stmt).ToSQL(b.context)
}

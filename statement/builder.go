package statement

import "github.com/minodisk/sqlabble/generator"

// Builder is a container for storing options
// so that you can build a query with the same options
// over and over again.
type Builder struct {
	context generator.Context
}

// NewBuilder returns a Builder with a specified options.
func NewBuilder(options generator.Options) Builder {
	return Builder{
		context: options.ToContext(),
	}
}

// Build converts a statement into a query and a slice of values.
func (b Builder) Build(statement Statement) (string, []interface{}) {
	return statement.node().ToSQL(b.context)
}

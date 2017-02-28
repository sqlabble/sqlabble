package builder

import (
	"github.com/minodisk/sqlabble/statement"
	"github.com/minodisk/sqlabble/token"
)

// Typical builders commonly used to build queries.
var (
	Standard         = NewBuilder(token.StandardFormat)
	StandardIndented = NewBuilder(token.StandardIndentedFormat)
	MySQL            = NewBuilder(token.MySQLFormat)
	MySQLIndented    = NewBuilder(token.MySQLIndentedFormat)
)

// Builder is a container for storing options
// so that you can build a query with the same options
// over and over again.
type Builder struct {
	Format token.Format
}

// NewBuilder returns a Builder with a specified options.
func NewBuilder(format token.Format) Builder {
	return Builder{
		Format: format,
	}
}

// Build converts a statement into a query and a slice of values.
func (b Builder) Build(n statement.Statement) (string, []interface{}) {
	// if n, ok := stmt.(statement.Nodeizer); ok {
	root := statement.Traverse(n)
	tokenizers, values := statement.Nodize(root)
	sql := token.Generate(tokenizers.Tokenize(0), b.Format)
	if len(values) == 0 {
		values = nil
	}
	return sql, values
	// }

	// tokenizer, values := statement.Nodeize(stmt)
	// query := token.Generate(tokenizer.Tokenize(0), b.Format)
	// if len(values) == 0 {
	// 	values = nil
	// }
	// return query, values
}

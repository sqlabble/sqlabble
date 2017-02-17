package builder

import "github.com/minodisk/sqlabble/node"

// Options is an option on how to build the query.
// It includes format settings and settings to
// accommodate SQL server specific specifications.
type Options struct {
	Prefix, Indent, Quote string
	FlatSets              bool
}

// ToContext converts Options to node.Context.
func (o Options) ToContext() node.Context {
	return node.NewContext(
		o.Prefix,
		o.Indent,
		o.Quote,
		o.FlatSets,
	)
}

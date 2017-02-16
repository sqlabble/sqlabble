package generator

// Options is an option on how to build the query.
// It includes format settings and settings to
// accommodate SQL server specific specifications.
type Options struct {
	Prefix, Indent string
	FlatSets       bool
}

// ToContext converts Options to generator.Context.
func (o Options) ToContext() Context {
	return newContext(o)
}

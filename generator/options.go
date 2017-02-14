package generator

type Options struct {
	Prefix, Indent string
	FlatSets       bool
}

func (o Options) ToContext() Context {
	return newContext(o)
}

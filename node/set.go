package node

// Set is a set of nodes like UNION.
type Set struct {
	separator  Expression
	generators []Node
}

// NewSet returns Set.
func NewSet(separator Expression, generators ...Node) Set {
	return Set{
		separator:  separator,
		generators: generators,
	}
}

// ToSQL returns a query and a slice of values.
func (us Set) ToSQL(ctx Context) (string, []interface{}) {
	res := []Node{}
	for i, g := range us.generators {
		if needsParentheses(ctx, g) {
			g = NewParentheses(g)
		}
		if i == 0 {
			res = append(res, g)
			continue
		}
		res = append(res, us.separator, g)
	}
	return NewNodes(res...).ToSQL(ctx)
}

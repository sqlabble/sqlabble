package generator

type Unions struct {
	separator  Expression
	generators []Node
}

func NewUnions(separator Expression, generators ...Node) Unions {
	return Unions{
		separator:  separator,
		generators: generators,
	}
}

func (us Unions) ToSQL(ctx Context) (string, []interface{}) {
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
	return NewParallelNodes(res...).ToSQL(ctx)
}

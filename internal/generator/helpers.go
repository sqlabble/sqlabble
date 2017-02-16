package generator

func placeholders(i int) string {
	s := ""
	for ; i > 0; i-- {
		if i > 1 {
			s += "?, "
			continue
		}
		s += "?"
	}
	return s
}

func needsParentheses(ctx Context, generator Node) bool {
	if !ctx.flatSets {
		return true
	}

	gs, ok := generator.(ParallelNodes)
	if !ok {
		return true
	}

	for _, g := range gs {
		if _, ok := g.(Unions); !ok {
			return true
		}
	}

	return false
}

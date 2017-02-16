package statement

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

func valuesNodes(p vals) []vals {
	ts := []vals{}
	for p != nil {
		ts = append([]vals{p}, ts...)
		p = p.previous()
	}
	return ts
}

func clauseNodes(c clause) []clause {
	cs := []clause{}
	for c != nil {
		cs = append([]clause{c}, cs...)
		c = c.previous()
	}
	return cs
}

func tableNodes(c joiner) []joiner {
	cs := []joiner{}
	for c != nil {
		cs = append([]joiner{c}, cs...)
		c = c.previous()
	}
	return cs
}

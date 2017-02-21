package statement

// func Node(stmt Statement) node.Node {
// 	return stmt.node()
// }

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

func valuesNodes(p Vals) []Vals {
	ts := []Vals{}
	for p != nil {
		ts = append([]Vals{p}, ts...)
		p = p.previous()
	}
	return ts
}

func clauseNodes(c Clause) []Clause {
	cs := []Clause{}
	for c != nil {
		cs = append([]Clause{c}, cs...)
		c = c.previous()
	}
	return cs
}

func tableNodes(c Joiner) []Joiner {
	cs := []Joiner{}
	for c != nil {
		cs = append([]Joiner{c}, cs...)
		c = c.previous()
	}
	return cs
}

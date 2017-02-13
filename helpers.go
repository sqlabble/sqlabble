package sqlabble

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

func valuesNodes(p valuesNode) []valuesNode {
	ts := []valuesNode{}
	for p != nil {
		ts = append([]valuesNode{p}, ts...)
		p = p.previous()
	}
	return ts
}

func clauseNodes(c clauseNode) []clauseNode {
	cs := []clauseNode{}
	for c != nil {
		cs = append([]clauseNode{c}, cs...)
		c = c.previous()
	}
	return cs
}

func tableNodes(c tableNode) []tableNode {
	cs := []tableNode{}
	for c != nil {
		cs = append([]tableNode{c}, cs...)
		c = c.previous()
	}
	return cs
}

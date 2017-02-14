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

func tableNodes(c tableOrTableAs) []tableOrTableAs {
	cs := []tableOrTableAs{}
	for c != nil {
		cs = append([]tableOrTableAs{c}, cs...)
		c = c.previous()
	}
	return cs
}

package grammar

func PrevValues(p Values) []Values {
	ts := []Values{}
	for p != nil {
		ts = append([]Values{p}, ts...)
		p = p.Prev()
	}
	return ts
}

func Clauses(c Clause) []Clause {
	cs := []Clause{}
	for c != nil {
		cs = append([]Clause{c}, cs...)
		c = c.Prev()
	}
	return cs
}

func Tables(c Table) []Table {
	cs := []Table{}
	for c != nil {
		cs = append([]Table{c}, cs...)
		c = c.Prev()
	}
	return cs
}

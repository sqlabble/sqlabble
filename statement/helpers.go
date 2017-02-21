package statement

import "github.com/minodisk/sqlabble/token"

// func Node(stmt Statement) node.Node {
// 	return stmt.node()
// }

func nodeizeClauses(c Clause) (token.Tokenizer, []interface{}) {
	clauses := collectClauses(c)
	ts := make(token.Tokenizers, len(clauses))
	vals := []interface{}{}
	for i, clause := range clauses {
		var vs []interface{}
		ts[i], vs = clause.self()
		vals = append(vals, vs...)
	}
	return ts, vals
}

func nodeizeJoiners(j Joiner) (token.Tokenizer, []interface{}) {
	joiners := collectJoiners(j)
	ts := make(token.Tokenizers, len(joiners))
	values := []interface{}{}
	for i, j := range joiners {
		var vals []interface{}
		ts[i], vals = j.self()
		values = append(values, vals...)
	}
	return ts, values
}

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

func collectClauses(c Clause) []Clause {
	cs := []Clause{}
	for c != nil {
		cs = append([]Clause{c}, cs...)
		c = c.previous()
	}
	return cs
}

func collectJoiners(c Joiner) []Joiner {
	cs := []Joiner{}
	for c != nil {
		cs = append([]Joiner{c}, cs...)
		c = c.previous()
	}
	return cs
}

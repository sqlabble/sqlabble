package statement

import "github.com/minodisk/sqlabble/token"

func Nodeize(stmt Statement) (token.Tokenizer, []interface{}) {
	return stmt.nodeize()
}

func nodeizeClauses(c Clause) (token.Tokenizer, []interface{}) {
	clauses := collectClauses(c)
	ts := make(token.Tokenizers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		ts[i], vals = c.self()
		values = append(values, vals...)
	}
	return ts, values
}

func collectClauses(c Clause) []Clause {
	cs := []Clause{}
	for c != nil {
		cs = append([]Clause{c}, cs...)
		c = c.previous()
	}
	return cs
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

func collectJoiners(c Joiner) []Joiner {
	cs := []Joiner{}
	for c != nil {
		cs = append([]Joiner{c}, cs...)
		c = c.previous()
	}
	return cs
}

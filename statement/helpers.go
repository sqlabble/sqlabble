package statement

import "github.com/minodisk/sqlabble/tokenizer"

func Nodeize(stmt Statement) (tokenizer.Tokenizer, []interface{}) {
	return stmt.nodeize()
}

type Prever interface {
	previous() Prever
	nodeizeSelf() (tokenizer.Tokenizer, []interface{})
}

func nodeizePrevs(c Prever) (tokenizer.Tokenizer, []interface{}) {
	cs := []Prever{}
	for c != nil {
		cs = append([]Prever{c}, cs...)
		c = c.previous()
	}

	ts := make(tokenizer.Tokenizers, len(cs))
	values := []interface{}{}
	for i, c := range cs {
		var vals []interface{}
		ts[i], vals = c.nodeizeSelf()
		values = append(values, vals...)
	}
	return ts, values
}

func nodeizeClauses(c Clause) (tokenizer.Tokenizer, []interface{}) {
	clauses := collectClauses(c)
	ts := make(tokenizer.Tokenizers, len(clauses))
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

func nodeizeJoiners(j Joiner) (tokenizer.Tokenizer, []interface{}) {
	joiners := collectJoiners(j)
	ts := make(tokenizer.Tokenizers, len(joiners))
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

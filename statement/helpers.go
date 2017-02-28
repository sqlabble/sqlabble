package statement

import "github.com/minodisk/sqlabble/tokenizer"

func Nodeize(stmt Statement) (tokenizer.Tokenizer, []interface{}) {
	return stmt.nodeize()
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

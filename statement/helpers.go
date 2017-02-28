package statement

import "github.com/minodisk/sqlabble/tokenizer"

func Nodeize(stmt Statement) (tokenizer.Tokenizer, []interface{}) {
	return stmt.nodeize()
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

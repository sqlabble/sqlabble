package statement

import "github.com/minodisk/sqlabble/tokenizer"

type Nodeizer interface {
	nodeize() (tokenizer.Tokenizer, []interface{})
}

func Nodize(n Nodeizer) (tokenizer.Tokenizers, []interface{}) {
	var tokenizers tokenizer.Tokenizers
	values := []interface{}{}

	if childer, ok := n.(Childer); ok {
		if child := childer.child(); child != nil {
			t1, vals1 := n.nodeize()
			t2, vals2 := Nodize(child)
			first, _ := t1.FirstLine()
			_, last := t1.LastLine()
			t12 := tokenizer.
				NewContainer(first).
				SetMiddle(t2.Prefix(childer.separator()...)).
				SetLast(last)

			tokenizers = tokenizer.NewTokenizers(t12)
			values = append(append(values, vals1...), vals2...)
		}
	}

	if tokenizers == nil {
		t, vals := n.nodeize()
		tokenizers = tokenizer.NewTokenizers(t)
		values = vals
	}

	if nexter, ok := n.(Nexter); ok {
		if next := nexter.next(); next != nil {
			ts, vals := Nodize(next)
			tokenizers = append(tokenizers, ts...)
			values = append(values, vals...)
		}
	}

	return tokenizers, values
}

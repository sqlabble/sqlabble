package statement

import "github.com/minodisk/sqlabble/tokenizer"

type Nodeizer interface {
	nodeize() (tokenizer.Tokenizer, []interface{})
}

func Nodize(n Nodeizer) (tokenizer.Tokenizers, []interface{}) {
	var tokenizers tokenizer.Tokenizers
	values := []interface{}{}

	ns := []Nodeizer{}

	{
		p := n
		for {
			prever, ok := p.(Prever)
			if !ok {
				break
			}
			p = prever.prev()
			if p == nil {
				break
			}
			ns = append([]Nodeizer{p}, ns...)
		}
	}
	ns = append(ns, n)
	{
		p := n
		for {
			nexter, ok := p.(Nexter)
			if !ok {
				break
			}
			p = nexter.next()
			if p == nil {
				break
			}
			ns = append(ns, p)
		}
	}

	for _, n := range ns {
		if childer, ok := n.(Childer); ok {
			for _, child := range childer.children() {
				t1, vals1 := n.nodeize()
				t2, vals2 := Nodize(child)
				first, _ := t1.FirstLine()
				_, last := t1.LastLine()
				t12 := tokenizer.
					NewContainer(first).
					SetMiddle(t2.Prefix(childer.separator()...)).
					SetLast(last)

				tokenizers = append(tokenizers, t12)
				values = append(append(values, vals1...), vals2...)
			}
			continue
		}

		t1, values1 := n.nodeize()
		tokenizers = append(tokenizers, t1)
		values = append(values, values1...)
	}

	return tokenizers, values
}

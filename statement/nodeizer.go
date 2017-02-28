package statement

import "github.com/minodisk/sqlabble/tokenizer"

func Nodize(stmt Statement) (tokenizer.Tokenizers, []interface{}) {
	var tokenizers tokenizer.Tokenizers
	values := []interface{}{}

	ns := []Statement{}

	{
		p := stmt
		for {
			prever, ok := p.(Prever)
			if !ok {
				break
			}
			p = prever.prev()
			if p == nil {
				break
			}
			ns = append([]Statement{p}, ns...)
		}
	}

	if c, ok := stmt.(Lister); ok {
		ns = append(ns, c.list()...)
	} else {
		ns = append(ns, stmt)
	}

	{
		p := stmt
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
			t1, vals1 := n.nodeize()
			values = append(values, vals1...)
			first, _ := t1.FirstLine()
			_, last := t1.LastLine()

			children := childer.children()
			ts := make(tokenizer.Tokenizers, len(children))
			for i, child := range children {
				var vals []interface{}
				ts[i], vals = Nodize(child)
				values = append(values, vals...)
			}

			if separatorer, ok := childer.(Separatorer); ok {
				seps := separatorer.separator()
				ts = ts.Prefix(seps...)
			}

			t12 := tokenizer.
				NewContainer(first).
				SetMiddle(ts).
				SetLast(last)
			tokenizers = append(tokenizers, t12)
			continue
		}

		t1, values1 := n.nodeize()
		tokenizers = append(tokenizers, t1)
		values = append(values, values1...)
	}

	return tokenizers, values
}

package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Args []Statement

func (ps Args) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if len(ps) == 0 {
		return tokenizer.NewLine(token.FuncLParen, token.FuncRParen), nil
	}

	var tk tokenizer.Tokenizer
	values := []interface{}{}
	for i, p := range ps {
		t, vals := p.nodeize()
		if i == 0 {
			tk = t
			values = append(values, vals...)
			continue
		}
		tk = tokenizer.ConcatTokenizers(
			tk,
			t,
			tokenizer.NewLine(token.Comma),
		)
		values = append(values, vals...)
	}
	return tk.Prepend(token.FuncLParen).Append(token.FuncRParen), values
}

type ValOrColOrFuncOrSubs []ValOrColOrFuncOrSub

func (vs ValOrColOrFuncOrSubs) ToArgs() Args {
	args := make(Args, len(vs))
	for i, v := range vs {
		args[i] = v
	}
	return args
}

type Func struct {
	name keyword.Func
	args Args
}

func (c Func) Eq(value ValOrColOrFuncOrSub) ComparisonOperation {
	e := NewEq(value)
	e.column = c
	return e
}

func (c Func) NotEq(value ValOrColOrFuncOrSub) ComparisonOperation {
	n := NewNotEq(value)
	n.column = c
	return n
}

func (c Func) Gt(value ValOrColOrFuncOrSub) ComparisonOperation {
	g := NewGt(value)
	g.column = c
	return g
}

func (c Func) Gte(value ValOrColOrFuncOrSub) ComparisonOperation {
	g := NewGte(value)
	g.column = c
	return g
}

func (c Func) Lt(value ValOrColOrFuncOrSub) ComparisonOperation {
	l := NewLt(value)
	l.column = c
	return l
}

func (c Func) Lte(value ValOrColOrFuncOrSub) ComparisonOperation {
	l := NewLte(value)
	l.column = c
	return l
}

func (c Func) Like(value ValOrColOrFuncOrSub) ComparisonOperation {
	l := NewLike(value)
	l.column = c
	return l
}

func (c Func) RegExp(value ValOrColOrFuncOrSub) ComparisonOperation {
	r := NewRegExp(value)
	r.column = c
	return r
}

func (c Func) Between(from, to ValOrColOrFuncOrSub) Between {
	b := NewBetween(from, to)
	b.column = c
	return b
}

func (c Func) In(arg ValsOrSub) ContainingOperation {
	i := NewIn(arg)
	i.column = c
	return i
}

func (c Func) NotIn(arg ValsOrSub) ContainingOperation {
	n := NewNotIn(arg)
	n.column = c
	return n
}

func (c Func) EqAll(arg Subquery) NonScalarOperation {
	n := NewEqAll(arg)
	n.column = c
	return n
}

func (c Func) NotEqAll(arg Subquery) NonScalarOperation {
	n := NewNotEqAll(arg)
	n.column = c
	return n
}

func (c Func) GtAll(arg Subquery) NonScalarOperation {
	n := NewGtAll(arg)
	n.column = c
	return n
}

func (c Func) GteAll(arg Subquery) NonScalarOperation {
	n := NewGteAll(arg)
	n.column = c
	return n
}

func (c Func) LtAll(arg Subquery) NonScalarOperation {
	n := NewLtAll(arg)
	n.column = c
	return n
}

func (c Func) LteAll(arg Subquery) NonScalarOperation {
	n := NewLteAll(arg)
	n.column = c
	return n
}

func (c Func) EqAny(arg Subquery) NonScalarOperation {
	n := NewEqAny(arg)
	n.column = c
	return n
}

func (c Func) NotEqAny(arg Subquery) NonScalarOperation {
	n := NewNotEqAny(arg)
	n.column = c
	return n
}

func (c Func) GtAny(arg Subquery) NonScalarOperation {
	n := NewGtAny(arg)
	n.column = c
	return n
}

func (c Func) GteAny(arg Subquery) NonScalarOperation {
	n := NewGteAny(arg)
	n.column = c
	return n
}

func (c Func) LtAny(arg Subquery) NonScalarOperation {
	n := NewLtAny(arg)
	n.column = c
	return n
}

func (c Func) LteAny(arg Subquery) NonScalarOperation {
	n := NewLteAny(arg)
	n.column = c
	return n
}

func (f Func) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, v := f.args.nodeize()
	return t.Prepend(token.Word(f.name)), v
}

// isValOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (f Func) isValOrFuncOrSub() bool {
	return true
}

// isValOrColOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrColOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (f Func) isValOrColOrFuncOrSub() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (f Func) isValOrColOrAliasOrFuncOrSubOrFormula() bool {
	return true
}
